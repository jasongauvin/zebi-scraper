package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/zebi-scraper/api/models"
	"github.com/jasongauvin/zebi-scraper/api/repositories"
	"github.com/jasongauvin/zebi-scraper/api/services"
)

// GetCustomer used the token and check request authorization in header for return customer
func GetCustomer(c *gin.Context) models.Customer {
	var bearer = c.GetHeader("Authorization")
	var token = strings.TrimPrefix(bearer, "Bearer ")

	_, claims, err := services.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return models.Customer{}
	}

	customer := models.Customer{
		Email: claims.Email,
	}

	repositories.FindCustomerByEmail(&customer)

	return customer
}
