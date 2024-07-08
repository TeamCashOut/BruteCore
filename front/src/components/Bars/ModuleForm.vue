<template>
  <div class="relative overflow-x-auto shadow-md sm:rounded-lg pt-4">
    <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
        <tr>
          <th scope="col" class="px-6 py-3">ID</th>
          <th scope="col" class="px-6 py-3">Title</th>
          <th scope="col" class="px-6 py-3">Time of creation</th>
          <th scope="col" class="px-6 py-3">Author</th>
          <th scope="col" class="px-6 py-3">Data Type</th>
          <th scope="col" class="px-6 py-3">Version</th>
          <th scope="col" class="px-6 py-3">
            <a @click.prevent="add_file = true" href="#" type="button" data-modal-target="editUserModal" data-modal-show="editUserModal" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Добавить</a>
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in modules " :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600" >
          <td class="px-6 py-4">{{ item.id }}</td>  
          <td class="px-6 py-4 text-gray-900 text-base fort-semibold whitespace-nowrap dark:text-white">{{ item.name }}</td>
          <td class="px-6 py-4">{{ item.create_time }}</td>
          <td class="px-6 py-4">{{ item.author }}</td>
          <td class="px-6 py-4">{{ item.data_type }}</td>
          <td class="px-6 py-4">{{ item.version }}</td>
          <td class="px-6 py-4">
            <a @click.prevent="DeleteModule(item.id)" href="#" class=" font-medium text-red-600 dark:text-red-500 hover:underline">Удалить модуль</a>
          </td>
        </tr>
      </tbody>
    </table>
  </div>

  <!-- Main modal -->
  <div v-if="add_file" @click="add_file = false" class="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-40" aria-hidden="true"></div>
  <!-- Modal Insert Module -->
  <div v-if="add_file" id="folder-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-md max-h-full">
    <div class="relative p-4 w-full max-w-md max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Добавление модуля</h3>
          <button @click="add_file = false" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
            </svg>
            <span class="sr-only">Close modal</span>
          </button>
        </div>
        <!-- Modal body -->
        <div class="p-4 md:p-5">
          <form class="space-y-4" action="#">
            <input @change="handleFileChange" class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" id="file_input" type="file">            
            <button @click="uploadFile"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Добавить модуль</button>              
            <div v-if="uploading">
              <p>Прогресс: {{ progress }}%</p>
              <progress :value="progress" max="100"></progress>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
  
  <!-- AlertForm --> 
  <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />
</template>

<script>

import axios from 'axios'
import AlertForm from '../Common/AlertForm.vue'

export default {
  components: {
    AlertForm,
  },
  inject: ['server_addr'],
  data() {
    return {
      msg_txt: null,
      msg_id: null,
      modules: null,
      
      add_file: false,
      selectedFile: null,
      uploading: false,
      progress: 0
    }
  },
  async mounted() {
    await this.mount()
  },
  methods: {
    async mount() {
      const response = await this.HTTP('GET', '/api/modl/GetModules', null)
      if (response.success) {
        this.modules = response.modules
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
    async DeleteModule(index) {
      const response = await this.HTTP('GET', '/api/modl/DeleteModule/?id='+index, null)
      if (response.success) {
        await this.mount()
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }      
    },
    // New Beginning
    handleFileChange(event) {
      this.selectedFile = event.target.files[0];
    },
    uploadFile() {
      if (!this.selectedFile) {
        this.msg_txt = 'Sele​​​​​​ct file to upload' 
        this.msg_id = 0
        return
      }

      this.uploading = true;
      const formData = new FormData();
      formData.append("file", this.selectedFile)
      axios.post(this.server_addr+'/api/modl/UploadModule', formData, {
        headers:{
          'Authorization': localStorage.accessToken,
          "Content-Type": "multipart/form-data",
        },
        onUploadProgress: (progressEvent) => {
          this.progress = Math.round(
            (progressEvent.loaded * 100) / progressEvent.total
          );
        }
      })
      .then(async (response) => {
        if (response.success) {
          this.msg_txt = 'The module was loaded successfully.'
          this.msg_id = 1
          await this.mount()
        } else {
          this.msg_txt = response.msg_txt
          this.msg_id = 0
        }
      })
      .catch((error) => {
        this.msg_txt = error.response.data.msg_txt
        this.msg_id = 0
      })
      .finally(() => {
        this.add_file = false
        this.selectedFile = null
        this.uploading = false
        this.progress = 0
      });
    }
  }
}
</script>
