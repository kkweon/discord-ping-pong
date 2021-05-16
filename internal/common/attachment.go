package common

type DiscordAttachment struct {
	// attachment id
	ID string `json:"id"`
	// name of file attached
	Filename string `json:"filename"`
	// the attachment's media type
	ContentType string `json:"content_type,omitempty"`
	// size of file in bytes
	Size int `json:"size"`
	// source url of file
	URL string `json:"url"`
	// a proxied url of file
	ProxyURL string `json:"proxy_url"`
	// height of file (if image)
	Height int `json:"height,omitempty"`
	// width of file (if image)
	Width int `json:"width,omitempty"`
}
