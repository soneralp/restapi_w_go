package metrics

import (
	"example.com/rest-api/db"
)

var metricsChan = make(chan MetricData, 100)

type MetricData struct {
	Endpoint string
	UserID   interface{}
}

func init() {
	go processMetrics()
}

func trackRequest(endpoint string, userID interface{}) {
	metricsChan <- MetricData{Endpoint: endpoint, UserID: userID}
}

func processMetrics() {
	for metric := range metricsChan {
		saveMetric(metric)
	}
}

func saveMetric(metric MetricData) {
	var query string
	var args []interface{}

	if metric.UserID != nil {
		query = `INSERT INTO metrics (endpoint, user_id) VALUES (?, ?)`
		args = []interface{}{metric.Endpoint, int(metric.UserID.(float64))}
	} else {
		query = `INSERT INTO metrics (endpoint) VALUES (?)`
		args = []interface{}{metric.Endpoint}
	}

	_, err := db.DB.Exec(query, args...)
	if err != nil {
		println("Metric kayıt hatası:", err.Error())
	}
}
