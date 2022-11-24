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

func radiansToDegrees(rad float32) float32 {
	return rad * (180.0 / math.Pi)
}

func (p *Point) MotionX() {
	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	}

	corInRad := degreesToRadians(p.corner)
	fmt.Println("corInRad: ", corInRad)
	cos := math.Cos(float64(corInRad))
	fmt.Println("cos: ", cos)

	p.x += p.speed * float32(cos)
}

func (p *Point) MotionY() {
	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	}

	corInRad := degreesToRadians(p.corner)
	fmt.Println("corInRad: ", corInRad)
	sin := math.Sin(float64(corInRad))
	fmt.Println("sin: ", sin)

	p.y += p.speed * float32(sin)
}

//sinInRad := math.Sin(float64(p.corner))
//fmt.Println("sinInRad: ", sinInRad)
//sinInDeg := radiansToDegrees(float32(sinInRad))
//fmt.Println("sinInDeg: ", sinInDeg)
//sinInDeg := radiansToDegrees(float32(math.Sin(float64(p.corner))))

//p.y += p.speed * sinInDeg

//func (p *Point) MotionParameters() {
//	if p.maxSpeed-p.speed < p.acceleration {
//		p.speed = p.maxSpeed
//	} else if p.speed < p.maxSpeed {
//		p.speed += p.acceleration
//	}
//
//	p.x += p.speed
//	p.y += p.speed
//}
