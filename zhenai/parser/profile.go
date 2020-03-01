package parser

import (
	"demoCrawler/engine"
	"demoCrawler/model"
	"regexp"
	"strconv"
)

var ageRe	= regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)岁</div>`)
var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}

	profile.Marriage = extractString(contents, marriageRe)

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}

	profile.Name = name

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
