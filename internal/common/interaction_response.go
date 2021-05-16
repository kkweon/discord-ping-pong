package common

type DiscordInteractionResponse struct {
	// the type of response
	Type DiscordInteractionCallbackType `json:"type"`
	// 	an optional response message
	Data DiscordInteractionApplicationCommandCallbackData `json:"data"`
}
