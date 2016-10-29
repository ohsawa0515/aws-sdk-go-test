package main

import (
	"fmt"
	"log"
)

func runEc2() {

	instances, err := ListEC2Ids()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(instances)
}

func main() {
	runEc2()
}
