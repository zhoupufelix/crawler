package main

import (
	"regexp"
	"fmt"
)

const text = `
My email is ccmouse@gmail.com dsa
hahahahaaa 
email1 is abc@qq.com
email2 is cc@163.com
email2 is cc@163.com.cn
`

func main(){
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	//match := re.FindString(text)
	//match := re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text,-1)
	for _,m := range match{
		fmt.Println(m)
	}
}
