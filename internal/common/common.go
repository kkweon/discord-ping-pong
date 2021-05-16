package common

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

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

type DiscordApplicationCommandOptionChoice struct {
	// 1-100 character choice name
	Name string `json:"name"`
	// value of the choice, up to 100 characters if string
	Value string `json:"value"`
}

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
	Choices []DiscordApplicationCommandOptionChoice `json:"choices,omitempty"`
	//	array of ApplicationCommandOption	if the option is a subcommand or subcommand group type, this nested options will be the parameters
	Options []DiscordApplicationCommandOption `json:"options,omitempty"`
}

type DiscordApplicationCommand struct {
	//	snowflake	unique id of the command
	ID string `json:"id,omitempty"`
	// snowflake	unique id of the parent application
	ApplicationID string `json:"application_id,omitempty"`
	//	1-32 lowercase character name matching ^[\w-]{1,32}$
	Name string `json:"name"`
	//1-100 character description
	Description       string                            `json:"description"`
	Options           []DiscordApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission bool                              `json:"default_permission,omitempty"`
}

type DiscordApplicationCommandInteractionDataResolved struct {
	//  ID: User	the IDs and User objects
	Users string `json:"users"`
	//  ID: Member	the IDs and partial Member objects
	Members string `json:"members"`
	//  ID: Role	the IDs and Role objects
	Roles string `json:"roles"`
	//  ID: Channel	the IDs and partial Channel objects
	Channels string `json:"channels"`
}

type DiscordInteractionType int

const (
	DiscordInteractionTypePing               DiscordInteractionType = 1
	DiscordInteractionTypeApplicationCommand                        = 2
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
		Resolved DiscordApplicationCommandInteractionDataResolved `json:"resolved"`
		//  of ApplicationCommandInteractionDataOption	the params + values from the user
		Options []DiscordApplicationCommandInteractionDataOption `json:"options"`
	} `json:"data"`
	// 	the guild it was sent from
	GuildID string `json:"guild_id"`
	// 	the channel it was sent from
	ChannelID string `json:"channel_id"`
	//  member object	guild member data for the invoking user, including permissions
	Member struct {
		//  object	the user this guild member represents
		User discordgo.User `json:"user"`
		// 	this users guild nickname
		Nick string `json:"nick"`
		//  of snowflakes	array of role object ids
		Roles []string `json:"roles"`
		//  timestamp	when the user joined the guild
		JoinedAt time.Time `json:"joined_at"`
		//  timestamp	when the user started boosting the guild
		PremiumSince time.Time `json:"premium_since"`
		// 	whether the user is deafened in voice channels
		Deaf bool `json:"deaf"`
		// 	whether the user is muted in voice channels
		Mute bool `json:"mute"`
		// 	whether the user has not yet passed the guild's Membership Screening requirements
		Pending bool `json:"pending"`
		// 	total permissions of the member in the channel, including overrides, returned when in the interaction object
		Permissions string `json:"permissions"`
	} `json:"member"`
	//  object	user object for the invoking user, if invoked in a DM
	User discordgo.User `json:"user"`
	// 	a continuation token for responding to the interaction
	Token string `json:"token"`
	// 	read-only property, always 1
	Version int `json:"version"`
}

type DiscordInteractionCallbackType int

const (
	// DiscordInteractionCallbackTypePong ACK a Ping
	DiscordInteractionCallbackTypePong DiscordInteractionCallbackType = 1
	// DiscordInteractionCallbackTypeChannelMessageWithSource respond to an interaction with a message
	DiscordInteractionCallbackTypeChannelMessageWithSource DiscordInteractionCallbackType = 4
	// DiscordInteractionCallbackTypeDeferredChannelMessageWithSource ACK an interaction and edit a response later, the user sees a loading state
	DiscordInteractionCallbackTypeDeferredChannelMessageWithSource DiscordInteractionCallbackType = 5
)

type DiscordImage struct {
	// source url of image (only supports http(s) and attachments)
	URL string `json:"url"`
	// a proxied url of the image
	ProxyURL string `json:"proxy_url"`
	// height of image
	Height int `json:"height"`
	// width of image
	Width int `json:"width"`
}

type DiscordEmbed struct {
	// title of embed
	Title string `json:"title"`
	// type of embed (always "rich" for webhook embeds)
	Type string `json:"type"`
	// description of embed
	Description string `json:"description"`
	// url of embed
	URL string `json:"url"`
	// timestamp	timestamp of embed content
	Timestamp time.Time `json:"timestamp"`
	// color code of the embed
	Color int `json:"color"`
	// footer object	footer information
	Footer struct {
		// footer text
		Text string `json:"text"`
		// url of footer icon (only supports http(s) and attachments)
		IconURL string `json:"icon_url"`
		// a proxied url of footer icon
		ProxyIconURL string `json:"proxy_icon_url"`
	} `json:"footer"`
	// image object	image information
	Image DiscordImage `json:"image"`
	// thumbnail object	thumbnail information
	Thumbnail DiscordImage `json:"thumbnail"`
	// video object	video information
	Video DiscordImage `json:"video"`
	// provider object	provider information
	Provider struct {
		// name of provider
		Name string `json:"name"`
		// url of provider
		URL string `json:"url"`
	} `json:"provider"`
	// author object	author information
	Author struct {
		// name of author
		Name string `json:"name"`
		// url of author
		URL string `json:"url"`
		// url of author icon (only supports http(s) and attachments)
		IconURL string `json:"icon_url"`
		// a proxied url of author icon
		ProxyIconURL string `json:"proxy_icon_url"`
	} `json:"author"`
	// of embed field objects	fields information
	Fields []struct {
		// name of the field
		Name string `json:"name"`
		// value of the field
		Value string `json:"value"`
		// whether or not this field should display inline
		Inline bool `json:"inline"`
	} `json:"fields"`
}

type DiscordAllowedMention struct {
	// of allowed mention types	An array of allowed mention types to parse from the content.
	Parse []string `json:"parse"`
	// of snowflakes	Array of role_ids to mention (Max size of 100)
	Roles []string `json:"roles"`
	// of snowflakes	Array of user_ids to mention (Max size of 100)
	Users []string `json:"users"`
	// For replies, whether to mention the author of the message being replied to (default false)
	RepliedUser bool `json:"replied_user"`
}

type DiscordInteractionApplicationCommandCallbackData struct {
	//	is the response TTS
	TTS bool `json:"tts"`
	// message content
	Content string `json:"content"`
	// of embeds	supports up to 10 embeds
	Embeds []DiscordEmbed `json:"embeds"`
	// allowed mentions object
	AllowedMentions DiscordAllowedMention `json:"allowed_mentions"`
	// set to 64 to make your response ephemeral
	Flags int `json:"flags"`
}

type DiscordInteractionResponse struct {
	// the type of response
	Type DiscordInteractionCallbackType `json:"type"`
	// 	an optional response message
	Data DiscordInteractionApplicationCommandCallbackData `json:"data"`
}
