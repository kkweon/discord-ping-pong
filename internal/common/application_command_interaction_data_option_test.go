package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscordOptionTypeString_MarshalJSON(t *testing.T) {
	helloworld := "helloworld"
	msg := DiscordApplicationCommandInteractionDataOption{
		Name: "name",
		Type: DiscordApplicationCommandOptionTypeString,
		Value: DiscordOptionType{
			StringValue: &helloworld,
		},
	}

	b, err := json.Marshal(msg)
	assert.NoError(t, err)

	expected := `{"name":"name","type":3,"value":"helloworld"}`
	assert.Equal(t, expected, string(b))
}

func TestDiscordOptionTypeInt_MarshalJSON(t *testing.T) {
	value := 1
	msg := DiscordApplicationCommandInteractionDataOption{
		Name: "name",
		Type: DiscordApplicationCommandOptionTypeString,
		Value: DiscordOptionType{
			IntValue: &value,
		},
	}

	b, err := json.Marshal(msg)
	assert.NoError(t, err)

	expected := `{"name":"name","type":3,"value":1}`
	assert.Equal(t, expected, string(b))
}

func TestDiscordOptionType_UnmarshalJSON(t *testing.T) {
	input := `{"name":"name","type":3,"value":1024}`
	value := 1024
	expected := DiscordApplicationCommandInteractionDataOption{
		Name: "name",
		Type: 3,
		Value: DiscordOptionType{
			IntValue: &value,
		},
	}

	var actual DiscordApplicationCommandInteractionDataOption
	err := json.Unmarshal([]byte(input), &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
