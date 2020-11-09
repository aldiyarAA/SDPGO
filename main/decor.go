package main

import "fmt"

func decor() {

	var input int

	fmt.Print("Now you have the opportunity to brew an elixir for dehydration for more protection\n" +
		"\nWhat ingredients will you use " +
		"\n\n" +
		"1.alcohol \n" +
		"2.penicillin\n" +
		"3.peroxide hydrogen\n")
	fmt.Scan(&input)

	if input == 1 {
		case1()
	}
	if input == 2 {
		case2()
	}
	if input == 3 {
		case3()
	}

}
func case1() {
	elixir := &water{}
	var what string
	what = "alcohol"
	elixirWith := &alcohol{
		elixir: elixir,
	}
	fmt.Printf("elixir with %s", what)
	fmt.Printf(" has  %d points", elixirWith.getPoint())
	sum += elixirWith.getPoint()
	builder()

}
func case2() {
	elixir := &water{}
	var what string
	what = "penicillin"
	elixirWith := &penicillin{
		elixir: elixir,
	}
	fmt.Printf("elixir with %s", what)
	fmt.Printf(" has %d points", elixirWith.getPoint())
	sum += elixirWith.getPoint()
	builder()
}
func case3() {
	elixir := &water{}
	var what string
	what = "peroxide hydrogen"
	elixirWith := &peroxideHydrogen{
		elixir: elixir,
	}
	fmt.Printf("elixir with %s", what)
	fmt.Printf(" has %d points", elixirWith.getPoint())
	sum += elixirWith.getPoint()
	builder()
}

type elixir interface {
	getPoint() int
}

type water struct {
}

func (w *water) getPoint() int {
	return 5
}

type peroxideHydrogen struct {
	elixir elixir
}

func (c *peroxideHydrogen) getPoint() int {
	elixirPoint := c.elixir.getPoint()
	return elixirPoint + 10
}

type alcohol struct {
	elixir elixir
}

func (a *alcohol) getPoint() int {
	elixirPoint := a.elixir.getPoint()
	return elixirPoint + 20
}

type penicillin struct {
	elixir elixir
}

func (p *penicillin) getPoint() int {
	elixirPoint := p.elixir.getPoint()
	return elixirPoint + 15
}

type allIngredient struct {
	elixir elixir
}

func (a *allIngredient) getPoint() int {
	elixirPoint := a.elixir.getPoint()
	return elixirPoint + 30
}
