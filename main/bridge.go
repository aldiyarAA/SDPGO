package main

import "fmt"

type Mask interface {
	maskChoice()
	set(fastening)
}

type clothMask struct {
	fastening fastening
}

func (c *clothMask) maskChoice() {
	fmt.Println("clothMask choosen")
	c.fastening.setFastening()
}

func (c *clothMask) set(f fastening) {
	c.fastening = f
}

type respiratory struct {
	fastening fastening
}

func (c *respiratory) maskChoice() {
	fmt.Println("ffp choosen")
	c.fastening.setFastening()
}

func (c *respiratory) set(f fastening) {
	c.fastening = f
}

type fastening interface {
	setFastening()
}

type onEars struct {
}

func (o *onEars) setFastening() {
	fmt.Println("onEars fastening method choosen")
}

type onHead struct {
}

func (o *onHead) setFastening() {
	fmt.Println("onHead fastening method choosen")
}

func choiceForMedMask() {
	var n int
	medecine := &clothMask{}
	fmt.Print("Choose wearing type:\n" +
		"1. On ears\n" +
		"2. on head\n")
	fmt.Scan(&n)
	if n == 1 {
		ears := &onEars{}
		medecine.set(ears)
		medecine.maskChoice()
		sum += 7
		decor()
	}
	if n == 2 {
		head := &onHead{}
		medecine.set(head)
		medecine.maskChoice()
		sum += 9
		decor()
	}
}

func choiceForRespMask() {
	var n int
	resperator := &respiratory{}
	fmt.Print("Choose wearing type:\n" +
		"1. On ears\n" +
		"2. on head\n")
	fmt.Scan(&n)
	if n == 1 {
		ears := &onEars{}
		resperator.set(ears)
		resperator.maskChoice()
		sum += 12
		decor()
	}
	if n == 2 {
		head := &onHead{}
		resperator.set(head)
		resperator.maskChoice()
		sum += 14
		decor()
	}
}

func bridge() {

	var input int

	fmt.Print("\nIt is time to choose main Protection \n" +
		"\nWhat type of mask you want to use " +
		"\n\n" +
		"1.Medicine Mask \n" +
		"2.Resperatory Mask\n")
	fmt.Scan(&input)

	if input == 1 {

		choiceForMedMask()

	}
	if input == 2 {

		choiceForRespMask()
	}

}
