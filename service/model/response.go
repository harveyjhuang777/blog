package model

type ResponseStatus int

const (
	ResponseStatusFail ResponseStatus = 0
	ResponseStatusOK   ResponseStatus = 1
)

type Response struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
}
