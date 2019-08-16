package news

import (
	"testing"
)

const key = "6bbf7d4138154b0499e1087b3b8867b2"

func TestHeadlines (t *testing.T) {
	n := NewNewsClient(key)

	_, err := n.GetHeadlines("de")
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestQuery (t *testing.T) {
	n := NewNewsClient(key)
	_, err := n.RunQuery(
		Query{
			Sources: "spiegel-online",
			Language: "de",
		},
	)
	if err != nil {
		t.Fatalf("%v", err)
	}
}