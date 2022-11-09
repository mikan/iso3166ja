package iso3166ja

import "testing"

func TestParseCountry(t *testing.T) {
	tests := []struct {
		Line    string
		Country Country
	}{
		{
			"|{{flagicon|Iceland}} [[アイスランド]]||lang=\"en\"|Iceland||&lt;code style=\"speak-as:spell-out\"&gt;352&lt;/code&gt;||&lt;code style=\"speak-as:spell-out\"&gt;ISL&lt;/code&gt;||&lt;code style=\"speak-as:spell-out\"&gt;IS&lt;/code&gt;||北ヨーロッパ||[[ISO 3166-2:IS]]",
			Country{
				Name:     "Iceland",
				NameJA:   "アイスランド",
				Numeric:  352,
				Alpha3:   "ISL",
				Alpha2:   "IS",
				RegionJA: "北ヨーロッパ",
				ISO31662: "ISO 3166-2:IS",
			},
		},
		{
			"|{{flagicon|British Indian Ocean Territory}} [[イギリス領インド洋地域]]||lang=\"en\"|British Indian Ocean Territory||<code style=\"speak-as:spell-out\">086</code>||<code style=\"speak-as:spell-out\">IOT</code>||<code style=\"speak-as:spell-out\">IO</code>|| style=\"white-space:nowrap;\" |インド洋地域||[[ISO 3166-2:IO]]",
			Country{
				Name:     "British Indian Ocean Territory",
				NameJA:   "イギリス領インド洋地域",
				Numeric:  86,
				Alpha3:   "IOT",
				Alpha2:   "IO",
				RegionJA: "インド洋地域",
				ISO31662: "ISO 3166-2:IO",
			},
		},
	}
	for i, test := range tests {
		actual, err := ParseCountry(test.Line)
		if err != nil {
			t.Fatalf("#%d unepected error: %v", i, err)
		}
		if actual.Name != test.Country.Name {
			t.Errorf("#%d expected name is %s, got %s", i, test.Country.Name, actual.Name)
		}
		if actual.NameJA != test.Country.NameJA {
			t.Errorf("#%d expected name_ja is %s, got %s", i, test.Country.NameJA, actual.NameJA)
		}
		if actual.Numeric != test.Country.Numeric {
			t.Errorf("#%d expected numeric is %d, got %d", i, test.Country.Numeric, actual.Numeric)
		}
		if actual.Alpha3 != test.Country.Alpha3 {
			t.Errorf("#%d expected alpha_3 is %s, got %s", i, test.Country.Alpha3, actual.Alpha3)
		}
		if actual.Alpha2 != test.Country.Alpha2 {
			t.Errorf("#%d expected alpha_2 is %s, got %s", i, test.Country.Alpha2, actual.Alpha2)
		}
		if actual.RegionJA != test.Country.RegionJA {
			t.Errorf("#%d expected region_ja is %s, got %s", i, test.Country.RegionJA, actual.RegionJA)
		}
		if actual.ISO31662 != test.Country.ISO31662 {
			t.Errorf("#%d expected iso3166_2 is %s, got %s", i, test.Country.ISO31662, actual.ISO31662)
		}
	}
}
