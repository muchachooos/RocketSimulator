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
	c := make(chan string, 1)

	firstPoint := Point{
		title:        "F",
		corner:       70,
		speed:        0,
		acceleration: 6,
		maxSpeed:     287,
		x:            0,
		y:            0,
	}

	go firstPoint.MovePoint(c)

	go LordOfTime(c)

	j := make(chan string, 1)
	<-j
}

func LordOfTime(cc chan string) {
	for {
		cc <- "Go"
		time.Sleep(time.Second)
	}
}

func (p *Point) MovePoint(ccc chan string) {
	for {
		<-ccc
		p.IncreaseSpeed()

		p.MotionX()
		p.MotionY()
		fmt.Println(" c:", p.corner, " accel:", p.acceleration, " maxS:", p.maxSpeed, " speed:", p.speed, " X:", p.x, " Y:", p.y)
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
