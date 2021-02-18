package controllers

import (
	"placement/services"
)

var (
	companyService services.CompanyService
)

func init() {
	companyService = services.NewCompanyService()
}

// func RegisterCompany(c *fiber.Ctx) error {
// 	var body *models.Company
// 	err := json.Unmarshal(c.Body(), body)
// 	if err != nil {
// 		error := models.ErrorResponse{
// 			Message: "Something went wrong",
// 			Status:  http.StatusBadRequest,
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.Name == "" {
// 		error := models.ErrorResponse{
// 			Message: "name cant be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "name",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.Email == "" {
// 		error := models.ErrorResponse{
// 			Message: "email cant be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "email",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.RegistrationNumber == "" {
// 		error := models.ErrorResponse{
// 			Message: "registerion number cant be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "registerionNumber",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.GSTNumber == "" {
// 		error := models.ErrorResponse{
// 			Message: "gst number cant be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "gstNumber",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.WebSiteURL == "" {
// 		error := models.ErrorResponse{
// 			Message: "web siteurl cant be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "webSiteURL",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if len(body.PhoneNumber) != 10 {
// 		error := models.ErrorResponse{
// 			Message: "phone number should be 10 digits long",
// 			Status:  http.StatusBadRequest,
// 			Key:     "phoneNumber",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.Address == "" {
// 		error := models.ErrorResponse{
// 			Message: "addres can't be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "address",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.Password == "" {
// 		error := models.ErrorResponse{
// 			Message: "password can't be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "password",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.ConfirmPassword == "" {
// 		error := models.ErrorResponse{
// 			Message: "confirm password can't be empty",
// 			Status:  http.StatusBadRequest,
// 			Key:     "confirmPassword",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	if body.Password != body.ConfirmPassword {
// 		error := models.ErrorResponse{
// 			Message: "password and confirm password should be same",
// 			Status:  http.StatusBadRequest,
// 			Key:     "password,confirmPassword",
// 		}
// 		return c.Status(error.Status).JSON(error)
// 	}
// 	body, err = companyService.RegisterCompany(body)
// }
