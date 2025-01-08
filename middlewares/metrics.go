package middlewares

import (
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

// MetricsMiddleware her isteği kayıt altına alan middleware
func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := c.Request.URL.Path
		userId := c.GetInt64("userId")

		go func() {
			if userId != 0 {
				trackRequest(endpoint, userId)
			} else {
				trackRequest(endpoint, nil)
			}
		}()

		c.Next()
	}
}

// trackRequest metrik kaydını gerçekleştiren fonksiyon
func trackRequest(endpoint string, userID interface{}) {
	var query string
	var args []interface{}

	if userID != nil {
		query = `INSERT INTO metrics (endpoint, user_id) VALUES (?, ?)`
		args = []interface{}{endpoint, userID}
	} else {
		query = `INSERT INTO metrics (endpoint) VALUES (?)`
		args = []interface{}{endpoint}
	}

	_, err := db.DB.Exec(query, args...)
	if err != nil {
		println("Metric kayıt hatası:", err.Error())
	}
}
