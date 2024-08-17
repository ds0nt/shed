import axios from 'axios';

const API = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    timeout: 5000, // Set a timeout value if needed
});

API.interceptors.response.use(
    function (response) {
        // If the request succeeds, you can use it directly.
        return response;
    },
    function (error) {
        // Any status codes that fall outside the range of 2xx cause this function to trigger
        // You can handle based on your own business
        console.error('Error fetching data:', error);
        throw error;
    }
);

import { ref } from 'vue';
export function useAsyncFn(asyncFn) {
    const data = ref(null);
    const error = ref(null);
    const loading = ref(false);

    const execute = async (...args) => {
        loading.value = true;
        error.value = null;

        try {
            data.value = await asyncFn(...args);
        } catch (err) {
            error.value = err;
        } finally {
            loading.value = false;
        }
    };

    return { data, error, loading, execute };
}

// Example async function to fetch data
export const fetchChats = async () => {
    try {
        const response = await API.get('/chats'); // Replace with your API endpoint
        return response.data;
    } catch (error) {
        console.error('Error fetching data:', error);
        throw error;
    }
};
export const useFetchChats = () => useAsyncFn(fetchChats);

export default API;