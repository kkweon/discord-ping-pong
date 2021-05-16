package common

type DiscordApplicationCommandOption struct {
	// value of ApplicationCommandOptionType
	Type DiscordApplicationCommandOptionType `json:"type"`
	//	1-32 lowercase character name matching ^[\w-]{1,32}$
	Name string `json:"name"`
	//	1-100 character description
	Description string `json:"description"`
	//	boolean	if the parameter is required or optional--default false
	Required bool `json:"required"`
	//	array of ApplicationCommandOptionChoice	choices for string and int types for the user to pick from
	Choices []*DiscordApplicationCommandOptionChoice `json:"choices,omitempty"`
	//	array of ApplicationCommandOption	if the option is a subcommand or subcommand group type, this nested options will be the parameters
	Options []*DiscordApplicationCommandOption `json:"options,omitempty"`
}
