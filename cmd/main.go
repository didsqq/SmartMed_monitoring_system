package main

import (
	smartmed "github.com/didsqq/SmartMed_monitoring_system"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/handler"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/repository"
	"github.com/didsqq/SmartMed_monitoring_system/pkg/service"
	_ "github.com/lib/pq"
	"github.com/rs/cors" // Импортируем библиотеку для CORS
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Устанавливаем формат логов
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Инициализация конфигурации
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	// Подключение к базе данных
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("error initializing db: %s", err.Error())
	}

	// Создание репозиториев, сервисов и обработчиков
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Настройка CORS
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"*"},                             // Разрешить все домены (замените на конкретные домены, если необходимо)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},  // Разрешенные методы
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Разрешенные заголовки
		AllowCredentials: true,                                      // Разрешить отправку cookies (по необходимости)
	}

	// Создаем обработчик CORS
	corsHandler := cors.New(corsOptions).Handler(handlers.InitRoutes())

	// Инициализируем сервер
	srv := new(smartmed.Server)

	// Запуск сервера с маршрутизатором и CORS
	if err := srv.Run(viper.GetString("port"), corsHandler); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err.Error())
	}
}

// Функция для загрузки конфигурации
func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}
