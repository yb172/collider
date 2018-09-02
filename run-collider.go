package main

import (
	"fmt"
	"log"

	"github.com/yb172/collider/collider"
	"github.com/yb172/collider/loader"
)

func main() {
	materials, err := loader.LoadFromYaml("./materials.yaml")
	if err != nil {
		log.Fatalf("Error while importing materials: %v", err)
	}
	results, err := collider.Collide(materials)
	if err != nil {
		log.Fatalf("%v. Please see your boss", err)
	}
	fmt.Println(results)
}
