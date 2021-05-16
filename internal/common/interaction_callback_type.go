package common

type DiscordInteractionCallbackType int

const (
	// DiscordInteractionCallbackTypePong ACK a Ping
	DiscordInteractionCallbackTypePong DiscordInteractionCallbackType = 1
	// DiscordInteractionCallbackTypeChannelMessageWithSource respond to an interaction with a message
	DiscordInteractionCallbackTypeChannelMessageWithSource DiscordInteractionCallbackType = 4
	// DiscordInteractionCallbackTypeDeferredChannelMessageWithSource ACK an interaction and edit a response later, the user sees a loading state
	DiscordInteractionCallbackTypeDeferredChannelMessageWithSource DiscordInteractionCallbackType = 5
)
