package main

import (
	"github.com/thanhpk/baseconv"
)

func X() {
	x, _ := baseconv.DecodeHex("70B1D707EAC2EDF4C6389F440C7294B51FFF57B")
	println(x)
}
