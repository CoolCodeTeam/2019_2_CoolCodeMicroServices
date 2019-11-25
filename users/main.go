package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/delivery"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/repository"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/usecase"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/users/users_service"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils"
	"github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/grpc_utils"
	middleware "github.com/go-park-mail-ru/2019_2_CoolCodeMicroServices/utils/middlwares"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"net"
	"net/http"
	"os"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "1"
	DB_NAME     = "postgres"
)

var (
	redisAddr = flag.String("addr", "redis://localhost:6379", "redis addr")
)

func connectDatabase() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return db, err
	}
	if db == nil {
		return db, errors.New("Can not connect to database")
	}
	return db, nil

}

func startUsersGRPCService(port string, service grpc_utils.UsersServiceServer) {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		//
	}
	s := grpc.NewServer()

	grpc_utils.RegisterUsersServiceServer(s, service)

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			logrus.Fatalf("Chats gRPC service failed at port %d", 5000)
			os.Exit(1)
		}
	}()
}

func main() {
	//Connect databases
	redis := connectRedis()
	db, err := connectDatabase()

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

	sessionRepository := repository.NewSessionRedisStore(redis)
	users := useCase.NewUserUseCase(repository.NewUserDBStore(db), sessionRepository)
	handlersUtils := utils.NewHandlersUtils(logrusLogger)
	usersApi := delivery.NewUsersHandlers(users, sessionRepository, handlersUtils)

	// Стартуем наш gRPC сервер для прослушивания tcp
	startUsersGRPCService("5000", users_service.NewGRPCUsersService(users))

	//Cлушаем HTTP
	middlewares := middleware.HandlersMiddlwares{
		Users:  users,
		Logger: logrusLogger,
	}

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://boiling-chamber-90136.herokuapp.com",
			"https://boiling-chamber-90136.herokuapp.com",
			"http://localhost:3000",
			"http://95.163.209.195:8000",
			"http://localhost:8000"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowCredentials(),
	)

	r := mux.NewRouter()
	handler := middlewares.PanicMiddleware(middlewares.LogMiddleware(r, logrusLogger))
	r.HandleFunc("/users", usersApi.SignUp).Methods("POST")
	r.HandleFunc("/users/login", usersApi.Login).Methods("POST")
	r.Handle("/users/{id:[0-9]+}", middlewares.AuthMiddleware(usersApi.EditProfile)).Methods("PUT")
	r.Handle("/users/logout", middlewares.AuthMiddleware(usersApi.Logout)).Methods("DELETE")
	r.Handle("/users/photos", middlewares.AuthMiddleware(usersApi.SavePhoto)).Methods("POST")
	r.Handle("/users/photos/{id:[0-9]+}", middlewares.AuthMiddleware(usersApi.GetPhoto)).Methods("GET")
	r.Handle("/users/{id:[0-9]+}", middlewares.AuthMiddleware(usersApi.GetUser)).Methods("GET")
	r.Handle("/users/names/{name:[\\s\\S]+}", middlewares.AuthMiddleware(usersApi.FindUsers)).Methods("GET")
	r.HandleFunc("/users", usersApi.GetUserBySession).Methods("GET") //TODO:Добавить в API
	logrus.Info("Server started")
	err = http.ListenAndServe(":8001", corsMiddleware(handler))
	if err != nil {
		logrusLogger.Error(err)
		return
	}

}

func connectRedis() *redis.Pool {
	redisConn := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.DialURL(*redisAddr)
		},
	}
	return redisConn
}
