package handlers

import (
	"errors"
	"log"
	jwtGen "skrik/JWT"
	"skrik/config"
	"skrik/database"
	"skrik/models"
	smtpModule "skrik/smtp"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// GetUser gets a user by ID
// @Summary Get user
// @Description Get a user by ID
// @Tags users
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param userid query string true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account [get]
func GetUser(c *fiber.Ctx) error {
	var request models.User
	user := c.Query("userid")
	if user == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"}) //nolint:errcheck
	}

	if err := database.DataBase.First(&request, "ID= ?", user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}

	return c.Status(fiber.StatusOK).JSON(request) //nolint:errcheck
}

// @Summary Create a new user
// @Description Create a new user with the input data
// @Tags users
// @Accept  json
// @Produce  plain
// @Param user body models.User true "User data"
// @Success 200 {string} string "JWT Token"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 409 {object} models.ErrorResponse "There's already user with that email"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account/register [post]
func CreateUser(c *fiber.Ctx) error {
	var request models.User
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}
	request.Password = string(hashedPassword)
	newUser := models.User{Email: request.Email, Username: request.Username, Password: request.Password}
	if err := database.DataBase.Create(&newUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{Message: "There's already user with that email"}) //nolint:errcheck
		}
		log.Println("couldn't create database record", err)
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"}) //nolint:errcheck
	}

	token, err := jwtGen.GenerateJWT(newUser.ID)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "couldn't create JWT token"}) //nolint:errcheck
	}
	VerifyJWT, err := jwtGen.GenerateVerificationJWT(newUser.ID)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "couldn't create email verification token"}) //nolint:errcheck
	}

	err = smtpModule.SendEmail(newUser.Email, "email verification", "localhost:8080/api/account/verify?token=", VerifyJWT)
	if err != nil {
		log.Println("couldn't send an email", err) //nolint:errcheck
	}

	return c.Status(fiber.StatusOK).SendString(token) //nolint:errcheck
}

// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param user body models.User true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account [put]
func UpdateUser(c *fiber.Ctx) error {
	var request map[string]interface{}
	var user models.User
	//userid := c.Query("userid")
	tokendata := c.Locals("user").(*jwt.Token)
	claims := tokendata.Claims.(jwt.MapClaims)
	userid := claims["userid"]
	if userid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"}) //nolint:errcheck
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}
	if err := database.DataBase.First(&user, "ID=?", userid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}
	for key, value := range request {
		if value == "" || value == nil {
			delete(request, key)
		}
	}

	if err := database.DataBase.Model(&user).Updates(request).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Error updating user"}) //nolint:errcheck
	}

	return c.Status(fiber.StatusOK).JSON(request) //nolint:errcheck
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Security BearerAuth
// @Param userid query string true "User ID"
// @Success 200 {string} string "User deleted"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account/ [delete]
func DeleteUser(c *fiber.Ctx) error {
	var user models.User
	userid := c.Query("userid")
	if userid == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad Request"}) //nolint:errcheck
	}

	if err := database.DataBase.Delete(&user, userid).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}

	return c.Status(fiber.StatusOK).SendString("Delete " + userid) //nolint:errcheck
}

// @Summary Login
// @Description it logs in..
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 200 {string} string "Successful"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account/login [post]
func Login(c *fiber.Ctx) error {
	var request models.User
	var user models.User

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}

	if err := database.DataBase.First(&user, "Email = ?", request.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "User not found"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "Internal Server Error"}) //nolint:errcheck
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{Message: "password deosn't match"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{Message: "panic"})
	}

	token, err := jwtGen.GenerateJWT(user.ID)
	if err != nil {
		log.Println("couldn't create JWT token", err)
	}
	//nolint:errcheck
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
	//
}

func VerifyEmail(c *fiber.Ctx) error {
	var user models.User
	EnvConfig := config.GetConfig()
	tokenQuery := c.Query("token")
	token, err := jwt.Parse(tokenQuery, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvConfig.Jwt_keyword), nil
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "invalid token"}) //nolint:errcheck
	}
	if !token.Valid {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "panic"}) //nolint:errcheck
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["userid"] == nil || claims["userid"] == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "empty id"}) //nolint:errcheck
	}
	userID := claims["userid"]
	if err := database.DataBase.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "panic"}) //nolint:errcheck
	}
	user.IsVerified = true
	if err := database.DataBase.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "couldn't verify"}) //nolint:errcheck
	}
	return c.Status(fiber.StatusOK).SendString("verified")
}

// @Summary Reset
// @Description it resets password....
// @Tags users
// @Accept  json
// @Produce  plain
// @Param user body models.User true "User data"
// @Success 200 {string} string "email sent"
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account/reset [post]
func ResetPassword(c *fiber.Ctx) error {
	var request models.User
	var user models.User
	if err := c.BodyParser(&request); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad request"}) //nolint:errcheck
	}

	if err := database.DataBase.First(&user, "email = ?", request.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "no such email"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad request111"}) //nolint:errcheck
	}

	token, err := jwtGen.GenerateVerificationJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "Bad request"}) //nolint:errcheck
	}
	smtpModule.SendEmail(user.Email, "Password recovery link", "localhost:8080/api/account/reset?token=", token) //nolint:errcheck

	return c.Status(fiber.StatusOK).SendString("email sent")
}

// @Summary reset password
// @Description it resets password..
// @Tags users
// @Accept  json
// @Produce  json
// @Param token query string true "token"
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse "Bad Request"
// @Failure 404 {object} models.ErrorResponse "User not found"
// @Failure 500 {object} models.ErrorResponse "Internal Server Error"
// @Router /api/account/resetendpoint [post]
func ResetEndpoint(c *fiber.Ctx) error {
	EnvConfig := config.GetConfig()
	var user models.User
	var request models.User
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "couldn't parse body"}) //nolint:errcheck
	}
	tokenQuery := c.Query("token")
	token, err := jwt.Parse(tokenQuery, func(token *jwt.Token) (interface{}, error) {
		return []byte(EnvConfig.Jwt_keyword), nil
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "invalid token"}) //nolint:errcheck
	}
	if !token.Valid {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "panic"}) //nolint:errcheck
	}
	claims := token.Claims.(jwt.MapClaims)
	if claims["userid"] == nil || claims["userid"] == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "empty id"}) //nolint:errcheck
	}
	userID := claims["userid"]

	if err := database.DataBase.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Message: "couldn't find user"}) //nolint:errcheck
		}
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Message: "panic"}) //nolint:errcheck
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "couldn't create password hash"}) //nolint:errcheck
	}
	user.Password = string(hashedPassword)

	if err := database.DataBase.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Message: "couldn't update password"}) //nolint:errcheck
	}
	return c.Status(fiber.StatusOK).JSON(&user) //nolint:errcheck it sends user back cuz i wanted (and might want in future) to see if it changes password and nothing else. zxcursed
}

type connections struct {
	userid map[int]*websocket.Conn
	mu     sync.Mutex
}

var connManager = connections{
	userid: make(map[int]*websocket.Conn),
}

func Test(app *fiber.App) {
	app.Get("api/ws", jwtGen.JwtProtected(), websocket.New(wsHandler))
}

func wsHandler(c *websocket.Conn) {
	connManager.mu.Lock()
	localsToken := c.Locals("user").(*jwt.Token)
	claims := localsToken.Claims.(jwt.MapClaims)
	useridFloat, ok := claims["userid"].(float64)
	if !ok {
		c.Close()
		connManager.mu.Unlock()
		return
	}
	userid := int(useridFloat)
	connManager.userid[userid] = c
	connManager.mu.Unlock()

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			connManager.mu.Lock()
			delete(connManager.userid, userid)
			connManager.mu.Unlock()
			c.Close()
			break
		}
		go msgHandler(msg, userid)
	}
}

func msgHandler(msg []byte, userid int) {
	temp := strings.SplitN(string(msg), ":", 2)
	if len(temp) < 2 {
		return
	}
	recieverids := strings.Split(temp[0], ",")
	for _, ids := range recieverids {
		recieverid, err := strconv.Atoi(ids)
		if err != nil {
			continue
		}

		connManager.mu.Lock()
		recConnection, exists := connManager.userid[recieverid]
		connManager.mu.Unlock()
		if !exists {
			continue
		}

		if err := recConnection.WriteMessage(websocket.TextMessage, []byte(temp[1])); err != nil {
			connManager.mu.Lock()
			delete(connManager.userid, recieverid)
			recConnection.Close()
			connManager.mu.Unlock()
		}
	}
	return
}
