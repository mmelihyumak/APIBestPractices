package repository

import (
	"goapibestpractices/models"
	"log"
)

var err error

func AddBrand(brand models.Brand){
	_, err = models.DB.Exec("INSERT INTO brands (name, segment) VALUES ($1, $2)", brand.Name, brand.Segment)
	if err != nil {
		log.Fatalf("Failed to insert record: %v\n", err)
	}
}

func GetAllBrands()([]models.Brand, error){
	rows, err := models.DB.Query("SELECT * FROM brands")
	if (err != nil){
		return nil, err
	}
	defer rows.Close()

	var brands []models.Brand
	for (rows.Next()){
		var brand models.Brand

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

func GetBrandByName(brandName string)(models.Brand, error){
	var brand models.Brand

	row, err := models.DB.Query("SELECT * FROM brands WHERE name = ($1)", brandName)
	if err != nil{
		log.Fatal("Name alanı ile ilgili kayıt bulunamadı")
		return brand, err
	}

	err = row.Scan(&brand.Name, &brand.Segment)
	if err != nil{
		log.Fatal("Bilinmeyen bir hata oluştu")
		return brand, nil
	}

	return brand, nil
}