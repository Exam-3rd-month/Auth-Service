package storage

import (
	"database/sql"
	"fmt"
	"log"

	"auth-service/internal/config"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func ConnectDB(config config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		config.Database.User,
		config.Database.DBName,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	log.Printf("--------------------------- Connected to the database %s --------------------------------\n", config.Database.DBName)

	return db, nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// func checkPassword(hashedPassword, password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// 	return err == nil
// }
