package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"todoApp/internal/server/grpc"
	"todoApp/internal/server/rest"
	"todoApp/pkg/handler"
	"todoApp/pkg/repository"
	"todoApp/pkg/service"
)

// @title TodoApp API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	restServer := new(rest.Server)
	go func() {
		if err := restServer.Run(viper.GetString("rest_port"), handlers.InitRoutes()); err != nil {
			if err != http.ErrServerClosed {
				logrus.Fatalf("error uccured while running http restServer %s", err.Error())
			}
		}
	}()
	grpcServer := grpc.NewServer(services, logrus.New())
	go func() {
		if err := grpcServer.Run(viper.GetString("grpc_port")); err != nil {
			logrus.Fatalf("error uccured while running tcp grpcServer %s", err.Error())
		}
	}()

	logrus.Println("Server Started")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server is Shutting Down")

	if err := restServer.ShutDown(context.Background()); err != nil {
		logrus.Errorf("error occured while server shutting down: %s", err.Error())
	}

	grpcServer.ShutDown()

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured while closing db: %s", err.Error())
	}

	logrus.Println("Server closed successfully")
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
