package common

type DiscordAuthor struct {
	// name of author
	Name string `json:"name"`
	// url of author
	URL string `json:"url"`
	// url of author icon (only supports http(s) and attachments)
	IconURL string `json:"icon_url"`
	// a proxied url of author icon
	ProxyIconURL string `json:"proxy_icon_url"`
}
