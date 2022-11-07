package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	PostgresUser = "postgres"
	PostgresPassword = "1234"
	PostgresHost = "localhost"
	PostgresPort = 5432
	PostgresDatabase = "avtomobil"
)

func main() {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPassword,
		PostgresDatabase,
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}

	a := NewDBManager(db)

	id, err := a.CreateAvtomobil(&Avtomobil{
		CategoryName: "vechicle",
		Name: "Tahoi",
		Price: 32000.0,
		ImageUrl: "test_url",
		Images: []*AvtomobilImage{
			{
				ImageUrl: "test_url1",
				SequenceNumber: 1,
			},
			{
				ImageUrl: "test_url2",
				SequenceNumber: 2,
			},
			{
				ImageUrl: "test_url3",
				SequenceNumber: 3,
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to create avtomobil: %v", err)
	}

	Avtomobil, err := a.GetAvtomobil(id)
	if err != nil {
		log.Fatalf("failed to get avtomobil: %v", err)
	}
	fmt.Println(Avtomobil)

	resp, err := a.GetAllAvtomobils(&GetAvtomobilParams{
		Limit: 10,
		Page: 1,
	})

	if err != nil {
		log.Fatalf("failed to get product: %v", err)
	}
	fmt.Printf("%v", resp)

	err = a.UpdateAvtomobil(&avtomobile{
		ID: 4,
		CategoryName: "vechicle",
		Name: "Onix",
		Price: 20000.0,
		ImageUrl: "test_url",
		Images: []*AvtomobilImage{
			{
				ImageUrl: "test_url1",
				SequenceNumber: 1,
			},
			{
				ImageUrl: "test_url2",
				SequenceNumber: 2,
			},
		},
	})
	if err != nil {
		log.Fatalf("failed to update automobile: %v", err)
	}

	err = a.DeleteAvtomobil(Avtomobil.ID)
	if err != nil {
		log.Fatalf("failed to delete avtomobile: %v", err)
	}
}