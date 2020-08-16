package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearch(t *testing.T) {
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
	search, err := bi.Search(*req)
	if err != nil {
		t.Fatalf("Failed to search bikes: %+v", err)
	}
	bts, err := json.MarshalIndent(search, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Search Result:\n%s\n", string(bts))
}

const (
	testSearchBikes = `
	{
		"bikes": [
			{
			"date_stolen": 1595368800,
			"description": "",
			"frame_colors": [
				"Silver, gray or bare metal",
				"Black"
			],
			"frame_model": "Pitch",
			"id": 796170,
			"is_stock_img": false,
			"large_img": null,
			"location_found": null,
			"manufacturer_name": "Specialized",
			"external_id": null,
			"registry_name": null,
			"registry_url": null,
			"serial": "WSBC920550413P",
			"status": null,
			"stolen": true,
			"stolen_location": "Geneva, IL - US",
			"thumb": null,
			"title": "2019 Specialized Pitch",
			"url": "https://bikeindex.org/bikes/796170",
			"year": 2019
			},
			{
			"date_stolen": 1595356703,
			"description": null,
			"frame_colors": [
				"Black"
			],
			"frame_model": "Synapse",
			"id": 793914,
			"is_stock_img": false,
			"large_img": null,
			"location_found": null,
			"manufacturer_name": "Cannondale",
			"external_id": null,
			"registry_name": null,
			"registry_url": null,
			"serial": "MC10654",
			"status": null,
			"stolen": true,
			"stolen_location": "Aurora, IL - US",
			"thumb": null,
			"title": "2018 Cannondale Synapse",
			"url": "https://bikeindex.org/bikes/793914",
			"year": 2018
			},
			{
			"date_stolen": 1593309600,
			"description": null,
			"frame_colors": [
				"Purple"
			],
			"frame_model": "Comp ",
			"id": 777025,
			"is_stock_img": false,
			"large_img": "https://files.bikeindex.org/uploads/Pu/273390/large_Screenshot_20200627-162807_Gallery.jpg",
			"location_found": null,
			"manufacturer_name": "Specialized",
			"external_id": null,
			"registry_name": null,
			"registry_url": null,
			"serial": "WSBC614014175N",
			"status": null,
			"stolen": true,
			"stolen_location": "Aurora, IN - US",
			"thumb": "https://files.bikeindex.org/uploads/Pu/273390/small_Screenshot_20200627-162807_Gallery.jpg",
			"title": "2018 Specialized Comp",
			"url": "https://bikeindex.org/bikes/777025",
			"year": 2018
			}
		]
	}`
)
