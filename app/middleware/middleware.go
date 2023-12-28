package middleware

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/znxxx/Pitchaya_agnos_backend/models"
)

type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *customResponseWriter) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func LoggingMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := bytes.NewBuffer(nil)
		writer := &customResponseWriter{c.Writer, body}
		c.Writer = writer

		requestBody, err := c.GetRawData()
		if err != nil {
			fmt.Println("Error reading request body:", err)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))

		c.Next()

		requestLog := models.RequestLog{
			IP:            c.ClientIP(),
			RequestMethod: c.Request.Method,
			RequestPath:   c.Request.URL.Path,
			StatusCode:    c.Writer.Status(),
			RequestBody:   string(requestBody),
			Timestamp:     time.Now(),
		}
		logRequest(db, requestLog)

		responseLog := models.ResponseLog{
			StatusCode:   c.Writer.Status(),
			ResponseBody: body.String(),
			Timestamp:    time.Now(),
		}
		logResponse(db, responseLog)
	}
}

func logRequest(db *sql.DB, log models.RequestLog) {
	_, err := db.Exec(`
        INSERT INTO request_logs (ip, request_method, request_path, status_code, request_body, timestamp)
        VALUES ($1, $2, $3, $4, $5, $6)`,
		log.IP, log.RequestMethod, log.RequestPath, log.StatusCode, log.RequestBody, log.Timestamp)
	if err != nil {
		fmt.Println("Error logging request:", err)
	}
}

func logResponse(db *sql.DB, log models.ResponseLog) {
	_, err := db.Exec(`
        INSERT INTO response_logs (status_code, response_body, timestamp)
        VALUES ($1, $2, $3)`,
		log.StatusCode, log.ResponseBody, log.Timestamp)
	if err != nil {
		fmt.Println("Error logging response:", err)
	}
}
