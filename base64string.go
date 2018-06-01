package base64string

import (
	"encoding/base64"
	"encoding/json"
)

// Base64String -- automaticly base64 enocdes string when JSON marshaled
// NOTE: only handles encode when marshaled in case of JSON marshal!!
type Base64String string

var maskedReponse = "********"

// String -- always returns masked response
func (t Base64String) String() string {
	return maskedReponse
}

// Get -- getter function for real non-encode payload
func (t Base64String) Get() string {
	return string(t)
}

// MarshalJSON -- encode string in to base64 encode string
func (t Base64String) MarshalJSON() ([]byte, error) {

	s := string(t)
	enc := encode(&s)

	return json.Marshal(enc)
}

// UnmarshalJSON -- decode base64 encode string back to string
func (t *Base64String) UnmarshalJSON(b []byte) (err error) {

	var tmp string
	if err = json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	decodedString := decode(&tmp)
	*t = Base64String(decodedString)

	return nil
}

// decode -- base64 encoded string into string
func decode(s *string) string {
	if s == nil {
		return ""
	}
	decoded, err := base64.StdEncoding.DecodeString(*s)
	if err != nil {
		return ""
	}
	return string(decoded)
}

// encode -- string into base64 encoded string
func encode(s *string) string {
	if s == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(*s))
}
