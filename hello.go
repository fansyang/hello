package main

import (
	"fmt"
	"github.com/fansyang/hello/morestrings"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/mozillazg/go-pinyin"
	"github.com/skip2/go-qrcode"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	id := uuid.New().String()
	fmt.Println("UUID:", id)

	hans := "中国人"

	// 默认
	a := pinyin.NewArgs()
	fmt.Println("default:", pinyin.Pinyin(hans, a)) // [[zhong] [guo] [ren]]
	// Convert 跟 Pinyin 的唯一区别就是 a 参数可以是 nil
	fmt.Println("pinyin.Convert():", pinyin.Convert(hans, nil))
	// LazyPinyin 汉字转拼音，与 `Pinyin` 的区别是： 返回值类型不同，并且不支持多音字模式，每个汉字只取第一个音
	// LazyConvert 跟 LazyPinyin 的唯一区别就是 a 参数可以是 nil
	fmt.Println(pinyin.LazyPinyin(hans, a)) // [zhong guo ren]
	// Separator 默认配置： `Slug` 中 Join 所用的分隔符，默认为"-"
	fmt.Println(pinyin.Slug(hans, a)) // zhong-guo-ren

	// Style 默认配置：风格
	a.Style = pinyin.Finals
	fmt.Println(pinyin.Pinyin(hans, a)) // [[ong] [uo] [en]]

	a.Style = pinyin.FinalsTone
	fmt.Println(pinyin.Pinyin(hans, a)) // [[ōng] [uó] [én]]

	a.Style = pinyin.FinalsTone2
	fmt.Println(pinyin.Pinyin(hans, a)) // [[o1ng] [uo2] [e2n]]

	a.Style = pinyin.FinalsTone3
	fmt.Println(pinyin.Pinyin(hans, a)) // [[ong1] [uo2] [en2]]

	a.Style = pinyin.FirstLetter
	fmt.Println(pinyin.Pinyin(hans, a)) // [[z] [g] [r]]

	a.Style = pinyin.Initials
	fmt.Println(pinyin.Pinyin(hans, a)) // [[zh] [g] [r]]

	a.Style = pinyin.Normal
	fmt.Println(pinyin.Pinyin(hans, a)) // [[zhong] [guo] [ren]]

	// 包含声调
	a.Style = pinyin.Tone
	fmt.Println(pinyin.Pinyin(hans, a)) // [[zhōng] [guó] [rén]]

	// 声调用数字表示
	a.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, a)) // [[zho1ng] [guo2] [re2n]]

	a.Style = pinyin.Tone3
	fmt.Println(pinyin.Pinyin(hans, a)) // [[zhong1] [guo2] [ren2]]

	// Heteronym 默认配置：是否启用多音字模式
	b := pinyin.NewArgs()
	b.Heteronym = true
	fmt.Println(pinyin.Pinyin(hans, b)) // [[zhong zhong] [guo] [ren]]

	b.Style = pinyin.Tone2
	fmt.Println(pinyin.Pinyin(hans, b)) // [[zho1ng zho4ng] [guo2] [re2n]]

	fmt.Println(pinyin.Version)
	fmt.Println(pinyin.Author)
	fmt.Println(pinyin.License)
	fmt.Println(pinyin.Copyright)

	// Fallback 默认配置: 如何处理没有拼音的字符(忽略这个字符)
	han := "中国人abc"
	c := pinyin.NewArgs()
	c.Fallback = func(r rune, a pinyin.Args) []string {
		return []string{string(r)}
	}
	fmt.Println(pinyin.Pinyin(han, c)) // [[zhong] [guo] [ren] [a] [b] [c]]

	// 使用go-qrcode库生成一个二维码图片
	qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, "./qrcode.png")
}
