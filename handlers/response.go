package handlers

import (
	"encoding/json"
	"errors"
	"io"
)

func DecodeJson(reader io.Reader, target interface{}) (err error) {
	decode := json.NewDecoder(reader)
	err = decode.Decode(target)

	if err != nil {
		return errors.New("Failed to decode json data")
	}
	return nil
}
