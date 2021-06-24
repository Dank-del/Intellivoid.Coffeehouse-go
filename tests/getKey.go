package tests

import (
	"io/ioutil"
	"log"
)

func returnKey() string {
	content, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
