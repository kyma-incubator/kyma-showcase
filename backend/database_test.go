package main

import (
	"context"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Run("should check that connection is already established and return error", func(t *testing.T) {
		//given
		client, _ := redismock.NewClientMock()
		d := Database{address: "localhost:8081", password: "", connection: client, ctx: context.Background()}

		//when
		err := d.Connect()

		//then
		assert.Equal(t, nil, err)
	})
	t.Run("should connect to database successfully with correct Database fields", func(t *testing.T) { //todo:test refactor

		//given
		d := Database{address: "localhost:8081", password: "", ctx: context.Background()}

		//when
		d.Connect()

		//then
		assert.NotEqual(t, nil, d.connection)
	})
}

func TestGetFromDB(t *testing.T) {
	t.Run("should acquire data correctly using GetFromDB function", func(t *testing.T) {

		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		var ctx = context.TODO()
		database := Database{connection: client}
		test := "1"
		testval := "2"
		client.Set(ctx, test, testval, 0)
		clientMock.ExpectGet(test).SetVal(testval)

		//when
		val, err := database.GetFromDB(test)
		assert.NoError(t, err)

		//then
		assert.Equal(t, testval, val)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}

func TestInsertToDB(t *testing.T) {
	t.Run("should insert data correctly using the InsertToDB function", func(t *testing.T) {

		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		test := "1"
		testval := "2"

		//when
		clientMock.ExpectSet(test, testval, 0).SetVal(testval)
		err := database.InsertToDB(test, testval)
		assert.NoError(t, err)

		clientMock.ExpectGet(test).SetVal(testval)
		val, err := database.GetFromDB(test)
		assert.NoError(t, err)

		//then
		assert.Equal(t, testval, val)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}

func TestGetAllKeysDB(t *testing.T) {
	t.Run("should acquire all keys using the GetAllKeys function", func(t *testing.T) {

		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		var ctx = context.TODO()
		database := Database{connection: client}
		client.Set(ctx, "1", "2", 0)
		client.Set(ctx, "2", "3", 0)
		client.Set(ctx, "3", "4", 0)
		clientMock.ExpectKeys("*")

		//when
		val, err := database.GetAllKeys()
		assert.NoError(t, err)

		//then
		var i = 0
		for _, key := range val {
			i++
			assert.Equal(t, i, key)
		}
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
