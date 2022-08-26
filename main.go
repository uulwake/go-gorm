package main

import (
	"fmt"
	"go-gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "secret"
		dbname   = "gorm-sql"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Item{}, &models.Order{}, &models.Outbound{})

	item := models.Item{Name: "item1", Qty: 10, Weight: 20.1}
	db.Create(&item)
	fmt.Println(item)

	db.Transaction(func(tx *gorm.DB) error {
		order := models.Order{RecipientName: "name1", RecipientAddress: "addr1", Shipper: "JNA"}
		if err := tx.Create(&order).Error; err != nil {
			log.Println(err)
			return err
		}

		outbound := models.Outbound{ItemID: item.ID, OrderID: order.ID, Qty: 3}
		if err := tx.Create(&outbound).Error; err != nil {
			log.Println(err)
			return err
		}

		return nil
	})
}
