package common

type DiscordApplicationCommandOptionType int

const (
	DiscordApplicationCommandOptionTypeSubCommand      DiscordApplicationCommandOptionType = 1
	DiscordApplicationCommandOptionTypeSubCommandGroup DiscordApplicationCommandOptionType = 2
	DiscordApplicationCommandOptionTypeString          DiscordApplicationCommandOptionType = 3
	DiscordApplicationCommandOptionTypeInteger         DiscordApplicationCommandOptionType = 4
	DiscordApplicationCommandOptionTypeBoolean         DiscordApplicationCommandOptionType = 5
	DiscordApplicationCommandOptionTypeUser            DiscordApplicationCommandOptionType = 6
	DiscordApplicationCommandOptionTypeChannel         DiscordApplicationCommandOptionType = 7
	DiscordApplicationCommandOptionTypeRole            DiscordApplicationCommandOptionType = 8
	DiscordApplicationCommandOptionTypeMentionable     DiscordApplicationCommandOptionType = 9
)
