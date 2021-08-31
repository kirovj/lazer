package lazer

import (
	"strconv"
	"testing"
)

func TestLazer(t *testing.T) {
	log := Default()

	for i := 0; i < 1001; i++ {
		log.Info("test  " + strconv.Itoa(i))
	}
	select {}
}
