package main

import (
	"OpenWeatherMap-API/pkg/handlers"
	"OpenWeatherMap-API/pkg/models"
	"flag"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func DBInit(user, password, dbname, port string) (*gorm.DB, error) {
	dsn := "host=localhost" +
		" user=" + user +
		" password=" + password +
		" dbname=" + dbname +
		" port=" + port +
		" sslmode=disable" +
		" TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.City{},
		&models.Marks{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	DBName := flag.String("dbname", "OpenWeatherMap", "Enter the name of DB")
	DBUser := flag.String("dbuser", "postgres", "Enter the name of a DB user")
	DBPassword := flag.String("dbpassword", "developer", "Enter the password of user")
	DBPort := flag.String("dbport", "5432", "Enter the port of DB")
	flag.Parse()

	db, err := DBInit(*DBUser, *DBPassword, *DBName, *DBPort)
	if err != nil {
		log.Fatal("db connection error:", err)
	}
	log.Println("successfully connected to DB")

	h := handlers.NewHandler(db)

	router := gin.Default()

	router.GET("/city", h.GetOneWeather)

	err = router.Run(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
