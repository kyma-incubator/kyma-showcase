package database

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	t.Run("should return nil when connection is already initialized", func(t *testing.T) {
		//given
		client, _ := redismock.NewClientMock()
		d := Database{address: "localhost:8081", password: "", connection: client, ctx: context.Background()}

		//when
		err := d.Connect()

		//then
		assert.NoError(t, err)
	})
	t.Run("should return error when trying to ping database", func(t *testing.T) {

		//given
		d := Database{address: "localhost:8081", password: "", ctx: context.Background()}

		//when
		err := d.Connect()

		//then
		assert.Error(t, err)
		assert.NotEqual(t, nil, d.connection)
	})
}

func TestInsert(t *testing.T) {
	const key, value = "key", "value"
	t.Run("should return error when connection is not initialized", func(t *testing.T) {
		//given
		_, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		database := Database{connection: nil}

		//when
		err := database.Insert(key, value)

		//then
		assert.Error(t, err, errors.New("INSERT to db: connection not initialized"))
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return error when error occurred in adding value to database", func(t *testing.T) {
		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		clientMock.ExpectSet(key, value, 0)

		//when
		err := database.Insert(key, value)

		//then
		assert.Error(t, err)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should insert data correctly to database with no errors", func(t *testing.T) {

		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		clientMock.ExpectSet(key, value, 0).SetVal(value)

		//when
		err := database.Insert(key, value)

		//then
		assert.NoError(t, err)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}

func TestGet(t *testing.T) {
	const key = "key"
	t.Run("should return error when connection is not initialized", func(t *testing.T) {
		//given
		_, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		database := Database{connection: nil}

		//when
		_, err := database.Get(key)

		//then
		assert.Error(t, err, errors.New("GET from db: connection not initialized"))
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return error when key doesn't exists in database", func(t *testing.T) {
		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		clientMock.ExpectGet(key).RedisNil()

		//when
		_, err := database.Get(key)

		//then
		assert.Equal(t, "GET from db:key "+key+" does not exist", err.Error())
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return error when error occurred in getting value from database", func(t *testing.T) {

		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		clientMock.ExpectGet(key)

		//when
		_, err := database.Get(key)

		//then
		assert.Contains(t, err.Error(), "GET from db:error: ")
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return error when key exists but value is empty", func(t *testing.T) {

		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		clientMock.ExpectGet(key).SetVal("")

		//when
		val, err := database.Get(key)

		//then
		assert.Equal(t, "", val)
		assert.Equal(t, "GET from db:for key "+key+" value is empty", err.Error())
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return no error when data is got correctly from database", func(t *testing.T) {
		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(true)
		database := Database{connection: client}
		value := "value"
		clientMock.ExpectGet(key).SetVal(value)

		//when
		_, err := database.Get(key)

		//then
		assert.NoError(t, err)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}

func TestGetAll(t *testing.T) {
	t.Run("should return error when connection is not initialized", func(t *testing.T) {
		//given
		_, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		database := Database{connection: nil}

		//when
		_, err := database.GetAll()

		//then
		assert.Error(t, err, errors.New("GETALL from db: connection not initialized"))
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return error when error occurred in getting keys from database", func(t *testing.T) {
		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		database := Database{connection: client}
		clientMock.ExpectKeys("*")

		//when
		val, err := database.GetAll()

		//then
		assert.Error(t, err)
		var n []string = nil
		assert.Equal(t, n, val)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return no error when all keys have been got from database", func(t *testing.T) {
		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		database := Database{connection: client}
		clientMock.ExpectKeys("*").SetVal([]string{"1", "2", "3"})

		//when
		_, err := database.GetAll()

		//then
		assert.NoError(t, err)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("should return no error when database is empty", func(t *testing.T) {
		//given
		client, clientMock := redismock.NewClientMock()
		clientMock.ClearExpect()
		clientMock.MatchExpectationsInOrder(false)
		database := Database{connection: client}
		clientMock.ExpectKeys("*").SetVal([]string{})

		//when
		_, err := database.GetAll()

		//then
		assert.NoError(t, err)
		if err := clientMock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
