package main

import (
	"fmt"
	"time"
)

type Point struct {
	speed int
	x     int
	y     int
}

func main() {
	c := make(chan string, 1)
	o := make(chan string, 1)

	firstPoint := Point{speed: 3, x: 1, y: 0}
	secondPoint := Point{speed: 7, x: 10, y: 0}

	go firstPoint.MovePoint(c)
	go secondPoint.MovePoint(o)

	go LordOfTime(c, o)

	j := make(chan string, 1)
	<-j
}

func LordOfTime(cc, oo chan string) {
	for {
		cc <- "Go"
		oo <- "Go"
		time.Sleep(time.Second)
	}
}

func (p *Point) MovePoint(ccc chan string) {
	for {
		<-ccc
		p.MotionParameters()
		fmt.Println(p)
	}
}

func (p *Point) MotionParameters() {
	p.x += p.speed
	p.y += p.speed
}
