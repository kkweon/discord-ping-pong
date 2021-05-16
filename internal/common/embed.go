package common

import "time"

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
	Footer *DiscordFooter `json:"footer,omitempty"`
	// image object	image information
	Image *DiscordImage `json:"image,omitempty"`
	// thumbnail object	thumbnail information
	Thumbnail *DiscordImage `json:"thumbnail,omitempty"`
	// video object	video information
	Video *DiscordImage `json:"video,omitempty"`
	// provider object	provider information
	Provider *DiscordProvider `json:"provider,omitempty"`
	// author object	author information
	Author *DiscordAuthor `json:"author,omitempty"`
	// of embed field objects	fields information
	Fields []DiscordEmbedField `json:"fields,omitempty"`
}
