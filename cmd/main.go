package main

import (
	"Gogogo/configs"
	"Gogogo/internal/service/news"
	"Gogogo/internal/utils"
	"Gogogo/reflection"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Инициализация подключений
	utils.InitDatabase(config)
	utils.InitRabbitMQ(config)
	utils.InitRedis(config)

	// Создание Gin роутера
	router := news.SetupRouter()

	// Регистрация сервисной схемы для Moleculer
	endpointInfo := reflection.CollectEndpointInfo(router, "news", "https://localhost:8080")

	// Сериализация в JSON
	jsonData, err := json.MarshalIndent(endpointInfo, "", "    ")
	if err != nil {
		fmt.Println("Ошибка при сериализации в JSON:", err)
		return
	}

	// Вывод JSON
	// fmt.Println(string(jsonData))

	reflection.RegisterEndpoints(string(jsonData))

	// Запуск сервера
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
