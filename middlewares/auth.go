package middlewares

import (
	"emvn/consts"
	auth_usecase "emvn/internal/usecase/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func getTokenFromBearer(reqToken string) string {
	if reqToken == "" {
		return ""
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	return splitToken[1]
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqTokenHeader := c.GetHeader("Authorization")
		if reqTokenHeader == "" {
			c.JSON(consts.CodeTokenRequired.HttpStatus, gin.H{
				"code":    consts.CodeTokenRequired.Code,
				"message": consts.CodeTokenRequired.Message,
			})
			c.Abort()
			return
		}
		reqToken := getTokenFromBearer(reqTokenHeader)
		if reqToken == "" {
			c.JSON(consts.CodeTokenExpired.HttpStatus, gin.H{
				"code":    consts.CodeTokenExpired.Code,
				"message": consts.CodeTokenExpired.Message,
			})
			c.Abort()
			return
		}

		acToken := auth_usecase.AccessToken{}
		err := acToken.Verify(c.Request.Context(), reqToken)
		if err != nil {
			c.JSON(consts.CodeInvalidToken.HttpStatus, gin.H{
				"code":    consts.CodeInvalidToken.Code,
				"message": consts.CodeInvalidToken.Message,
			})
			c.Abort()
			return
		}

		// We can query user from db and set it to context
		// But for now, we just set the uid to context
		c.Set(consts.GinAuthUid, acToken.Sub)
		c.Next()
	}
}
