package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Item struct {
	Value  int
	Pom    *PomFile
	CsProj *CsProj
}

type PomFile struct {
	OtherThing string
}

type CsProj struct {
}

func main() {
	decoder := json.NewDecoder(os.Stdin)

	for {
		var item Item
		err := decoder.Decode(&item)

		if err == nil {
			fmt.Println("=======================")
			fmt.Println(item.Value)

			if item.Pom != nil {
				fmt.Println(item.Pom.OtherThing)
			} else if item.CsProj != nil {
				// do cs proj
			}
		} else {
			log.Fatal(err)
		}
	}
}
