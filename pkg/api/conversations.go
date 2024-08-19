package api

import (
	"net/http"

	"github.com/ds0nt/shed/domain/conversations"
	"github.com/ds0nt/shed/pkg/log"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// listConversationsHandler returns a list of chat conversations with an AI
func (s *Service) listConversationsHandler(c echo.Context) error {
	conversations := []*conversations.Conversation{}
	// Get the list of chat conversations from the storage mechanism
	err := s.Store.ListJSON(c.Request().Context(), "conversations", &conversations)
	if err != nil {
		log.Errorf("Failed to get chat conversations: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to get chat conversations")
	}

	// Return the list of chat conversations as JSON
	return c.JSON(http.StatusOK, conversations)
}

// createConversationHandler creates a new chat conversation with an AI
func (s *Service) createConversationHandler(c echo.Context) error {
	// Get the chat conversation from the request
	chat := &conversations.Conversation{}
	if err := c.Bind(chat); err != nil {
		log.Errorf("Failed to bind chat: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to bind chat")
	}

	guid, err := uuid.NewV7()
	if err != nil {
		log.Errorf("Failed to generate chat ID: %v", err)
	}
	chat.Id = guid.String()

	key := conversations.NewConversationKey("admin", chat.Id)

	// Save the chat conversation
	if err := s.Store.CreateJSON(c.Request().Context(), "conversations", key.String(), chat); err != nil {
		log.Errorf("Failed to create chat conversation: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to create chat conversation")
	}

	// Return success
	return c.String(http.StatusOK, "Chat conversation created")
}

// getConversationHandler returns a chat conversation with an AI
func (s *Service) getConversationHandler(c echo.Context) error {
	// Get the chat conversation ID from the request
	id := c.Param("id")
	key := conversations.NewConversationKeyFromString(id)

	// Get the chat conversation from the storage mechanism
	chat := &conversations.Conversation{}
	if err := s.Store.ReadJSON(c.Request().Context(), "conversations", key.String(), chat); err != nil {
		log.Errorf("Failed to get chat conversation: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to get chat conversation")
	}

	// Return the chat conversation as JSON
	return c.JSON(http.StatusOK, chat)
}

// sendMessageHandler sends a message in a chat conversation with an AI
func (s *Service) sendMessageHandler(c echo.Context) error {
	// Get the chat conversation ID from the request
	id := c.Param("id")
	key := conversations.NewConversationKeyFromString(id)

	// Get the chat conversation from the storage mechanism
	chat := &conversations.Conversation{}
	if err := s.Store.ReadJSON(c.Request().Context(), "conversations", key.String(), chat); err != nil {
		log.Errorf("Failed to get chat conversation: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to get chat conversation")
	}

	// Get the message from the request
	message := &conversations.Message{}
	if err := c.Bind(message); err != nil {
		log.Errorf("Failed to bind message: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to bind message")
	}

	// Add the message to the chat conversation
	chat.Messages = append(chat.Messages, *message)

	// Save the chat conversation
	if err := s.Store.UpdateJSON(c.Request().Context(), "conversations", key.String(), chat); err != nil {
		log.Errorf("Failed to update chat conversation: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to update chat conversation")
	}

	// Return success
	return c.String(http.StatusOK, "Message sent")
}
