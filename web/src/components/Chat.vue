<template>
  <div class="flex flex-col h-full bg-gray-100">
    <h1 class="text-2xl font-bold p-4 bg-green-500 text-white">Chat , Me {{user?.email}} </h1>
    <div class="flex-grow overflow-y-auto p-4 space-y-4 bg-gray-200">
      <div v-for="message in messages" :key="message.id" class="flex flex-col space-y-2 max-w-lg mx-auto" :class="{ 'items-end': isMe(message.user_id), 'items-start': !isMe(message.user_id) }">
        <div v-if="!isMe(message.user_id)" class="flex items-end space-x-2">
          <img class="w-8 h-8 rounded-full" :src="message.avatar" alt="Avatar">
          <span class="font-semibold text-gray-700">{{ message.user_id }}</span>
        </div>
        <div v-else class="font-semibold text-green-700">{{ message.user_id }}</div>
        <span class="p-2 rounded-lg shadow" :class="{'bg-green-200': isMe(message.user_id), 'bg-white': !isMe(message.user_id), 'rounded-tl-none': isMe(message.user_id), 'rounded-tr-none': !isMe(message.user_id)}">
          <span v-if="message.type === 'text'">{{ message.content }}</span>
          <img v-else-if="message.type === 'image'" :src="message.content" class="max-w-full h-auto rounded-lg" />
          <video v-else-if="message.type === 'video'" :src="message.content" controls class="max-w-full h-auto rounded-lg" />
          <a v-else :href="message.content" download class="underline text-green-600">Download file</a>
        </span>
        <span class="text-sm text-gray-500">{{ message.timestamp }}</span>
        <button v-if="isMe(message.user_id)" @click="deleteMessage(message.id)" class="text-red-500">Delete</button>
      </div>
    </div>
    <div class="border-t border-gray-300 p-4 bg-green-200 flex items-center">
      <form class="flex-grow space-x-4" @submit.prevent="sendMessage">
        <input v-model="newMessage" type="text" placeholder="Enter a message" class="flex-grow rounded-full border border-gray-300 p-2 sm:text-sm md:text-base lg:text-lg" />
        <input type="file" @change="onFileChange($event)" class="hidden" />
        <label for="file-upload" class="cursor-pointer" >
          <font-awesome-icon icon="paper-plane" class="text-gray-500 hover:text-gray-700 sm:text-sm md:text-base lg:text-lg" />
        </label>
        <button type="submit" class="bg-green-500 text-white rounded-full p-2 hover:bg-green-600 sm:text-sm md:text-base lg:text-lg">
          <font-awesome-icon icon="microphone" />
        </button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {onMounted, onUnmounted, Ref, ref} from 'vue';

export default {
  setup() {
    const messages = ref([
      { id: 1, user_id: '65d4dabb26ca83160d86ffa4', type: 'text', content: 'Hello, how are you?', timestamp: 'timestamp', avatar: 'https://via.placeholder.com/150' },
      { id: 2, user_id: '65d4dabb26ca83160d86ffa5', type: 'text', content: 'I\'m good, thanks! How about you?', timestamp: 'timestamp', avatar: '' },
    ]);

    const newMessage = ref('');
    const newFile: Ref<File | null> = ref(null);
    let socket: WebSocket | null = null;

    let getUser = localStorage.getItem('user');
    let user = getUser ? JSON.parse(getUser) : null;

    let currentUserID = user?.id;

    const isMe = (user_id: string) => {
      return user_id === currentUserID;
    };

    const sendMessage = () => {
      const timestamp = new Date().toUTCString();

      if(!currentUserID) return console.log('User not found');

      if(!newMessage.value && !newFile.value) return console.log('Message or file is required');

      let messageData = {
        id: messages.value.length + 1,
        user_id: currentUserID,
        type: 'text',
        content: newMessage.value,
        timestamp: timestamp , // replace with actual timestamp
        avatar: '',
      }
      if (newMessage.value) {
        newMessage.value = '';
      }

      if (newFile.value) {
        messageData.type = getFileType(newFile.value);
        messageData.content = URL.createObjectURL(newFile.value);
        newFile.value = null;
      }

      messages.value.push(messageData);

      if (socket) {
        socket.send(JSON.stringify({ data: messages.value[messages.value.length - 1] }));
      }
    };

    const deleteMessage = (id: number) => {
      messages.value = messages.value.filter(message => message.id !== id);
    };

    const getFileType = (file: File) => {
      if (file.type.startsWith('image/')) {
        return 'image';
      } else if (file.type.startsWith('video/')) {
        return 'video';
      } else {
        return 'file';
      }
    };

    const onFileChange = (event: Event) => {
      const target = event.target as HTMLInputElement;
      newFile.value = target.files?.[0] || null;
    };

    onMounted(() => {
      socket = new WebSocket('ws://localhost:8080/ws?chat_id=65d4dabb26ca83160d86ffa4');
      socket.onopen = () => {
        console.log('Connected to server');
      };
      socket.onmessage = (event) => {
        const chatMessage = JSON.parse(event.data);
        messages.value.push(chatMessage.data);
      };
      socket.onclose = () => {
        console.log('Disconnected from server');
      };
    });

    onUnmounted(() => {
      if (socket) {
        socket.close();
      }
    });

    return { messages, newMessage, newFile, sendMessage, deleteMessage, onFileChange,isMe, user };
  },
};
</script>
