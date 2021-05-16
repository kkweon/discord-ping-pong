package common

// DiscordApplicationCommandInteractionDataOption
type DiscordApplicationCommandInteractionDataOption struct {
	// the name of the parameter
	Name string `json:"name"`
	// value of ApplicationCommandOptionType
	Type DiscordApplicationCommandOptionType `json:"type"`
	// the value of the pair
	Value DiscordOptionType `json:"value"`
	// of ApplicationCommandInteractionDataOption	present if this option is a group or subcommand
	Options []DiscordApplicationCommandInteractionDataOption `json:"options,omitempty"`
}
