package utils

import (
	"Gogogo/configs"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

// DBConnections структура для хранения соединений с БД
type DBConnections struct {
	Auth *sql.DB // База данных для авторизации
	News *sql.DB //База данных для новостей
}

// Глобальная переменная для хранения всех подключений
var DB DBConnections

// InitDatabases инициализирует подключения к обеим базам данных
func InitDatabase(config configs.Config) {
	// Строки подключения для обеих баз данных
	authConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.AuthUser, config.AuthPassword, config.AuthName, config.AuthHost, config.AuthPort)

	newsConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.NewsUser, config.NewsPassword, config.NewsName, config.NewsHost, config.NewsPort)

	// Подключаемся к первой базе (auth)
	var err error
	DB.Auth, err = sql.Open("postgres", authConnStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных Auth: ", err)
	}

	// Проверка выполнения запроса в БД Auth
	var authDBPingStatus string
	err = DB.Auth.QueryRow("SELECT 1 FROM users").Scan(&authDBPingStatus)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса к базе данных Auth:", err)
	}
	log.Println("Подключение к базе данных Auth успешно установлено. Статус:", authDBPingStatus)

	DB.News, err = sql.Open("postgres", newsConnStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных Edu: ", err)
	}

	// Проверка выполнения запроса в БД Edu
	var newsDBPingStatus string
	err = DB.News.QueryRow("SELECT 1 FROM articles").Scan(&newsDBPingStatus)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса к базе данных Edu:", err)
	}
	log.Println("Подключение к базе данных Edu успешно установлено. Статус:", newsDBPingStatus)

	log.Println("Подключение к базам данных успешно установлено. Статус: OK")
}

// CloseDB закрывает соединения с базами данных
func CloseDB() {
	if DB.Auth != nil {
		err := DB.Auth.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных Auth: %v", err)
		} else {
			log.Println("Соединение с базой данных Auth закрыто.")
		}
	}
	if DB.News != nil {
		err := DB.News.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных Edu: %v", err)
		} else {
			log.Println("Соединение с базой данных Edu закрыто.")
		}
	}
}
