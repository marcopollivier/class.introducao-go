package main

import (
	"fmt"
	"math"
)

type geometria interface {
	area() float64
	perim() float64
}

type quadrado struct {
	largura, altura float64
}

type circulo struct {
	raio float64
}

func (q quadrado) area() float64 {
	return q.largura * q.altura
}
func (q quadrado) perim() float64 {
	return 2*q.largura + 2*q.altura
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perim() float64 {
	return 2 * math.Pi * c.raio
}

func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	q := quadrado{largura: 3, altura: 4}
	c := circulo{raio: 5}

	medir(q)
	medir(c)
}
