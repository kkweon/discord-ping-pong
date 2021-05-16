package common

type DiscordInteractionApplicationCommandCallbackData struct {
	//	is the response TTS
	TTS bool `json:"tts"`
	// message content
	Content string `json:"content"`
	// of embeds	supports up to 10 embeds
	Embeds []*DiscordEmbed `json:"embeds,omitempty"`
	// allowed mentions object
	AllowedMentions *DiscordAllowedMention `json:"allowed_mentions,omitempty"`
	// set to 64 to make your response ephemeral
	Flags int `json:"flags"`
}
