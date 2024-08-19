import { useFetchConversations } from '@/api/api';
import { onMounted, reactive, ref } from 'vue';

export let fetchConversationRefs = useFetchConversations()

export const drawer = ref(false);


export function toggleDrawer() {
    drawer.value = !drawer.value;
}
