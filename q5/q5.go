package q5

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Sale struct {
	Code     string
	Products []*Product
}

func CalculateTotalSales(sales map[string]*Sale) float64 {
	var resultado float64
	for _, venda:= range sales {
		for _, produto:= range venda.Products {
			resultado += produto.Price
		}
	}
	return resultado
}
