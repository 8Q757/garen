package response

import (
	"github.com/gin-gonic/gin"
	"github.com/viticis/garen/common/net"
	"net/http"
)

type response struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Result(status net.Status, data interface{}, c *gin.Context) {

	c.JSON(http.StatusOK, response{
		Code: status.Code,
		Msg:  status.Msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	SuccessWithData(nil, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	Result(net.OK, data, c)
}

func Fail(status net.Status, c *gin.Context) {
	FailWithData(status, nil, c)
}

func FailWithData(status net.Status, data interface{}, c *gin.Context) {
	Result(status, data, c)
}
