package common

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
