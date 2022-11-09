package iso3166ja

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strconv"
)

// DefaultCSVColumns はデフォルトの CSV 出力列を定義します。
var DefaultCSVColumns = []string{"name", "name_ja", "numeric", "alpha2", "alpha3", "region_ja", "iso3166_2"}

// FormatCSV は Country のスライスを CSV に変換します。
func FormatCSV(c []Country, columns ...string) ([]byte, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(columns); err != nil {
		return nil, fmt.Errorf("failed to write csv header: %w", err)
	}
	for i, v := range c {
		var line []string
		for _, column := range columns {
			switch column {
			case "name":
				line = append(line, v.Name)
			case "name_ja":
				line = append(line, v.NameJA)
			case "numeric":
				line = append(line, strconv.Itoa(v.Numeric))
			case "alpha2":
				line = append(line, v.Alpha2)
			case "alpha3":
				line = append(line, v.Alpha3)
			case "region_ja":
				line = append(line, v.RegionJA)
			case "iso3166_2":
				line = append(line, v.ISO31662)
			}
		}
		if err := w.Write(line); err != nil {
			return nil, fmt.Errorf("failed to write csv line %d %v: %w", i+1, line, err)
		}
	}
	w.Flush()
	return buf.Bytes(), nil
}

// FormatArrayJSON は Country のスライスを配列の JSON に変換します。
func FormatArrayJSON(c []Country) ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

// FormatMapJSON は Country のスライスをキーが Alpha-2 のマップの JSON に変換します。
func FormatMapJSON(c []Country) ([]byte, error) {
	m := make(map[string]Country)
	for _, v := range c {
		country := v
		country.Alpha2 = ""
		m[v.Alpha2] = country
	}
	return json.MarshalIndent(m, "", "  ")
}
