package q3

import "errors"

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func UpdateStock(product *Product, sales map[string]int) error {
	if product == nil {

		return errors.New("Not implemented yet")
	}
	for codigo, quantidade := range sales {
		if quantidade <= 0 {
			return errors.New("Not implemented yet")
		}
		if codigo == product.Code {
			novaquantidade := product.Quantity - quantidade
			if novaquantidade < 0 {
				return errors.New("Not implemented yet")
			}
			product.Quantity = novaquantidade
		}
	}
	return nil
}
