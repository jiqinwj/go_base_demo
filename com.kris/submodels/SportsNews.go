package submodels

import (

	"github.com/pquerna/ffjson/ffjson"
	"go_base_demo/com.kris/models"
)

type SportsNews struct {
	Tags []string  //array
	models.NewsModel
}
func (sn SportsNews) ToJSON() string {
	result,err:=ffjson.Marshal(sn)
	if err!=nil {
		return err.Error()
	} else {
		return string(result)
	}
}