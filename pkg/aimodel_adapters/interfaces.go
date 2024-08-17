package aimodeladapters

import "github.com/ds0nt/shed/domain/conversations"

type Chatter interface {
	Chat(conversations.Conversation) string
}
