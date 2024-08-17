package conversations

import "fmt"

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
	fmt.Sscanf(s, "%s:%d", &owner, &id)
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
