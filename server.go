package main

import (
	"os"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	h "bfp/avi/api/handlers"
	m "bfp/avi/api/models"
	"bfp/avi/api/config"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/contrib/jwt"
)

func main() {
	db := *InitDB()
	router := gin.Default()
	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	private := r.Group("/api/v1")
	public := r.Group("/api/v1")
	private.Use(jwt.Auth(config.GetString("TOKEN_KEY")))

	//manage users
	userHandler := h.NewUserHandler(db)
	private.GET("/users", userHandler.Index)
	public.POST("/user", userHandler.Create)
	public.POST("/login", userHandler.Auth)
	private.POST("/change_password", userHandler.ChangePassword)

	//manage incidents
	incidentHandler := h.NewIncidentHandler(db)
	private.GET("/incidents", incidentHandler.Index)
	private.GET("/incidents/:incident_id", incidentHandler.Show)
	private.POST("/incident", incidentHandler.Create)

	//manage fire status
	fireStatusHandler := h.NewFireStatusHandler(db)
	private.POST("/fire_status", fireStatusHandler.Create)
	private.PUT("/fire_status/update/:id", fireStatusHandler.Update)

	//manage fire stations
	fireStationHandler := h.NewFireStationHandler(db)
	private.GET("/fire_stations", fireStationHandler.Index)
	private.POST("/fire_station", fireStationHandler.Create)
	private.PUT("/fire_stations/:station_id", fireStationHandler.Update)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func InitDB() *gorm.DB {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetString("DB_USER"), config.GetString("DB_PASS"),
		config.GetString("DB_HOST"), config.GetString("DB_PORT"),
		config.GetString("DB_NAME"))
	log.Printf("\nDatabase URL: %s\n", dbURL)

	_db, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Error connecting to the database:  %s", err))
	}
	_db.DB()
	_db.LogMode(true)
	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.User{},&m.Incident{},&m.FireStatus{},&m.FireStation{})
	_db.Set("gorm:table_options", "ENGINE=InnoDB")
	return _db
}

func GetPort() string {
    var port = os.Getenv("PORT")
    // Set a default port if there is nothing in the environment
    if port == "" {
        port = "8000"
        fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
    }
    fmt.Println("port -----> ", port)
    return ":" + port
}