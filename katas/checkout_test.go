package katas

import "testing"

func getOffers() []Offer {
	return []Offer{
		{productId: "A", quantity: 3, price: 130},
		{productId: "B", quantity: 2, price: 45},
	}
}

func getProducts() []Product {
	return []Product{
		{id: "A", price: 50},
		{id: "B", price: 30},
		{id: "C", price: 20},
		{id: "D", price: 10},
	}
}

func TestCalculateTotalWithSingleA(t *testing.T) {
	var offers = getOffers()
	var products = getProducts()
	var basket = []string{"A"}
	const expected = float64(50)
	var total = CalculateTotal(basket, offers, products)

	if expected != total {
		t.Errorf("CalculateTotal failed, expected %f actual %f", expected, total)
	}
}

func TestCalculateTotalWithOfferA(t *testing.T) {
	var offers = getOffers()
	var products = getProducts()
	var basket = []string{"A", "A", "A"}
	const expected = float64(130)
	var total = CalculateTotal(basket, offers, products)

	if expected != total {
		t.Errorf("CalculateTotal failed, expected %f actual %f", expected, total)
	}
}

func TestCalculateTotalWithOfferAAndB(t *testing.T) {
	var offers = getOffers()
	var products = getProducts()
	var basket = []string{"A", "A", "A", "B"}
	const expected = float64(160)
	var total = CalculateTotal(basket, offers, products)

	if expected != total {
		t.Errorf("CalculateTotal failed, expected %f actual %f", expected, total)
	}
}

func TestCalculateTotalWithOfferAAndOfferB(t *testing.T) {
	var offers = getOffers()
	var products = getProducts()
	var basket = []string{"A", "A", "A", "B", "B", "B"}
	const expected = float64(205)
	var total = CalculateTotal(basket, offers, products)

	if expected != total {
		t.Errorf("CalculateTotal failed, expected %f actual %f", expected, total)
	}
}
