package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ds0nt/shed/domain/conversations"
	"github.com/ds0nt/shed/pkg/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var _ = log.InitLogger()

func TestMain(m *testing.M) {
	err := os.RemoveAll("data")
	if err != nil {
		log.Error("Failed to remove data directory")
	}
	os.Exit(m.Run())
}

func TestService_listConversationsHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	// Create a new conversation
	ctx := context.Background()
	conversation := conversations.Conversation{
		Name: "Test Conversation",
		Id:   "test:test",
	}
	key := conversations.ConversationKey{
		Owner: "test",
		ID:    "test",
	}
	err := s.Store.CreateJSON(ctx, "conversations", key.String(), conversation)
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
	assert.Contains(t, rec.Body.String(), key.String())
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

	id, err := uuid.NewV7()
	assert.NoError(t, err)
	ctx := context.Background()
	key := "admin:" + id.String()

	conversation := conversations.Conversation{
		Id:       id.String(),
		Name:     "Test Conversation",
		Messages: []conversations.Message{{ID: id.String(), Text: "TEST"}},
	}
	err = s.Store.CreateJSON(ctx, "conversations", key, conversation)
	assert.NoError(t, err)

	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/conversations/"+key, nil)
	rec := httptest.NewRecorder()

	// Call the getConversationHandler
	c := s.Echo.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(key)
	err = s.getConversationHandler(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), conversation.Id)
}
