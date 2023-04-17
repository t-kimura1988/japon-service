package types

import "net/http"

type GrpcCodeConvert struct {
	Code int32
}

func (err *GrpcCodeConvert) ToHttp() int {
	switch err.Code {
	case 5:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
