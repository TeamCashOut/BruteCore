<script setup>
  import { useDark, useToggle } from '@vueuse/core';

  const isDark = useDark();
  const toggleDark = useToggle(isDark);
</script>

<template>
  <!-- Mode Switcher -->  
  <div class="max-w-screen-xl flex-wrap justify-between mx-auto flex items-center bg-white-800 p-2.5 m-1 fixed top-2 right-2 rounded-lg shadow dark:border dark:bg-gray-800 dark:border-gray-700">
    <button @click="toggleDark()">
      <svg v-if="isDark" class="h-6 w-6" viewBox="0 0 20 20" fill="currentColor">
        <path fill="#FFFFFF" fill-rule="evenodd" d="M10 2a1 1 0 011 1v1a1 1 0 11-2 0V3a1 1 0 011-1zm4 8a4 4 0 11-8 0 4 4 0 018 0zm-.464 4.95l.707.707a1 1 0 001.414-1.414l-.707-.707a1 1 0 00-1.414 1.414zm2.12-10.607a1 1 0 010 1.414l-.706.707a1 1 0 11-1.414-1.414l.707-.707a1 1 0 011.414 0zM17 11a1 1 0 100-2h-1a1 1 0 100 2h1zm-7 4a1 1 0 011 1v1a1 1 0 11-2 0v-1a1 1 0 011-1zM5.05 6.464A1 1 0 106.465 5.05l-.708-.707a1 1 0 00-1.414 1.414l.707.707zm1.414 8.486l-.707.707a1 1 0 01-1.414-1.414l.707-.707a1 1 0 011.414 1.414zM4 11a1 1 0 100-2H3a1 1 0 000 2h1z" clip-rule="evenodd"/>          
      </svg>
      <svg v-else class="h-6 w-6" viewBox="0 0 20 20" fill="currentColor">          
        <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"/>
      </svg>
    </button>
  </div>
  
  <!-- LoginForm -->  
  <div class="flex flex-col items-center justify-center px-6 py-40">
    <div class="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
      <div class="p-6 space-y-4 md:space-y-6 sm:p-8">
        <h1 class="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
          Sign in to your account
        </h1>
        <form class="space-y-4 md:space-y-6" action="#">
          <div>
            <label for="login" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Login</label>
            <input v-model="username" type="text" name="login" id="login" placeholder="username" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required="">
          </div>
          <div>
            <label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Password</label>
            <input v-model="password" type="password" name="password" id="password" placeholder="••••••••" class="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required="">
          </div>
          
          <button @click.prevent="Login" type="button" class="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">Sign in</button>                  
        </form>
      </div>
    </div>
  </div>   
  
  <!-- AlertForm --> 
  <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />
  
</template>

<script>

import axios from 'axios';
import AlertForm from './Common/AlertForm.vue'

export default {
  components: { AlertForm },
  inject: ['server_addr'],
  data() {
    return {
      username  : '',
      password: '',
      msg_txt: null,
      msg_id: null, 
    };
  },
  mounted() {
    const accessToken = localStorage.getItem('accessToken');
    const refreshToken = localStorage.getItem('refreshToken');

    if (accessToken && refreshToken) {
      this.$router.push('/Dashboard');
    }
  },
  methods: {    
    async Login() {
      this.msg_txt = null
      this.msg_id = null
      try {
        const response = await axios.post(this.server_addr+'/api/auth/Login', {
          username: this.username,
          password: this.password
        })

        if (response.data.success) {
          localStorage.setItem('accessToken', response.data.accessToken)
          localStorage.setItem('refreshToken', response.data.refreshToken)
          this.$router.push('/Dashboard')
        } else {
          this.msg_txt = response.data.msg_txt   
          this.msg_id = 0       
        }

      } catch(err) {
        if (err.response.data.msg_txt !== undefined) {
          this.msg_txt = err.response.data.msg_txt
          this.msg_id = 0   
        } else {
          this.msg_txt = err.message
          this.msg_id = 0   
        }
      }
    },
  },  
};
</script>