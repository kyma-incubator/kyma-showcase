package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

//go:generate mockery --name=Database
// Database struct contains address(String), password(string), connection, ctx.
type Database struct {
	address    string
	password   string
	connection *redis.Client
	ctx        context.Context
}

// NewDatabaseConnection returns a Database structure using the given arguments.
func NewDatabaseConnection(address, password string) Database {
	return Database{
		address:  address,
		password: password,
		ctx:      context.Background(),
	}
}

// Connect returns an error if connection already has been initialized, otherwise creates a new connection.
func (d *Database) Connect() error {
	if d.connection != nil {
		log.Info("Connection to database is already initialized")
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

// InsertToDB adds a database entry using provided key(String) and value(String).
func (d Database) InsertToDB(key string, value string) error {
	if d.connection == nil {
		return errors.New("INSERTTODB: connection not initialized")
	}
	_, err := d.connection.Set(d.ctx, key, value, 0).Result()
	return err
}

// GetFromDB requests and receives a value(String) for the given key(String).
func (d Database) GetFromDB(key string) (interface{}, error) {
	if d.connection == nil {
		return nil, errors.New("GETFROMDB: connection not initialized")
	}

	val, err := d.connection.Get(d.ctx, key).Result()

	switch {
	case err == redis.Nil:
		err = errors.New("GETFROMDB:key " + key + " does not exist")
	case err != nil:
		err = errors.New("GETFROMDB:error: " + err.Error() + " occurred in getting data from db")
	case val == "":
		err = errors.New("GETFROMDB:for key " + key + " value is empty")
	}
	return val, err
}

// GetAllKeys returns keys([]string containing all the keys in the database.
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
