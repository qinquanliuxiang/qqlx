package data_test

import (
	"context"
	"testing"
	"qqlx/base/data"
	"qqlx/test"
)

func TestCacge(t *testing.T) {
	// defer test.Close1()
	// defer test.Close2()
	if err := test.Cache.SetInt64(context.Background(), "test", 100, &data.NeverExpires); err != nil {
		t.Logf("set err: %v", err)
	}
	// if err := test.Cache.Del(context.Background(), "test"); err != nil {
	// 	t.Logf("del err: %v", err)
	// }
}
