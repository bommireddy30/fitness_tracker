package handler

import (
	"Connection/database"
	"Connection/dto/requestdto"
	"Connection/dto/responsedto"
	"log"

	"github.com/gofiber/fiber/v2"
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
	// valiate the request body
	if len(reqDto.UserName) < 3 || len(reqDto.Password) < 8 {
		respDto.ErrorDescription = "UserName and Password validation Failed"
		return c.Status(fiber.StatusExpectationFailed).JSON(respDto)
	}
	// make Database call
	err := database.CreateUser(reqDto)
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
