package common

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
