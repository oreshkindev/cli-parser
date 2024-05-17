package main

import (
	"cli-parser/database"
	"cli-parser/internal"
	"context"
	"log"
	"os"
	"time"
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

	// синхронизируем данные каждые 10 минут
	ticker := time.NewTicker(10 * time.Minute)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err = sync(manager); err != nil {
				return err
			}
		}
	}
}

func sync(manager *internal.Manager) error {
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

	// синхронизируем товарные позиции
	if err = manager.Product.Sync(); err != nil {
		return err
	}

	// синхронизируем изображения товарных позиции
	if err = manager.ProductImage.Sync(); err != nil {
		return err
	}

	return nil
}
