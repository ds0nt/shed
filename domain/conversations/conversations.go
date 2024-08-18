package conversations

import (
	"fmt"
	"strconv"
	"strings"
)

type ConversationKey struct {
	Owner string
	ID    int
}

func NewConversationKey(owner string, id int) ConversationKey {
	return ConversationKey{Owner: owner, ID: id}
}
func (k *ConversationKey) String() string {
	return fmt.Sprintf("%s:%d", k.Owner, k.ID)
}
func NewConversationKeyFromString(s string) ConversationKey {
	var owner string
	var id int
	parts := strings.Split(s, ":")
	if len(parts) == 2 {
		owner = parts[0]
		id, _ = strconv.Atoi(parts[1])
	}
	return ConversationKey{Owner: owner, ID: id}
}

type Conversation struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

type Message struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	From      string `json:"from"`
	Timestamp int64  `json:"timestamp"`
}
