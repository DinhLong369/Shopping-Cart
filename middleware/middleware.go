package middleware

import (
	"Shopping-cart/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// JWTAuthMiddleware kiểm tra token trong Authorization header
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Tách giá trị Authorization thành 2 phần: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse và xác minh token bằng secret key
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Kiểm tra thuật toán ký
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// Trả về secret key để thư viện JWT xác minh chữ ký
			return config.SecretKey, nil
		})

		// Nếu lỗi hoặc token không hợp lệ thì trả lỗi 401
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Nếu token hợp lệ, lấy payload (claims) trong token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Lấy ID người dùng từ claims
			userIDFloat, ok := claims["id_user"].(float64) // Thử ép kiểu sang float64
			if !ok {
				logrus.Error("Invalid id_user type in token")
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token payload"})
				c.Abort()
				return
			}
			// Lưu userID vào context để controller có thể sử dụng
			c.Set("userID", uint(userIDFloat))
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Cho request đi tiếp
		c.Next()
	}
}

// 1️⃣ Request gửi đến server → middleware JWT được Gin tự động gọi.
// 2️⃣ Middleware lấy header Authorization.
// 3️⃣ Middleware parse + verify token:

// Nếu token hợp lệ → gán userID vào context → request được cho đi tiếp (gọi c.Next()).

// Nếu token không hợp lệ → trả lỗi 401 Unauthorized → không cho request đi tiếp (gọi c.Abort()).
// 4️⃣ Controller hoặc middleware tiếp theo sẽ lấy userID từ context nếu cần.

// Client gửi request --> Middleware kiểm tra token:
//     |
//     +-- Token hợp lệ --> cho đi tiếp --> Controller xử lý --> Response OK
//     |
//     +-- Token không hợp lệ --> Response 401 Unauthorized
