package server

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ostheperson/go-auth-service/internal/domain"
)

type dbclient struct {
	db *gorm.DB
}

func New(env *domain.Env) domain.Service {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.DB_USERNAME,
		env.DB_PASSWORD,
		env.DB_HOST,
		env.DB_PORT,
		env.DB_DATABASE,
	)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	s := &dbclient{db: db}
	return s
}

func (s *dbclient) GetClient() *gorm.DB {
	return s.db
}

func (s *dbclient) Health() map[string]string {
	sqlDB, err := s.db.DB()
	if err != nil {
		panic("failed to get underlying *sql.DB instance")
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
