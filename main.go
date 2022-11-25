package main

import "RocketSimulator/point"

func main() {
	chanOne := make(chan struct{}, 1)

	chanTwo := make(chan struct{}, 1)

	firstPoint := point.NewPoint("First", 70, 5, 0, 7, 220, 0, 0)

	secondPoint := point.NewPoint("Sec", 39, 3, 0, 4, 181, 0, 0)

	go firstPoint.MovePoint(chanOne)

	go secondPoint.MovePoint(chanTwo)

	go LordOfTime(chanOne, chanTwo)

	j := make(chan string, 1)
	<-j
}
