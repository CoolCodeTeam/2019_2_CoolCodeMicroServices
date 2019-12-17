package main

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/chats_service"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/delivery"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/repository"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/chats/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	middleware "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"net"
	"net/http"
	"os"
)

const (
	users_address = "127.0.0.1:5000"
)

func startChatsGRPCService(port string, service grpc_utils.ChatsServiceServer) {
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		//
	}
	s := grpc.NewServer()

	grpc_utils.RegisterChatsServiceServer(s, service)

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("Chats gRPC service failed at port %d", 5001)
			os.Exit(1)
		}
	}()
}

func main() {
	//Init logrus
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	f, err := os.OpenFile("logs.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrusLogger.Error("Can`t open file:" + err.Error())
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stderr, f)
	logrusLogger.SetOutput(mw)

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrusLogger.Error("Can`t get viper config:" + err.Error())
	}

	//FIXME:
	//consulCfg := viper.GetStringMapString("consul")
	//consul := utils.GetConsul(consulCfg["url"])
	//configs := utils.LoadConfig(consul, consulCfg["prefix"])

	consul := utils.GetConsul("95.163.209.195:8010")
	configs := utils.LoadConfig(consul, "chats")

	dbconfig := utils.DBConfig{
		configs["db_name"],
		configs["db_user"],
		configs["db_password"],
		configs["db_host"],
	}
	port := ":" + configs["port"]

	//Connect database
	db, err := utils.ConnectDatabase(dbconfig)

	//Connect to users
	conn := utils.ConnectGRPC(users_address)
	defer conn.Close()

	client := grpc_utils.NewUsersServiceClient(conn)
	users := grpc_utils.NewUsersGRPCProxy(client)
	chats := useCase.NewChatsUseCase(repository.NewChatsDBRepository(db), users)
	handlersUtils := utils.NewHandlersUtils(logrusLogger)
	chatsApi := delivery.NewChatHandlers(users, chats, handlersUtils)

	// Стартуем наш gRPC сервер для прослушивания tcp
	startChatsGRPCService("5001", chats_service.NewGRPCChatsService(chats))

	middlewares := middleware.HandlersMiddlwares{
		Users:  users,
		Logger: logrusLogger,
	}

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://boiling-chamber-90136.herokuapp.com",
			"https://boiling-chamber-90136.herokuapp.com",
			"http://localhost:3000",
			"http://95.163.209.195:8000",
			"http://localhost:8000",
			"http://coolcode.site",
		}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	handler := middlewares.PanicMiddleware(middlewares.LogMiddleware(r, logrusLogger))
	r.Handle("/chats", middlewares.AuthMiddleware(chatsApi.PostChat)).Methods("POST")
	r.Handle("/chats/users/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.GetChatsByUser)).Methods("GET")
	r.Handle("/chats/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.GetChatById)).Methods("GET")
	r.Handle("/chats/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.RemoveChat)).Methods("DELETE")

	r.Handle("/channels/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.GetChannelById)).Methods("GET")
	r.Handle("/channels/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.EditChannel)).Methods("PUT")
	r.Handle("/channels/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.RemoveChannel)).Methods("DELETE")
	//TODO: r.Handle("/channels/{id:[0-9]+}/members", middlewares.AuthMiddleware(chatsApi.LogoutFromChannel)).Methods("DELETE")
	r.Handle("/workspaces/{id:[0-9]+}/channels", middlewares.AuthMiddleware(chatsApi.PostChannel)).Methods("POST")

	r.Handle("/workspaces/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.GetWorkspaceById)).Methods("GET")
	r.Handle("/workspaces/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.EditWorkspace)).Methods("PUT")
	//TODO: r.Handle("/workspaces/{id:[0-9]+}/members", middlewares.AuthMiddleware(chatsApi.LogoutFromWorkspace)).Methods("DELETE")
	r.Handle("/workspaces/{id:[0-9]+}", middlewares.AuthMiddleware(chatsApi.RemoveWorkspace)).Methods("DELETE")
	r.Handle("/workspaces", middlewares.AuthMiddleware(chatsApi.PostWorkspace)).Methods("POST")
	r.Handle("/metrics", promhttp.Handler())
	logrus.Infof("Chats http server started on %s port: ", port)
	err = http.ListenAndServe(port, corsMiddleware(handler))
	if err != nil {
		logrusLogger.Error(err)
		return
	}

}
