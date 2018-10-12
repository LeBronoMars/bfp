package main

import (
	"os"
	"fmt"
	"log"
	"time"
	"net/http"
	"strings"

	"bfp/avi/api/config"
	_ "github.com/go-sql-driver/mysql"
	h "bfp/avi/api/handlers"
	m "bfp/avi/api/models"
	
	"github.com/jinzhu/gorm"
	"github.com/dgrijalva/jwt-go"

	"github.com/itsjamie/gin-cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := *InitDB()
	router := gin.Default()

	config := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Authorization",
		Methods:         "GET, POST, PUT, DELETE",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          24 * time.Hour,
	}
	router.Use(cors.Middleware(config))

	LoadAPIRoutes(router, &db)
}

func LoadAPIRoutes(r *gin.Engine, db *gorm.DB) {
	private := r.Group("/api/v1")
	public := r.Group("/api/v1")
	private.Use(Auth(config.GetString("TOKEN_KEY"), db))

	// manage users
	userHandler := h.NewUserHandler(db)
	private.GET("/users", userHandler.Index)
	public.POST("/user", userHandler.Create)
	public.POST("/login", userHandler.Auth)
	private.PUT("/users/:id", userHandler.Update)
	private.PUT("/change-password", userHandler.ChangePassword)

	// manage incidents
	incidentHandler := h.NewIncidentHandler(db)
	private.GET("/incidents", incidentHandler.Index)
	private.GET("/incidents/:incident_id", incidentHandler.Show)
	private.POST("/incident", incidentHandler.Create)

	// manage fire status
	fireStatusHandler := h.NewFireStatusHandler(db)
	private.POST("/fire_status", fireStatusHandler.Create)
	private.PUT("/fire_status/update/:id", fireStatusHandler.Update)

	// manage fire stations
	fireStationHandler := h.NewFireStationHandler(db)
	private.GET("/stations", fireStationHandler.Index)
	private.POST("/station", fireStationHandler.Create)
	private.PUT("/stations/:station_id", fireStationHandler.Update)

	// manage status
	statusHandler := h.NewStatusHandler(db)
	private.GET("/statuses", statusHandler.Index)
	private.POST("/status", statusHandler.Create)
	private.PUT("/statuses/:id", statusHandler.Update)

	// regional office
	regionalOfficeHandler := h.NewRegionalOfficeHandler(db)
	private.GET("/regional-offices", regionalOfficeHandler.Index)
	private.POST("/regional-office", regionalOfficeHandler.Create)
	private.PUT("/regional-offices/:id", regionalOfficeHandler.Update)

	var port = os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("PORT ---> ",port)
	r.Run(fmt.Sprintf(":%s", port))
}

func Auth(secret string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "" && strings.Contains(c.Request.Header.Get("Authorization"), "Bearer") {
		
			tokenString := c.Request.Header.Get("Authorization")
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			    return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				response := &Response{Message: err.Error()}
				c.JSON(http.StatusUnauthorized, response)
				c.Abort()
			} else {
				claims, _ := token.Claims.(jwt.MapClaims)
				email := fmt.Sprintf("%s", claims["iss"])
				
				user := m.User{}	
				res := db.Where("email = ?", email).First(&user)
				if res.RowsAffected > 0 {
					if user.Status == "blocked" {
						response := &Response{Message: "You're account is currently blocked. Please contact the administration."}
						c.JSON(http.StatusUnauthorized, response)
						c.Abort();
					}
				} else {
					c.Next();
				}
			}
		} else {
			response := &Response{Message: "Authorization is required"}
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
		}
	}
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
	_db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&m.User{},
																&m.Incident{},
																&m.FireStatus{},
																&m.FireStation{},
																&m.RegionalOffice{},
																&m.Status{})
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

type Response struct {
	Message string `json:"message"`
}