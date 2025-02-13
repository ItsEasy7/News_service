package utils

import (
	"Gogogo/configs"
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	// креды для подключения
}

// RedisClient глобальная переменная для клиента Redis
var RedisClient *redis.Client

// InitRedis инициализирует подключение к Redis
func InitRedis(config configs.Config) {

	// Получаем строку подключения к Redis
	redisURL := config.CACHER
	if redisURL == "" {
		log.Fatal("Переменная окружения CACHER не установлена")
	}

	// Настройка клиента Redis
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Fatalf("Ошибка при парсинге URL для Redis: %v", err)
	}

	// Инициализация клиента Redis
	RedisClient = redis.NewClient(opt)

	// Проверяем подключение
	ctx := context.Background()
	_, err = RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка при подключении к Redis: %v", err)
	}

	// Проверяем статус в Redis
	status, err := RedisClient.Get(ctx, "status").Result()
	if err != nil {
		if err == redis.Nil {
			log.Println("Ключ 'status' не найден в Redis")
		} else {
			log.Fatalf("Ошибка при получении ключа 'status' из Redis: %v", err)
		}
	}

	log.Println("Подключение к Redis успешно установлено. Статус:", status)
}
