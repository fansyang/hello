package main

import (
	"fmt"

	"github.com/fansyang/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	pinyin "github.com/mozillazg/go-pinyin"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	id := uuid.New().String()
	fmt.Println("UUID:", id)

	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println(pinyin.Pinyin(hans, a))

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a))

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a))
}
