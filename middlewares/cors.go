package middlewares

import "github.com/gin-gonic/gin"

// CORS allows Cross-origin resource sharing.
// Allow all origins, all methods, all headers, and credentials.
// In production, you should restrict the origin to your domain though a environment variable.
// For example, you can use the following code:
// c.Header("Access-Control-Allow-Origin", config.GetConfig().Frontend.Domain)
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Range")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "600")
		c.Next()
	}
}
