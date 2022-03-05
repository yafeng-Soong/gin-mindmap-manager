package response

import (
	"net/http"
	"paper-manager/database"
	"paper-manager/model/errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type PageResponse struct {
	CurrentPage int64       `json:"currentPage"`
	PageSize    int64       `json:"pageSize"`
	Total       int64       `json:"total"`
	Pages       int64       `json:"pages"` // 总页数
	Data        interface{} `json:"data"`
}

type ResponseEnums struct {
	Code int
	Msg  string
}

var (
	SUCCESS = &ResponseEnums{Code: 200, Msg: "操作成功"}
	ERROR   = &ResponseEnums{Code: 400, Msg: "操作失败"}
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS.Code, data, SUCCESS.Msg, c)
}

func FailWithError(err *errors.MyError, c *gin.Context) {
	if err.Data == nil {
		err.Data = err.Msg
	}
	Result(err.Code, err.Data, err.Msg, c)
}

func ServerError(msg string, c *gin.Context) {
	Result(ERROR.Code, msg, ERROR.Msg, c)
}

func NewPageResponse(page *database.Page) *PageResponse {
	return &PageResponse{
		CurrentPage: page.CurrentPage,
		PageSize:    page.Pages,
		Pages:       page.Pages,
		Total:       page.Total,
	}
}
