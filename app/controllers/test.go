package controllers

import (
	"github.com/justin-zhengyi-wu/revel-app/app/dao"
	"github.com/justin-zhengyi-wu/revel-app/app/models"
	"github.com/revel/revel"
)

type Test struct {
	*revel.Controller
}

func (c Test) Index() revel.Result {
	testDao := dao.TestDao{}
	list, err := testDao.List()
	if err != nil {
		message := models.Message{
			Code: models.InnerError,
			Msg:  "fail",
			Data: []models.Test{},
		}
		return c.RenderJSON(message)
	}
	if list == nil {
		list = []models.Test{}
	}
	return c.RenderJSON(models.Message{Code: models.OK, Msg: "success", Data: list})
}
