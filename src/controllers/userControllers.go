package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/janakhpon/gopherscom/src/models"
	"github.com/janakhpon/gopherscom/src/utils"
	"golang.org/x/crypto/bcrypt"
)

func GetUserList(c *gin.Context) {
	var userList []models.User
	err := dbConnect.Model(&userList).Select()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userList,
	})
	return
}

func GetUser(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	user := &models.User{ID: id}
	err := dbConnect.Select(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "failed to fetch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "succeed",
		"data": user,
	})
	return
}

func UserSignup(c *gin.Context) {
	var userBody models.User
	c.BindJSON(&userBody)

	user := models.User{
		ID:        uuid.New().String(),
		NAME:      userBody.NAME,
		EMAIL:     userBody.EMAIL,
		PASSWORD:  userBody.PASSWORD,
		CREATEDAT: time.Now(),
		UPDATEDAT: time.Now(),
	}

	if user.NAME == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"err": "Required name!",
		})
		return
	}
	if user.EMAIL == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"err": "Required email!",
		})
		return
	}
	if user.PASSWORD == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"err": "Required password!",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PASSWORD), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	user.PASSWORD = string(hash)

	resuser := &models.User{}
	err = dbConnect.Model(resuser).Where("email = ?", user.EMAIL).Select()

	// func (u *DB_User) AnotherGetItemByName(db *pg.DB) error {
	// 	err := db.Model(u).Relation("PersonalInfo").Where("PersonalInfo.name = ?", u.PersonalInfo.Name).Select()
	// 		return err
	// 	}
	// 	return nil
	// }

	if resuser.EMAIL != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "email alreday exists",
		})
		return
	}

	insertError := dbConnect.Insert(&user)
	if insertError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": insertError,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    &user,
	})

	return
}

func UserSignin(c *gin.Context) {
	var userBody models.User
	c.BindJSON(&userBody)
	user := models.User{}
	// user.ID = primitive.NewObjectID()
	user.NAME = userBody.NAME
	user.EMAIL = userBody.EMAIL
	user.PASSWORD = userBody.PASSWORD
	// user.CREATEDAT = time.Now()
	// user.UPDATEDAT = time.Now()
	if user.NAME == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"err": "Required name!",
		})
		return
	}
	if user.EMAIL == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"err": "Required email!",
		})
		return
	}
	if user.PASSWORD == "" {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"err": "Required password!",
		})
		return
	}
	var jwt models.JWT

	resuser := &models.User{}
	err := dbConnect.Model(resuser).Where("email = ?", user.EMAIL).Select()

	// func (u *DB_User) AnotherGetItemByName(db *pg.DB) error {
	// 	err := db.Model(u).Relation("PersonalInfo").Where("PersonalInfo.name = ?", u.PersonalInfo.Name).Select()
	// 		return err
	// 	}
	// 	return nil
	// }

	fmt.Println(resuser)

	if resuser.EMAIL == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "User not found",
		})
	}

	// err := collectionUsers.FindOne(context.TODO(), bson.M{"email": user.EMAIL}).Decode(&resuser)
	// if resuser == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"err": err,
	// 		"msg": "user not found!",
	// 	})
	// 	return
	// }

	hashedPassword := resuser.PASSWORD
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.PASSWORD))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
			"msg": "Passwords do not match!",
		})
		return
	}
	user.ID = resuser.ID
	user.CREATEDAT = resuser.CREATEDAT
	user.UPDATEDAT = resuser.UPDATEDAT
	token, refreshtoken, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err,
		})
		return
	}
	jwt.Token = token
	c.JSON(http.StatusAccepted, gin.H{
		"token":        token,
		"refreshtoken": refreshtoken,
	})

}

func TokenVerifyMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("SECRET")), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ID := claims["_id"]
				name := claims["name"]
				email := claims["email"]
				created_at := claims["created_at"]
				updated_at := claims["updated_at"]
				fmt.Println(claims)
				c.JSON(http.StatusAccepted, gin.H{
					"msg":       "Authorized",
					"ID":        ID,
					"NAME":      name,
					"EMAIL":     email,
					"UPDATEDAT": updated_at,
					"CREATEDAT": created_at,
				})
				c.Next()
			} else {
				fmt.Println("Unauthorized")
				fmt.Println(error)
				c.JSON(http.StatusNotImplemented, gin.H{
					"err": "invalid token",
				})
				c.Abort()
			}
		} else {
			c.JSON(http.StatusNotImplemented, gin.H{
				"err": "invalid token",
			})
			c.Abort()
			return
		}
	}
}

func RefreshToken(c *gin.Context) {
	authHeader := c.GetHeader("refreshToken")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		authToken := bearerToken[1]
		token, error := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			email := claims["email"]
			resuser := &models.User{}
			err := dbConnect.Model(resuser).Where("email = ?", email).Select()

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"err": err,
				})
				return
			}
			user := models.User{
				ID:        resuser.ID,
				NAME:      resuser.NAME,
				EMAIL:     resuser.EMAIL,
				PASSWORD:  resuser.PASSWORD,
				CREATEDAT: resuser.CREATEDAT,
				UPDATEDAT: resuser.UPDATEDAT,
			}

			token, refreshtoken, err := utils.GenerateToken(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": err,
				})
				return
			}

			c.JSON(http.StatusAccepted, gin.H{
				"token":        token,
				"refreshToken": refreshtoken,
			})

		} else {
			fmt.Println("Unauthorized")
			fmt.Println(error)
			c.JSON(http.StatusNotImplemented, gin.H{
				"err": "invalid refreshToken",
			})
			return
		}
	} else {
		c.JSON(http.StatusNotImplemented, gin.H{
			"err": "invalid refreshToken",
		})
		return
	}
}