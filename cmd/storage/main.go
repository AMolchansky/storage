package main

import (
	"fmt"
	"log"
)
import "github.com/AMolchansky/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("Hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded: ", restoredFile)
}
