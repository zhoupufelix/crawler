package main

import (
	"regexp"
	"crawler/fetcher"
	"fmt"
	"log"
)

const text = `
My email is ccmouse@gmail.com dsa
hahahahaaa 
email1 is abc@qq.com
email2 is cc@163.com
email2 is cc@163.com.cn
`

func main(){
	contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun/shanghai")
	if err!=nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/shanghai/[\d]+)">下一页</a>`)
	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)
	//match := re.FindAllStringSubmatch(text,-1)

	//re:= regexp.MustCompile(cityRe)
	matches:=re.FindSubmatch(contents,-1)
	fmt.Print(matches)

}
