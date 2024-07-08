<script setup>
  import { useDark, useToggle } from '@vueuse/core';

  const isDark = useDark();
  const toggleDark = useToggle(isDark);
</script>

<template>
  <button data-drawer-target="sidebar-multi-level-sidebar" data-drawer-toggle="sidebar-multi-level-sidebar" aria-controls="sidebar-multi-level-sidebar" type="button" class="inline-flex items-center p-2 mt-2 ms-3 text-sm text-gray-500 rounded-lg sm:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
    <span class="sr-only">Open sidebar</span>
    <svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
      <path clip-rule="evenodd" fill-rule="evenodd" d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"></path>
    </svg>
  </button>

  <aside id="sidebar-multi-level-sidebar" class="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0" aria-label="Sidebar">
    <div class="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
        <ul class="space-y-2 font-medium">
          <li>
            <a href="#" class="flex flex-col items-center justify-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
              <div class="mb-2 flex justify-center items-center">
                <img :src="require('@/assets/1705848093544.png')" style="width: 50%; height: 50%;">
              </div>
              <span class="whitespace-nowrap text-xl">BruteCore</span>
            </a>
          </li>
          <li @click="ChoseBar(2)">
            <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
              <svg class="w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 20">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 4c0 1.657-3.582 3-8 3S1 5.657 1 4m16 0c0-1.657-3.582-3-8-3S1 2.343 1 4m16 0v6M1 4v6m0 0c0 1.657 3.582 3 8 3s8-1.343 8-3M1 10v6c0 1.657 3.582 3 8 3s8-1.343 8-3v-6"/>
              </svg>
              <span class="flex-1 ms-3 whitespace-nowrap">Combolist</span>
            </a>
          </li>
          <li @click="ChoseBar(3)">
              <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">                
                <svg class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 14H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v1M5 19h5m-9-9h5m4-4h8a1 1 0 0 1 1 1v12H9V7a1 1 0 0 1 1-1Zm6 8a2 2 0 1 1-4 0 2 2 0 0 1 4 0Z"/>
                </svg>
                <span class="flex-1 ms-3 whitespace-nowrap">Proxy</span>
              </a>
          </li>
          <li @click="ChoseBar(4)">
            <button type="button" class="flex items-center w-full p-2 text-base text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700" aria-controls="dropdown-example" data-collapse-toggle="dropdown-example">
              <svg class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 6h8M6 6v8m0-8V3.5A2.5 2.5 0 1 0 3.5 6H6Zm8 0v8m0-8h2.5A2.5 2.5 0 1 0 14 3.5V6Zm0 8H6m8 0h2.5a2.5 2.5 0 1 1-2.5 2.5V14Zm-8 0H3.5A2.5 2.5 0 1 0 6 16.5V14Z"/>
              </svg>
              <span class="flex-1 ms-3 text-left rtl:text-right whitespace-nowrap">Modules</span>                
            </button>              
          </li>
          <li @click="ChoseBar(5)">
            <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
              <svg class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="none" viewBox="0 0 20 20">
                <path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M1 10c1.5 1.5 5.25 3 9 3s7.5-1.5 9-3m-9-1h.01M2 19h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1ZM14 5V3a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v2h8Z"/>
              </svg>
              <span class="flex-1 ms-3 whitespace-nowrap">Sessions</span>                
            </a>
          </li>
          <li @click="toggleDark()">
            <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
              <svg v-if="isDark" class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3V1m0 18v-2M5.05 5.05 3.636 3.636m12.728 12.728L14.95 14.95M3 10H1m18 0h-2M5.05 14.95l-1.414 1.414M16.364 3.636 14.95 5.05M14 10a4 4 0 1 1-8 0 4 4 0 0 1 8 0Z"/>
              </svg>
              <svg v-else class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
                <path d="M17.293 13.293A8 8 0 016.707 2.707a8.001 8.001 0 1010.586 10.586z"/>
              </svg>
              <span class="flex-1 ms-3 whitespace-nowrap">Change Topic</span>
            </a>
          </li>
          <li @click="Logout()">
            <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
              <svg class="flex-shrink-0 w-5 h-5 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 16">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 8h11m0 0L8 4m4 4-4 4m4-11h3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-3"/>
              </svg>
              <span class="flex-1 ms-3 whitespace-nowrap">Exit</span>
            </a>
          </li>          
        </ul>

        <div class="fixed items-center bottom-1  p-2 mb-3 w-auto text-red-800 rounded-lg bg-red-50 dark:bg-gray-800 dark:text-red-400">
        <ul class="space-y-2 font-medium flex-grow">
          <li>
            <a href="#" class="cursor-default flex items-center rounded-lg dark:text-white">
              <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 21 21">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16.5A2.493 2.493 0 0 1 6.51 18H6.5a2.468 2.468 0 0 1-2.4-3.154 2.98 2.98 0 0 1-.85-5.274 2.468 2.468 0 0 1 .921-3.182 2.477 2.477 0 0 1 1.875-3.344 2.5 2.5 0 0 1 3.41-1.856A2.5 2.5 0 0 1 11 3.5m0 13v-13m0 13a2.492 2.492 0 0 0 4.49 1.5h.01a2.467 2.467 0 0 0 2.403-3.154 2.98 2.98 0 0 0 .847-5.274 2.468 2.468 0 0 0-.921-3.182 2.479 2.479 0 0 0-1.875-3.344A2.5 2.5 0 0 0 13.5 1 2.5 2.5 0 0 0 11 3.5m-8 5a2.5 2.5 0 0 1 3.48-2.3m-.28 8.551a3 3 0 0 1-2.953-5.185M19 8.5a2.5 2.5 0 0 0-3.481-2.3m.28 8.551a3 3 0 0 0 2.954-5.185"/>
              </svg>
              <span class="flex ms-3 whitespace-nowrap">CPU: {{ CPU }}</span>
            </a>
          </li>
          <li>
            <a href="#" class="cursor-default flex items-center rounded-lg dark:text-white">
              <svg class="w-5 h-5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 2a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1M1 9h14M2 5h12a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1Z"/>
              </svg>
              <span class="flex ms-3 whitespace-nowrap">RAM: {{ RAM }}</span>
            </a>
          </li>
          <li class ="cursor-default">
            <a href="#" class="cursor-default flex items-center rounded-lg dark:text-white">
              <svg class="w-5 h-5
              " aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 19 20">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 9.376v.786l8 3.925 8-3.925v-.786M1.994 14.191v.786l8 3.925 8-3.925v-.786M10 1.422 2 5.347l8 3.925 8-3.925-8-3.925Z"/>
              </svg>
                <span class="flex ms-3 whitespace-nowrap">DSK: {{ Disk }}</span>
              </a>
          </li>    
        </ul>
        </div>
    </div>
    
  </aside>

  <div class="p-4 sm:ml-64">
    <h1 class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">{{title}}</h1>
    <hr class="border border-gray-500 mt-3">
    <CombolistForm v-if="layer == 2" />
    <ProxyconfigForm v-if="layer == 3" />
    <ModulesForm v-if="layer == 4" />
    <SessionForm v-if="layer == 5" />    
  </div>
</template>

<script>

import axios from 'axios'
import CombolistForm from './Bars/CombolistForm.vue'
import ProxyconfigForm from './Bars/ProxyconfigForm.vue'
import ModulesForm from './Bars/ModuleForm.vue'
import SessionForm from './Bars/SessionForm.vue'

export default {
  Components: {
    CombolistForm,
    ProxyconfigForm,
    ModulesForm,
    SessionForm
  },
  inject: ['server_addr'],
  data() {
    return {
      msg_txt: null,
      title: 'Sessions',
      layer: 5,

      RAM: '0',
      CPU: '0',
      Disk: '0'
    };
  },
  async mounted() {
    const accessToken = localStorage.getItem('accessToken')
    const refreshToken = localStorage.getItem('refreshToken')
    
    if (!accessToken || !refreshToken) {
      this.GoToLogin()
    }

    try {
      const response = await axios.post(this.server_addr+'/api/auth/Validate', {}, {
        headers : {
          'Authorization': localStorage.accessToken,
          'Content-Type': 'application/json'
        }
      })

      if (!response.data.success) {
        this.GoToLogin()
      } else {
        let arr = localStorage.accessToken.split('.')
        let exp = JSON.parse(decodeURIComponent(escape(atob(arr[1])))).exp
        let cur = Math.floor(new Date().getTime() / 1000);
        if (exp - cur < 10800) {
          const response = await axios.post(this.server_addr+'/api/auth/Refresh', {
            "refreshToken": localStorage.refreshToken
          }, {
            headers : {
              'Authorization': localStorage.accessToken,
              'Content-Type': 'application/json'
            }
          })

          if (response.data.success) {
            localStorage.setItem('accessToken', response.data.accessToken)
            localStorage.setItem('refreshToken', response.data.refreshToken)
            this.$router.push('/Dashboard')
          } else {
            this.GoToLogin()
          }
        }
      }
    } catch(err) {
      this.GoToLogin()
    }

    this.intervalId = setInterval(async () => {
      const response = await axios.get(this.server_addr+'/Health')
      if (response.data.success) {
        this.RAM = response.data.ram
        this.CPU = response.data.cpu
        this.Disk = response.data.disk
      }
    }, 5000);
  },
  beforeUnmount() {
    clearInterval(this.intervalId); // Clear the timer when the component is destroyed
  },
  methods: {    
    GoToLogin() {
      localStorage.removeItem('accessToken')
      localStorage.removeItem('refreshToken')
      this.$router.push('/Login')
    },
    ChoseBar(a) {
      this.layer = a
      switch(a) {
        case 1:
          case 2:
          this.title = 'Combo Sheets'
          break
          case 3:
          this.title = 'Proxy Presets'
          break
          case 4:
          this.title = 'Modules'
          break
          case 5:
          this.title = 'Sessions'
          break
      }
    },
    async Logout() {
      try {    
        const response = await axios.post(this.server_addr+'/api/auth/Logout', {}, {
          headers : {
            'Authorization': localStorage.accessToken,
            'Content-Type': 'application/json'
          }
        })

        if (response.data.success) {
          this.GoToLogin()
        } else {
          this.msg_txt = response.data.msg_txt          
        }
      } catch(err) {
        if (err.response.data.msg_txt !== undefined) {
          this.msg_txt = err.response.data.msg_txt
        } else {
          this.msg_txt = err.message
        }
      }
    },
  }
};
</script>
