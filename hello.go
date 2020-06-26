package main

import (
	"fmt"

	"github.com/fansyang/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	id := uuid.New().String()
	fmt.Println("UUID:", id)
}
