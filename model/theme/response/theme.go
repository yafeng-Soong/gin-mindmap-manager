package response

import (
	"paper-manager/model/theme"
	"paper-manager/utils"
)

type ThemeResponse struct {
	Id          int    `json:"id"`
	Creator     string `json:"creator"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

func NewThemeResponse(t theme.Theme) *ThemeResponse {
	return &ThemeResponse{
		Id:          t.Id,
		Name:        t.Name,
		Description: t.Description,
		CreateTime:  t.CreateTime.Format(utils.FORMAT),
		UpdateTime:  t.UpdateTime.Format(utils.FORMAT),
	}
}
