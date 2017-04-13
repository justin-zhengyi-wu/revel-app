package dao

import (
	"database/sql"
	"fmt"
	"github.com/justin-zhengyi-wu/revel-app/app"
	"github.com/justin-zhengyi-wu/revel-app/app/models"
	"github.com/revel/revel"
)

type User struct {
}

func (this *User) List() ([]models.User, error) {
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

func (this *User) Add(item models.User) (sql.Result, error) {
	sql := "insert into user (status) value (%d)"
	sql = fmt.Sprintf(sql, item.Status)
	result, err := app.DB.Exec(sql)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (this *User) Update(item models.User) (int64, error) {
	sql := "update user set status=%d where id=%d"
	sql = fmt.Sprintf(sql, item.Status, item.Id)
	result, err := app.DB.Exec(sql)
	if err != nil {
		return 0, err
	}
	rows, err2 := result.RowsAffected()
	if err2 != nil {
		return 0, err
	}
	return rows, err
}

func (this *User) Delete(itemId int64) (bool, error) {
	sql := "delete from user where id=%d"
	sql = fmt.Sprintf(sql, itemId)
	result, err := app.DB.Exec(sql)
	if err != nil {
		return false, err
	}
	_, err2 := result.RowsAffected()
	if err2 != nil {
		return false, err
	}
	return true, err

}

func (this *User) FindById(itemId int64) (models.User, error) {
	sql := "select ID, status from user where id=?"
	var user models.User
	err := app.DB.QueryRow(sql, itemId).Scan(&user.Id, &user.Status)
	return user, err
}
