package request

import (
	"paper-manager/model/common/request"
)

type ThemeQueryVo struct {
	request.PageInfo
	Name    string `json:"name"`
	Removed bool   `json:"removed"`
}
