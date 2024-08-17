<template>
  <v-app>
    <app-bar />
    <v-main>
      <Chat />
      <v-container fluid>
      </v-container>
    </v-main>
    <app-drawer />    
  </v-app>
</template>

<script setup>
import AppBar from './AppBar.vue';
import AppDrawer from './AppDrawer.vue';
import Chat from './Chat.vue';

</script>

