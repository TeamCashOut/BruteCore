<template>
  <div v-if="SelectedSessionID === null" class="relative overflow-x-auto shadow-md sm:rounded-lg pt-4">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
              <th scope="col" class="px-6 py-3">ID</th>
              <th scope="col" class="px-6 py-3">Title</th>
              <th scope="col" class="px-6 py-3">Created Time</th>
              <th scope="col" class="px-6 py-3">Status</th>
              <th scope="col" class="px-6 py-3">
                <a @click.prevent="AddSession" href="#" type="button" data-modal-target="editUserModal" data-modal-show="editUserModal" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Добавить</a>
              </th>
              <th scope="col" class="px-6 py-3">Go to</th>
          </tr>
      </thead>
      <tbody>
        <tr v-for="item in sessions " :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600" >
          <td class="px-6 py-4">{{ item.id }}</td>  
          <td class="px-6 py-4 text-gray-900 text-base fort-semibold whitespace-nowrap dark:text-white">{{ item.name }}</td>
          <td class="px-6 py-4">{{ item.create_time }}</td>
          <td class="px-6 py-4">
            <div v-if="item.status == 11" class="flex items-center">
              <div class="h-2.5 w-2.5 rounded-full bg-gray-600 me-2" />Created
            </div>
            <div v-if="item.status == 22" class="flex items-center">
              <div class="h-2.5 w-2.5 rounded-full bg-green-600 me-2" />In progress
            </div>
            <div v-if="item.status == 23" class="flex items-center">
              <div class="h-2.5 w-2.5 rounded-full bg-blue-600 me-2" />Completed
            </div>
            <div v-if="item.status == 24" class="flex items-center">
              <div class="h-2.5 w-2.5 rounded-full bg-orange-600 me-2" />Stopped
            </div>
            <div v-if="item.status == 13" class="flex items-center">
              <div class="h-2.5 w-2.5 rounded-full bg-red-900 me-2" />Error
            </div>
          </td>

          <td class="px-6 py-4">
            <a @click.prevent="DeleteSession(item.id)" href="#" class=" font-medium text-red-600 dark:text-red-500 hover:underline">Delete Session</a>
          </td>
          <td class="px-4 py-4">
            <button @click.prevent="SelectedSessionID = item.id" type="button" class="w-16 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
              <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 8 14">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 13 5.7-5.326a.909.909 0 0 0 0-1.348L1 1"/>
              </svg>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
  
  <!-- AlertForm --> 
  <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />

  <!-- DetailedSessionForm -->
  <DetailedSessionForm :session_id="SelectedSessionID" @nullate_session_id="SelectedSessionID = null; mount()" />
  
</template>

<script>

import axios from 'axios'
import AlertForm from '../Common/AlertForm.vue'
import DetailedSessionForm from '../Common/DetailedSessionForm.vue'

export default {
  components: {
    AlertForm,
    DetailedSessionForm
  },
  inject: ['server_addr'],
  data() {
    return {
      msg_txt: null,
      msg_id: null,
      sessions: null,
      SelectedSessionID: null,
    }
  },
  async mounted() {
    await this.mount()
  },
  methods: {
    async mount() {
      const response = await this.HTTP('GET', '/api/sess/GetSessions', null)
      if (response.success) {
        this.sessions = response.sessions
      } else {
        this.msg_txt = response.msg_txt 
        this.msg_id = 0
      } 
    },
    async HTTP(method, addr, body) {
      const accessToken = localStorage.getItem('accessToken')
      if (!accessToken) {
        this.GoToLogin()
      }

      try {
        let headobj = {
          headers: {
            'Authorization': localStorage.accessToken,
            'Content-Type': 'application/json'
          }
        }
        let response
        if (method === 'get' || method === 'GET') {
          response = await axios.get(this.server_addr+addr, headobj)
        } else {
          response = await axios.post(this.server_addr+addr, body, headobj)
        }       
        return response.data
      } catch(err) {
        if (err.response.data.msg_txt !== undefined) {
          return {
            success: false,
            msg_txt: err.response.data.msg_txt,
          }
        } else {
          return {
            success: false,
            msg_txt: err.response.data.msg_txt,
          }
        }        
      }
    },
    async AddSession() {
      const response = await this.HTTP('POST', '/api/sess/CreateSession', null)
      if (response.success) {
        await this.mount()
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
    },
    async DeleteSession(index) {
      const response = await this.HTTP('GET', '/api/sess/DeleteSession/?id='+index, null)
      if (response.success) {
        await this.mount()
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }      
    }
  }
}

</script>