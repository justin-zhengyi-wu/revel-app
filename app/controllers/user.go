package controllers

import (
	"github.com/justin-zhengyi-wu/revel-app/app/dao"
	"github.com/justin-zhengyi-wu/revel-app/app/models"
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	testDao := dao.User{}
	list, err := testDao.List()
	if err != nil {
		message := models.Message{
			Code: models.InnerError,
			Msg:  "fail",
			Data: []models.User{},
		}
		return c.RenderJSON(message)
	}
	if list == nil {
		list = []models.User{}
	}
	return c.RenderJSON(models.Message{
		Code: models.OK,
		Msg:  "success",
		Data: list,
	})
}
