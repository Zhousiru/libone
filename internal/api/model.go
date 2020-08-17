package api

import "net/http"

const (
	respStatusOK int = iota
	respStatusErr
)

var statusMap map[int]status

type status struct {
	HTTPCode  int
	StatusMsg string
}

func init() {
	statusMap = make(map[int]status)

	statusMap[respStatusOK] = status{HTTPCode: http.StatusOK, StatusMsg: "ok"}
	statusMap[respStatusErr] = status{HTTPCode: http.StatusInternalServerError, StatusMsg: "err"}
}
