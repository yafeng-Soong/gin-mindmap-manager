package theme

import (
	"time"
)

type Theme struct {
	Id          int
	CreatorId   int
	Name        string
	Description string
	State       int
	CreateTime  time.Time
	UpdateTime  time.Time
}

type ThemeState struct {
	State string
	Code  int
}

var (
	NORMAL  = &ThemeState{State: "正常", Code: 0}
	REMOVED = &ThemeState{State: "被删除", Code: 1}
)

func (t *Theme) TableName() string {
	return "theme"
}
