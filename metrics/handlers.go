package metrics

import (
	"example.com/rest-api/db"
	"github.com/gin-gonic/gin"
)

// GetMetrics tüm metrikleri döndüren handler
func GetMetrics(c *gin.Context) {
	rows, err := db.DB.Query(`
		SELECT 
			m.endpoint, 
			COUNT(*) as count,
			COUNT(DISTINCT m.user_id) as unique_users
		FROM metrics m 
		GROUP BY m.endpoint
	`)
	if err != nil {
		c.JSON(500, gin.H{"error": "Database error"})
		return
	}
	defer rows.Close()

	metrics := make([]gin.H, 0)
	for rows.Next() {
		var endpoint string
		var count int
		var uniqueUsers int
		if err := rows.Scan(&endpoint, &count, &uniqueUsers); err != nil {
			c.JSON(500, gin.H{"error": "Scan error"})
			return
		}
		metrics = append(metrics, gin.H{
			"endpoint":       endpoint,
			"total_requests": count,
			"unique_users":   uniqueUsers,
		})
	}

	c.JSON(200, metrics)
}
