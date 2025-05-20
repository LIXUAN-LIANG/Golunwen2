import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import {router} from './router'
import axios from "@/axios/index.js";


const app = createApp(App);
//与vue实例关联
app.config.globalProperties.$axios = axios;
app.use(createPinia())
app.use(router)

app.mount('#app')
