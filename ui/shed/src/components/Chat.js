import { ref, reactive } from 'vue';

export const chatInput = ref("");
export const chatMessages = reactive([]);
export const chatUsers = reactive([]);
export const currentChatUser = ref(null);
export const chatBoxActiveStatus = ref(false);
export const timeStamps = ref(true);