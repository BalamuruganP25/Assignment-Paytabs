package messages

import (
	"log"
)

type UserApicredentials struct {
	ApiIp   string
	ApiPort string
}

var Apidetails UserApicredentials

type logservicedetails struct {
	Logfoldername string
	Logfilepath   string
}

var Logserviceconfig logservicedetails

var Logger *log.Logger

type ErrorResponse struct {
	ErrorCode    int    `json:"errorcode"`
	ErrorMessage string `json:"errormessage"`
}

var ErrorDetails ErrorResponse

