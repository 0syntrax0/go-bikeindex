package gobikeindex

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/go-querystring/query"
)

// SendStolenNotification -
type SendStolenNotification struct {
	Message string `url:"Message" json:"message"`
}

// SendStolenNotification Send a stolen bike notification | BIKE ID MUST BE FROM A STOLEN BIKE
// Requires `read_user` in the access token you use to send the notification
// Your application has to be approved to be able to do this. Email support@bikeindex.org to get access.
// Before your application is approved you can send notifications to yourself (to a bike that you own that’s stolen).
func (bi *BikeIndex) SendStolenNotification(id int64, options SendStolenNotification) (*SendStolenNotification, *Exception, error) {
	// convert options to HTTP form
	form, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	url := bi.buildURL("bikes", strconv.FormatInt(id, 10)) //".json"
	res, err := bi.put(form, url)
	if err != nil {
		return nil, nil, err
	}
	decoder := json.NewDecoder(res.Body)

	// handle NULL response
	if res.StatusCode != http.StatusOK {
		exception := new(Exception)
		err = decoder.Decode(exception)
		return nil, exception, err
	}

	ssn := new(SendStolenNotification)
	err = decoder.Decode(ssn)
	return ssn, nil, err
}
