package handlers

import (
	"net/http"
	"fmt"
	"time"
	"io"
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	jwt_lib "github.com/dgrijalva/jwt-go"
	m "bfp/avi/api/models"
	"bfp/avi/api/config"
	"gopkg.in/redis.v3"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

//get all users
func (handler UserHandler) Index(c *gin.Context) {
	
}

//create new user
func (handler UserHandler) Create(c *gin.Context) {
	firstname := c.PostForm("firstname")
	lastname := c.PostForm("lastname")
	username := c.PostForm("username")
	password := c.PostForm("password")
	userlevel := c.PostForm("user_level")
	userrole := c.PostForm("user_role")

	if (strings.TrimSpace(firstname) == "") {
		respond(http.StatusBadRequest,"Please supply the user's first name",c,true)
	} else if (strings.TrimSpace(lastname) == "") {
		respond(http.StatusBadRequest,"Please supply the user's last name",c,true)
	} else if (strings.TrimSpace(username) == "") {
		respond(http.StatusBadRequest,"Please supply the account's username",c,true)
	} else if (strings.TrimSpace(password) == "") {
		respond(http.StatusBadRequest,"Please supply the account's password",c,true)
	} else if (strings.TrimSpace(userlevel) == "") {
		respond(http.StatusBadRequest,"Please supply the account's user level ",c,true)
	} else if (strings.TrimSpace(userrole) == "") {
		respond(http.StatusBadRequest,"Please supply the account's user role ",c,true)
	} else {
		//check if username already existing
		user := m.User{}	
		handler.db.Table("users").Where("username = ?",username).Find(&user)

		if (user.Username != "") {
			respond(http.StatusBadRequest,"Username already taken",c,true)
		} else {
			now := time.Now().UTC()
			encryptedPassword := encrypt([]byte(config.GetString("CRYPT_KEY")), password)
			result := handler.db.Exec("INSERT INTO users VALUES(null,?,?,?,?,?,?,?,?,?)",firstname,lastname,"active",userrole,userlevel,username,encryptedPassword,now,now)

			if (result.RowsAffected == 1) {
				user := m.User{}	
				handler.db.Table("users").Last(&user)
				authenticatedUser := m.AuthenticatedUser{}
				authenticatedUser.Id = user.Id
				authenticatedUser.FirstName = user.FirstName
				authenticatedUser.LastName = user.LastName
				authenticatedUser.Status = user.Status
				authenticatedUser.Userrole = user.Userrole
				authenticatedUser.Userlevel = user.Userlevel
				authenticatedUser.Username = user.Username
				authenticatedUser.DateCreated = user.DateCreated
				authenticatedUser.DateUpdated = user.DateUpdated
				authenticatedUser.Token = generateJWT(username)
				c.JSON(http.StatusCreated, authenticatedUser)
			} else {
				respond(http.StatusBadRequest,"Unable to create new user, Please try again",c,true)
			}
		}
	}
}

//generate JWT
func generateJWT(username string) string {
	// Create the token
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	// Set some claims
	token.Claims["ID"] = username
	token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, _ := token.SignedString([]byte(config.GetString("TOKEN_KEY")))
    return tokenString
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	// key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

func AddTokenToRedis(c *gin.Context) {
    client := redis.NewClient(&redis.Options{
        Addr:     ":6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    token := c.Request.Header.Get("Authorization")
    err := client.Set(token, token, time.Duration(86400)*time.Second).Err()
    if err != nil {
        panic(err)
    } else {
    	fmt.Println("Successfully written in redis")
    	result, err := client.Get(token).Result()
    	if (err == nil) {
    		fmt.Println("RESULT ---> " + result)
    	}
    }
    defer client.Close()
}

func IsTokenValid(c *gin.Context) bool {
    client := redis.NewClient(&redis.Options{
        Addr:     ":6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    defer client.Close()
    token := c.Request.Header.Get("Authorization")
    result, _ := client.Get(token).Result()
	if (result != "") {
		return false
	}
	return true
}