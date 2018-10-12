package handlers

import (
	"net/http"
    "strings"
    "strconv"

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
	var query = handler.db

	userRoleParam,userRoleParamExist := c.GetQuery("user_role")
	userLevelParam,userLevelParamExist := c.GetQuery("user_level")
	startParam,startParamExist := c.GetQuery("start")
	limitParam,limitParamExist := c.GetQuery("limit")

	// user role param exist
	if userRoleParamExist {
		query = query.Where("user_role = ?", userRoleParam)
	} 

	// user level param exist
	if userLevelParamExist {
		query = query.Where("user_level = ?", userLevelParam)
	} 

	//start param exist
	if startParamExist {
		start,_ := strconv.Atoi(startParam)
		if start != 0 {
			query = query.Offset(start)				
		} else {
			query = query.Offset(0)
		}
	} 

	//limit param exist
	if limitParamExist {
		limit,_ := strconv.Atoi(limitParam)
		query = query.Limit(limit)
	} else {
		query = query.Limit(10)
	}

	query.Find(&users)
	c.JSON(http.StatusOK,users)
	return
}

//create new user
func (handler UserHandler) Create(c *gin.Context) {
	var user m.User
	err := c.Bind(&user)

	if err == nil {
		existingUser := m.User{}
		if handler.db.Where("email = ?",user.Email).First(&existingUser).RowsAffected < 1 {
			if (contains(USER_ROLES, user.UserRole)) {

				station_id := user.StationId
				fs := m.FireStation{}

				if handler.db.Where("id = ? ",station_id).First(&fs).RowsAffected > 0 {
					encryptedPassword := encrypt([]byte(config.GetString("CRYPT_KEY")), "123")
					user.Password = encryptedPassword
					result := handler.db.Create(&user)

					if result.RowsAffected > 0 {
						//authentication successful
						authenticatedUser := m.AuthenticatedUser{}
						authenticatedUser.Id = user.Id
						authenticatedUser.FirstName = user.FirstName
						authenticatedUser.LastName = user.LastName
						authenticatedUser.Status = user.Status
						authenticatedUser.Email = user.Email
						authenticatedUser.IsPasswordDefault = user.IsPasswordDefault
						authenticatedUser.UserRole = user.UserRole
						authenticatedUser.UserLevel = user.UserLevel
						authenticatedUser.DateCreated = user.CreatedAt
						authenticatedUser.DateUpdated = user.UpdatedAt
						authenticatedUser.Token = generateJWT(user.Email)
						c.JSON(http.StatusCreated, authenticatedUser)
					} else {
						respond(http.StatusBadRequest,result.Error.Error(),c,true)
					}
				} else {
					respond(http.StatusNotFound, "Fire station not found!",c,true)
				}
			} else {
				respond(http.StatusUnprocessableEntity, "Invalid user role. User role must be one of the following SUPER_ADMIN, ADMIN, or USER",c,true)
			}
		} else {
			respond(http.StatusUnprocessableEntity, "Email already taken",c,true)	
		}
	} else {
		respond(http.StatusBadRequest,err.Error(),c,true)
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
					authenticatedUser.UserRole = user.UserRole
					authenticatedUser.UserLevel = user.UserLevel
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
	oldPassword := c.PostForm("old_password")
	newPassword := c.PostForm("new_password")

	if (len(c.PostForm("old_password")) == 0) {
		respond(http.StatusUnprocessableEntity, "Old password parameter is missing.", c, true)
		return
	} else if (len(c.PostForm("new_password")) == 0) {
		respond(http.StatusUnprocessableEntity, "New password parameter is missing.", c, true)
		return
	} else {
		email := getEmailFromSession(c)

		user := m.User{}
		query := handler.db.Where("email = ?", email).Find(&user)
		
		if query.RowsAffected > 0 {

			decryptedPassword := decrypt([]byte(config.GetString("CRYPT_KEY")), user.Password)

			if decryptedPassword == oldPassword {
				if (decryptedPassword == newPassword) {
				respond(http.StatusUnprocessableEntity, "New password and old password cannot be the same.", c, true)
				} else {
					encryptedPassword := encrypt([]byte(config.GetString("CRYPT_KEY")), newPassword)
					user.Password = encryptedPassword
					user.IsPasswordDefault = false
					result := handler.db.Save(&user)

					if result.RowsAffected > 0 {
						respond(http.StatusBadRequest, "Password successfully changed!", c, false)
					} else {
						respond(http.StatusBadRequest, "Unable to change password", c, true)
					}
				}
			} else {
				respond(http.StatusUnprocessableEntity, "Invalid old password", c, true)
			}
		} else {
			respond(http.StatusNotFound, "User not found!", c, true)
		}
	}
}

func (handler UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	user := m.User{}
	qry := handler.db.Where("id = ? AND status = ?",id,"active").First(&user)

	if qry.RowsAffected > 0 {
		// check if email is available
		if (len(c.PostForm("email")) > 0) {
			userWithExistingEmail := m.User{}
			emailQuery := handler.db.Where("id != ? AND email = ?", id, c.PostForm("email")).First(&userWithExistingEmail)

			if (emailQuery.RowsAffected > 0) {
				respond(http.StatusUnprocessableEntity, "Email already taken!",c,true)	
				return
			} else {
				// update email
				user.LastName = c.DefaultPostForm("last_name", user.LastName)
			}
		}

		// update station id
		if (len(c.PostForm("station_id")) > 0) {
			stationId, _ := strconv.Atoi(c.PostForm("station_id"))

			// check if user station exist
			fs := m.FireStation{}

			if handler.db.Where("id = ? ", stationId).First(&fs).RowsAffected < 1 {
				respond(http.StatusNotFound, "Station not found!",c,true)	
				return
			} else {
				user.StationId = stationId
			}
		}
		
		// update first name
		user.FirstName = c.DefaultPostForm("first_name", user.FirstName)

		// update last name
		user.LastName = c.DefaultPostForm("last_name", user.LastName)

		// update contact no
		user.ContactNo = c.DefaultPostForm("contact_no", user.ContactNo)

		result := handler.db.Save(&user)

		if result.RowsAffected > 0 {
			c.JSON(http.StatusOK, user)
		} else {
			respond(http.StatusBadRequest,result.Error.Error(), c, true)	
		}
	}
}
