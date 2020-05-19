package main

type stock struct {
	price  int
	weight int
}

type StockSpanner struct {
	stack []stock
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (spanner *StockSpanner) Next(price int) int {
	weight := 1
	for i := len(spanner.stack) - 1; i >= 0 && spanner.stack[i].price <= price; i-- {
		weight += spanner.stack[i].weight
		spanner.stack = spanner.stack[:i]
	}
	spanner.stack = append(spanner.stack, stock{price: price, weight: weight})
	return weight
}

type StockSpannerSlow struct {
	prices []int
	days   int
	stack  []int
	length int
}

func ConstructorSlow() StockSpannerSlow {
	return StockSpannerSlow{
		prices: make([]int, 1),
		days:   0,
		stack:  make([]int, 2),
		length: 0,
	}
}

func (spanner *StockSpannerSlow) Next(price int) int {
	spanner.days++
	spanner.prices = append(spanner.prices, price)

	if spanner.length == 0 {
		spanner.length++
		spanner.stack[1] = 1
		return 1
	}

	for spanner.length > 0 && spanner.prices[spanner.stack[spanner.length]] <= price {
		spanner.length--
	}
	spanner.length++
	if len(spanner.stack) > spanner.length {
		spanner.stack[spanner.length] = spanner.days
	} else {
		spanner.stack = append(spanner.stack, spanner.days)
	}

	return spanner.stack[spanner.length] - spanner.stack[spanner.length-1]
}

type StockSpannerSlow2 struct {
	prices []int
	stack  []int
}

func ConstructorSlow2() StockSpannerSlow2 {
	return StockSpannerSlow2{
		prices: make([]int, 0),
		stack:  make([]int, 0),
	}
}

func (spanner *StockSpannerSlow2) Next(price int) int {
	spanner.prices = append(spanner.prices, price)
	spanner.stack = append(spanner.stack, len(spanner.prices))

	if len(spanner.prices) == 1 {

		return 1
	}

	length := len(spanner.stack)
	i := length - 1
	for i >= 0 && spanner.prices[spanner.stack[i]-1] <= price {
		i--
	}
	spanner.stack[i+1] = spanner.stack[length-1]
	spanner.stack = spanner.stack[:i+2]
	length = i + 2

	if length > 1 {
		return spanner.stack[length-1] - spanner.stack[length-2]
	}

	return spanner.stack[0]
}
