package common

type DiscordApplicationCommand struct {
	//	snowflake	unique id of the command
	ID string `json:"id,omitempty"`
	// snowflake	unique id of the parent application
	ApplicationID string `json:"application_id,omitempty"`
	//	1-32 lowercase character name matching ^[\w-]{1,32}$
	Name string `json:"name"`
	//1-100 character description
	Description       string                             `json:"description"`
	Options           []*DiscordApplicationCommandOption `json:"options,omitempty"`
	DefaultPermission bool                               `json:"default_permission,omitempty"`
}
