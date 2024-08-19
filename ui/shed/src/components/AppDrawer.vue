<template>
    <v-navigation-drawer v-model="drawer" app>
        <v-toolbar flat>
            <v-text-field flat solo hide-details label="Search" prepend-inner-icon="mdi-magnify"></v-text-field>
            <v-btn @click="appStore.createConversation">
             >New Chat</v-btn>
        </v-toolbar>
        <v-divider></v-divider>
        <v-list dense>
            <v-list-item v-for="item in appStore.conversations" :key="item.title" link>
                <v-list-item-icon>
                    <v-icon>{{ item.icon }}</v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                    <v-list-item-title>{{ item.title }}</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>

<script setup>
import { onMounted } from 'vue';
import { drawer } from './AppDrawer'
import { useAppStore } from '@/stores/app'

const appStore = useAppStore()

onMounted(appStore.refreshConversations)
</script>