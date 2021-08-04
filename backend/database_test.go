package main

import (
	"context"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDatabaseConnection(t *testing.T) {
	t.Run("Tests if data is properly assigned to the database structure", func(t *testing.T) {
		//given
		address := "localhost:8081"
		password := ""

		//when
		FakeDatabase := NewDatabaseConnection(address, password)

		//then
		assert.Equal(t, Database{address: "localhost:8081", password: "", ctx: context.Background()}, FakeDatabase)
		assert.NotEqual(t, Database{address: "address", password: "", ctx: context.Background()}, FakeDatabase)
		assert.NotEqual(t, Database{address: "localhost:8081", password: "password", ctx: context.Background()}, FakeDatabase)
	})
}

func TestConnect(t *testing.T){
	t.Run("Tests if connection is already established", func(t *testing.T){
		//given
		d:= Database {address: "localhost:8081", password: "", ctx:  context.Background()}
		d.Connect()

		//when
		err := d.Connect()

		//then
		assert.Equal(t, nil,err)
	})
	t.Run("Tests connecting to database with correct Database fields", func(t *testing.T){
		//given
		d:= Database {address: "localhost:8081", password: "", ctx:  context.Background()}
		//when
		d.Connect()
		//then
		assert.NotEqual(t, /*redis.NewClient(&redis.Options{Addr:"localhost:8081", Password: "", DB: 0}),*/nil, d.connection)
	})
}



func TestGetFromDB(t *testing.T) {
	//given
	client, clientMock := redismock.NewClientMock()
	clientMock.MatchExpectationsInOrder(true)
	var ctx = context.TODO()
	clientMock.ClearExpect()
	//key, value
	database := Database{connection: client}
	test := "1"
	testval := "2"
	client.Set(ctx, test, testval, 0)
	clientMock.ExpectGet(test).SetVal(testval)
	val, err := database.GetFromDB(test)
	if err != nil {
	}
	assert.Equal(t, testval, val)
	if err := clientMock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestInsertToDB(t *testing.T) { //todo:fajnie jakby uzyc redisowego geta
	//given
	client, clientMock := redismock.NewClientMock()
	clientMock.MatchExpectationsInOrder(true)
	//key, value
	database := Database{connection: client}
	//when
	test := "1"
	testval := "2"
	clientMock.ExpectSet(test, testval, 0).SetVal(testval)
	err := database.InsertToDB(test, testval)
	if err != nil {
	}
	//then
	clientMock.ExpectGet(test).SetVal(testval)
	val, err := database.GetFromDB(test)
	if err != nil {
	}
	assert.Equal(t, testval, val)
	if err := clientMock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetAllKeysDB(t *testing.T) {
	//given
	client, clientMock := redismock.NewClientMock()
	clientMock.MatchExpectationsInOrder(false)
	var i = 0
	var ctx = context.TODO()
	//key, value
	database := Database{connection: client}
	//when
	client.Set(ctx, "1", "2", 0)
	client.Set(ctx, "2", "3", 0)
	client.Set(ctx, "3", "4", 0)
	//then
	clientMock.ExpectKeys("*")
	val, err := database.GetAllKeys()
	if err != nil {
	}
	for _, key := range val {
		i++
		assert.Equal(t, i, key)
	}
	if err := clientMock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}



//func TestGetAndInsertDB(t *testing.T) { //todo:działa ale fajnie jakby to rozbić
//	//given
//	client, clientMock := redismock.NewClientMock()
//	clientMock.MatchExpectationsInOrder(true)
//	//key, value
//	database := Database{connection: client}
//	//when
//	test := "1"
//	testval := "2"
//	clientMock.ExpectSet("1", "2", 0).SetVal(testval)
//	err := database.InsertToDB(test, testval)
//	if err != nil {
//	}
//	//then
//	clientMock.ExpectGet("1").SetVal(testval)
//	val, err := database.GetFromDB(test)
//	if err != nil {
//	}
//	assert.Equal(t, testval, val)
//	if err := clientMock.ExpectationsWereMet(); err != nil {
//		t.Error(err)
//	}
//}






//func TestInsertToDB(t *testing.T){
//	//t.Run("", func(t *testing.T){
//	//
//	//})
//	t.Run("", func(t *testing.T){
//
//	})
//}

	//then
	//assert.Equal(t, val, "1" )
	//newsID := 123456789
	//key := fmt.Sprintf("news_redis_cache_%d", newsID)
	//
	//// mock ignoring `call api()`
	//
	//mock.ExpectGet(key).RedisNil()
	//mock.Regexp().ExpectSet(key, `[a-z]+`, 30 * time.Minute).SetErr(errors.New("FAIL"))
	//
	//_, err := NewsInfoForCache(db, newsID)
	//if err == nil || err.Error() != "FAIL" {
	//	t.Error("wrong error")
	//}
	//
	//if err := mock.ExpectationsWereMet(); err != nil {
	//	t.Error(err)
	//}

//
//
//	t.Run("", func(t *testing.T){
//	})
//}


//func (d Database) InsertToDB(key string, value string) error {
//	if d.connection == nil {
//		return errors.New("INSERTTODB: connection not initialized")
//	}
//	_, err := d.connection.Set(d.ctx, key, value, 0).Result()
//	return err
//}

//func TestDBGetAllHandler(t *testing.T) {
//	t.Run("should return error when database is empty", func(t *testing.T) {
//
//		//given
//		req, err := http.NewRequest("GET", "v1/images", nil)
//
//		require.NoError(t, err)
//		recorder := httptest.NewRecorder()
//		dbManagerMock := mocks.DBManager{}
//		testSubject := NewHandler(&dbManagerMock)
//		dbManagerMock.On("GetAllKeys").Return(nil, errors.New("Empty database, returned nil"))
//
//		//when
//		testSubject.DBGetAllHandler(recorder, req)
//
//		//then
//		assert.Equal(t, http.StatusNotFound, recorder.Code)
//
//	})
//}
//todo:miniredis
//func newMockRedis() *redismock.ClientMock {
//	mr, err := miniredis.Run()
//	if err != nil {
//		panic(err)
//	}
//	client := redis.NewClient(&redis.Options{
//		Addr: mr.Addr(),
//	})
//	return redismock.NewNiceMock(client)
//}

//
//
//func TestGetfromDB(t *testing.T){
//
//	//given
//	r :=newMockRedis()
//
//	r.Set("1", "2", 1*time.Minute)
//	r.Get()
//	r.Insert
//	//when
//
//	//then
//
//	r.D
//
//
//
//
//
//	assert.Equal(t,)
//
//	defer srv.Close()
//}