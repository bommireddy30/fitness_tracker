package handler

import (
	"Connection/database"
	"Connection/dto/requestdto"
	"Connection/dto/responsedto"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {

	logKey := "CreateUser: "
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// Construct the initial Response body
	respDto := responsedto.DefaultResp{}
	respDto.ErrorDescription = "Generic error"
	respDto.Status = "RS_ERROR"
	// parse the request body
	bodyStr := c.Body()
	reqDto := requestdto.CreateUserReqDto{}
	if err := c.BodyParser(&reqDto); err != nil {
		// body parsing failed
		log.Println(logKey, "BodyParser failed with an error:", err.Error(), "Body is", bodyStr)
		respDto.ErrorDescription = "Invalid request"
		respDto.Status = "INVALID_REQ"
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}

	// Creating userId and Password
	reqDto.UserId = reqDto.UserName + "-" + strings.Split(uuid.New().String(), "-")[0]
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqDto.Password), bcrypt.MinCost)
	if err != nil {
		log.Println("Failed to covert password to hash:", err.Error())
	}
	reqDto.Password = string(hashedPassword)

	// valiate the request body
	if len(reqDto.UserName) < 3 || len(reqDto.Password) < 8 {
		respDto.ErrorDescription = "UserName and Password validation Failed"
		return c.Status(fiber.StatusExpectationFailed).JSON(respDto)
	}
	// make Database call
	err = database.CreateUser(reqDto)
	if err != nil {
		log.Println(logKey, "database.CreateUser failed with  an error:", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}
	// construct response body
	respDto.ErrorDescription = "User Created Successfully"
	respDto.Status = "RS_OK"
	// return response
	return c.Status(fiber.StatusCreated).JSON(respDto)
}

func UserLogin(c *fiber.Ctx) error {
	logkey := "User login Handler"
	c.Accepts("json", "text")
	c.Accepts("application/json")

	// Create a default response body
	respDto := responsedto.LoginResp{}
	respDto.ErrorDescription = "Generic Error"
	respDto.Status = "ERROR_STATUS"

	// Parse request body
	reqDto := requestdto.UserLogin{}
	if err := c.BodyParser(&reqDto); err != nil {
		log.Println(logkey, "RequestBody pasing failed with error:", err.Error())
		respDto.ErrorDescription = "Invalid request"
		return c.Status(fiber.StatusBadRequest).JSON(respDto)
	}
	// Validate request Body
	if reqDto.UserName == "" || len(reqDto.Password) < 8 {
		respDto.ErrorDescription = "Invalid UserName or Password"
		return c.Status(fiber.StatusBadRequest).JSON(reqDto)
	}
	// Get user from database using userName
	userDto, err := database.UserLogin(reqDto.UserName)
	if err != nil {
		log.Println(logkey, "database.UserLogin failed with an error:", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(respDto)
	}
	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(userDto.Password), []byte(reqDto.Password))
	if err != nil {
		respDto.ErrorDescription = "Invalid Password"
		respDto.Status = "PASSWORD+_INVALID"
		return c.Status(fiber.StatusNotFound).JSON(respDto)
	}
	// Genarate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userid"] = userDto.UserId
	claims["userName"] = userDto.UserName
	claims["mail"] = userDto.Mail
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()

	tkn, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		log.Println(logkey, "Failed to Construct JWT with error:", err.Error())
		respDto.ErrorDescription = "Failed to Genarate JWT"
		return c.Status(fiber.StatusNotImplemented).JSON(respDto)
	}
	// Return JWT
	respDto.ErrorDescription = ""
	respDto.Status = "RS_OK"
	respDto.Token = tkn

	return c.Status(fiber.StatusOK).JSON(respDto)
}
