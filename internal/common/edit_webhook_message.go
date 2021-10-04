package common

// DiscordEditWebhookMessage is the type represented in
// https://discord.com/developers/docs/resources/webhook#edit-webhook-message
type DiscordEditWebhookMessage struct {
	// the message contents (up to 2000 characters)
	Content string `json:"content,omitempty"`
	// of up to 10 embed objects	embedded rich content
	Embeds []*DiscordEmbed `json:"embeds,omitempty"`
	// contents	the contents of the file being sent/edited
	File string `json:"file,omitempty"`
	// JSON encoded body of non-file params (multipart/form-data only)
	PayloadJSON string `json:"payload_json,omitempty"`
	// mention object	allowed mentions for the message
	AllowedMentions *DiscordAllowedMention `json:"allowed_mentions,omitempty"`
	// of attachment objects	attached files to keep
	Attachments *DiscordAttachment `json:"attachments,omitempty"`
}
