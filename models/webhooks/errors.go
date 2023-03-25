package models

type Error struct {
	Code      int        `json:"code"`
	Title     string     `json:"title"`
	Message   string     `json:"message,omitempty"`
	ErrorData *ErrorData `json:"error_data,omitempty"`
}

type ErrorData struct {
	Details string `json:"details"`
}
