package common

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalInteractionResponse(t *testing.T) {
	jsonSerialized := `{
            "type": 4,
            "data": {
                "tts": false,
                "content": "Congrats on sending your command!",
                "embeds": [],
                "allowed_mentions": { "parse": [] }
            }
        }`

	var resp DiscordInteractionResponse
	err := json.Unmarshal([]byte(jsonSerialized), &resp)

	assert.NoError(t, err)
	assert.Equal(t, DiscordInteractionCallbackTypeChannelMessageWithSource, resp.Type)
}
