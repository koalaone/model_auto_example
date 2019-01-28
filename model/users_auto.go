package model

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/koalaone/model_auto/mysql"
)

type User struct {
	ID    int    `json:"id"`    // 自动编号
	Uid   string `json:"uid"`   // 用户编码
	Uname string `json:"uname"` // 用户名
	Age   int    `json:"age"`   // 年龄
}

var UserReal = &User{}

func (t *User) TableName() string {
	return "users"
}

func GetUserByWheres(wheres map[string]interface{}) (*User, error) {
	if len(wheres) == 0 {
		return nil, errors.New("param[wheres] length is zero")
	}

	var object User
	var resultObjs []User
	err := mysql.SearchObject(&object, wheres, &resultObjs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if len(resultObjs) == 0 {
		return nil, nil
	}

	resultObj := resultObjs[0]
	if err != nil {
		return nil, err
	}

	return &resultObj, nil
}

func GetUserMultiByWheres(wheres map[string]interface{}) ([]User, error) {
	if len(wheres) == 0 {
		return nil, errors.New("param[wheres] length is zero")
	}

	var object User
	var resultObjs []User
	err := mysql.SearchObject(&object, wheres, &resultObjs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return resultObjs, nil
}

func GetUserMultiByInWheresLimit(wheres, ins map[string]interface{}, orders string,
	limit, offset int) ([]User, error) {

	if limit < 0 {
		return nil, errors.New("param[limit] less than zero")
	}

	if offset < 0 {
		return nil, errors.New("param[offset] less than zero")
	}

	var object User
	var resultObjs []User

	err := mysql.SearchObjectByOrder(&object, wheres, ins, orders, limit, offset, &resultObjs)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	if len(resultObjs) == 0 {
		return nil, nil
	}

	return resultObjs, nil
}

func CreateUser(value *User) error {
	if value == nil {
		return errors.New("param[value] is nil")
	}

	err := mysql.CreateObject(value)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(wheres, updates map[string]interface{}) error {
	if len(wheres) == 0 {
		return errors.New("param[wheres] length is zero")
	}
	if len(updates) == 0 {
		return errors.New("param[updates] length is zero")
	}

	var object User
	err := mysql.UpdateObject(&object, wheres, updates)
	if err != nil {
		return err
	}

	return nil
}
