package middlewares

import (
	"net/http"
	"strings"

	"example.com/utils"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(context *gin.Context) {

		client_secret_key := "sk_test_QhJyugNaceJASK94yE7mFv6GlZcZGgErlTVUMWjzbR"

		client, err := clerk.NewClient(client_secret_key)

		if err != nil {
			context.JSON(http.StatusGatewayTimeout, gin.H{
				"message": "Invalid or expired client session",
			})

			return
		}

		token := strings.TrimPrefix(context.GetHeader("Authorization"), "Bearer ")

		sessionId, err := utils.GetSessionIDFromToken(token)

		if err != nil {
			context.JSON(http.StatusGatewayTimeout, gin.H{
				"message": "Invalid or expired session",
			})

			return
		}

		session, err := client.Sessions().Read(sessionId)

		if err != nil {
			context.JSON(http.StatusGatewayTimeout, gin.H{
				"message": "Invalid or expired session",
			})

			return
		}

		userId := session.UserID

		requestedUserId := context.Param("clerkId")

		if userId != requestedUserId {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid user request",
			})

			return

		}

		context.Next()
	}
}
