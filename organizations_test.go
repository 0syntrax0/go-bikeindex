package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	testOrganizationsCreatevar = `
	{
		"name":"Organization Name",
		"kind":"school"
	}`
)

func TestOrganizationsCreate(t *testing.T) {
	// create test server
	srv := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, testOrganizationsCreatevar)
		}))
	defer srv.Close()

	bi := NewBikeIndexClient("", "")
	bi.APIURL = srv.URL // point to the test server

	req := &Organization{
		Name:    "Organization Name",
		Website: "www.example.com",
		Kind:    "school",
	}
	st := "Illinois"
	location := OrganizationLocation{
		Name:   "Location Name",
		Street: "123 Washinton Ave",
		City:   "Chicago",
		State:  &st,
	}
	req.Locations = append(req.Locations, location)
	o, exp, err := bi.OrganizationsCreate(req)
	if err != nil {
		t.Fatalf("Error creating organization: %+v", err)
	}
	if exp != nil {
		t.Fatalf("Exception creating organization: %+v", exp)
	}
	bts, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Organization create result:\n%s\n", string(bts))
}
