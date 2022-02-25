package middleware

import (
	tokenLibrary "address-book-go/internal/token/library"
	"address-book-go/pkg/er"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Bearer")
		if token == "" {
			authErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "token is required.", nil)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		claims, err := tokenLibrary.ParseToken(token)
		if err != nil {
			authErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "token is not valid.", err)
			c.AbortWithStatusJSON(authErr.GetStatus(), authErr.GetMsg())
			return
		}

		var jwtAccountId, parseAccountIdOk = claims["account_id"].(string)
		//var jwtIat, jwtIatOk = claims["iat"].(float64)

		if !parseAccountIdOk {
			parseJwtInfoErr := er.NewAppErr(401, er.UnauthorizedError, "token is not valid.", nil)
			c.AbortWithStatusJSON(parseJwtInfoErr.GetStatus(), parseJwtInfoErr.GetMsg())
			return
		}

		accountId, _ := strconv.Atoi(jwtAccountId)

		// Jwt token state management
		//env := api.GetEnv()

		// TODO 後踢前
		/*
			tokenCache := tokenRepo.NewRedisTokenRepo(env.Redis)
			serverIat, err := tokenCache.GetCodeCertTokenIat(accountId)

			if jwtIat < serverIat {
				iatErr := er.NewAppErr(401, er.UnauthorizedError, "Token is expired", nil)
				c.AbortWithStatusJSON(iatErr.GetStatus(), iatErr.GetMsg())
				return
			}
		*/

		c.Set("claims", claims)
		c.Set("account_id", accountId)

		c.Next()
	}
}
