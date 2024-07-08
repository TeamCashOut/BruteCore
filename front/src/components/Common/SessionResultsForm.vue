<template>
  <div v-if="isShow">
    <div class="flex justify-between items-center mb-4 px-5 pt-2">
      <button @click.prevent="$emit('SetFalseShowResults')" type="button" class="w-48 h-11 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-2 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
        <svg class="w-6 h-6 text-gray-800 dark:text-white inline" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 8 14">        
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 1 1.3 6.326a.91.91 0 0 0 0 1.348L7 13"/>          
        </svg>
        Назад к сессии
      </button>

      <ul class="grid gap-2 md:grid-cols-7 h-11 ml-2">
        <li v-for="item in result" :key="item.id" class="font-thin w-44 mr-16">
          <input :id="'hosting-' + item.id" :name="'hosting'" :value="item.id" v-model="sendObj.selectedItem" type="checkbox" class="hidden peer">
          <label :for="'hosting-' + item.id" class="inline-flex justify-between w-full h-11 p-5 text-gray-500 bg-white border border-gray-200 rounded-lg cursor-pointer dark:hover:text-gray-300 dark:border-gray-700 dark:peer-checked:text-blue-500 peer-checked:border-blue-600 peer-checked:text-blue-600 hover:text-gray-600 hover:bg-gray-100 dark:text-gray-400 dark:bg-gray-800 dark:hover:bg-gray-700 items-center">
            {{ item.name }}({{ item.count }})
          </label>
        </li>
      </ul>
    </div>

    <hr class="border border-gray-500 mt-3">

    <div class="relative overflow-x-auto shadow-md sm:rounded-lg mt-3">
      <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" class="px-6 py-2">Status</th>
            <th scope="col" class="px-6 py-2">Data</th>
            <th scope="col" class="px-6 py-3">Logs</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in respObj.ResultData" :key="item" class="odd:bg-white odd:dark:bg-gray-900 even:bg-gray-50 even:dark:bg-gray-800 border-b dark:border-gray-700">
            <th class="px-6 py-4">{{ item.status }}</th>
            <td class="px-6 py-4">{{ item.data }}</td>
            <td class="px-6 py-4">{{ item.log }}</td>
          </tr>
        </tbody>
      </table>
    </div>


    <div class="flex mt-3">
      <label for="category" class="w-28 h-10 flex items-center mb-2 text-sm font-medium text-gray-900 dark:text-white">Data format</label>
      <select v-model="sendObj.downldFormat" id="category" class="w-24 h-10 mr-2 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
        <option value="txt">txt</option>
        <option value="xlsx">xlsx</option>                
      </select>
      <button @click="DownloadResults('DownloadAll')" type="button" class="w-24 h-10 mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-2 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
        Download all
      </button>
      <button @click="DownloadResults('DownloadSelected')" type="button" class="w-36 h-10 mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-2 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
        Download selected
      </button>
      
      <nav aria-label="Page navigation example">
        <ul class="flex items-center -space-x-px h-10 text-base">
          <li @click="sendObj.page > 1 ? sendObj.page-- : null">
            <a href="#" class="flex items-center justify-center px-4 h-10 ms-0 leading-tight text-gray-500 bg-white border border-e-0 border-gray-300 rounded-s-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
              <span class="sr-only">Previous</span>
              <svg class="w-3 h-3 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 1 1 5l4 4"/>
              </svg>
            </a>
          </li>
          <li v-for="number in surf(sendObj.page, respObj.PagesCount)" :key="number">
            <a v-if="number !== sendObj.page" @click="sendObj.page = number" href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">{{ number }}</a>
            <a v-if="number === sendObj.page" href="#" aria-current="page" class="z-10 flex items-center justify-center px-4 h-10 leading-tight text-blue-600 border border-blue-300 bg-blue-50 hover:bg-blue-100 hover:text-blue-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white">{{ number }}</a>
          </li>
          <li @click="sendObj.page < respObj.PagesCount ? sendObj.page++ : null">
            <a href="#" class="flex items-center justify-center px-4 h-10 leading-tight text-gray-500 bg-white border border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
              <span class="sr-only">Next</span>
              <svg class="w-3 h-3 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 6 10">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 9 4-4-4-4"/>
              </svg>
            </a>
          </li>
        </ul>
      </nav>

      <select v-model="sendObj.entries_count" id="category" class="px-3 py-1 ml-3 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-16 h-10 p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
        <option :value="10">10</option>
        <option :value="25">25</option>
        <option :value="50">50</option>      
        <option :value="100">100</option>
      </select>
    </div>
    
    <!-- AlertForm --> 
    <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />
  </div>
</template>

<script>

import axios from 'axios'
import AlertForm from '../Common/AlertForm.vue'

function getSurroundingInts(selectedNumber, maxNumber) {
  const result = [];
  let startIndex
  let endIndex

  if (maxNumber <= 5) {
    startIndex = 1
    endIndex = maxNumber
  } else if (selectedNumber === 1 || selectedNumber === 2) {
    startIndex = 1
    endIndex = 5
  } else if (selectedNumber === maxNumber-1 || selectedNumber === maxNumber-2 || selectedNumber === maxNumber) {
    startIndex = maxNumber - 4
    endIndex = maxNumber
  } else {
    startIndex = Math.max(selectedNumber - 2, 1);
    endIndex = Math.min(selectedNumber + 2, maxNumber);
  }

  for (let i = startIndex; i <= endIndex; i++) {
    result.push(i);
  }

  return result;
}

export default {
  inject: ['server_addr'],
  props: ['sessionId', 'isShow'],
  watch: {
    async isShow() {
      await this.mount()
    },
    sendObj: {
      async handler() {
        await this.getResults()
      },
      deep: true
    }
  },
  components: {
    AlertForm
  },
  data() {
    return {
      msg_txt: null,
      msg_id: 0,
      sendObj: {
        selectedItem: [],
        entries_count: 10,
        page: 1,
        downldFormat: 'txt',
      },
      
      respObj: {
        PagesCount: 10,
        ResultData: {}
      },

      result: {},
      surf: getSurroundingInts,
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
    async mount() {
      const response = await this.HTTP('GET', '/api/sess/GetStatistic/?id='+this.sessionId, null)
      if (response.success) {
        this.result = response.result.filter(item => item.name !== "Errors")
        if (response.statistic.lastlines > 0) {
          this.result.push({
            id: 999,
            name:"Remainder",
            count: response.statistic.lastlines
          })
        }
      } else {
        this.msg_txt = response.msg_txt 
        this.msg_id = 0
      }
    },
    async getResults() {      
      if (this.sendObj.selectedItem.length === 0) {
        return
      }

      if (![10, 25, 50, 100].includes(this.sendObj.entries_count)) {
        return
      }

      if (this.sendObj.page < 1 || this.sendObj.page > this.respObj.PagesCount) {
        return
      }
      
      const response = await this.HTTP('GET', `/api/sess/GetResults/?session_id=${this.sessionId}&page=${this.sendObj.page}&entriescount=${this.sendObj.entries_count}&params=${this.sendObj.selectedItem.join(',')}`, null)
      if (response.success) {
        this.respObj = {
          PagesCount: response.pagescount,
          ResultData: response.data
        }

        if (this.sendObj.page > this.respObj.PagesCount) {
          this.sendObj.page = 1
        }        
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
    },
    async DownloadResults(opt) {
      let url = `${this.server_addr}/api/dwnl/${opt}/?session_id=${this.sessionId}&format=${this.sendObj.downldFormat}`
      if (opt === 'DownloadSelected') {
        if (this.sendObj.selectedItem.length === 0) {
          this.msg_txt = 'Select what you want to download'
          this.msg_id = 0
        }
        url = url+'&params='+this.sendObj.selectedItem.join(',')
      }
      url = url+'&auth='+localStorage.accessToken
      
      try {
        const response = await axios({method: 'get', url, responseType: 'arraybuffer'})
        url = window.URL.createObjectURL(new Blob([response.data]))
        const link = document.createElement('a')
        link.href = url
        link.setAttribute('download', 'results.zip')
        document.body.appendChild(link)
        link.click()
      } catch {
        this.msg_txt = 'error occured'
        this.msg_id = 0
      }
    }
  },
}

</script>
