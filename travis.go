package travis

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// GetPayload just reads from the request and returns the unmarshalled Payload
func GetPayload(r *http.Request) (*Payload, error) {

	var b requestBody
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &b)
	if err != nil {
		return nil, err
	}

	if b.Payload == nil {
		return nil, errors.New("missing data")
	}

	return b.Payload, nil

}
