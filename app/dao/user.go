package dao

import (
	"github.com/justin-zhengyi-wu/revel-app/app"
	"github.com/justin-zhengyi-wu/revel-app/app/models"
	"github.com/revel/revel"
)

type User struct {
}

func (this *User) List() (list []models.User, err error) {
	sql := "select id, status from test"
	rows, err := app.DB.Query(sql)
	if err != nil {
		revel.ERROR.Println("Sql error", err)
		return nil, err
	}

	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	results := []models.User{}
	for rows.Next() {
		result := models.User{}
		err := rows.Scan(&result.Id, &result.Status)
		if err != nil {
			revel.INFO.Println(err)
		}
		results = append(results, result)
	}
	err = rows.Err()
	if err != nil {
		revel.INFO.Println(err)
	}
	return results, err
}
