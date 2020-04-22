package katas

type Offer struct {
	productId string
	quantity  int
	price     float64
}

type Product struct {
	id    string
	price float64
}

func findProduct(productId string, products []Product) Product {
	for _, product := range products {
		if product.id == productId {
			return product
		}
	}
	return Product{}
}

func findOffer(productId string, offers []Offer) Offer {
	for _, offer := range offers {
		if offer.productId == productId {
			return offer
		}
	}
	return Offer{}
}

func CalculateTotal(basket []string, offers []Offer, products []Product) float64 {
	var total = float64(0)
	var runningProductQuantity = make(map[string]int)
	runningProductQuantity["A"] = 0
	runningProductQuantity["B"] = 0
	runningProductQuantity["C"] = 0
	runningProductQuantity["D"] = 0

	for _, basketItem := range basket {
		var product = findProduct(basketItem, products)
		var offer = findOffer(basketItem, offers)

		runningProductQuantity[basketItem]++
		var value = runningProductQuantity[basketItem]

		if offer.productId == basketItem && value == offer.quantity {
			total += product.price - ((float64(offer.quantity) * product.price) - offer.price)
			runningProductQuantity[basketItem] = 0
		} else {
			total += product.price
		}
	}

	return total
}
