package middleware

import (
	"github.com/jasongauvin/zebi-scraper/api/helpers"
	"net/http"
	"strings"

	"github.com/jasongauvin/zebi-scraper/api/models"
	"github.com/jasongauvin/zebi-scraper/api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/zebi-scraper/api/services"
)

// CheckAuthorization used the find token and check request authorization in header
func CheckAuthorization(c *gin.Context) {
	var bearer = c.GetHeader("Authorization")

	if !strings.HasPrefix(bearer, "Bearer") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	var token = helpers.ExtractToken(bearer)

	var _, claims, err = services.DecodeToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	if token == "" {
		c.JSON(http.StatusBadRequest, "No token")
		return
	}

	customer := models.Customer{
		Email: claims.Email,
	}

	if err := repositories.FindCustomerByEmail(&customer); err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	c.Next()
}
