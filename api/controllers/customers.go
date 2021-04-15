package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/gogoleplate/api/config"
	"github.com/jasongauvin/gogoleplate/api/models"
	"github.com/jasongauvin/gogoleplate/api/repositories"
	"github.com/jasongauvin/gogoleplate/api/services"
)

// Register takes a json object with password as parameter, hash password and persists a user into the DB
func Register(c *gin.Context) {
	var err error
	var customerForm models.CustomerForm

	if err := c.ShouldBindJSON(&customerForm); err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect informations")
		return
	}

	// Check email and password
	err = models.ValidateCustomer(&customerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	customer := models.Customer{
		Name:           customerForm.Name,
		Email:          customerForm.Email,
		HashedPassword: services.HashPassword(customerForm.Password),
		Role:           customerForm.Role,
	}

	if err = repositories.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, "Error while saving user informations in database")
		return
	}

	c.JSON(http.StatusOK, "user successfully created!")
}

// Login takes customer email and password a json params and returns a token or an error depending on the credentials given.
func Login(c *gin.Context) {
	var customerForm models.CustomerForm
	if err := c.ShouldBindJSON(&customerForm); err != nil {
		c.JSON(http.StatusBadRequest, "Incorrect customer informations")
		return
	}

	// Check email and password
	err := models.ValidateCustomer(&customerForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	customer := models.Customer{
		Email: customerForm.Email,
	}

	err = repositories.FindCustomerByEmail(&customer)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Customer not find.")
		return
	}

	// Verify password
	hashedPwd := services.HashPassword(customerForm.Password)
	if hashedPwd != customer.HashedPassword {
		c.JSON(http.StatusUnauthorized, "Email or password incorrect.")
		return
	}

	// Generate connection token
	token, err := services.GenerateToken(customerForm.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Couldn't create your authorization")
		return
	}
	validTime, _ := strconv.ParseInt(config.GoDotEnvVariable("TOKEN_VALID_DURATION"), 10, 64)

	c.SetCookie("Bearer", token, 60*int(validTime), "/", config.GoDotEnvVariable("DOMAIN"), false, false)
	c.JSON(http.StatusOK, token)
}