package selection

import (
	"errors"
	"net/http"
)

var ErrorTookWayTooLong = errors.New("took too long")

func Racer(first, second string) string {
	select {
	case <-ping(first):
		return first
	case <-ping(second):
		return second
	}
}

func ping(url string) chan struct{} {
	raceChannel := make(chan struct{})
	go func() {
		http.Get(url)
		close(raceChannel)
	}()
	return raceChannel
}
