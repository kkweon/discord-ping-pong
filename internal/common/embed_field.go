package common

type DiscordEmbedField struct {
	// name of the field
	Name string `json:"name"`
	// value of the field
	Value string `json:"value"`
	// whether or not this field should display inline
	Inline bool `json:"inline"`
}
