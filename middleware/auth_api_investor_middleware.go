package middleware

import (
	"net/http"
	api_investor "service-transaction/api/user_investor"
	"service-transaction/auth"
	"service-transaction/core"
	"service-transaction/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthApiInvestorMiddleware(authService auth.Service, userService core.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized API Investor", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// token, err := authService.ValidateToken(tokenString)
		// if err != nil {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// claim, ok := token.Claims.(jwt.MapClaims)

		// if !ok || !token.Valid {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// adminID := claim["admin_id"].(string)

		if tokenString == "" {
			response := helper.APIResponse("Unauthorized API Investor", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		investorID, err := api_investor.VerifyTokenInvestor(tokenString)

		if err != nil { //wrong token
			response := helper.APIResponse("Unauthorized API Investor", http.StatusUnauthorized, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUserInvestor", api_investor.InvestorId{UnixInvestor: investorID})
		c.Next()
	}
}
