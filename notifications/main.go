package main

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/delivery"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/notifications_service"
	useCase "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	middleware "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
	users_address = "localhost:5000"
	chats_adress  = "localhost:5001"
)

func startNotificationsGRPCService(port string, service grpc_utils.NotificationsServiceServer) {
	lis, err := net.Listen("tcp", ":5002")
	if err != nil {
		//
	}
	s := grpc.NewServer()

	grpc_utils.RegisterNotificationsServiceServer(s, service)

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("Chats gRPC service failed at port %d", 5002)
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

	viper.AddConfigPath("./notifications")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrusLogger.Error("Can`t get viper config:" + err.Error())
	}

	//FIXME:
	//consulCfg := viper.GetStringMapString("consul")
	//consul := utils.GetConsul(consulCfg["url"])
	//configs := utils.LoadConfig(consul, consulCfg["prefix"])

	consul := utils.GetConsul("95.163.209.195:8010")
	configs := utils.LoadConfig(consul, "notifications")

	port := ":" + configs["port"]

	handlersUtils := utils.NewHandlersUtils(logrusLogger)
	notificationsUseCase := useCase.NewNotificationUseCase()
	users := grpc_utils.NewUsersGRPCProxy(grpc_utils.NewUsersServiceClient(utils.ConnectGRPC(users_address)))
	notificationApi := delivery.NewNotificationHandlers(users,
		grpc_utils.NewChatsGRPCProxy(grpc_utils.NewChatsServiceClient(utils.ConnectGRPC(chats_adress))), notificationsUseCase, handlersUtils)

	startNotificationsGRPCService("5002", notifications_service.NewNotificationsGRPCService(notificationsUseCase))

	middlewares := middleware.HandlersMiddlwares{
		Users:  users,
		Logger: logrusLogger,
	}

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://boiling-chamber-90136.herokuapp.com",
			"https://boiling-chamber-90136.herokuapp.com",
			"http://localhost:3000",
			"http://localhost:8000",
			"http://95.163.209.195:8000",
			"http://coolcode.site"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	handler := middlewares.PanicMiddleware(middlewares.LogMiddleware(r, logrusLogger))
	r.Handle("/notifications/chats/{id:[0-9]+}", middlewares.AuthMiddleware(notificationApi.HandleNewWSConnection))
	r.Handle("/notifications/channels/{id:[0-9]+}", middlewares.AuthMiddleware(notificationApi.HandleNewWSConnection))
	r.Handle("/metrics", promhttp.Handler())
	logrus.Infof("Notifications http server started on %s port: ", port)
	err = http.ListenAndServe(port, corsMiddleware(handler))
	if err != nil {
		logrusLogger.Error(err)
		return
	}
}
