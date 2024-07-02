package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Setting struct {
	Host     string
	User     string
	Password string
	Port     string
	DBNAME   string
}

var DB *gorm.DB

func ImportSetting() Setting {
	var result Setting
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return Setting{}
	}
	result.Host = os.Getenv("poshost")
	result.User = os.Getenv("posuser")
	result.Password = os.Getenv("pospw")
	result.Port = os.Getenv("posport")
	result.DBNAME = os.Getenv("dbname")
	return result
}

func ConnectDB(s Setting) (*gorm.DB, error) {
	var connStr = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s",
		s.Host, s.User, s.Password, s.Port, s.DBNAME)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "api1.",
		},
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
