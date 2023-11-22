package main

import (
	"log"
	"os"
	"pois_debug_client/config"
	"pois_debug_client/handle"
	"strconv"
)

func main() {

	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	if len(os.Args) < 2 {
		log.Fatal("wrong calling method, insufficient parameters")
	}
	switch os.Args[1] {
	case "genfile":
		err := handle.GenPoisInitialParams()
		if err != nil {
			log.Println(err)
		}
	case "genchals":
		if len(os.Args) < 3 {
			log.Fatal("please enter the number of requests​")
		}
		num, err := strconv.Atoi(os.Args[2])
		err = handle.RequestToGetChallenge(num)
		if err != nil {
			log.Println(err)
		}
	case "commit":
		if len(os.Args) < 3 {
			log.Fatal("please enter the number of requests​")
		}
		num, err := strconv.Atoi(os.Args[2])
		err = handle.RequestToProveCommitProof(num)
		if err != nil {
			log.Println(err)
		}
	case "space":
		if len(os.Args) < 3 {
			log.Fatal("please enter the number of requests​")
		}
		num, err := strconv.Atoi(os.Args[2])
		err = handle.RequestToProveSpaceProof(num)
		if err != nil {
			log.Println(err)
		}
	case "delete":
		if len(os.Args) < 3 {
			log.Fatal("please enter the number of requests​")
		}
		num, err := strconv.Atoi(os.Args[2])
		err = handle.RequestToProveDeletionProof(num)
		if err != nil {
			log.Println(err)
		}
	}
}
