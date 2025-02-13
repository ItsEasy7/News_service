package utils

import (
	"Gogogo/configs"
	"log"

	"github.com/streadway/amqp"
)

// RabbitMQConn глобальная переменная для соединения с RabbitMQ
var RabbitMQClient *amqp.Connection

// InitRabbitMQ инициализирует соединение с RabbitMQ, используя конфигурацию из configs
func InitRabbitMQ(config configs.Config) {
	// Получаем строку подключения из конфигурации
	rabbitURL := config.RabbitMQURL
	if rabbitURL == "" {
		log.Fatal("Конфигурация RabbitMQURL не установлена")
	}

	// Подключаемся к RabbitMQ
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Ошибка при подключении к RabbitMQ: %v", err)
	}

	// Присваиваем глобальной переменной
	RabbitMQClient = conn

	log.Println("Подключение к RabbitMQ успешно установлено. Статус: OK")
}
