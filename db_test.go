package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinQuery(t *testing.T) {
	server := ServerProcessor{}
	server.ConnectDataBase()
	int2 := server.dbManager.DBGorm.FindMinIdx()

	// assert.err

	// t, expected, actual
	assert.Equal(t, int32(2), int2)

}
