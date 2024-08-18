package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/ds0nt/shed/domain/conversations"
	"github.com/ds0nt/shed/pkg/log"
	"github.com/stretchr/testify/assert"
)

var _ = log.InitLogger()

func TestService_listConversationsHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	// Create a new conversation
	ctx := context.Background()
	conversation := conversations.Conversation{
		Id:       0,
		Name:     "Test Conversation",
		Messages: []conversations.Message{{ID: 0, Text: "TEST"}},
	}
	err := s.Store.CreateJSON(ctx, "conversations", "testkey", conversation)
	assert.NoError(t, err)

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/conversations", nil)
	rec := httptest.NewRecorder()

	// Call the listConversationsHandler
	c := s.Echo.NewContext(req, rec)
	err = s.listConversationsHandler(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), strconv.Itoa(conversation.Id))
}

func TestService_createConversationHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	// Create a new request
	req := httptest.NewRequest(http.MethodPost, "/conversations", nil)
	rec := httptest.NewRecorder()

	// Call the createConversationHandler
	c := s.Echo.NewContext(req, rec)
	err := s.createConversationHandler(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestService_getConversationHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	ctx := context.Background()
	key := "admin:0"
	conversation := conversations.Conversation{
		Id:       0,
		Name:     "Test Conversation",
		Messages: []conversations.Message{{ID: 0, Text: "TEST"}},
	}
	err := s.Store.CreateJSON(ctx, "conversations", key, conversation)
	assert.NoError(t, err)

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/conversations/"+key, nil)
	rec := httptest.NewRecorder()

	// Call the getConversationHandler
	c := s.Echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("admin:0")
	err = s.getConversationHandler(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), strconv.Itoa(conversation.Id))
}
