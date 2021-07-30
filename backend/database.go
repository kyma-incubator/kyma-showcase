package main

import (
	"context"
	"errors"
	//"fmt"
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

func (d *Database) Connect() error { //niejawnie wykonac Connect (?) w innej metodzie, na razie zostawic
	if d.connection != nil {
		//return errors.New("connection already exists")
		return nil //ta opcja
	}

	d.connection = redis.NewClient(&redis.Options{
		Addr:     d.address,
		Password: d.password,
		DB:       0,
	})

	if err := d.connection.Ping(d.ctx).Err(); err != nil {
		return err
	}
	return nil
}

func (d Database) InsertToDB(key string, value string) error {
	if d.connection == nil {
		return errors.New("connection not initialized")
	}
	_, err := d.connection.Set(d.ctx, key, value, 0).Result()
	return err
}

func (d Database) GetFromDB(key string) (interface{}, error) {
	if d.connection == nil {
		return nil, errors.New("connection not initialized")
	}

	val, err := d.connection.Get(d.ctx, key).Result() ///zwracanie errorów -> gdy jest error to nie zwraca wartości?!

	switch {
	case err == redis.Nil:
		err = errors.New("key"+key+"does not exist")
	case err != nil:
		err = errors.New("func get error")
	case val == "":
		err = errors.New("for key "+key+ " value is empty")
	}
	return val, err
}

func (d Database) GetAllKeys() ([]string, error) {
	if d.connection == nil {
		return nil, errors.New("connection not initialized")
	}

	keys, err := d.connection.Keys(d.ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}
