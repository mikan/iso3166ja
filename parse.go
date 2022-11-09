package iso3166ja

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	numberPattern = regexp.MustCompile(`[0-9]+`)
	upperPattern  = regexp.MustCompile(`[A-Z]+`)
)

// Country は国コードとその名称を提供します。
type Country struct {
	Name     string `json:"name,omitempty"`
	NameJA   string `json:"name_ja,omitempty"`
	Numeric  int    `json:"numeric,omitempty"`
	Alpha3   string `json:"alpha3,omitempty"`
	Alpha2   string `json:"alpha2,omitempty"`
	RegionJA string `json:"region_ja,omitempty"`
	ISO31662 string `json:"iso3166_2,omitempty"`
}

// Parse は記事の wiki text を解析して Country のスライスを生成します。
func Parse(wikiText string) ([]Country, error) {
	var countries []Country
	scanner := bufio.NewScanner(strings.NewReader(wikiText))
	seeking := true // 最初の "|-" を走査中
	for scanner.Scan() {
		line := scanner.Text()
		if seeking {
			if line == "|-" {
				seeking = false
			}
			continue
		}
		if line == "|-" || line == "|- class=\"sortbottom\"" || strings.HasPrefix(line, "| colspan=\"7\" ") {
			continue
		}
		if line == "|}" {
			break
		}
		c, err := ParseCountry(line)
		if err != nil {
			println(err.Error())
			continue
		}
		countries = append(countries, *c)
	}
	return countries, nil
}

// ParseCountry は1行の wiki text から Country を構築します。
func ParseCountry(wikiText string) (*Country, error) {
	var s []string
	for _, v := range strings.Split(wikiText, "|") {
		if strings.Contains(v, "style=\"white-space:nowrap;\"") {
			continue
		}
		if len(v) > 0 {
			s = append(s, v)
		}
	}
	if !strings.HasSuffix(s[1], "]]") && strings.HasSuffix(s[2], "]]") {
		s = append(s[:1], s[2:]...)
	}
	if len(s) != 9 {
		for i, v := range s {
			println(fmt.Sprintf("DEBUG: #%d %s\n", i, v))
		}
		return nil, fmt.Errorf("unexpected pipe count: %s", wikiText)
	}
	nameJA := s[1]
	if strings.Contains(nameJA, "[[") {
		nameJA = nameJA[strings.LastIndex(nameJA, "[[")+2:]
	}
	if strings.Contains(nameJA, "]]") {
		nameJA = nameJA[:strings.LastIndex(nameJA, "]]")]
	}
	n, err := strconv.Atoi(numberPattern.FindString(s[4]))
	if err != nil {
		return nil, fmt.Errorf("unexpected numeric: %s", s[4])
	}
	c := Country{
		Name:     s[3],
		NameJA:   nameJA,
		Numeric:  n,
		Alpha3:   upperPattern.FindString(s[5]),
		Alpha2:   upperPattern.FindString(s[6]),
		RegionJA: s[7],
		ISO31662: strings.ReplaceAll(strings.ReplaceAll(s[8], "[[", ""), "]]", ""),
	}
	return &c, nil
}
