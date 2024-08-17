import { ref } from 'vue';

export const drawer = ref(false);
export const items = ref([
    { title: 'Chat 1', icon: 'mdi-chat' },
    { title: 'Chat 2', icon: 'mdi-chat' },
    { title: 'Chat 3', icon: 'mdi-chat' },
]);

export function toggleDrawer() {
    drawer.value = !drawer.value;
}