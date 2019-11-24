package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/messages/delivery"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/messages/repository"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/messages/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	middleware "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	consulapi "github.com/hashicorp/consul/api"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

type DBConfig struct {
	DBName     string
	DBUser     string
	DBPassword string
}

var (
	consulAddr = flag.String("consul", "95.163.209.195:8010", "consul addr")

	consul          *consulapi.Client
	consulLastIndex uint64 = 0

	globalCfg   = make(map[string]string)
	globalCfgMu = &sync.RWMutex{}

	cfgPrefix      = "messages/"
	prefixStripper = strings.NewReplacer(cfgPrefix, "")
)

const (
	users_address         = "localhost:5000"
	chats_adress          = "localhost:5001"
	notifications_address = "localhost:5002"
)

func connectDatabase(config DBConfig) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		config.DBUser, config.DBPassword, config.DBName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return db, err
	}
	if db == nil {
		return db, errors.New("Can not connect to database")
	}
	return db, nil

}

func connectGRPC(address string) *grpc.ClientConn {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalf("can not connect to usersGRPC %v", err)
	}
	return conn
}

func loadConfig() {
	qo := &consulapi.QueryOptions{
		WaitIndex: consulLastIndex,
	}
	kvPairs, qm, err := consul.KV().List(cfgPrefix, qo)
	if err != nil {
		fmt.Println(err)
		return
	}

	if consulLastIndex == qm.LastIndex {
		return
	}

	newConfig := make(map[string]string)

	for _, item := range kvPairs {
		if item.Key == cfgPrefix {
			continue
		}
		key := prefixStripper.Replace(item.Key)
		newConfig[key] = string(item.Value)
	}

	globalCfgMu.Lock()
	globalCfg = newConfig
	consulLastIndex = qm.LastIndex
	globalCfgMu.Unlock()
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

	consulConfig := consulapi.DefaultConfig()
	consulConfig.Address = *consulAddr
	consul, err = consulapi.NewClient(consulConfig)
	if err != nil {
		logrusLogger.Error("Can`t get consul config:" + err.Error())
	}
	loadConfig()

	dbconfig := DBConfig{
		globalCfg["db_name"],
		globalCfg["db_user"],
		globalCfg["db_password"],
	}
	port := ":" + globalCfg["port"]

	//Connect database
	db, err := connectDatabase(dbconfig)

	//connect to users
	usersConn := connectGRPC(users_address)
	defer usersConn.Close()

	//connect to chats
	chatsConn := connectGRPC(chats_adress)
	defer chatsConn.Close()

	//connect to notifications
	notificationsConn := connectGRPC(notifications_address)
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
		handlers.AllowedOrigins([]string{"http://boiling-chamber-90136.herokuapp.com", "https://boiling-chamber-90136.herokuapp.com", "http://localhost:3000"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	handler := middlewares.PanicMiddleware(middlewares.LogMiddleware(r, logrusLogger))
	r.Handle("/channels/{id:[0-9]+}/messages", middlewares.AuthMiddleware(messagesApi.SendMessage)).Methods("POST")
	r.Handle("/channels/{id:[0-9]+}/messages", middlewares.AuthMiddleware(messagesApi.GetMessagesByChatID)).Methods("GET")
	r.Handle("/chats/{id:[0-9]+}/messages", middlewares.AuthMiddleware(messagesApi.SendMessage)).Methods("POST").
		HeadersRegexp("Content-Type", "application/(text|json)")
	r.Handle("/chats/{id:[0-9]+}/messages", middlewares.AuthMiddleware(messagesApi.GetMessagesByChatID)).Methods("GET")
	r.Handle("/messages/{text:[((a-z)|(A-Z))0-9_-]+}", middlewares.AuthMiddleware(messagesApi.FindMessages)).Methods("GET")
	r.Handle("/messages/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.DeleteMessage)).Methods("DELETE")
	r.Handle("/messages/{id:[0-9]+}", middlewares.AuthMiddleware(messagesApi.EditMessage)).Methods("PUT")
	logrus.Info("Server started")
	err = http.ListenAndServe(port, corsMiddleware(handler))
	if err != nil {
		logrusLogger.Error(err)
		return
	}
}
