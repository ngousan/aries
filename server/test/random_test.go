package test

import (
	"aoBlog/utils"
	"fmt"
	"testing"
)

func TestGetRandom(t *testing.T) {
	s, _ := utils.CreateRandomCode(4)
	fmt.Println("random: ", s)
}
