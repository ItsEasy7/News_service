package main

import (
	"Gogogo/configs"
	"Gogogo/internal/service/news"
	"Gogogo/internal/utils"
	"log"
)

func main() {
	// Загрузка конфигурации
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Инициализация подключений
	utils.InitDatabase(config)
	utils.InitRabbitMQ(config)
	utils.InitRedis(config)

	// Создание Gin роутера
	router := news.SetupRouter()

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
