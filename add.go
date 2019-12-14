package main

import (
	"fmt"
	"os"
)

func add(filename, args string) error {
	w, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer w.Close()

	_, err = fmt.Fprintf(w, args)
	fmt.Printf("Task added :%s\n", args)
	return err
}
