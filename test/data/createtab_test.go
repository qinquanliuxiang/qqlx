package data_test

import (
	"testing"
	"qqlx/model"
	"qqlx/test"
)

func TestCreateTable(t *testing.T) {
	defer test.Close1()
	defer test.Close2()
	test.DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Policy{})
}
