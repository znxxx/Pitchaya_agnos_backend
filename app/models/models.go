package models

import "time"

type RequestLog struct {
	ID            uint      `db:"id"`
	IP            string    `db:"ip"`
	RequestMethod string    `db:"request_method"`
	RequestPath   string    `db:"request_path"`
	StatusCode    int       `db:"status_code"`
	RequestBody   string    `db:"request_body"`
	Timestamp     time.Time `db:"timestamp"`
}

type ResponseLog struct {
	ID           uint      `db:"id"`
	StatusCode   int       `db:"status_code"`
	ResponseBody string    `db:"response_body"`
	Timestamp    time.Time `db:"timestamp"`
}
