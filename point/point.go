package point

import (
	"fmt"
	"math"
)

type point struct {
	title        string
	corner       float32
	changCorner  float32
	speed        float32
	acceleration float32
	maxSpeed     float32
	x            float32
	y            float32
}

func NewPoint(title string, corner, changCorner, speed, acceleration, maxSpeed, x, y float32) point {
	return point{
		title:        title,
		corner:       corner,
		changCorner:  changCorner,
		speed:        speed,
		acceleration: acceleration,
		maxSpeed:     maxSpeed,
		x:            x,
		y:            y,
	}
}

func (p *point) MovePoint(channel chan struct{}) {
	for {
		<-channel
		p.increaseSpeed()
		p.changingCorner()
		p.motionX()
		p.motionY()
		fmt.Println(p.title, " c:", p.corner, " accel:", p.acceleration, " maxS:", p.maxSpeed, " speed:", p.speed, " X:", p.x, " Y:", p.y)
	}
}

func (p *point) increaseSpeed() {
	if p.maxSpeed-p.speed < p.acceleration {
		p.speed = p.maxSpeed
	}

	if p.speed < p.maxSpeed {
		p.speed += p.acceleration
	}
}

func (p *point) changingCorner() {
	p.corner = p.corner + p.changCorner
}

func (p *point) motionX() {
	cos := math.Cos(float64(degreesToRadians(p.corner))) //Перводим угол в радианы и высчитываем косинус

	p.x += p.speed * float32(cos)
}

func (p *point) motionY() {
	sin := math.Sin(float64(degreesToRadians(p.corner))) //Перводим угол в радианы и высчитываем синус

	p.y += p.speed * float32(sin)
}
