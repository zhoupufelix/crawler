package parser

import (
	"crawler/engine"
	"regexp"
	"strconv"
	"crawler/model"
)

var  ageRe = regexp.MustCompile(`<td width="180"><span class="grayL">年龄：</span>([\d]+)</td>`)
var marriageRe = regexp.MustCompile(`<td width="180"><span class="grayL">婚况：</span>([^<]+)</td>`)

func ParseProfile(contents []byte)engine.ParseResult{
	profile := model.Profile{}
	match:=ageRe.FindSubmatch(contents)
	if match != nil{
		age,err:= strconv.Atoi(extractString(contents,ageRe))
		if err != nil {
			profile.Age = age
		}
	}
	profile.Marriage = extractString(contents,marriageRe)
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp)string{
	match :=  marriageRe.FindSubmatch(contents)
	if len(match) >= 2{
		return string(match[1])
	}else{
		return ""
	}
}







