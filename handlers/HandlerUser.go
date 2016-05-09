package handlers

import (
	"net/http"
	"time"
    "strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
	"bfp/avi/api/config"
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

//user authentication
func (handler UserHandler) Auth(c *gin.Context) {
	if IsTokenValid(c) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if (strings.TrimSpace(username) == "") {
			respond(http.StatusBadRequest,"Please supply the user's username",c,true)
		} else if (strings.TrimSpace(password) == "") {
			respond(http.StatusBadRequest,"Please supply the user's password",c,true)
		} else {
			//check if username already existing
			user := m.User{}	
			handler.db.Table("users").Where("username = ?",username).Find(&user)

			if user.Username == "" {
				respond(http.StatusBadRequest,"Account not found!",c,true)
			} else {
				decryptedPassword := decrypt([]byte(config.GetString("CRYPT_KEY")), user.Password)
				//invalid password
				if decryptedPassword != password {
					respond(http.StatusBadRequest,"Account not found!",c,true)
				} else {
					//authentication successful
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
					c.JSON(http.StatusOK, authenticatedUser)
				}					
			}
		}
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)	
	}
}
