package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"wanna-go-out-api/domain"
)

type Orm struct {
	Db *gorm.DB
}

func NewORM(env *Env) *Orm {
	orm := &Orm{}

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", env.DbHost, env.DbUser, env.DbPassword, env.DbSid, env.DbPort)

	DB, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	err = DB.AutoMigrate(&domain.Message{})
	if err != nil {
		log.Panic(err.Error())
	}

	orm.Db = DB

	return orm
}
