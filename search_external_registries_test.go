package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchExternalRegistries(t *testing.T) {
	// create test server
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testSearchBikes)
		}))
	defer srv.Close()

	bi := NewBikeIndexClient("", "")
	bi.APIURL = srv.URL // point to the test server
	req := &SearchReq{
		Serial: "p1cc",
	}
	search, err := bi.SearchExternalRegistries(*req)
	if err != nil {
		t.Fatalf("Failed to Search External Registries bikes: %+v", err)
	}
	bts, err := json.MarshalIndent(search, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Search External Registries Result:\n%s\n", string(bts))
}
