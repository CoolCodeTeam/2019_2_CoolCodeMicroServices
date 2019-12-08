package main

import (
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/messages/delivery"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/messages/repository"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/messages/usecase"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils"
	"github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	middleware "github.com/CoolCodeTeam/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
)

const (
	users_address         = "localhost:5000"
	chats_adress          = "localhost:5001"
	notifications_address = "localhost:5002"
)

var (
	uuid_regexp = "[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}"
)

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

	viper.AddConfigPath("./messages")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrusLogger.Error("Can`t get viper config:" + err.Error())
	}

	consulCfg := viper.GetStringMapString("consul")

	consul := utils.GetConsul(consulCfg["url"])
	configs := utils.LoadConfig(consul, consulCfg["prefix"])

	dbconfig := utils.DBConfig{
		configs["db_name"],
		configs["db_user"],
		configs["db_password"],
		configs["db_host"],
	}
	port := ":" + configs["port"]

	//Connect database
	db, err := utils.ConnectDatabase(dbconfig)

	//connect to users
	usersConn := utils.ConnectGRPC(users_address)
	defer usersConn.Close()

	//connect to chats
	chatsConn := utils.ConnectGRPC(chats_adress)
	defer chatsConn.Close()

	//connect to notifications
	notificationsConn := utils.ConnectGRPC(notifications_address)
	defer chatsConn.Close()

	chatsClient := grpc_utils.NewChatsServiceClient(chatsConn)
	chats := grpc_utils.NewChatsGRPCProxy(chatsClient)

	users := grpc_utils.NewUsersGRPCProxy(grpc_utils.NewUsersServiceClient(usersConn))

	messages := useCase.NewMessageUseCase(repository.NewMessageDbRepository(db), chats)
	handlersUtils := utils.NewHandlersUtils(logrusLogger)
	messagesApi := delivery.NewMessageHandlers(messages, users,
		handlersUtils, grpc_utils.NewNotificationsGRPCProxy(grpc_utils.NewNotificationsServiceClient(notificationsConn)))

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
			"http://coolcode.site"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	handler := middlewares.PanicMiddleware(middlewares.LogMiddleware(r, logrusLogger))
	r.Handle("/messages/channels/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.SendMessage)).Methods("POST")
	r.Handle("/messages/channels/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.GetMessagesByChatID)).Methods("GET")
	r.Handle("/messages/chats/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.SendMessage)).Methods("POST").
		HeadersRegexp("Content-Type", "application/(text|json)")
	r.Handle("/messages/chats/{id:[0-9]+}/files", middlewares.AuthMiddleware(messagesApi.SendFile)).Methods("POST").
		HeadersRegexp("Content-Type", "multipart/form-data")
	r.Handle("/messages/chats/{id:[0-9]+}/files/{uid:[0-9a-fA-F]{8}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{4}\\-[0-9a-fA-F]{12}}",
		middlewares.AuthMiddleware(messagesApi.GetFile)).Methods("GET")
	r.Handle("/messages/chats/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.GetMessagesByChatID)).Methods("GET")
	r.Handle("/messages/{text:[\\s\\S]+}", middlewares.AuthMiddleware(messagesApi.FindMessages)).Methods("GET")
	r.Handle("/messages/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.DeleteMessage)).Methods("DELETE")
	r.Handle("/messages/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.EditMessage)).Methods("PUT")
	r.Handle("/messages/{id:[0-9]+}/likes", middlewares.AuthMiddleware(messagesApi.Like)).Methods("POST")
	r.Handle("/metrics", promhttp.Handler())
	logrus.Infof("Messages http server started on %s port: ", port)
	err = http.ListenAndServe(port, corsMiddleware(handler))

	if err != nil {
		logrusLogger.Error(err)
		return
	}
}
