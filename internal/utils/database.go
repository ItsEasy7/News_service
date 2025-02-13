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
	Edu  *sql.DB // База данных для образования
	News *sql.DB //База данных для новостей
}

// Глобальная переменная для хранения всех подключений
var DB DBConnections

// InitDatabases инициализирует подключения к обеим базам данных
func InitDatabase(config configs.Config) {
	// Строки подключения для обеих баз данных
	authConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.AuthHost, config.AuthPort, config.AuthUser, config.AuthPassword, config.AuthName)

	eduConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.EduHost, config.EduPort, config.EduUser, config.EduPassword, config.EduName)

	newsConnStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		config.NewsHost, config.NewsPort, config.NewsUser, config.NewsPassword, config.NewsName)

	// Подключаемся к первой базе (auth)
	var err error
	DB.Auth, err = sql.Open("postgres", authConnStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных Auth: ", err)
	}

	// Проверка выполнения запроса в БД Auth
	var authDBPingStatus string
	err = DB.Auth.QueryRow("SELECT * FROM test").Scan(&authDBPingStatus)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса к базе данных Auth:", err)
	}
	log.Println("Подключение к базе данных Auth успешно установлено. Статус:", authDBPingStatus)

	// Подключаемся ко второй базе (edu)
	DB.Edu, err = sql.Open("postgres", eduConnStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных Edu: ", err)
	}

	// Проверка выполнения запроса в БД Edu
	var eduDBPingStatus string
	err = DB.Edu.QueryRow("SELECT * FROM test").Scan(&eduDBPingStatus)
	if err != nil {
		log.Fatal("Ошибка при выполнении запроса к базе данных Edu:", err)
	}
	log.Println("Подключение к базе данных Edu успешно установлено. Статус:", eduDBPingStatus)

	DB.News, err = sql.Open("postgres", newsConnStr)
	if err != nil {
		log.Fatal("Ошибка при подключении к базе данных Edu: ", err)
	}

	// Проверка выполнения запроса в БД Edu
	var newsDBPingStatus string
	err = DB.News.QueryRow("SELECT * FROM test").Scan(&newsDBPingStatus)
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
	if DB.Edu != nil {
		err := DB.Edu.Close()
		if err != nil {
			log.Printf("Ошибка при закрытии соединения с базой данных Edu: %v", err)
		} else {
			log.Println("Соединение с базой данных Edu закрыто.")
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
