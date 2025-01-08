package metrics

import (
	"example.com/rest-api/db"
)

// trackRequest metrik kaydını gerçekleştiren fonksiyon
func trackRequest(endpoint string, userID interface{}) {
	var query string
	var args []interface{}

	if userID != nil {
		query = `INSERT INTO metrics (endpoint, user_id) VALUES (?, ?)`
		args = []interface{}{endpoint, int(userID.(float64))}
	} else {
		query = `INSERT INTO metrics (endpoint) VALUES (?)`
		args = []interface{}{endpoint}
	}

	_, err := db.DB.Exec(query, args...)
	if err != nil {
		println("Metric kayıt hatası:", err.Error())
	}
}
