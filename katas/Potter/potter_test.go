package Potter

import "testing"

/*
	Rules

	1. One copy of a book costs 8 EUR.
	2. 2 different books from the series = 5% discount
	3. 3 different books from the series = 10% discount
	4. 4 different books from the series = 20% discount
	5. 5 different books from the series = 25% discount

	Note: that if you buy, say, four books, of which 3 are
	different titles, you get a 10% discount on the 3 that
	form part of a set, but the fourth book still costs 8 EUR.
*/

func TestCalculatePriceWithOneBook(t *testing.T) {
	const expected = float64(8)
	var total = CalculatePrice([]int{1})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithTwoEqualBooks(t *testing.T) {
	const expected = float64(16)
	var total = CalculatePrice([]int{1, 1})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithTwoDifferentBooks(t *testing.T) {
	const expected = float64(8 * 2 * 0.95)
	var total = CalculatePrice([]int{0, 1})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithThreeDifferentBooks(t *testing.T) {
	const expected = float64(8 * 3 * 0.90)
	var total = CalculatePrice([]int{0, 1, 2})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithFourDifferentBooks(t *testing.T) {
	const expected = float64(8 * 4 * 0.8)
	var total = CalculatePrice([]int{0, 1, 2, 4})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithFiveDifferentBooks(t *testing.T) {
	const expected = float64(8 * 5 * 0.75)
	var total = CalculatePrice([]int{0, 1, 2, 3, 4})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithSingleDiscountAndExtraBook(t *testing.T) {
	const expected = float64((16 * 0.95) + 8)
	var total = CalculatePrice([]int{0, 1, 1})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithDoubleDiscount(t *testing.T) {
	const expected = float64(2 * (8 * 2 * 0.95))
	var total = CalculatePrice([]int{0, 0, 1, 1})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithDoubleDiscountAndExtraBooks(t *testing.T) {
	const expected = float64((8 * 4 * 0.8) + (8 * 2 * 0.95))
	var total = CalculatePrice([]int{0, 0, 1, 2, 2, 3})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithFiveDifferentBooksAndExtraBook(t *testing.T) {
	const expected = float64(8 + (8 * 5 * 0.75))
	var total = CalculatePrice([]int{0, 1, 1, 2, 3, 4})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithFiveDifferentBooksAndThreeExtraBooks(t *testing.T) {
	const expected = float64(2 * (8 * 4 * 0.8))
	var total = CalculatePrice([]int{0, 0, 1, 1, 2, 2, 3, 4})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}

func TestCalculateWithFiveDifferentBooksAndALotOfExtraBooks(t *testing.T) {
	const expected = float64(3*(8*5*0.75) + 2*(8*4*0.8))
	var total = CalculatePrice([]int{0, 0, 0, 0, 0,
		1, 1, 1, 1, 1,
		2, 2, 2, 2,
		3, 3, 3, 3, 3,
		4, 4, 4, 4})

	if total != expected {
		t.Errorf("TestCalculatePrice failed. Expected: %f, Actual %f", expected, total)
	}
}
