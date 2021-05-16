package common

type DiscordEditWebhookMessage struct {
	// the message contents (up to 2000 characters)
	Content string `json:"content"`
	// of up to 10 embed objects	embedded rich content
	Embeds DiscordEmbed `json:"embeds"`
	// contents	the contents of the file being sent/edited
	File string `json:"file"`
	// JSON encoded body of non-file params (multipart/form-data only)
	PayloadJSON string `json:"payload_json"`
	// mention object	allowed mentions for the message
	AllowedMentions DiscordAllowedMention `json:"allowed_mentions"`
	// of attachment objects	attached files to keep
	Attachments DiscordAttachment `json:"attachments"`
}
