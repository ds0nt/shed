<template>
    <v-app-bar app dark>
        <v-app-bar-nav-icon @click="toggleDrawer"></v-app-bar-nav-icon>
        <v-toolbar-title>Shed</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-btn icon>
            <v-icon>mdi-account-circle</v-icon>
        </v-btn>
    </v-app-bar>
</template>

<script setup>
import { toggleDrawer } from './AppDrawer'
</script>