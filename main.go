package main

import (
	"log"
	"pois_debug_client/client"
)

func main() {
	minerId := "12D3KooWPBNzvjQtLEaDHBLACgZYBW9uHQyUtSvNoFnBm8GVPAJ4"
	key, err := client.GetMinerPoisKey(minerId)
	if err != nil {
		log.Println(err)
	}
	log.Println("get pois key success", key)
}
