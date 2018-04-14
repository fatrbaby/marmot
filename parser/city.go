package parser

import (
	"github.com/fatrbaby/cobweb/engine"
	"regexp"
)

type City struct {
	Name []byte
	Link []byte
}

const (
	CityListPattern = `<a href="(http://www.zhenai.com/zhenghun/[\w]+)"[^>]*>([^<]+)</a>`
	CityPattern     = `href="(http://www.zhenai.com/zhenghun/[^"]+)"`
	PersonPattern   = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
)

var (
	CityListMatcher = regexp.MustCompile(CityListPattern)
	CityMather      = regexp.MustCompile(CityPattern)
	PersonMatcher   = regexp.MustCompile(PersonPattern)
)

func CityParser(contents []byte, _ string) engine.ParsedResult {
	// find cities
	var results = engine.ParsedResult{}

	matches := CityMather.FindAllSubmatch(contents, -1)

	for _, match := range matches {
		results.Spiders = append(results.Spiders, engine.Spider{Url: string(match[1]), Parser: CityParser})
	}

	// find person on city page
	matches = PersonMatcher.FindAllSubmatch(contents, -1)

	for _, match := range matches {
		results.Spiders = append(results.Spiders, engine.Spider{
			Url: string(match[1]),
			Parser: ProfileParserBridge(string(match[2])),
		})
	}

	return results
}

func CityListParser(contents []byte, _ string) engine.ParsedResult {
	matches := CityListMatcher.FindAllSubmatch(contents, -1)
	var results = engine.ParsedResult{}

	for _, match := range matches {
		results.Spiders = append(results.Spiders, engine.Spider{Url: string(match[1]), Parser: CityParser})
	}

	return results
}
