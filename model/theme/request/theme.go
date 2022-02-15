package request

import (
	"paper-manager/model/common/request"
)

type ThemeQueryVo struct {
	request.PageInfo
	Name    string `json:"name"`
	Removed bool   `json:"removed"`
}

type ThemeUpdateVo struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
