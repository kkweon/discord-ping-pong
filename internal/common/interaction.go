package common

import (
	"github.com/bwmarrin/discordgo"
)

type DiscordInteraction struct {
	// 	id of the interaction
	ID string `json:"id"`
	// 	id of the application this interaction is for
	ApplicationID string `json:"application_id"`
	// 	the type of interaction
	Type DiscordInteractionType `json:"type"`
	// 	the command data payload
	Data struct {
		// 	the ID of the invoked command
		ID string `json:"id"`
		// 	the name of the invoked command
		Name string `json:"name"`
		// 	converted users + roles + channels
		Resolved *DiscordApplicationCommandInteractionDataResolved `json:"resolved,omitempty"`
		//  of ApplicationCommandInteractionDataOption	the params + values from the user
		Options []*DiscordApplicationCommandInteractionDataOption `json:"options,omitempty"`
	} `json:"data"`
	// 	the guild it was sent from
	GuildID string `json:"guild_id"`
	// 	the channel it was sent from
	ChannelID string `json:"channel_id"`
	//  member object	guild member data for the invoking user, including permissions
	Member *DiscordMember `json:"member,omitempty"`
	//  object	user object for the invoking user, if invoked in a DM
	User discordgo.User `json:"user"`
	// 	a continuation token for responding to the interaction
	Token string `json:"token"`
	// 	read-only property, always 1
	Version int `json:"version"`
}
