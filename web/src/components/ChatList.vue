<template>
  <div class="flex flex-col h-full overflow-y-auto">
    <div class="px-4 py-3">
      <h2 class="text-lg font-semibold text-gray-900">Chats</h2>
    </div>
    <ul class="px-4 py-2 space-y-2">

      <li v-for="chat in chats" :key="chat.id" class="flex items-center space-x-3" @click="onChatClick(chat)">
        <img :src="chat.avatar" alt="avatar" class="w-10 h-10 rounded-full">
        <div class="flex-1 min-w-0">
          <p class="text-sm font-medium text-gray-900 truncate">{{ chat.name }}</p>
          <p class="text-sm text-gray-500 truncate">{{ chat.lastMessage }}</p>
        </div>
        <time datetime="2022-01-01" class="text-sm text-gray-500">{{ chat.time }}</time>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import {onMounted, ref} from 'vue';
import {useRouter} from "vue-router";

export default {
  setup() {

    const router = useRouter();

    const chats = ref([
      {
        id: "65d4dabb26ca83160d86ffa4",
        name: 'John Doe',
        avatar: 'https://example.com/avatar.jpg',
        lastMessage: 'Hello, how are you?',
        time: '10:00 AM'
      },
      // Add more chat data here
    ]);


    onMounted(() => {
      // Fetch chat data from the server
    });

    const onChatClick = (chat: any) => {

      router.push(`/chat-list/${chat.id}`);

      // Navigate to the chat room
    };

    return { chats, onChatClick };
  },
};
</script>