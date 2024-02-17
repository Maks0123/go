package models

import "errors"

// Product структура для товару
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ProductsList struct {
	products []Product
}

// Products список товарів
var products = ProductsList{
	products: []Product{
		{ID: 1, Name: "Кава"},
		{ID: 2, Name: "Чай"},
	},
}

func GetData() *ProductsList {
	return &products
}

func (pl *ProductsList) Get() []Product {
	return products.products
}

func (pl *ProductsList) Add(p Product) {
	pl.products = append(pl.products, p)
}

func (pl *ProductsList) Delete(id int) {
	for i, p := range pl.products {
		if p.ID == id {
			pl.products = append((pl.products)[:i], (pl.products)[i+1:]...)
			break
		}
	}
}

func (pl *ProductsList) Update(p Product) error {
	for i, pr := range pl.products {
		if pr.ID == p.ID {
			(pl.products)[i] = p
			return nil
		}
	}

	return errors.New("product not found")
}

func (pl *ProductsList) Find(id int) *Product {
	for _, p := range pl.products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}

/////////////////////////////////////////////////

// Customer структура для користувача
type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
type CustomersList struct {
	customers []Customer
}

// Customerss список користувачів
var customers = CustomersList{
	customers: []Customer{
		{ID: 1, Name: "Jack", Code: "1111111"},
		{ID: 2, Name: "John", Code: "2222222"},
		{ID: 3, Name: " j. John", Code: "33333333"},
	},
}

func GetCustomersData() *CustomersList {
	return &customers
}

func (pl *CustomersList) Get() []Customer {
	return customers.customers
}

func (pl *CustomersList) Add(p Customer) {
	pl.customers = append(pl.customers, p)
}

func (pl *CustomersList) Delete(id int) {
	for i, p := range pl.customers {
		if p.ID == id {
			pl.customers = append((pl.customers)[:i], (pl.customers)[i+1:]...)
			break
		}
	}
}

func (pl *CustomersList) Update(p Customer) error {
	for i, pr := range pl.customers {
		if pr.ID == p.ID {
			(pl.customers)[i] = p
			return nil
		}
	}

	return errors.New("customer not found")
}

func (pl *CustomersList) Find(id int) *Customer {
	for _, p := range pl.customers {
		if p.ID == id {
			return &p
		}
	}
	return nil
}


