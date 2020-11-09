package main

import "fmt"

func builder() {
	var input int
	fmt.Print("Finally, you can build your own breathing apparatus and use it to save lives\n" +
		"You can choose from two Breathe ventilation devices\n" +
		"1.medical \n" +
		"2.handmade\n")
	fmt.Scan(&input)
	if input == 1 {
		medical()
	}
	if input == 2 {
		handmade()
	}

}

type pointchecker interface {
	getpoint() int
}

func medical() {
	normalBuilder := getBuilder("normal")
	director := newDirector(normalBuilder)
	normalBreathe := director.buildBreathe()
	fmt.Printf("Medical Breathe hose Type: %s\n", normalBreathe.hosetype)
	fmt.Printf("Medical Breathe balloon Type: %s\n", normalBreathe.balloontype)
	fmt.Printf("Medical Breathe switch Type: %s\n", normalBreathe.switchtype)
	sum += 20
	result()
}
func handmade() {
	handmadeBuilder := getBuilder("handmade")
	director := newDirector(handmadeBuilder)
	handmadeBreathe := director.buildBreathe()

	fmt.Printf("\nHandmade Breathe hose Type: %s\n", handmadeBreathe.hosetype)
	fmt.Printf("Handmade Breathe balloon Type: %s\n", handmadeBreathe.balloontype)
	fmt.Printf("Handmade Breathe switch Type: %s\n", handmadeBreathe.switchtype)
	sum += 10
	result()
}

type Breathe struct {
	balloontype string
	hosetype    string
	switchtype  string
}

type iBuilder interface {
	setBalloontype()
	setHosetype()
	setSwitchtype()
	getBreathe() Breathe
}

func getBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}

	if builderType == "handmade" {
		return &handmadeBuilder{}
	}
	return nil
}

type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b iBuilder) {
	d.builder = b
}

func (d *director) buildBreathe() Breathe {
	d.builder.setHosetype()
	d.builder.setBalloontype()
	d.builder.setSwitchtype()
	return d.builder.getBreathe()
}

type handmadeBuilder struct {
	balloontype string
	hosetype    string
	switchtype  string
}

func newHandmadeBuilder() *handmadeBuilder {
	return &handmadeBuilder{}
}

func (b *handmadeBuilder) setBalloontype() {
	b.balloontype = "Household"
}

func (b *handmadeBuilder) setHosetype() {
	b.hosetype = "\nplastic"
}

func (b *handmadeBuilder) setSwitchtype() {
	b.switchtype = "blue"
}

func (b *handmadeBuilder) getBreathe() Breathe {
	return Breathe{
		hosetype:    b.hosetype,
		balloontype: b.balloontype,
		switchtype:  b.switchtype,
	}
}

type normalBuilder struct {
	balloontype string
	hosetype    string
	switchtype  string
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setBalloontype() {
	b.balloontype = "Medical"
}

func (b *normalBuilder) setHosetype() {
	b.hosetype = "\nrubber"
}

func (b *normalBuilder) setSwitchtype() {
	b.switchtype = "red"
}

func (b *normalBuilder) getBreathe() Breathe {
	return Breathe{
		hosetype:    b.hosetype,
		balloontype: b.balloontype,
		switchtype:  b.switchtype,
	}
}
