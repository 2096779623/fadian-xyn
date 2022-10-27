package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var without_emoji bool
	var name string
	flag.BoolVar(&without_emoji, "without-emoji", false, "去除表情")
	flag.StringVar(&name, "name", "", "名字")
	flag.Parse()
	if len(name) == 0 {
		err := fmt.Errorf("name %q arg missing", name)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if !without_emoji == true {
		var text string = "和{name}抱在一起只是感受着{name}温热的体温和头发的香味和指尖划过皮肤的触感就已经硬到不行了🥰但是还是没有进一步的行为而是互相说了晚安就睡觉😴是每个女孩子生活中都会出现至少一次的典型女同性恋时刻但是对{name}可能会多十倍!"
		fmt.Println(strings.Replace(text, "{name}", name, -1))
		fmt.Println("source: https://twitter.com/shizuki_lena/status/1585326000686391296")
	} else if !without_emoji == false {
		var text string = "和{name}抱在一起只是感受着{name}温热的体温和头发的香味和指尖划过皮肤的触感就已经硬到不行了但是还是没有进一步的行为而是互相说了晚安就睡觉是每个女孩子生活中都会出现至少一次的典型女同性恋时刻但是对{name}可能会多十倍!"
		fmt.Println(strings.Replace(text, "{name}", name, -1))
		fmt.Println("source: https://twitter.com/shizuki_lena/status/1585326000686391296")
	}

}
