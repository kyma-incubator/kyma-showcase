package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Database struct {
	address    string
	password   string
	connection *redis.Client
	ctx        context.Context
}

func NewDatabaseConnection(address, password string) Database {
	return Database{
		address:  address,
		password: password,
		ctx:      context.Background(),
	}
}

func (d Database) Connect() error {
	d.connection = redis.NewClient(&redis.Options{
		Addr:     d.address,
		Password: d.password,
		DB:       0,
	})

	fmt.Printf("%s, %s\n", d.address, d.password)
	fmt.Println(d.connection)
	if err := d.connection.Ping(d.ctx).Err(); err != nil {
		return err
	}
	pong, err := d.connection.Ping(d.ctx).Result()
	fmt.Printf("pong %s, err %s\n", pong, err)
	//d.connection = rdb
	return nil
}

func (d Database) InsertToDB(key string, value string) error {
	if d.connection == nil {
		return errors.New("Connection not initalized1")
	}

	_, err := d.connection.Set(d.ctx, key, value, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (d Database) GetFromDB(key string) (interface{}, error) {
	if d.connection == nil {
		return nil, errors.New("Connection not initalized2")
	}

	val, err := d.connection.Get(d.ctx, key).Result()

	switch {
	case err == redis.Nil:
		fmt.Printf("%s key does not exist", key)
	case err != nil:
		fmt.Println(err)
	case val == "":
		fmt.Println("value is empty")
	}

	return val, err
}

func (d Database) GetAllKeys() ([]string, error) {
	if d.connection == nil {
		return nil, errors.New("Connection not initalized3")
	}

	keys, err := d.connection.Keys(d.ctx, "*").Result()

	if err != nil {
		return nil, err
	}
	return keys, nil
}
