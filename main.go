package main

import (
	"fmt"
	"time"
)

type Point struct {
	x int
	y int
}

func main() {
	c := make(chan string, 1)

	j := make(chan string, 1)

	firstPoint := Point{x: 0, y: 0}

	go firstPoint.movePoint(c)

	go LordOfTime(c)

	<-j
}

func LordOfTime(cc chan string) {
	for {
		cc <- "writeEbala"
		time.Sleep(time.Second)
	}
}

func (p *Point) movePoint(ccc chan string) {
	for {
		<-ccc
		p.x++
		p.y++
		fmt.Println(p)
	}
}
