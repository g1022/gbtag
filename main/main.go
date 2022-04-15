package main

import (
	"github.com/g1022/gbtag"
	"log"
)

func main() {
	log.Println(gbtag.ToString())
	log.Println(gbtag.ToJson())
	log.Println("hello gbtag")
}
