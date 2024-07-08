<template>
  
  <v-contextmenu @show="parentMethod('show', -1)" ref="contextmenu">
    <v-contextmenu-item :auto-hide="true" @click="parentMethod('insert folder')">Add folder</v-contextmenu-item>
    <v-contextmenu-item :auto-hide="true" @click="parentMethod('insert file')">Add combo-list</v-contextmenu-item>
  </v-contextmenu>
  
  <div class="border border-gray-500 mt-3 dark:caret-lime-800 w-auto h-96 overflow-y-auto flex flex-col"> 
    <ul>          
      <ComboItem class="space-y-2 font-medium" :node="node" v-for="node in tree" :key="node" 
        @call-parent-method="parentMethod" :isShowId="showId" :isSelectId="selectId"/>
    </ul>
    <div v-contextmenu:contextmenu class="flex-grow" />
  </div>

  <div class="border border-gray-500 mt-3 dark:caret-lime-800 w-auto h-96 overflow-y-auto hover:bg-gray-100 dark:hover:bg-gray-700 p-2 space-y-2 font-medium text-gray-900 dark:text-white">        
    <div v-if="selectedObj != null" class="flex">      
      <!-- Left side -->
      <div class="w-1/3">
        <h1 class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Параметры</h1>
        <table>
          <tbody>
            <tr><td>Serial number in the database:</td><td>{{ selectedObj.info.id }}</td></tr>
            <tr><td>Name:</td><td>{{ selectedObj.info.name }}</td></tr>
            <tr><td>Time of creation:</td><td>{{ selectedObj.info.create_time }}</td></tr>
            <tr><td>Element type:</td><td>{{ selectedObj.info.type }}</td></tr>
            <tr v-if="selectedObj.info.data_type != null"><td>Data type:</td><td>{{ selectedObj.info.data_type }}</td></tr>
            <tr><td>Status:</td><td>{{ selectedObj.info.status }}</td></tr>
            <tr><td>Total number of lines:</td><td>{{ selectedObj.info.lines_count }}</td></tr>
          </tbody>
        </table>          
      </div>

      <!-- Right side -->
      <div v-if="selectedObj.info.links != null" class="w-2/3 relative overflow-x-auto">
        <h1 class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Sources</h1>
        <table>
          <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
            <tr>
              <th scope="col" class="px-6 py-1">ID</th>
              <th scope="col" class="px-6 py-1">Status</th>
              <th scope="col" class="px-6 py-1">Type</th>
              <th scope="col" class="px-6 py-1">Qty-in line</th>
              <th scope="col" class="px-6 py-1">Path</th>
              <th scope="col" class="px-6 py-1">Error</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in selectedObj.info.links" :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
              <td class="px-6 py-1">{{ item.id }}</td>              
              <td class="px-6 py-1">{{ item.status }}</td>
              <td class="px-6 py-1">{{ item.type }}</td>
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
  <div v-if="insertItem.type != null" @click="insertItem.type = null; insertItem.name = null" class="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-40" aria-hidden="true"></div>
  <!-- Modal Insert Folder -->
  <div v-if="insertItem.type === 2" id="folder-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-md max-h-full">
    <div class="relative p-4 w-full max-w-md max-h-full">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Adding a folder</h3>
          <button @click="insertItem.type = null; insertItem.name = null" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
            </svg>
            <span class="sr-only">Close modal</span>
          </button>
        </div>
        <!-- Modal body -->
        <div class="p-4 md:p-5">
          <form class="space-y-4" action="#">
            <div>
              <input v-model="insertItem.name" type="text" id="name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white" placeholder="Наименование папки" required>
            </div>              
            <button @click.prevent="OPRAddItem"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Создать папку</button>              
          </form>
        </div>
      </div>
    </div>
  </div>
  
  <!-- Modal Insert File -->
  <div v-if="insertItem.type === 3" id="folder-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-2xl max-h-full">
    <div class="relative p-4 w-full max-w-2xl max-h-full mx-auto">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <!-- Modal header -->
        <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
          <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Adding a file/h3>
          <button @click="insertItem.type = null; insertItem.name = null" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
            <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
            </svg>
            <span class="sr-only">Close modal</span>
          </button>
        </div>
        <!-- Modal body -->
        <form class="p-4 md:p-5">
          <div class="grid gap-4 mb-4 grid-cols-2">
            <div class="col-span-2 sm:col-span-1">
              <label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Designation</label>
              <input v-model="insertItem.name" type="text" name="name" id="name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Введите наименование" required>
            </div>
            
            <div class="col-span-2 sm:col-span-1">
              <label for="category" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Data type</label>
              <select v-model="insertItem.data_type" id="category" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
                <option :value="Number(4)">EMAIL</option>
                <option :value="Number(5)">USERNAME</option>
                <option :value="Number(6)">PASSWORD</option>
                <option :value="Number(7)">PIN</option>
                <option :value="Number(8)">USERNAME:PASSWORD</option>
                <option :value="Number(9)">EMAIL:PASSWORD</option>
                <option :value="Number(10)">DATA</option>                
              </select>
            </div>
            
            <div class="col-span-2">
              <label for="description" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Sources</label>
              <table class="w-full">
                <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                  <tr>
                    <th scope="col" class="px-6 py-1">Type</th>
                    <th scope="col" class="px-6 py-1">Path</th>
                    <th scope="col" class="px-6 py-1">
                      <a @click.prevent="itemAction('add')" href="#" type="button" data-modal-target="editUserModal" data-modal-show="editUserModal" class="font-medium text-blue-600 dark:text-blue-500 hover:underline">Добавить</a>
                    </th>
                  </tr>
                </thead>
                <tbody>                  
                  <tr v-for="item in insertItem.links" :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                    <td class="px-6 py-1">
                      <select v-model="item.type" id="category" class="px-6 py-1 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
                        <!--<option value="15">Local file</option>-->
                        <option :value="Number(16)">Remote file</option>
                        <option :value="Number(17)">Link</option>
                      </select>
                    </td>              
                    <td class="px-6 py-1 dark:text-white">
                      <input v-model="item.path" type="text" name="name" id="name" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Введите источник" required>
                    </td>
                    <td class="px-6 py-4">
                      <a @click.prevent="itemAction('remove', item.id)" href="#" class="font-medium text-red-600 dark:text-red-500 hover:underline">Удалить</a>
                    </td>
                  </tr>
                </tbody>
              </table>              
            </div>

          </div>
          <button @click.prevent="OPRAddItem"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Создать файл</button>
        </form>
      </div>
    </div>
  </div>
  <!-- AlertForm --> 
  <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />
</template>

<script>

import axios from 'axios'
import ComboItem from '../Common/ComboItem.vue'
import AlertForm from '../Common/AlertForm.vue'

function deleteNodeById(arr, idToDelete) {
  return arr.filter(node => {  
    if (node.id === idToDelete) {
      return false // Delete the current node
    }

    if (node.child && node.child.length > 0) {      
      node.child = deleteNodeById(node.child, idToDelete) // Recursively call deleteNodeById on child nodes
      
      if (node.child.length === 0) { // Check if there are any child nodes left after deletion
        delete node.child // If not, delete the child property
      }
    }

    return true // Leave the node in the array
  })
}

function addNode(tree, parentId, newNode) {
  if (parentId === undefined || parentId === null) {
    tree.push(newNode)
    return
  } 

  tree.forEach(node => {
    if (node.id === parentId) {
      // Add a new node to the child array of the parent node
      if (node.child === undefined || node.child === null) {
        node.child = []
      }
      node.child.push(newNode)      
    } else if (node.child && node.child.length > 0) {
      // Recursively call addNode on child nodes
      addNode(node.child, parentId, newNode)
    }
  });
}

export default {
  components: {
    ComboItem,
    AlertForm,
  },
  inject: ['server_addr'],
  data() {
    return {
      msg_txt: null,
      msg_id: null,
      tree: null,
      
      showId: 0,
      selectId: 0,
      selectedObj: null,
      
      insertItem: {},
    };
  },
  async mounted() {
    const response = await this.HTTP('GET', '/api/list/GetComboLists', null)
    if (response.success) {
      this.tree = response.tree
    } else {
      this.msg_txt = response.msg_txt 
      this.msg_id = 0
    } 
  },
  methods: {
    parentMethod(action, id) {
      switch(action) {
        case "delete": this.OPRDeleteItem(id); break;
        case "insert folder": 
          this.insertItem.type = 2
          this.insertItem.p_id = id
          this.insertItem.child = null
          break
        case "insert file": 
          this.insertItem.type = 3
          this.insertItem.p_id = id
          this.insertItem.data_type = 4
          this.insertItem.links = []
          break;
        case "show": 
          this.showId = id 
          if (this.showId !== -1) {
            this.$refs.contextmenu.hide()
          } 
          break;
        case "select": this.OPRSelectItem(id); break;          
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
    async OPRDeleteItem(index) {
      const response = await this.HTTP('GET', '/api/list/DeleteComboList/?id='+index, null)
      if (response.success) {
        this.tree = deleteNodeById(this.tree, index)
        this.msg_txt = "Element removed"
        this.msg_id = 1  
      } else {
        this.msg_txt = response.msg_txt          
        this.msg_id = 0
      }
    },
    async OPRSelectItem(index) {
      this.selectId = index
      const response = await this.HTTP('GET', '/api/list/GetInfoComboList/?id='+index, null)
      if (response.success) {
        this.selectedObj = response
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }      
    },
    async OPRAddItem() {
      this.msg_txt = null
      this.msg_id = null

      if (this.insertItem.name === null || this.insertItem.name === '') {
        this.msg_txt = 'Please provide a name'
      }

      if (this.insertItem.type === 3) {
        if (this.insertItem.data_type === null) {
          this.msg_txt = 'Data type not specified'
        }

        if (this.insertItem.links.length === 0) {
          this.msg_txt = 'Please provide at least one download link'
        }
      }

      if (this.msg_txt !== null) {
        this.msg_id = 0
        return
      }

      const response = await this.HTTP('POST', '/api/list/UploadComboList', this.insertItem)
      if (response.success) {
        this.insertItem.id = response.id
        this.insertItem.create_time = response.create_time      
        addNode(this.tree, this.insertItem.p_id, this.insertItem)
        
        this.msg_txt = 'Item added to list'
        this.msg_id = 1
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
      this.insertItem = {}        
    },
    itemAction(action, id) {
      if (action === 'remove') {
        const indexToRemove = this.insertItem.links.findIndex(obj => obj.id === id)
        if (indexToRemove !== -1) {
          this.insertItem.links.splice(indexToRemove, 1)
        }        
      }

      if (action === 'add') {
        let maxId = Math.max(...this.insertItem.links.map(obj => obj.id))
        if(maxId == -Infinity) {
          maxId = 0
        } 
        this.insertItem.links.push({
          id: maxId+1,
          type: 17
        })
      }
    }
  }
}
</script>
