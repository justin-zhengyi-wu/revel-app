package controllers

import (
	"github.com/justin-zhengyi-wu/revel-app/app/dao"
	"github.com/justin-zhengyi-wu/revel-app/app/models"
	"github.com/revel/revel"
	"strconv"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	userDao := dao.User{}
	list, err := userDao.List()
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

func (c User) Add() revel.Result {
	item := models.User{}
	c.Params.Bind(&item.Status, "status")
	userDao := dao.User{}
	_, err := userDao.Add(item)
	if err != nil {
		message := models.Message{
			Code: models.InnerError,
			Msg:  "fail",
			Data: models.User{},
		}
		return c.RenderJSON(message)
	}
	return c.RenderJSON(models.Message{
		Code: models.OK,
		Msg:  "success",
		Data: item,
	})
}

func (c User) Update() revel.Result {
	idStr := c.Params.Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	c.Validation.Required(id)
	if c.Validation.HasErrors() || err != nil {
		c.RenderJSON(models.Message{
			Code: models.ParamError,
			Msg:  "id为空",
		})
	}
	item := models.User{}
	c.Params.Bind(&item.Status, "status")
	userDao := dao.User{}
	_, err2 := userDao.Update(item)
	if err2 != nil {
		message := models.Message{
			Code: models.InnerError,
			Msg:  "fail",
			Data: models.User{},
		}
		return c.RenderJSON(message)
	}
	return c.RenderJSON(models.Message{
		Code: models.OK,
		Msg:  "success",
		Data: item,
	})
}

func (c User) Delete() revel.Result {
	idStr := c.Params.Get("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	c.Validation.Required(id)
	if c.Validation.HasErrors() || err != nil {
		c.RenderJSON(models.Message{
			Code: models.ParamError,
			Msg:  "id为空",
		})
	}
	userDao := dao.User{}
	done, err := userDao.Delete(id)
	if err != nil || done != true {
		message := models.Message{
			Code: models.InnerError,
			Msg:  "fail",
		}
		return c.RenderJSON(message)
	}
	return c.RenderJSON(models.Message{
		Code: models.OK,
		Msg:  "success",
		Data: true,
	})
}
