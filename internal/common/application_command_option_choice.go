package common

type DiscordApplicationCommandOptionChoice struct {
	// 1-100 character choice name
	Name string `json:"name"`
	// value of the choice, up to 100 characters if string
	Value string `json:"value"`
}
