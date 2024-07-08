import { createRouter, createWebHistory } from "vue-router"
import LoginForm from "./components/LoginForm.vue"
import DashBoard from "./components/DashBoard.vue"

export default createRouter({
  history: createWebHistory(),
  routes: [
    {path: '/Login', component: LoginForm, alias: '/'},
    {path: '/Dashboard', component: DashBoard},    
    {path: '/:catchAll(.*)', redirect: '/Dashboard'},
  ]
})