package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinQuery(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()
	int2 := server.dbManager.DBGorm.FindMinIdx()

	// assert.err

	// t, expected, actual
	assert.Equal(t, int32(2), int2)
}

func TestCountAttr(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	cnt := server.dbManager.DBGorm.AttrTotalCount()
	assert.Equal(t, 21, int(cnt))

}

func TestInsertAttrDefaultRow(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	defer server.dbManager.DBGorm.Release()

	server.dbManager.DBGorm.InsertDefaultAttr1()
	server.dbManager.DBGorm.InsertDefaultAttr2()

}
