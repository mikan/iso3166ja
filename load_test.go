package iso3166ja

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	raw, err := os.ReadFile("testdata/wikitext.txt")
	if err != nil {
		t.Fatal(err)
	}
	list, err := Parse(string(raw))
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}
	if len(list) != 249 {
		t.Errorf("expected len(Load())=%d, got=%d", 249, len(list))
	}
}
