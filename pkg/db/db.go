package db

import (
	"os"
	"log"
	"gorm.io/driver/postgres"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"github.com/dimitaryord/go-api/pkg/models"
)

func GetConnectionURI() (connectionURI string) {
	connectionURI, exists := os.LookupEnv("NEONDB_URL")	
	if !exists {
		log.Print("No db url provided.")
		return ""
	} 
	return connectionURI
}


func Init() *gorm.DB{
	if err := godotenv.Load(); err != nil {
		log.Print("No env file")
	} 
	connectionURI := GetConnectionURI()
	db, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&models.Person{})
	return db
}