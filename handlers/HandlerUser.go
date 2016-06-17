package handlers

import (
	"net/http"
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
	users := []m.User{}
	handler.db.Find(&users)
	c.JSON(http.StatusOK,users)
	return
}

//create new user
func (handler UserHandler) Create(c *gin.Context) {
	if IsTokenValid(c) {
		var user m.User
		err := c.Bind(&user)
		if err == nil {
			existingUser := m.User{}
			if handler.db.Where("email = ?",user.Email).First(&existingUser).RowsAffected < 1 {
				encryptedPassword := encrypt([]byte(config.GetString("CRYPT_KEY")), "123")
				user.Password = encryptedPassword
				result := handler.db.Create(&user)
				if result.RowsAffected > 0 {
					c.JSON(http.StatusCreated, user)
				} else {
					respond(http.StatusBadRequest,result.Error.Error(),c,true)
				}
			} else {
				respond(http.StatusForbidden,"Email already taken",c,true)	
			}
		} else {
			respond(http.StatusBadRequest,err.Error(),c,true)
		}
	} else {
		respond(http.StatusForbidden,"Sorry, but your session has expired!",c,true)	
	}
	return
}

//user authentication
func (handler UserHandler) Auth(c *gin.Context) {
	if IsTokenValid(c) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		if (strings.TrimSpace(email) == "") {
			respond(http.StatusBadRequest,"Email is required",c,true)
		} else if (strings.TrimSpace(password) == "") {
			respond(http.StatusBadRequest,"Password is required",c,true)
		} else {
			//check if email already existing
			user := m.User{}	
			query := handler.db.Where("email = ?",email).Find(&user)

			if query.RowsAffected < 1 {
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
					authenticatedUser.Email = user.Email
					authenticatedUser.IsPasswordDefault = user.IsPasswordDefault
					authenticatedUser.Userrole = user.Userrole
					authenticatedUser.Userlevel = user.Userlevel
					authenticatedUser.DateCreated = user.CreatedAt
					authenticatedUser.DateUpdated = user.UpdatedAt
					authenticatedUser.Token = generateJWT(email)
					c.JSON(http.StatusOK, authenticatedUser)
				}					
			}
		}
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)	
	}
}

func (handler UserHandler) ChangePassword (c *gin.Context) {
	if IsTokenValid(c) {
		email := c.PostForm("email")
		newPassword := c.PostForm("new_password")
		user := m.User{}
		query := handler.db.Where("email = ?",email).Find(&user)
		
		if query.RowsAffected > 0 {
			encryptedPassword := encrypt([]byte(config.GetString("CRYPT_KEY")), newPassword)
			user.Password = encryptedPassword
			result := handler.db.Save(&user)
			if result.RowsAffected > 0 {
				c.JSON(http.StatusOK,"Password successfully changed!")
			} else {
				respond(http.StatusBadRequest,"Unable to change password",c,true)
			}
		} else {
			respond(http.StatusBadRequest,"User not found!",c,true)
		}
	} else {
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)	
	}
	return
}
