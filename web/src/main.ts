import { createApp } from 'vue'
import './style.css'
import App from './App.vue'

import { library } from '@fortawesome/fontawesome-svg-core';
import { faPaperPlane, faMicrophone } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import ChatList from "./components/ChatList.vue";
import Chat from "./components/Chat.vue";
import Login from "./Login.vue";
import {createRouter, createWebHistory} from "vue-router";

library.add(faPaperPlane, faMicrophone);

const app = createApp(App);


const routes = [
    // { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound },
    {
        path:"/login",
        name:"Login",
        component: Login
    },
    { path: '/chat-list', name: 'ChatList', component: ChatList },
    { path: '/chat-list/:id', name: 'ChatDetail', component: Chat },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});




app.component('font-awesome-icon', FontAwesomeIcon);

app.use(router);

app.mount('#app')
