package common

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type DiscordMember struct {
	//  object	the user this guild member represents
	User discordgo.User `json:"user"`
	// 	this users guild nickname
	Nick string `json:"nick"`
	//  of snowflakes	array of role object ids
	Roles []string `json:"roles"`
	//  timestamp	when the user joined the guild
	JoinedAt time.Time `json:"joined_at"`
	//  timestamp	when the user started boosting the guild
	PremiumSince time.Time `json:"premium_since"`
	// 	whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`
	// 	whether the user is muted in voice channels
	Mute bool `json:"mute"`
	// 	whether the user has not yet passed the guild's Membership Screening requirements
	Pending bool `json:"pending"`
	// 	total permissions of the member in the channel, including overrides, returned when in the interaction object
	Permissions string `json:"permissions"`
}
