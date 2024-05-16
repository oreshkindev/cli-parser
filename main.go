package main

import (
	"cli-parser/database"
	"cli-parser/internal"
	"context"
	"log"
	"os"
)

var (
	connection *database.Database
	err        error
)

func main() {
	context := context.Background()

	if err = run(context); err != nil {
		log.Fatal(err)
	}
}

func run(context context.Context) error {

	// подключаемся к базе данных
	if connection, err = database.New(context, os.Getenv("DATABASE_URL")); err != nil {
		return err
	}
	defer connection.Close()

	// инициализируем менеджер
	manager := internal.New(context, connection)

	// синхронизируем бренды
	if err = manager.Brand.Sync(); err != nil {
		return err
	}

	// синхронизируем характеристики
	if err = manager.Characteristic.Sync(); err != nil {
		return err
	}

	// синхронизируем расширенные характеристики
	if err = manager.CharacteristicExtended.Sync(); err != nil {
		return err
	}

	return nil
}
