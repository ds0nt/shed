<template>
    <v-app>
      <v-container>
        <v-row class="mx-0">
          <div v-for="user in chatUsers" :key="user.id" class="chat-user">
            <v-avatar :color="user.onlineStatus ? 'green' : 'red'">{{ user.name[0] }}</v-avatar>
            <h5>{{ user.name }}</h5>
          </div>
        </v-row>
        <v-row class="mx-0 flex-grow-1 chat-messages">
          <v-list three-line>
            <v-list-item v-for="msg in chatMessages" :key="msg.id">
                <v-list-item-subtitle>{{ msg.username }}</v-list-item-subtitle>
                <v-list-item-title>{{ msg.message }}</v-list-item-title>
              <div v-if="timeStamps">{{ msg.time }}</div>
            </v-list-item>
          </v-list>
        </v-row>
        <v-row class="mx-0 justify-end">
          <v-textarea v-model="chatInput" label="Type your message" filled background-color="white" clearable></v-textarea>
          <v-btn color="success" text @click="sendMessage">Send</v-btn>
        </v-row>
      </v-container>
    </v-app>
  </template>
  
<script setup>
import { chatInput, chatMessages, chatUsers, currentChatUser, chatBoxActiveStatus, timeStamps } from './Chat';

const sendMessage = () => {
  if(chatInput.value) {
    chatMessages.push({message: chatInput.value, sender: currentChatUser.value, time: new Date(), status: 'unread'});
    chatInput.value = '';
  }
};
</script>