package repository

import (
	"strconv"

	"github.com/ruang-guru/playground/backend/basic-golang/cashier-app/db"
)

type ProductRepository struct {
	db db.DB
}

func NewProductRepository(db db.DB) ProductRepository {
	return ProductRepository{db}
}

func (u *ProductRepository) LoadOrCreate() ([]Product, error) {
<<<<<<< HEAD
	//beginanswer
	records, err := u.db.Load("products")
	if err != nil {
		records = [][]string{
			{"category", "product_name", "price"},
		}
		if err := u.db.Save("product", records); err != nil {
			return nil, err
		}
	}

	result := make([]Product, 0)
	for i := 1; i < len(records); i++ {
		price, err := strconv.Atoi(records[i][2])
		if err != nil {
			return nil, err
		}

		user := Product{
			Category:    records[i][0],
			ProductName: records[i][1],
			Price:       price,
		}
		result = append(result, user)
	}

	return result, nil
	//endanswer return []Product{}, nil
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	//beginanswer
	productItems, err := u.LoadOrCreate()
	if err != nil {
		return nil, err
	}
	return productItems, nil
	//endanswer return []Product{}, nil
=======
	return []Product{}, nil // TODO: replace this
}

func (u *ProductRepository) SelectAll() ([]Product, error) {
	return []Product{}, nil // TODO: replace this
>>>>>>> 6a266f35b3e5d854980b80d4d6208d897f7008b9
}
