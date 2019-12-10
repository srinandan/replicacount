package types

import (
	"log"
)

//ErrorMessage hold the return value when there is an error
type ErrorMessage struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message,omitempty"`
}

//CountResponse returns the successful response
type CountResponse struct {
	Replicas          int32 `json:"replicas,omitempty"`
	ReadyReplicas     int32 `json:"ready_replicas,omitempty"`
	AvailableReplicas int32 `json:"available_replicas,omitempty"`
}

//log levels, default is error
var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)
