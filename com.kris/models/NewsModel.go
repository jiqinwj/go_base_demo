package models

import "github.com/pquerna/ffjson/ffjson"

type NewsModel struct {
	NewsID int
	NewsTitle string
}

func (news NewsModel) ToJSON() string {
	result,err:=ffjson.Marshal(news)
	if err!=nil {
		return err.Error()
	} else {
		return string(result)
	}
}
