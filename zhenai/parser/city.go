package parser

import (
	"crawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[\d]+)" target="_blank">([^<]+)</a>`
const nextRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+/[\d]+)">下一页</a>`

func ParseCity(contents []byte)engine.ParseResult{
	//用户详情页传递
	re:= regexp.MustCompile(cityRe)
	matches:=re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}

	if matches != nil {
		for _,m := range matches{
			result.Items = append(result.Items,"User="+string(m[2]))
			name := string(m[2])
			result.Requests = append(result.Requests,engine.Request{
				Url : string(m[1]),
				ParserFunc: func(c []byte) engine.ParseResult {
					return ParseProfile(c,name)
				},
			})
		}
	}

	//下一页传递队列
	re = regexp.MustCompile(nextRe)
	match := re.FindSubmatch(contents)
	if match != nil  {
		result.Items = append(result.Items,string(match[1]))
		result.Requests = append(result.Requests,engine.Request{
			Url : string(match[1]),
			ParserFunc:ParseCity,
		})
	}


	return result
}