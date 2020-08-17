package gobikeindex

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
)

// RefreshToken required fields
type RefreshToken struct {
	GrantType    string `url:"grant_type" json:"grant_type"`
	ClientID     string `url:"client_id" json:"client_id"`
	RefreshToken string `url:"refresh_token" json:"refresh_token"`
}

// RefreshToken - Since tokens provide access to someone’s account, one of the ways OAuth2 keeps things secure is by having tokens expire
// so if someone compromises an access token, they don’t have unlimited access to an account. The Bike Index expires tokens after 1 hour.
// When your access token expires, you can get a new one by making a POST request with your app id and the refresh token:
func (bi *BikeIndex) RefreshToken() (*Exception, error) {
	options := RefreshToken{
		GrantType:    "refresh_token",
		ClientID:     bi.appID,
		RefreshToken: bi.refreshToken,
	}
	form, err := query.Values(options)
	if err != nil {
		return nil, err
	}
	log.Printf("@@@ %+v", form)
	log.Printf("@@@ %+v", options)

	res, err := bi.post(form, tokenURL) //.json
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(res.Body)

	// handle NULL response
	if res.StatusCode != http.StatusOK {
		exception := new(Exception)
		err = decoder.Decode(exception)
		return exception, err
	}

	rt := new(RefreshToken)
	err = decoder.Decode(rt)
	bi.refreshToken = rt.RefreshToken
	return nil, err
}

// func (bi *BikeIndex) BikeUpdate1(id int64, options BikeUpdate) (*BikeUpdate, *Exception, error) {
// 	form, err := query.Values(participant)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	res, err := twilio.post(form, twilio.buildUrl(fmt.Sprintf("Conferences/%s/Participants.json", conferenceSid)))
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	decoder := json.NewDecoder(res.Body)

// 	if res.StatusCode != http.StatusCreated {
// 		exception := new(Exception)
// 		err = decoder.Decode(exception)
// 		return nil, exception, err
// 	}

// 	conf := new(ConferenceParticipant)
// 	err = decoder.Decode(conf)
// 	return conf, nil, err
// }
