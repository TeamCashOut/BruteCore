<template>

  <v-contextmenu ref="contextmenu">
    <v-contextmenu-item :auto-hide="true" @click="createPreset">Create a preset</v-contextmenu-item>
  </v-contextmenu>
  
  <div class="border border-gray-500 mt-3 dark:caret-lime-800 w-auto h-96 overflow-y-auto flex flex-col"> 
    <ul>        
      <li @click="selectPreset(preset.id)" class="space-y-2 font-medium" v-for="preset in presets" :key="preset" ref="li" :id="preset.id">
        <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
          <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" viewBox="0 0 20 18" fill="none" xmlns="http://www.w3.org/2000/svg">
            <path d="M6.5 8C5.80777 8 5.13108 7.79473 4.55551 7.41015C3.97993 7.02556 3.53133 6.47893 3.26642 5.83939C3.00152 5.19985 2.9322 4.49612 3.06725 3.81719C3.2023 3.13825 3.53564 2.51461 4.02513 2.02513C4.51461 1.53564 5.13825 1.2023 5.81719 1.06725C6.49612 0.932205 7.19985 1.00152 7.83939 1.26642C8.47893 1.53133 9.02556 1.97993 9.41015 2.55551C9.79473 3.13108 10 3.80777 10 4.5" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <path d="M6.5 17H1V15C1 13.9391 1.42143 12.9217 2.17157 12.1716C2.92172 11.4214 3.93913 11 5 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
            <path d="M19.5 11H18.38C18.2672 10.5081 18.0714 10.0391 17.801 9.613L18.601 8.818C18.6947 8.72424 18.7474 8.59708 18.7474 8.4645C18.7474 8.33192 18.6947 8.20476 18.601 8.111L17.894 7.404C17.8002 7.31026 17.6731 7.25761 17.5405 7.25761C17.4079 7.25761 17.2808 7.31026 17.187 7.404L16.392 8.204C15.9647 7.93136 15.4939 7.73384 15 7.62V6.5C15 6.36739 14.9473 6.24021 14.8536 6.14645C14.7598 6.05268 14.6326 6 14.5 6H13.5C13.3674 6 13.2402 6.05268 13.1464 6.14645C13.0527 6.24021 13 6.36739 13 6.5V7.62C12.5081 7.73283 12.0391 7.92863 11.613 8.199L10.818 7.404C10.7242 7.31026 10.5971 7.25761 10.4645 7.25761C10.3319 7.25761 10.2048 7.31026 10.111 7.404L9.404 8.111C9.31026 8.20476 9.25761 8.33192 9.25761 8.4645C9.25761 8.59708 9.31026 8.72424 9.404 8.818L10.204 9.618C9.9324 10.0422 9.73492 10.5096 9.62 11H8.5C8.36739 11 8.24021 11.0527 8.14645 11.1464C8.05268 11.2402 8 11.3674 8 11.5V12.5C8 12.6326 8.05268 12.7598 8.14645 12.8536C8.24021 12.9473 8.36739 13 8.5 13H9.62C9.73283 13.4919 9.92863 13.9609 10.199 14.387L9.404 15.182C9.31026 15.2758 9.25761 15.4029 9.25761 15.5355C9.25761 15.6681 9.31026 15.7952 9.404 15.889L10.111 16.596C10.2048 16.6897 10.3319 16.7424 10.4645 16.7424C10.5971 16.7424 10.7242 16.6897 10.818 16.596L11.618 15.796C12.0422 16.0676 12.5096 16.2651 13 16.38V17.5C13 17.6326 13.0527 17.7598 13.1464 17.8536C13.2402 17.9473 13.3674 18 13.5 18H14.5C14.6326 18 14.7598 17.9473 14.8536 17.8536C14.9473 17.7598 15 17.6326 15 17.5V16.38C15.4919 16.2672 15.9609 16.0714 16.387 15.801L17.182 16.601C17.2758 16.6947 17.4029 16.7474 17.5355 16.7474C17.6681 16.7474 17.7952 16.6947 17.889 16.601L18.596 15.894C18.6897 15.8002 18.7424 15.6731 18.7424 15.5405C18.7424 15.4079 18.6897 15.2808 18.596 15.187L17.796 14.392C18.0686 13.9647 18.2662 13.4939 18.38 13H19.5C19.6326 13 19.7598 12.9473 19.8536 12.8536C19.9473 12.7598 20 12.6326 20 12.5V11.5C20 11.3674 19.9473 11.2402 19.8536 11.1464C19.7598 11.0527 19.6326 11 19.5 11ZM14 14.5C13.5055 14.5 13.0222 14.3534 12.6111 14.0787C12.2 13.804 11.8795 13.4135 11.6903 12.9567C11.5011 12.4999 11.4516 11.9972 11.548 11.5123C11.6445 11.0273 11.8826 10.5819 12.2322 10.2322C12.5819 9.8826 13.0273 9.6445 13.5123 9.54804C13.9972 9.45157 14.4999 9.50108 14.9567 9.6903C15.4135 9.87952 15.804 10.2 16.0787 10.6111C16.3534 11.0222 16.5 11.5055 16.5 12C16.5 12.663 16.2366 13.2989 15.7678 13.7678C15.2989 14.2366 14.663 14.5 14 14.5Z" fill="currentColor"/>
          </svg>

          <span class="ms-5">{{ preset.name }}</span>
          <a @click="OPRDeleteItem(preset.id)" href="#" class="ml-auto py-1 font-medium text-red-600 dark:text-red-500 hover:underline">Удалить пресет</a>
          <span class="ml-auto py-1">{{ preset.create_time }}</span>
        </a>          
      </li>    
    </ul>
    <div v-contextmenu:contextmenu class="flex-grow" />
  </div>

  <!-- Info about a specific preset -->
  <div class="border border-gray-500 mt-3 dark:caret-lime-800 w-auto h-96 overflow-y-auto hover:bg-gray-100 dark:hover:bg-gray-700 p-2 space-y-2 font-medium text-gray-900 dark:text-white">        
    <div v-if="selectedObj != null" class="flex">      
      <!-- Left side -->
      <div class="w-1/3">
        <h1 class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Options</h1>
        <table>
          <tbody>
            <tr>
              <td>Serial number in the database:</td>
              <td>{{ selectedObj.info.id }}</td>
            </tr>
            <tr>
              <td>Name:</td>
              <td>{{ selectedObj.info.name }}</td>
            </tr>
            <tr>
              <td>Creation time:</td>
              <td>{{ selectedObj.info.create_time }}</td>
            </tr>
            <tr>
              <td>Auto update:</td>
              <td v-if="selectedObj.info.use_update === false">Нет</td>
              <td v-if="selectedObj.info.use_update === true">Да</td>
            </tr>
            <tr v-if="selectedObj.info.use_update === true">
              <td>Update interval (seconds):</td>
              <td>{{ selectedObj.info.interval }}</td>
            </tr>            
            <tr>
              <td>Timeout(milliseconds):</td>
              <td>{{ selectedObj.info.timeout }}</td>
            </tr>
            <tr>
              <td>Total number of lines:</td>
              <td>{{ selectedObj.info.lines_count }}</td>
            </tr>
          </tbody>
        </table>          
      </div>

      <!-- Right side -->
      <div v-if="selectedObj.info.links !== null" class="w-2/3 relative overflow-x-auto">
        <h1 class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Источники</h1>
        <table>
          <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th class="px-6 py-1">ID</th>
              <th class="px-6 py-1">Статус</th>
              <th class="px-6 py-1">Source Type</th>
              <th class="px-6 py-1">Тип прокси</th>
              <th class="px-6 py-1">Кол-во строк</th>
              <th class="px-6 py-1">Путь</th>
              <th class="px-6 py-1">Ошибка</th>
            </tr>
          </thead>
          <tbody class="text-gray-900 dark:text-white">
            <tr v-for="item in selectedObj.info.links" :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
              <td class="px-6 py-1 font-medium">{{ item.id }}</td>              
              <td class="px-6 py-1">{{ item.status }}</td>
              <td class="px-6 py-1">{{ item.source_type }}</td>
              <td class="px-6 py-1">{{ item.proxy_type }}</td>
              <td class="px-6 py-1">{{ item.count }}</td>
              <td class="px-6 py-1">{{ item.path }}</td>
              <td class="px-6 py-1">{{ item.err_txt }}</td>              
            </tr>
          </tbody>
        </table>
      </div>
    </div>  
  </div>

  <!-- Main modal -->
  <div v-if="insertPreset != null" @click="insertPreset = null" class="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-40" aria-hidden="true"></div>
  <!-- Modal Insert Preset -->
  <div v-if="insertPreset != null" id="folder-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-2xl max-h-full">
    <div class="relative p-4 w-full max-w-2xl max-h-full mx-auto">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Добавление пресета</h3>
          <button @click="insertPreset = null" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
            </svg>
            <span class="sr-only">Close modal</span>
          </button>
        </div>

        <!-- Modal Body -->
        <form class="p-4 md:p-5">
          <div class="grid gap-4 mb-4 grid-cols-3">
            
            <!-- Title -->
            <div class="col-span-3 sm:col-span-1">
              <label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Наименование</label>
              <input v-model="insertPreset.name" type="text" name="name" id="name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Введите наименование">
            </div>

            <!-- Interval -->
            <div class="col-span-3 sm:col-span-1">
              <div class="flex items-center">
                <input v-model="insertPreset.use_update" id="disabled-checked-checkbox" type="checkbox" value="" class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600">
                <label for="disabled-checked-checkbox" class="ms-2 text-sm font-medium text-gray-900 dark:text-white">Интервал (секунды)</label>
              </div>
              <label for="interval" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"></label>
              <input v-model="insertPreset.interval" type="number" name="interval" id="interval" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
            </div>
            
            <!-- Timeout -->
            <div class="col-span-3 sm:col-span-1">
              <label for="timeout" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Таймаут (милесекунды)</label>
              <input v-model="insertPreset.timeout" type="number" name="timeout" id="timeout" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
            </div>
            
            <!-- Source table -->
            <div class="col-span-3">
              <label for="description" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Источники</label>
              <table class="w-full">
                <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                  <tr>
                    <th scope="col" class="px-2 py-1">Тип источника</th>
                    <th scope="col" class="px-2 py-1">Тип прокси</th>
                    <th scope="col" class="px-2 py-1">Путь</th>
                    <th scope="col" class="px-2 py-1"><a @click.prevent="itemAction('add')" href="#" type="button" data-modal-target="editUserModal" data-modal-show="editUserModal" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Добавить</a></th>
                  </tr>
                </thead>
                <tbody>                  
                  <tr v-for="item in insertPreset.links" :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <td class="px-2 py-1">
                      <select v-model="item.source_type" id="category" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">                        
                        <option :value="Number(15)">Local file</option>
                        <option :value="Number(16)">Remote file</option>
                        <option :value="Number(17)">Ссылка</option>
                      </select>
                    </td>   
                    <td class="px-2 py-1">
                      <select v-model="item.proxy_type" id="category" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
                        <option :value="Number(18)">HTTP/S</option>
                        <option :value="Number(19)">SOCKS4</option>
                        <option :value="Number(20)">SOCKS5</option>
                      </select>
                    </td>
                    <td class="px-2 py-1 dark:text-white">
                      <input v-if="item.source_type == 15" @change="item.path = $event.target.files[0]" class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" id="file_input" type="file">            
                      <input v-if="item.source_type == 16 || item.source_type == 17" v-model="item.path" type="text" name="name" id="name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Введите источник" required>
                    </td>
                    <td class="px-2 py-1">
                      <a @click.prevent="itemAction('remove', item.id)" href="#" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</a>
                    </td>
                  </tr>
                </tbody>
              </table>              
            </div>
            
          </div>

          <!-- Add button -->
          <button @click.prevent="OPRAddItem"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Создать пресет</button>
        </form>

      </div>
    </div>
  </div>

  <!-- AlertForm --> 
  <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />
</template>

<script>

import axios from 'axios'
import AlertForm from '../Common/AlertForm.vue'

const selectedClassDay = 'bg-gray-100'
const selectedClassDark = 'dark:bg-gray-700'

export default {
  components: {
    AlertForm,
  },
  inject: ['server_addr'],
  data() {
    return {
      msg_txt: null,
      msg_id: null,

      presets: [],
      insertPreset: null,
      selectedObj: null
    }
  },
  async mounted() {
    const response = await this.HTTP('GET', '/api/prox/GetProxyPresets', null)
    if (response.success) {
      this.presets = response.presets
    } else {
      this.msg_txt = response.msg_txt 
      this.msg_id = 0
    } 
  },
  methods: {
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
    createPreset() {
      this.insertPreset = {
        use_update: false,
        timeout: 15000,
        interval: 300,
        links: []
      }
    },
    async OPRDeleteItem(index) {
      const response = await this.HTTP('GET', '/api/prox/DeleteProxyPreset/?id='+index, null)
      if (response.success) {
        this.presets = this.presets.filter(obj => obj.id !== index);
        this.msg_txt = "Element removed"
        this.msg_id = 1  
      } else {
        this.msg_txt = response.msg_txt          
        this.msg_id = 0
      }
    },
    async OPRSelectItem(index) {
      const response = await this.HTTP('GET', '/api/prox/GetInfoProxyPreset/?id='+index, null)
      if (response.success) {
        this.selectedObj = response
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
    },
    selectPreset(id) {
      this.$refs.li.forEach(async (element) => {
        if (id == element.id) {
          element.classList.add(selectedClassDay)
          element.classList.add(selectedClassDark)
          await this.OPRSelectItem(id)
        } else {
          element.classList.remove(selectedClassDay)
          element.classList.remove(selectedClassDark)
        }
      })
    },
    itemAction(action, id) {
      if (action === 'remove') {
        const indexToRemove = this.insertPreset.links.findIndex(obj => obj.id === id)
        if (indexToRemove !== -1) {
          this.insertPreset.links.splice(indexToRemove, 1)
        }        
      }

      if (action === 'add') {
        let maxId = Math.max(...this.insertPreset.links.map(obj => obj.id))
        let preset = {
          id: (maxId == -Infinity) ? 1 : maxId+1,
          source_type: 17,
          proxy_type: 18
        }
        this.insertPreset.links.push(preset)
      }
    },
    async OPRAddItem() {
      this.msg_txt = null
      this.msg_id = null

      if (this.insertPreset.name === null || this.insertPreset.name === '') {
        this.msg_txt = 'Please provide a name'
      }

      if (this.insertPreset.timeout < 1000 || this.insertPreset.timeout == null) {
        this.msg_txt = 'Please specify a timeout'
      }

      if (this.insertPreset.interval < 10 && this.insertPreset.use_update == true) {
        this.msg_txt = 'Specify proxy auto-update interval'
      }

      if (this.insertPreset.links.length === 0) {
        this.msg_txt = 'Please provide at least one download link'
      }

      let readFileSync = async (file) => {
        return new Promise((resolve, reject) => {
          let reader = new FileReader()
          reader.onload = () => resolve(btoa(reader.result))
          reader.onerror = () => reject(new Error('Error reading file'))
          reader.readAsText(file)
        })
      }

      try {
        for (const el of this.insertPreset.links) {
          if (el.proxy_type !== 18 && el.proxy_type !== 19 && el.proxy_type !== 20) {
            throw new Error('Invalid proxy type specified') 
          }

          switch (el.source_type) {
            case 15:
              if (el.path === '' || el.path === undefined) {
                throw new Error('File not specified')
              }
              el.data = await readFileSync(el.path)
              el.path = el.path.name
              break
              case 16:
              if (el.path === '' || el.path === undefined) {
                throw new Error('Source path not specified')
              }
              break
              case 17:
              if (el.path === '' || el.path === undefined) {
                throw new Error('Source reference not specified')
              }
              break
            default:
              throw new Error('Unrecognized source type')
          }
        }
      } catch(e) {
        this.msg_txt = e.message
      }
      
      
      if (this.msg_txt !== null) {
        this.msg_id = 0
        return
      }

      const response = await this.HTTP('POST', '/api/prox/UploadProxyPreset', this.insertPreset)
      if (response.success) {
        this.insertPreset.id = response.id
        this.insertPreset.create_time = response.create_time   
        if (this.presets == null) {
          this.presets = []
        }   
        this.presets.push(this.insertPreset)
        
        this.msg_txt = 'Item added to list'
        this.msg_id = 1
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
      this.insertPreset = null      
    }
  }
}

</script>

