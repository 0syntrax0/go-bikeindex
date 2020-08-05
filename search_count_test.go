package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchCount(t *testing.T) {
	// create test server
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testSearchCountBikes)
		}))
	defer srv.Close()

	bi := NewBikeIndexClient("", "")
	bi.APIURL = srv.URL // point to the test server
	req := &SearchCountReq{
		Manufacturer: "Santa Cruz",
	}
	search, err := bi.SearchCount(*req)
	if err != nil {
		t.Fatalf("Failed to search bikes: %+v", err)
	}
	bts, err := json.MarshalIndent(search, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Search Count Result:\n%s\n", string(bts))
}

const (
	testSearchCountBikes = `
	{
		"proximity": 19,
		"stolen": 100,
		"non": 111
	}`
)
