import { createApp } from 'vue'
import App from './App.vue'
import './assets/tailwind.css'
import router from './router'
import contextmenu from "v-contextmenu";
import "v-contextmenu/dist/themes/default.css";

const app = createApp(App)

app.use(router)
app.use(contextmenu)
app.provide('server_addr', 'http://localhost:985')

app.mount('#app')
