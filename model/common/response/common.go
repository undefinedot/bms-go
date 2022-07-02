package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* handler返回响应的代码封装 */

const (
	SUCCESS = 0
	ERROR   = 7
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Ok(c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, "操作成功！", c)
}

func OkWithMsg(msg string, c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	result(SUCCESS, data, "查询成功", c)
}

func OkWithDetail(data interface{}, msg string, c *gin.Context) {
	result(SUCCESS, data, msg, c)
}

func Fail(c *gin.Context) {
	result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMsg(msg string, c *gin.Context) {
	result(ERROR, map[string]interface{}{}, msg, c)
}

// result 重复代码抽离出来
func result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}
