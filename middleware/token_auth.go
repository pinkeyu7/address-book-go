package middleware

import (
	"address-book-go/api"
	tokenLibrary "address-book-go/internal/token/library"
	tokenRepo "address-book-go/internal/token/repository"
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
		var jwtIat, jwtIatOk = claims["iat"].(float64)

		if !parseAccountIdOk || !jwtIatOk {
			parseJwtInfoErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "token is not valid.", nil)
			c.AbortWithStatusJSON(parseJwtInfoErr.GetStatus(), parseJwtInfoErr.GetMsg())
			return
		}

		accId, _ := strconv.Atoi(jwtAccountId)

		// Jwt token state management
		env := api.GetEnv()
		tc := tokenRepo.NewRedis(env.RedisCluster)
		serverIat, err := tc.GetTokenIat(accId)
		if err != nil {
			findErr := er.NewAppErr(http.StatusInternalServerError, er.UnknownError, "find token error.", err)
			c.AbortWithStatusJSON(findErr.GetStatus(), findErr.GetMsg())
			return
		}
		if jwtIat < serverIat {
			iatErr := er.NewAppErr(http.StatusUnauthorized, er.UnauthorizedError, "token is expired.", nil)
			c.AbortWithStatusJSON(iatErr.GetStatus(), iatErr.GetMsg())
			return
		}

		// Set claims
		c.Set("claims", claims)
		c.Set("account_id", accId)

		c.Next()
	}
}
