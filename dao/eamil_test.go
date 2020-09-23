package dao

import (
	"testing"
)

func TestCreteTable(t *testing.T) {
	err := createTable()
	if err != nil {
		t.Logf("fail, err:%v \n", err)
	} else {
		t.Log("success")
	}
}
