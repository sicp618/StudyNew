package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	db := initDB()
	defer db.Close()

	assert.NotNil(t, db)
}

func TestInitRedis(t *testing.T) {
	client := initRedis()

	assert.NotNil(t, client)
}

func TestInitService(t *testing.T) {
	db := initDB()
	defer db.Close()

	client := initRedis()

	r := initService(db, client)

	assert.NotNil(t, r)
}