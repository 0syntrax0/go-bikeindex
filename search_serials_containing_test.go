package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchSerialsContaining(t *testing.T) {
	// create test server
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testSearchBikes)
		}))
	defer srv.Close()

	bi := NewBikeIndexClient("", "")
	bi.APIURL = srv.URL // point to the test server
	req := &SearchReq{
		Manufacturer: "Santa Cruz",
	}
	search, err := bi.SearchSerialsContaining(*req)
	if err != nil {
		t.Fatalf("Failed to Search Serials Containing bikes: %+v", err)
	}
	bts, err := json.MarshalIndent(search, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Search Serials Containing Result:\n%s\n", string(bts))
}
