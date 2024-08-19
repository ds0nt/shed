import {
  fetchConversations,
  createConversation,
  getConversation,
  sendMessage,
} from '@/api/api'

const defaultXHRState = () => ({
  isLoading: false,
  error: null,
  data: null,
  response: null,
})

const state = () => ({

  general: {
    isLoading: false,
    loadError: null,
    selectedModel: null,

    // Do i need this?
    modelPrefs: {
      seed: null,
      numPredict: null,
      temperature: null,
      frequencyPenalty: null,
      presencePenalty: null,
      stop: null,
      topK: null,
      topP: null,
      mirostat: null,
      mirostatEta: null,
      mirostatTau: null,
      repeatLastN: null,
      tfsZ: null,
      contextLength: null,
      batchSize: null,
    },

    currentUser: null,
    currentConversation: null,

    isTyping: false,
    draftMessages: {},

    searchTerm: '',
  },

  conversations: [],
  xhr: {
    createConversations: defaultXHRState(),
    fetchConversations: defaultXHRState(),
  },
  messages: [],
})

const actions = {
  async createConversation(data) {
    this.xhr.createConversations.isLoading = true
    try {
      const response = await createConversation(data)
      this.xhr.createConversations.response = response
      this.xhr.createConversations.data = response.data
    } catch (error) {
      this.xhr.createConversations.error = error
    } finally  {
      this.xhr.createConversations.isLoading = false
     }
  },
  async refreshConversations() {
    this.setIsLoading(true);
    try {
      const conversations = await fetchConversations();
      this.setConversations(conversations);
      this.setIsLoading(false);
    } catch (error) {
      this.setError(error);
    }
  },
  setSelectedModel(model) {
    this.preferences.selectedModel = model
  },
  // ...existing actions
  setIsTyping(isTyping) {
    this.isTyping = isTyping
  },
  setError(error) {
    this.error = error
  },
  setIsLoading(isLoading) {
    this.isLoading = isLoading
  },
  setSearchTerm(searchTerm) {
    this.searchTerm = searchTerm
  },
  setDraftMessage({ conversationId, draftMessage }) {
    this.draftMessages = { ...this.draftMessages, [conversationId]: draftMessage }
  },

  // Set the list of conversations
  setConversations(conversations) {
    this.conversations = conversations
  },
  // Set the current conversation
  setCurrentConversation(conversation) {
    this.currentConversation = conversation
  },

  // Set the list of messages in the current conversation
  setMessages(messages) {
    this.messages = messages
  },
  // Set the selected model
  // Add a new message to a specific conversation
  async addMessage({ conversationId, message }) {
    const targetConversation = this.conversations.find(c => c.id === conversationId);
    if (targetConversation) {
      targetConversation.messages.push(message);
    } else {
      throw new Error(`Conversation ${conversationId} not found.`);
    }
  },

  // Handle incoming message - This might be different depending on how you get incoming messages.
  async handleIncomingMessage(message) {
    const targetConversation = this.conversations.find(c => c.id === message.conversationId);
    if (targetConversation) {
      targetConversation.messages.push(message);
      if (this.currentConversation?.id === message.conversationId) {
        // If the incoming message belongs to the current conversation, scroll to bottom or do whatever you want 
      }
    }
  },

  // Update a message in a specific conversation
  async updateMessage({ conversationId, messageId, newText }) {
    const targetConversation = this.conversations.find(c => c.id === conversationId);
    if (targetConversation) {
      const targetMessage = targetConversation.messages.find(m => m.id === messageId);
      if (targetMessage) {
        targetMessage.text = newText;
      }
    }
  },

  // Delete a message from a specific conversation
  async deleteMessage({ conversationId, messageId }) {
    const targetConversation = this.conversations.find(c => c.id === conversationId);
    if (targetConversation) {
      const messageIndex = targetConversation.messages.findIndex(m => m.id === messageId);
      if (messageIndex > -1) {
        targetConversation.messages.splice(messageIndex, 1);
      }
    }
  },

  // Select a model
  selectModel(modelId) {
    const selectedModel = models.find(m => m.id === modelId); // Assuming models is a list of available models.
    if (selectedModel) {
      this.setSelectedModel(selectedModel);
    } else {
      throw new Error(`Model ${modelId} not found.`);
    }
  }
}


const getters = {
  // Get the list of conversations
  getConversations(state) {
    return state.conversations
  },
  // Get the current conversation
  getCurrentConversation(state) {
    return state.currentConversation
  },
  // Get the list of messages in the current conversation
  getMessages(state) {
    return state.messages
  },
  // Get the selected model
  getSelectedModel(state) {
    return state.preferences.selectedModel
  },
  getIsTyping(state) {
    return state.isTyping
  },
  getError(state) {
    return state.error
  },
  getIsLoading(state) {
    return state.isLoading
  },
  getFilteredConversations(state) {
    return state.conversations.filter(c => c.name.includes(state.searchTerm))
  },
  getDraftMessage: (state) => (conversationId) => {
    return state.draftMessages[conversationId]
  },
}

import { defineStore } from 'pinia'
export const useAppStore = defineStore('app', {
  state,
  getters,
  actions
})