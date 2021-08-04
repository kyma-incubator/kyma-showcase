package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Database struct {
	address    string
	password   string
	connection *redis.Client //todo:zdefiniowac interface okreslajacy kontrakt miedzy nami a redisem
	ctx        context.Context
}

func NewDatabaseConnection(address, password string) Database {
	return Database{
		address:  address,
		password: password,
		ctx:      context.Background(),
	}
}

func (d *Database) Connect() error {
	if d.connection != nil {
		log.Info("Connection to database is already established")
		return nil
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
		return errors.New("INSERTTODB: connection not initialized")
	}
	_, err := d.connection.Set(d.ctx, key, value, 0).Result()
	return err
}

func (d Database) GetFromDB(key string) (interface{}, error) {
	if d.connection == nil {
		return nil, errors.New("GETFROMDB: connection not initialized")
	}

	val, err := d.connection.Get(d.ctx, key).Result()

	switch {
	case err == redis.Nil:
		err = errors.New("GETFROMDB:key " +  key + " does not exist")
	case err != nil:
		err = errors.New("GETFROMDB:error: " +  err.Error() + " occurred in getting data from db")
	case val == "":
		err = errors.New("GETFROMDB:for key " +  key + " value is empty")
	}
	return val, err
}

func (d Database) GetAllKeys() ([]string, error) {
	if d.connection == nil {
		return nil, errors.New("GETALLKEYS: connection not initialized")
	}

	keys, err := d.connection.Keys(d.ctx, "*").Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
