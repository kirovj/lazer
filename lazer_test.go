package lazer

import (
	"strconv"
	"testing"
	"time"
)

func TestLazer(t *testing.T) {

	log := Default()

	for i := 0; i < 100; i++ {
		log.Info("test  " + strconv.Itoa(i))
	}

	time.Sleep(10 * time.Second)

	for i := 0; i < 100; i++ {
		log.Error("error  " + strconv.Itoa(i))
	}

	select {}
}
