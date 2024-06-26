package main

import (
	"gihub.com/Yahar4/cmd/api"
	"log"
)

func main() {
	s := api.NewAPIServer(":9090", nil)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
