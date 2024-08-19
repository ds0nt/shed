package conversations

import (
	"fmt"
	"strings"
)

type ConversationKey struct {
	Owner string
	ID    string
}

func NewConversationKey(owner string, id string) ConversationKey {
	return ConversationKey{Owner: owner, ID: id}
}
func (k *ConversationKey) String() string {
	return fmt.Sprintf("%s:%s", k.Owner, k.ID)
}
func NewConversationKeyFromString(s string) ConversationKey {
	var owner string
	var id string
	parts := strings.Split(s, ":")
	if len(parts) == 2 {
		owner = parts[0]
		id = parts[1]
	}
	return ConversationKey{Owner: owner, ID: id}
}

type Conversation struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

type Message struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	From      string `json:"from"`
	Timestamp int64  `json:"timestamp"`
}
