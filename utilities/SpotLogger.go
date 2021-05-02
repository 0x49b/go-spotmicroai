package utilities

import (
	"log"
	"os"
)

func Logger(){
	file, err := os.OpenFile("spotmicro.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
}
