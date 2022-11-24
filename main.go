package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	title        string
	corner       float32
	speed        float32
	acceleration float32
	maxSpeed     float32
	x            float32
	y            float32
}

func main() {
	chanOne := make(chan string, 1)

	chanTwo := make(chan string, 1)

	firstPoint := Point{
		title:        "Fir",
		corner:       70,
		speed:        0,
		acceleration: 6,
		maxSpeed:     287,
		x:            0,
		y:            0,
	}

	secondPoint := Point{
		title:        "Sec",
		corner:       54,
		speed:        0,
		acceleration: 4,
		maxSpeed:     199,
		x:            0,
		y:            0,
	}

	go firstPoint.MovePoint(chanOne)

	go secondPoint.MovePoint(chanTwo)

	go LordOfTime(chanOne, chanTwo)

	j := make(chan string, 1)
	<-j
}

func LordOfTime(chanOne, chanTwo chan string) {
	for {
		chanOne <- "Go"
		chanTwo <- "Go"
		time.Sleep(time.Second)
	}
}

func (p *Point) MovePoint(channel chan string) {
	for {
		<-channel
		p.IncreaseSpeed()
		p.MotionX()
		p.MotionY()
		fmt.Println(p.title, " c:", p.corner, " accel:", p.acceleration, " maxS:", p.maxSpeed, " speed:", p.speed, " X:", p.x, " Y:", p.y)
	}
}

func (p *Point) IncreaseSpeed() {
	if p.speed < p.maxSpeed {
		p.speed += p.acceleration
	}
}

func degreesToRadians(deg float32) float32 {
	return deg * (math.Pi / 180.0)
}

func (p *Point) MotionX() {
	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	}

	cos := math.Cos(float64(degreesToRadians(p.corner)))

	p.x += p.speed * float32(cos)
}

func (p *Point) MotionY() {
	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	}

	sin := math.Sin(float64(degreesToRadians(p.corner)))

	p.y += p.speed * float32(sin)
}
