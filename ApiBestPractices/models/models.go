package models

import (
	"database/sql"
	"log"
)

var DB *sql.DB

type Brand struct {
	Name    string
	Segment int
}

func ConnectDb()(*sql.DB, error){
	connStr := "postgres://localhost:5432/bookings?sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Veritabanı bağlantısı başarılı")
	
	return DB, err
}

func AllBrands() ([]Brand, error) {
	rows, err := DB.Query("SELECT * FROM brands")
	if (err != nil){
		return nil, err
	}
	defer rows.Close()

	var brands []Brand
	for (rows.Next()){
		var brand Brand

		err := rows.Scan(&brand.Name, &brand.Segment)
		if (err != nil){
			return nil, err
		}

		brands = append(brands, brand)
	}

	if err = rows.Err(); err != nil {
        return nil, err
    }

    return brands, nil
}
