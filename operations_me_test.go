package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOperationsMe(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testOperationsMe)
		}))
	defer srv.Close()

	bi := NewBikeIndexClient("", "")
	bi.APIURL = srv.URL // point to the test server
	opMe, err := bi.OperationsMe()
	if err != nil {
		t.Fatalf("failed to fetch operations me: %v \n %+v", err, opMe)
	}
	bts, err := json.MarshalIndent(opMe, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal indent operations me: %+v", err)
	}
	t.Logf("Operations Me Result: \n%s\n", string(bts))
}

const (
	testOperationsMe = `
	{
		"id": 42069
	}`
)
