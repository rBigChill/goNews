package newsapi

import "fmt"

type Error struct {
	Status  string `json: "status,omitempty"`
	Code    string `json: "code,omitempty"`
	Message string `json: "message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Status: %s: [%s] %s", e.Status, e.Code, e.Message)
}
