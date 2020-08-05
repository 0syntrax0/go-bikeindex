package gobikeindex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// BikeDeleteImage deletes an image from the bike
func (bi *BikeIndex) BikeDeleteImage(bikeID, imageID int64) (*Exception, error) {
	resp, err := bi.delete(bi.buildURL("bikes", strconv.FormatInt(bikeID, 10), "images", strconv.FormatInt(imageID, 10)))
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		exc := new(Exception)
		err = json.Unmarshal(respBody, exc)
		return exc, err
	}
	return nil, nil
}
