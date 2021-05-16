package common

import (
	"encoding/json"
	"errors"
)

type DiscordOptionType struct {
	IntValue    *int
	StringValue *string
}

func (d *DiscordOptionType) UnmarshalJSON(bytes []byte) error {
	var s string
	err := json.Unmarshal(bytes, &s)
	if err == nil {
		d.StringValue = &s
		return nil
	}

	var i int
	err = json.Unmarshal(bytes, &i)
	if err == nil {
		d.IntValue = &i
		return nil
	}

	return err
}

func (d DiscordOptionType) MarshalJSON() ([]byte, error) {
	if d.IntValue != nil {
		return json.Marshal(*d.IntValue)
	} else if d.StringValue != nil {
		return json.Marshal(*d.StringValue)
	}
	return nil, errors.New("no value is set")
}
