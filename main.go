package main

import (
	"fmt"
	"time"
)

type Point struct {
	title        string
	speed        float32
	acceleration float32
	maxSpeed     float32
	x            float32
	y            float32
}

func main() {
	c := make(chan string, 1)
	o := make(chan string, 1)

	firstPoint := Point{
		title:        "F",
		speed:        0,
		acceleration: 3,
		maxSpeed:     10,
		x:            0,
		y:            0,
	}

	secondPoint := Point{
		title:        "S",
		speed:        0,
		acceleration: 0,
		maxSpeed:     0,
		x:            0,
		y:            0,
	}

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

	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	} else if p.speed < p.maxSpeed {
		p.speed += p.acceleration
	}

	p.x += p.speed
	p.y += p.speed
}
