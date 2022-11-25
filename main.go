package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct {
	title        string
	corner       float32
	changCorner  float32
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
		corner:       90,
		changCorner:  15,
		speed:        0,
		acceleration: 5,
		maxSpeed:     43,
		x:            0,
		y:            0,
	}

	secondPoint := Point{
		title:        "Sec",
		corner:       39,
		changCorner:  3,
		speed:        0,
		acceleration: 3,
		maxSpeed:     32,
		x:            0,
		y:            0,
	}

	go firstPoint.MovePoint(chanOne)

	go secondPoint.MovePoint(chanTwo)

	go LordOfTime(chanOne, chanTwo)

	j := make(chan string, 1)
	<-j
}

func LordOfTime(channel ...chan string) {
	for {
		for i := range channel {
			channel[i] <- "Go"
		}
		time.Sleep(time.Second)
	}
}

func (p *Point) MovePoint(channel chan string) {
	for {
		<-channel
		p.IncreaseSpeed()
		p.ChangingCorner()
		p.MotionX()
		p.MotionY()
		fmt.Println(p.title, " c:", p.corner, " accel:", p.acceleration, " maxS:", p.maxSpeed, " speed:", p.speed, " X:", p.x, " Y:", p.y)
	}
}

func (p *Point) IncreaseSpeed() {
	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	}

	if p.speed < p.maxSpeed {
		p.speed += p.acceleration
	}
}

func (p *Point) ChangingCorner() {
	p.corner = p.corner + p.changCorner
}

func degreesToRadians(deg float32) float32 {
	return deg * (math.Pi / 180.0)
}

func (p *Point) MotionX() {
	cos := math.Cos(float64(degreesToRadians(p.corner))) //Перводим угол в радианы и высчитываем косинус

	p.x += p.speed * float32(cos)
}

func (p *Point) MotionY() {
	sin := math.Sin(float64(degreesToRadians(p.corner))) //Перводим угол в радианы и высчитываем синус

	p.y += p.speed * float32(sin)
}
