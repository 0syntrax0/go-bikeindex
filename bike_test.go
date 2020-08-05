package gobikeindex

import (
	"encoding/json"
	"testing"
)

func TestBike(t *testing.T) {
	srv := createTestServer(testBike)
	bi := NewBikeIndexClient("", "")
	bi.APIURL = srv.URL // point to the test server
	bike, err := bi.BikeByID(69)
	if err != nil {
		t.Fatalf("Failed to fetch bike by id: %+v", err)
	}
	bts, err := json.MarshalIndent(bike, "", "  ")
	if err != nil {
		t.Fatalf("Failed: %s", err.Error())
	}
	t.Logf("Bike Result:\n%s\n", string(bts))
}

const (
	testBike = `
	{
		"bike": {
		  "date_stolen": null,
		  "description": "",
		  "frame_colors": [
			"White"
		  ],
		  "frame_model": "Pinscher",
		  "id": 2,
		  "is_stock_img": false,
		  "large_img": "https://files.bikeindex.org/uploads/Pu/105/large_2013-05-29_18.02.21.jpg",
		  "location_found": null,
		  "manufacturer_name": "Doberman",
		  "external_id": null,
		  "registry_name": null,
		  "registry_url": null,
		  "serial": "psr88190907",
		  "status": null,
		  "stolen": false,
		  "stolen_location": null,
		  "thumb": "https://files.bikeindex.org/uploads/Pu/105/small_2013-05-29_18.02.21.jpg",
		  "title": "2010 Doberman Pinscher",
		  "url": "https://bikeindex.org/bikes/2",
		  "year": 2010,
		  "registration_created_at": 1358542988,
		  "registration_updated_at": 1580486891,
		  "api_url": "https://bikeindex.org/api/v1/bikes/2",
		  "manufacturer_id": 1,
		  "paint_description": null,
		  "name": "Pinsch",
		  "frame_size": "m",
		  "rear_tire_narrow": false,
		  "front_tire_narrow": false,
		  "type_of_cycle": "Bike",
		  "test_bike": false,
		  "rear_wheel_size_iso_bsd": 559,
		  "front_wheel_size_iso_bsd": 559,
		  "handlebar_type_slug": null,
		  "frame_material_slug": "steel",
		  "front_gear_type_slug": "1",
		  "rear_gear_type_slug": "1",
		  "additional_registration": "",
		  "stolen_record": null,
		  "public_images": [
			{
			  "name": "Doberman Pinscher bicycle White",
			  "full": "https://files.bikeindex.org/uploads/Pu/105/2013-05-29_18.02.21.jpg",
			  "large": "https://files.bikeindex.org/uploads/Pu/105/large_2013-05-29_18.02.21.jpg",
			  "medium": "https://files.bikeindex.org/uploads/Pu/105/medium_2013-05-29_18.02.21.jpg",
			  "thumb": "https://files.bikeindex.org/uploads/Pu/105/small_2013-05-29_18.02.21.jpg",
			  "id": 105
			},
			{
			  "name": "Doberman Pinscher bicycle White",
			  "full": "https://files.bikeindex.org/uploads/Pu/30/the_pinsch.JPG",
			  "large": "https://files.bikeindex.org/uploads/Pu/30/large_the_pinsch.JPG",
			  "medium": "https://files.bikeindex.org/uploads/Pu/30/medium_the_pinsch.JPG",
			  "thumb": "https://files.bikeindex.org/uploads/Pu/30/small_the_pinsch.JPG",
			  "id": 30
			},
			{
			  "name": "Doberman Pinscher bicycle White",
			  "full": "https://files.bikeindex.org/uploads/Pu/49/2013-03-06_07.47.44.jpg",
			  "large": "https://files.bikeindex.org/uploads/Pu/49/large_2013-03-06_07.47.44.jpg",
			  "medium": "https://files.bikeindex.org/uploads/Pu/49/medium_2013-03-06_07.47.44.jpg",
			  "thumb": "https://files.bikeindex.org/uploads/Pu/49/small_2013-03-06_07.47.44.jpg",
			  "id": 49
			},
			{
			  "name": "Doberman Pinscher bicycle White",
			  "full": "https://files.bikeindex.org/uploads/Pu/29/car_bikes.JPG",
			  "large": "https://files.bikeindex.org/uploads/Pu/29/large_car_bikes.JPG",
			  "medium": "https://files.bikeindex.org/uploads/Pu/29/medium_car_bikes.JPG",
			  "thumb": "https://files.bikeindex.org/uploads/Pu/29/small_car_bikes.JPG",
			  "id": 29
			},
			{
			  "name": "Doberman Pinscher bicycle White",
			  "full": "https://files.bikeindex.org/uploads/Pu/31/pinsch_just_frame.JPG",
			  "large": "https://files.bikeindex.org/uploads/Pu/31/large_pinsch_just_frame.JPG",
			  "medium": "https://files.bikeindex.org/uploads/Pu/31/medium_pinsch_just_frame.JPG",
			  "thumb": "https://files.bikeindex.org/uploads/Pu/31/small_pinsch_just_frame.JPG",
			  "id": 31
			}
		  ],
		  "components": []
	`
)
