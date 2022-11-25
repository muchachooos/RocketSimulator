package main

import "time"

func LordOfTime(channel ...chan struct{}) {
	for {
		for i := range channel {
			channel[i] <- struct{}{}
		}
		time.Sleep(time.Second)
	}
}
