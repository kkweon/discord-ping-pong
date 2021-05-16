package common

type DiscordAllowedMention struct {
	// of allowed mention types	An array of allowed mention types to parse from the content.
	Parse []string `json:"parse,omitempty"`
	// of snowflakes	Array of role_ids to mention (Max size of 100)
	Roles []string `json:"roles,omitempty"`
	// of snowflakes	Array of user_ids to mention (Max size of 100)
	Users []string `json:"users,omitempty"`
	// For replies, whether to mention the author of the message being replied to (default false)
	RepliedUser bool `json:"replied_user,omitempty"`
}
