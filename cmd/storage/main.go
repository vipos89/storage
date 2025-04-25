package main

import (
	"fmt"
	"log"

	"github.com/vipos89/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println(st)
	file, err :=st.Upload("test.txt", []byte("test hello world"))
	if err != nil {
		log.Fatal(err)
	}
	restoredFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(restoredFile)
}