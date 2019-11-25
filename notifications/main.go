package main

import (
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/delivery"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/notifications_service"
	useCase "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/notifications/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	middleware "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	consulAddr = flag.String("consul", "95.163.209.195:8010", "consul addr")

	consul          *consulapi.Client
	consulLastIndex uint64 = 0

	globalCfg   = make(map[string]string)
	globalCfgMu = &sync.RWMutex{}

	cfgPrefix      = "notifications/"
	prefixStripper = strings.NewReplacer(cfgPrefix, "")
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

	port := ":" + globalCfg["port"]

	handlersUtils := utils.NewHandlersUtils(logrusLogger)
	notificationsUseCase := useCase.NewNotificationUseCase()
	users := grpc_utils.NewUsersGRPCProxy(grpc_utils.NewUsersServiceClient(connectGRPC("5000")))
	notificationApi := delivery.NewNotificationHandlers(users,
		grpc_utils.NewChatsGRPCProxy(grpc_utils.NewChatsServiceClient(connectGRPC("5001"))), notificationsUseCase, handlersUtils)

	startNotificationsGRPCService("5002", notifications_service.NewNotificationsGRPCService(notificationsUseCase))

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
	r.Handle("/chats/{id:[0-9]+}/notifications", middlewares.AuthMiddleware(notificationApi.HandleNewWSConnection))
	r.Handle("/metrics", promhttp.Handler())
	logrus.Info("Server started")
	err = http.ListenAndServe(port, corsMiddleware(handler))
	if err != nil {
		logrusLogger.Error(err)
		return
	}
}
