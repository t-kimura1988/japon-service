package exceptions

import "fmt"

type ServiceError struct {
	Message    string
	LogMessage string
	Code       int64
}

func (err *ServiceError) Error() string {
	return fmt.Sprintf("Service Error: Code=%d, Message=%s", err.Code, err.LogMessage)
}

func (err *ServiceError) Wrap(traceMsg string) {
	err.LogMessage = fmt.Sprintf("%s %s", err.LogMessage, traceMsg)
}
