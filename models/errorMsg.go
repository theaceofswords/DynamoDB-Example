package models

type ErrorMsg struct {
	ErrError   string `json:"error"`
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
}
