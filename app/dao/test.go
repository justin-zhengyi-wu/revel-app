package dao

import (
	"fmt"
	"github.com/justin-zhengyi-wu/revel-app/app"
	"github.com/justin-zhengyi-wu/revel-app/app/models"
	"github.com/revel/revel"
)

type TestDao struct {
}

func (this *TestDao) List() (list []models.Test, err error) {
	sql := "select id, status from test"
	rows, err := app.DB.Query(sql)
	if err != nil {
		revel.INFO.Println("Sql error", err)
	}
	fmt.Printf("%v", rows)
	defer rows.Close()
	results := []models.Test{}
	for rows.Next() {
		result := models.Test{}
		err = rows.Scan(&result.Id, &result.Status)
		results = append(results, result)
	}
	err = rows.Err()
	if err != nil {
		revel.INFO.Println(err)
	}
	return results, err
}
