package parser

import (
	"crawler/engine"
	"regexp"
	"crawler/model"
)

var  baseRe = regexp.MustCompile(`<div.*?class="m-btn purple".*?>([^<]+)</div>`)

func ParseProfile(contents []byte,name string)engine.ParseResult{
	profile := model.Profile{}
	match:=baseRe.FindAllSubmatch(contents,-1)

	if len(match) >= 9  {
		//profile.Age = string(match[1][1])
		//profile.Height = string(match[3][1])
		//profile.Weight = string(match[4][1])
		profile.Income    = string(match[6][1])
		profile.Marriage  = string(match[0][1])
		profile.Education = string(match[8][1])
		profile.Xinzuo    = string(match[2][1])
		profile.Hokou     = string(match[5][1])
		profile.Xinzuo    = string(match[2][1])
	}
	profile.Name = name
	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

//func extractString(contents []byte,re *regexp.Regexp)string{
//	match :=  marriageRe.FindSubmatch(contents)
//	if len(match) >= 2{
//		return string(match[1])
//	}else{
//		return ""
//	}
//}







