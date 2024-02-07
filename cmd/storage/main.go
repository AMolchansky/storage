package main

import "fmt"
import "github.com/AMolchansky/storage/internal/storage"

func main() {
	st := storage.NewStorage()

	fmt.Println("it works", st)
}
