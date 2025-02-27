package casbin_test

import (
	"testing"
	"qqlx/test"
)

func TestCasbinCreateRole(t *testing.T) {
	// test.Enforcer.AddPolicy("admin", "*", "*")

	// ok, err := test.Enforcer.AddPolicies([][]string{
	// 	{"admin", "user", "create"},
	// 	{"admin", "user", "delete"},
	// 	{"admin", "user", "update"},
	// 	{"admin", "user", "get"},
	// })
	// t.Logf("ok: %v, err: %v", ok, err)
	// policsy, err := test.Enforcer.GetFilteredPolicy(0, "admin")
	// t.Logf("policsy: %v, err: %v", policsy, err)
	ok, err := test.Enforcer.RemovePolicies([][]string{
		{"admin", "user", "create"},
		{"admin", "user", "delete"},
		{"admin", "user", "update"},
		{"admin", "user", "get"},
	})
	t.Logf("ok: %v, err: %v", ok, err)
	// test.Enforcer.SavePolicy()
	// ok, err := test.Enforcer.Enforce("qq", "user", "create")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if !ok {
	// 	t.Fatal("authentication failed")
	// }
}
