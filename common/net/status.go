package net

type Status struct {
	Code string
	Msg  string
}

var (
	OK                  = Status{Code: "200", Msg: "OK"}
	StatusUnauthorized  = Status{Code: "401", Msg: "StatusUnauthorized"}
	Forbidden           = Status{Code: "403", Msg: "Forbidden"}
	NotFound            = Status{Code: "404", Msg: "NotFound"}
	InternalServerError = Status{Code: "500", Msg: "InternalServerError"}
)
