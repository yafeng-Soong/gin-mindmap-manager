package response

import (
	"net/http"
	"paper-manager/model/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type PageResponse struct {
	CurrentPage int64       `json:"currentPage"`
	PageSize    int64       `json:"pageSize"`
	Total       int64       `json:"total"`
	Pages       int64       `json:"pages"` // 总页数
	Data        interface{} `json:"data"`
}

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS().Code, map[string]interface{}{}, SUCCESS().Msg, c)
}

func OkWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS().Code, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS().Code, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR().Code, map[string]interface{}{}, ERROR().Msg, c)
}

func FailWithMessage(code int, message string, c *gin.Context) {
	Result(code, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR().Code, data, message, c)
}

func FailWithError(err *errors.MyError, c *gin.Context) {
	if err.Data == nil {
		err.Data = err.Msg
	}
	Result(err.Code, err.Data, err.Msg, c)
}

func ServerError(msg string, c *gin.Context) {
	Result(ERROR().Code, msg, ERROR().Msg, c)
}

func NewPageResponse(currentPage int64, pageSize int64) *PageResponse {
	return &PageResponse{
		CurrentPage: currentPage,
		PageSize:    pageSize,
	}
}

func (p *PageResponse) CountPages() (ok bool) {
	pages := p.Total / p.PageSize
	if p.Total%p.PageSize != 0 {
		pages++
	}
	if p.CurrentPage > pages {
		ok = false
	} else {
		ok = true
	}
	p.Pages = pages
	return
}
