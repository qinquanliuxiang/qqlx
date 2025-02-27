package data_test

import (
	"fmt"
	"testing"
	"qqlx/model"
	"qqlx/test"
)

func TestAssociation(t *testing.T) {
	db := test.DB
	defer test.Close1()
	defer test.Close2()
	// db.Create(&model.Role{
	// 	Name:        "admin",
	// 	Description: "admin",
	// 	Policys: []*model.Policy{
	// 		{
	// 			Name:   "admin",
	// 			Path:   "/admin",
	// 			Method: "get",
	// 		},
	// 		{
	// 			Name:   "admin",
	// 			Path:   "/admin",
	// 			Method: "post",
	// 		},
	// 	},
	// })
	// var role model.Role
	// if err := db.Model(&model.Role{}).Take(&role, 3).Error; err != nil {
	// 	t.Fatal(err)
	// }
	// total := db.Model(&role).Association("Policys").Count()
	// t.Log(total)

	// 在原有基础上追加，多次执行会追加多条记录
	// if err := db.Model(&role).Association("Policys").Append(&model.Policy{
	// 	Name:   "admin",
	// 	Path:   "/admin",
	// 	Method: "put",
	// }); err != nil {
	// 	t.Fatal(err)
	// }

	// 替换,将现有 Role 关联所有 Policy 的记录替换为指定的 Policy
	// if err := db.Model(&role).Association("Policys").Replace([]*model.Policy{
	// 	{
	// 		Name:   "admin",
	// 		Path:   "/admin",
	// 		Method: "put",
	// 	},
	// 	{
	// 		Name:   "admin",
	// 		Path:   "/admin",
	// 		Method: "test",
	// 	},
	// }); err != nil {
	// 	t.Fatal(err)
	// }

	// 删除关联，只会删除中间表的记录。被删除表的记录不会被删除
	// if err := db.Model(&role).Association("Policys").Delete([]*model.Policy{
	// 	{
	// 		MetaData: &model.MetaData{
	// 			ID: 8,
	// 		},
	// 		Name:   "admin",
	// 		Path:   "/admin",
	// 		Method: "get",
	// 	},
	// 	{
	// 		MetaData: &model.MetaData{
	// 			ID: 9,
	// 		},
	// 		Name:   "admin",
	// 		Path:   "/admin",
	// 		Method: "test",
	// 	},
	// }); err != nil {
	// 	t.Fatal(err)
	// }

	// var user *model.User
	// if err := db.Model(user).Take(&user, "name = ?", "admin").Error; err != nil {
	// 	t.Fatal(err)
	// }

	var role *model.Role
	if err := db.Model(role).Take(&role, "name = ?", "admin").Error; err != nil {
		t.Fatal(err)
	}

	// if err := db.Model(&user).Association("Role").Append(role); err != nil {
	// 	t.Fatal(err)
	// }
	if err := db.Model(role).Association("Policys").Append(&model.Policy{
		Name:   "admin",
		Path:   "*",
		Method: "*",
	}); err != nil {
		panic(fmt.Errorf("failed to append policy, %w", err))
	}
}
