<template> <!-- You will need to remove border border-gray-500 --> 
  <div v-if="session_id !== null" class=" mt-3 dark:caret-lime-800 w-auto h-auto overflow-y-auto flex flex-col dark:text-white"> 
    <div v-if="!ShowResults">
    <div class="flex justify-between items-center mb-4 px-5 pt-2">
      <button @click.prevent="$emit('nullate_session_id')" type="button" class="w-44 h-11 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-2 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
        <svg class="w-6 h-6 text-gray-800 dark:text-white inline" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 8 14">        
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 1 1.3 6.326a.91.91 0 0 0 0 1.348L7 13"/>          
        </svg>
        Back to sessions
      </button>
      
      <div class="flex-1 flex items-center px-3">
        <div class="pr-2">
          Time of creation: {{ session.create_time }}
        </div>
        <input v-model="name_value" type="text" name="name" id="name" class="flex-1 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Введите наименование" required>          
        <svg @click="AlterField('name', name_value)" class="pl-2 w-12 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5.917 5.724 10.5 15 1.5"/>
        </svg>        
        <svg @click="name_value = session.name" class="pl-2 w-12 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
        </svg>  
      </div>
      
      <button @click="ShowResults = true" class="w-52 h-11 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-2 py-2.5 text-center dark:blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
        Go to results
        <svg class="w-6 h-6 text-gray-800 dark:text-white inline" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 8 14">
          <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 13 5.7-5.326a.909.909 0 0 0 0-1.348L1 1"/>
        </svg>        
      </button>
    </div>

    <!-- Buttons on the left side -->
    <div class="flex items-center mb-4 px-5">
      <button @click.prevent="ItemSelect(1)" class="mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select combo sheet/button>
      <button @click.prevent="ItemSelect(2)" class="mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select proxy preset</button>
      <button @click.prevent="ItemSelect(4)" class="mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select proxy from file</button>
      <button @click.prevent="ItemSelect(3)" class="mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select module</button>
      <button @click.prevent="AlterField('status', 22)" class="mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Launch</button>
      <button @click.prevent="AlterField('status', 24)" class="mr-2 text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Stop</button>      
      <label for="numberBox" class="mr-4 text-sm font-medium text-gray-400 dark:text-gray-500">Flows</label>
      <input v-model="worker_count_value" id="numberBox" type="number" class="w-36 px-4 py-2.5 text-gray-700 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:focus:ring-blue-600 dark:focus:border-blue-600">
      <svg @click="AlterField('worker_count', worker_count_value)" class="pl-2 w-12 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 12">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5.917 5.724 10.5 15 1.5"/>
      </svg>        
      <svg @click="worker_count_value = session.worker_count" class="pl-2 w-12 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
      </svg> 
    </div>
    <!-- Statistics -->
    <div class="px-5 pt-2 mt-3 h-auto overflow-y-auto p-2 space-y-2 font-medium text-gray-900">        
      <div class="flex">      
        <div class="px-5 w-4/5">
          <h1 class="self-center text-2xl font-semibold whitespace-nowrap dark:text-white">Options</h1>
          <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
            <tbody>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">Выбранный комбо лист</th>
                <td class="px-4 py-2">{{ session.database_name }}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">Выбранный proxy preset</th>
                <td class="px-4 py-2">{{ session.proxy_name }}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">Выбранный module</th>
                <td class="px-4 py-2">{{ session.module_name }}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">Статус</th>
                <td class="px-4 py-2 flex items-center">
                  <div v-if="session.status == 'Created'" class="h-2.5 w-2.5 rounded-full bg-gray-600 me-2" />
                  <div v-if="session.status == 'In progress'" class="h-2.5 w-2.5 rounded-full bg-green-600 me-2" />
                  <div v-if="session.status == 'Completed'" class="h-2.5 w-2.5 rounded-full bg-blue-600 me-2" />
                  <div v-if="session.status == 'Stopped'" class="h-2.5 w-2.5 rounded-full bg-orange-600 me-2" />
                  <div v-if="session.status == 'Error'" class="h-2.5 w-2.5 rounded-full bg-red-900 me-2" />
                  {{ session.status }}
                </td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">Время начала</th>
                <td class="px-4 py-2">{{ session.start_time }}</td>
              </tr>
              <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">Время завершения</th>
                <td class="px-4 py-2">{{ session.finish_time }}</td>
              </tr>
            </tbody>
          </table>      
        </div>

        <div class="px-5 w-4/5 relative overflow-x-auto">
          <div class="flex justify-between items-center">
            <h1 class="text-2xl font-semibold whitespace-nowrap dark:text-white">Статистика</h1>
            <svg @click="RefreshStatistic" class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 18 20">
              <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 1v5h-5M2 19v-5h5m10-4a8 8 0 0 1-14.947 3.97M1 10a8 8 0 0 1 14.947-3.97"/>
            </svg>
          </div>

          <div class="max-h-56 overflow-y-auto"> <!-- max-h-64 is 7 elements -->
            <table class="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
              <tbody>
                <tr v-for="item in session.result" :key="item" class="bg-white border-b dark:bg-gray-800 dark:border-gray-700">
                  <th scope="row" class="px-4 py-2 font-medium text-gray-900 whitespace-nowrap dark:text-white">{{ item.name }}</th>
                  <td class="px-4 py-2">{{ item.count }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

      </div>
      <!-- ProgressBar -->
      <div class="px-5">
        <div class="w-full bg-gray-200 dark:bg-gray-700">
          <div class="bg-blue-600 text-xs font-medium text-blue-100 text-center p-0.5 leading-none py-1" :style="{ width: session.statistic.percent }">{{ session.statistic.percent }}</div>
        </div>
      </div>
      <!-- Statistic -->
      <div class="px-5">
        <ul class="flex w-full text-sm font-medium text-gray-900 bg-white border border-gray-200 sm:flex dark:bg-gray-700 dark:border-gray-600 dark:text-white">
          <li class="flex-1 border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
            <div class="flex items-center px-3 py-1">
              <label for="horizontal-list-radio-license" class="w-full text-gray-900 dark:text-gray-300">Строк в комболисте: {{ session.statistic.combolistcount }}</label>
            </div>
          </li>
          <li class="flex-1 border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
            <div class="flex items-center px-3 py-1">
              <label for="horizontal-list-radio-id" class="w-full text-gray-900 dark:text-gray-300">Количество proxy: {{ session.statistic.proxycount }}</label>
            </div>
          </li>
          <li class="flex-1 border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
            <div class="flex items-center px-3 py-1">
              <label for="horizontal-list-radio-id" class="w-full text-gray-900 dark:text-gray-300">Пройдено строк: {{ session.statistic.passedcount }}</label>
            </div>
          </li>
          <li class="flex-1 border-b border-gray-200 sm:border-b-0 sm:border-r dark:border-gray-600">
            <div class="flex items-center px-3 py-1">
              <label for="horizontal-list-radio-military" class="w-full text-gray-900 dark:text-gray-300">Осталось строк: {{session.statistic.lastlines }}</label>
            </div>
          </li>
        </ul>
      </div>
      <!-- Error Message -->
      <div v-if="session.errors" class="px-5">
        <div id="alert-border-2" class="flex items-center p-4 mb-4 text-red-800 border-t-4 border-red-300 bg-red-50 dark:text-red-400 dark:bg-gray-800 dark:border-red-800" role="alert">
          <div class="flex p-4 mb-4 text-sm text-reed-800 rounded-lg bg-blue-50 dark:bg-gray-800 dark:text-red-400" role="alert">
            <svg class="flex-shrink-0 inline w-4 h-4 me-3 mt-[2px]" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 20 20">
              <path d="M10 .5a9.5 9.5 0 1 0 9.5 9.5A9.51 9.51 0 0 0 10 .5ZM9.5 4a1.5 1.5 0 1 1 0 3 1.5 1.5 0 0 1 0-3ZM12 15H8a1 1 0 0 1 0-2h1v-3H8a1 1 0 0 1 0-2h2a1 1 0 0 1 1 1v4h1a1 1 0 0 1 0 2Z"/>
            </svg>
            <span class="sr-only">Info</span>
            <div>
              <span class="font-medium">An error occurred:</span>
                <ul class="mt-1.5 list-disc list-inside">
                  <li v-for="err in session.errors.split(';')" :key="err">{{ err }}</li>
                </ul>
            </div>
          </div> 
        </div>
      </div>
    </div>

    <!-- Main modal -->
    <div v-if="select_item > 0" @click="select_item = 0" class="fixed top-0 left-0 w-full h-full bg-black opacity-50 z-40" aria-hidden="true"></div>
    <!-- Combo Player Selection Modal Window -->
    <div v-if="select_item === 1" id="list-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-md max-h-full">
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Modal content -->
        <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <!-- Modal header -->
          <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Выбор комбо листа</h3>
            <button @click="select_item = 0" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
              <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
              </svg>
              <span class="sr-only">Close modal</span>
            </button>
          </div>

          <!-- Modal body -->
          <div class="p-4 md:p-5">
            <form class="space-y-4" action="#">
              <div class="border border-gray-500 mt-3 dark:border-gray-600 w-auto h-96 overflow-y-auto flex flex-col dark:caret-lime-800"> 
                <ul>          
                  <ComboItem class="space-y-2 font-medium" :node="node" v-for="node in tree" :key="node" 
                    @call-parent-method="parentMethod" :isShowId="showId" :isSelectId="selectId"/>
                </ul>
              </div>
              <button @click.prevent="SelectAction('database_id')"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select комбо лист</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal window for choosing proxy preset -->
    <div v-if="select_item === 2" id="proxy-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-md max-h-full">
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Modal content -->
        <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <!-- Modal header -->
          <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Выбор proxy из presetа</h3>
            <button @click="select_item = 0" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
              <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
              </svg>
              <span class="sr-only">Close modal</span>
            </button>
          </div>

          <!-- Modal body -->
          <div class="p-4 md:p-5">
            <form class="space-y-4" action="#">
              <div class="border border-gray-500 mt-3 dark:border-gray-600 w-auto h-96 overflow-y-auto flex flex-col dark:caret-lime-800"> 
                <ul>        
                  <li @click="selectProxyPreset(preset.id)" class="space-y-2 font-medium" v-for="preset in tree" :key="preset" ref="pli" :id="preset.id">
                    <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
                      <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" viewBox="0 0 20 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M6.5 8C5.80777 8 5.13108 7.79473 4.55551 7.41015C3.97993 7.02556 3.53133 6.47893 3.26642 5.83939C3.00152 5.19985 2.9322 4.49612 3.06725 3.81719C3.2023 3.13825 3.53564 2.51461 4.02513 2.02513C4.51461 1.53564 5.13825 1.2023 5.81719 1.06725C6.49612 0.932205 7.19985 1.00152 7.83939 1.26642C8.47893 1.53133 9.02556 1.97993 9.41015 2.55551C9.79473 3.13108 10 3.80777 10 4.5" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        <path d="M6.5 17H1V15C1 13.9391 1.42143 12.9217 2.17157 12.1716C2.92172 11.4214 3.93913 11 5 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        <path d="M19.5 11H18.38C18.2672 10.5081 18.0714 10.0391 17.801 9.613L18.601 8.818C18.6947 8.72424 18.7474 8.59708 18.7474 8.4645C18.7474 8.33192 18.6947 8.20476 18.601 8.111L17.894 7.404C17.8002 7.31026 17.6731 7.25761 17.5405 7.25761C17.4079 7.25761 17.2808 7.31026 17.187 7.404L16.392 8.204C15.9647 7.93136 15.4939 7.73384 15 7.62V6.5C15 6.36739 14.9473 6.24021 14.8536 6.14645C14.7598 6.05268 14.6326 6 14.5 6H13.5C13.3674 6 13.2402 6.05268 13.1464 6.14645C13.0527 6.24021 13 6.36739 13 6.5V7.62C12.5081 7.73283 12.0391 7.92863 11.613 8.199L10.818 7.404C10.7242 7.31026 10.5971 7.25761 10.4645 7.25761C10.3319 7.25761 10.2048 7.31026 10.111 7.404L9.404 8.111C9.31026 8.20476 9.25761 8.33192 9.25761 8.4645C9.25761 8.59708 9.31026 8.72424 9.404 8.818L10.204 9.618C9.9324 10.0422 9.73492 10.5096 9.62 11H8.5C8.36739 11 8.24021 11.0527 8.14645 11.1464C8.05268 11.2402 8 11.3674 8 11.5V12.5C8 12.6326 8.05268 12.7598 8.14645 12.8536C8.24021 12.9473 8.36739 13 8.5 13H9.62C9.73283 13.4919 9.92863 13.9609 10.199 14.387L9.404 15.182C9.31026 15.2758 9.25761 15.4029 9.25761 15.5355C9.25761 15.6681 9.31026 15.7952 9.404 15.889L10.111 16.596C10.2048 16.6897 10.3319 16.7424 10.4645 16.7424C10.5971 16.7424 10.7242 16.6897 10.818 16.596L11.618 15.796C12.0422 16.0676 12.5096 16.2651 13 16.38V17.5C13 17.6326 13.0527 17.7598 13.1464 17.8536C13.2402 17.9473 13.3674 18 13.5 18H14.5C14.6326 18 14.7598 17.9473 14.8536 17.8536C14.9473 17.7598 15 17.6326 15 17.5V16.38C15.4919 16.2672 15.9609 16.0714 16.387 15.801L17.182 16.601C17.2758 16.6947 17.4029 16.7474 17.5355 16.7474C17.6681 16.7474 17.7952 16.6947 17.889 16.601L18.596 15.894C18.6897 15.8002 18.7424 15.6731 18.7424 15.5405C18.7424 15.4079 18.6897 15.2808 18.596 15.187L17.796 14.392C18.0686 13.9647 18.2662 13.4939 18.38 13H19.5C19.6326 13 19.7598 12.9473 19.8536 12.8536C19.9473 12.7598 20 12.6326 20 12.5V11.5C20 11.3674 19.9473 11.2402 19.8536 11.1464C19.7598 11.0527 19.6326 11 19.5 11ZM14 14.5C13.5055 14.5 13.0222 14.3534 12.6111 14.0787C12.2 13.804 11.8795 13.4135 11.6903 12.9567C11.5011 12.4999 11.4516 11.9972 11.548 11.5123C11.6445 11.0273 11.8826 10.5819 12.2322 10.2322C12.5819 9.8826 13.0273 9.6445 13.5123 9.54804C13.9972 9.45157 14.4999 9.50108 14.9567 9.6903C15.4135 9.87952 15.804 10.2 16.0787 10.6111C16.3534 11.0222 16.5 11.5055 16.5 12C16.5 12.663 16.2366 13.2989 15.7678 13.7678C15.2989 14.2366 14.663 14.5 14 14.5Z" fill="currentColor"/>
                      </svg>

                      <span class="ms-5">{{ preset.name }}</span>
                      <span class="ml-auto py-1">{{ preset.create_time }}</span>
                    </a>          
                  </li>    
                </ul>
              </div>
              <button @click.prevent="SelectAction('proxy_id')"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select proxy preset</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal window for choosing a module -->
    <div v-if="select_item === 3" id="folder-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-md max-h-full">
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Modal content -->
        <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <!-- Modal header -->
          <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Выбор модуля</h3>
            <button @click="select_item = 0" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
              <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
              </svg>
              <span class="sr-only">Close modal</span>
            </button>
          </div>

          <!-- Modal body -->
          <div class="p-4 md:p-5">
            <form class="space-y-4" action="#">
              <div class="border border-gray-500 mt-3 dark:border-gray-600 w-auto h-96 overflow-y-auto flex flex-col dark:caret-lime-800"> 
                <ul>        
                  <li @click="selectModule(modl.id)" class="space-y-2 font-medium" v-for="modl in tree" :key="modl" ref="mli" :id="modl.id">
                    <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">
                      <svg class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" viewBox="0 0 20 18" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M6.5 8C5.80777 8 5.13108 7.79473 4.55551 7.41015C3.97993 7.02556 3.53133 6.47893 3.26642 5.83939C3.00152 5.19985 2.9322 4.49612 3.06725 3.81719C3.2023 3.13825 3.53564 2.51461 4.02513 2.02513C4.51461 1.53564 5.13825 1.2023 5.81719 1.06725C6.49612 0.932205 7.19985 1.00152 7.83939 1.26642C8.47893 1.53133 9.02556 1.97993 9.41015 2.55551C9.79473 3.13108 10 3.80777 10 4.5" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        <path d="M6.5 17H1V15C1 13.9391 1.42143 12.9217 2.17157 12.1716C2.92172 11.4214 3.93913 11 5 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
                        <path d="M19.5 11H18.38C18.2672 10.5081 18.0714 10.0391 17.801 9.613L18.601 8.818C18.6947 8.72424 18.7474 8.59708 18.7474 8.4645C18.7474 8.33192 18.6947 8.20476 18.601 8.111L17.894 7.404C17.8002 7.31026 17.6731 7.25761 17.5405 7.25761C17.4079 7.25761 17.2808 7.31026 17.187 7.404L16.392 8.204C15.9647 7.93136 15.4939 7.73384 15 7.62V6.5C15 6.36739 14.9473 6.24021 14.8536 6.14645C14.7598 6.05268 14.6326 6 14.5 6H13.5C13.3674 6 13.2402 6.05268 13.1464 6.14645C13.0527 6.24021 13 6.36739 13 6.5V7.62C12.5081 7.73283 12.0391 7.92863 11.613 8.199L10.818 7.404C10.7242 7.31026 10.5971 7.25761 10.4645 7.25761C10.3319 7.25761 10.2048 7.31026 10.111 7.404L9.404 8.111C9.31026 8.20476 9.25761 8.33192 9.25761 8.4645C9.25761 8.59708 9.31026 8.72424 9.404 8.818L10.204 9.618C9.9324 10.0422 9.73492 10.5096 9.62 11H8.5C8.36739 11 8.24021 11.0527 8.14645 11.1464C8.05268 11.2402 8 11.3674 8 11.5V12.5C8 12.6326 8.05268 12.7598 8.14645 12.8536C8.24021 12.9473 8.36739 13 8.5 13H9.62C9.73283 13.4919 9.92863 13.9609 10.199 14.387L9.404 15.182C9.31026 15.2758 9.25761 15.4029 9.25761 15.5355C9.25761 15.6681 9.31026 15.7952 9.404 15.889L10.111 16.596C10.2048 16.6897 10.3319 16.7424 10.4645 16.7424C10.5971 16.7424 10.7242 16.6897 10.818 16.596L11.618 15.796C12.0422 16.0676 12.5096 16.2651 13 16.38V17.5C13 17.6326 13.0527 17.7598 13.1464 17.8536C13.2402 17.9473 13.3674 18 13.5 18H14.5C14.6326 18 14.7598 17.9473 14.8536 17.8536C14.9473 17.7598 15 17.6326 15 17.5V16.38C15.4919 16.2672 15.9609 16.0714 16.387 15.801L17.182 16.601C17.2758 16.6947 17.4029 16.7474 17.5355 16.7474C17.6681 16.7474 17.7952 16.6947 17.889 16.601L18.596 15.894C18.6897 15.8002 18.7424 15.6731 18.7424 15.5405C18.7424 15.4079 18.6897 15.2808 18.596 15.187L17.796 14.392C18.0686 13.9647 18.2662 13.4939 18.38 13H19.5C19.6326 13 19.7598 12.9473 19.8536 12.8536C19.9473 12.7598 20 12.6326 20 12.5V11.5C20 11.3674 19.9473 11.2402 19.8536 11.1464C19.7598 11.0527 19.6326 11 19.5 11ZM14 14.5C13.5055 14.5 13.0222 14.3534 12.6111 14.0787C12.2 13.804 11.8795 13.4135 11.6903 12.9567C11.5011 12.4999 11.4516 11.9972 11.548 11.5123C11.6445 11.0273 11.8826 10.5819 12.2322 10.2322C12.5819 9.8826 13.0273 9.6445 13.5123 9.54804C13.9972 9.45157 14.4999 9.50108 14.9567 9.6903C15.4135 9.87952 15.804 10.2 16.0787 10.6111C16.3534 11.0222 16.5 11.5055 16.5 12C16.5 12.663 16.2366 13.2989 15.7678 13.7678C15.2989 14.2366 14.663 14.5 14 14.5Z" fill="currentColor"/>
                      </svg>

                      <span class="ms-5">{{ modl.name }}</span>
                      <span class="ml-auto py-1">{{ modl.create_time }}</span>
                    </a>          
                  </li>    
                </ul>
              </div>
              <button @click.prevent="SelectAction('module_id')"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select модуля</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    <!-- Modal window for selecting proxy from file -->
    <div v-if="select_item === 4" id="proxy-modal" tabindex="-1" aria-hidden="true" class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 overflow-y-auto overflow-x-hidden z-50 w-full max-w-md max-h-full">
      <div class="flex items-center justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
        <!-- Modal content -->
        <div class="inline-block align-bottom bg-white dark:bg-gray-800 rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-2xl sm:w-full">
          <!-- Modal header -->
          <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
            <h3 class="text-xl font-semibold text-gray-900 dark:text-white">Выбор proxy from file</h3>
            <button @click="select_item = 0" type="button" class="end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="authentication-modal">
              <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
              </svg>
              <span class="sr-only">Close modal</span>
            </button>
          </div>

          <!-- Modal body -->
          <div class="p-4 md:p-5">
            <form class="space-y-4" action="#">
              <div class="col-span-3 sm:col-span-1">
                <label for="timeout" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Таймаут (милесекунды)</label>
                <input v-model="insertItem.timeout" type="number" name="timeout" id="timeout" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
              </div>
              <div class="col-span-3 sm:col-span-1">
                <label for="timeout" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Тип proxy</label>
                <select v-model="insertItem.proxy_type" id="category" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-800 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500">
                  <option :value="Number(18)">HTTP/S</option>
                  <option :value="Number(19)">SOCKS4</option>
                  <option :value="Number(20)">SOCKS5</option>
                </select>
              </div>
                
              <input @change="insertItem.file = $event.target.files[0]" class="block w-full text-sm text-gray-900 border border-gray-300 rounded-lg cursor-pointer bg-gray-50 dark:text-gray-400 focus:outline-none dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" id="file_input" type="file">            
              <button @click.prevent="UploadSessionProxy()"  type="button" class="w-full text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">Select proxy from file</button>
            </form>
          </div>
        </div>
      </div>
    </div>
    </div>

    <SessionResult :isShow="ShowResults" :sessionId="session_id" @SetFalseShowResults="ShowResults = false"/>
  </div>

  <!-- AlertForm --> 
  <AlertForm :msg_txt="msg_txt" :msg_id="msg_id" />
</template>


<script>

import axios from 'axios'
import SessionResult from './SessionResultsForm.vue'
import AlertForm from '../Common/AlertForm.vue'
import ComboItem from '../Common/ComboItem.vue'

const selectedClassDay = 'bg-gray-100'
const selectedClassDark = 'dark:bg-gray-700'

export default {
  components: {
    AlertForm,
    ComboItem,
    SessionResult
  },
  inject: ['server_addr'],
  props: ['session_id'],
  data() {
    return {
      msg_txt: null,
      msg_id: 0,
      name_value: null,
      worker_count_value: null,
      select_item: 0,
      tree: null,
      showId: -1,
      selectId: -1,
      ShowResults: false,
      
      session: {},
      insertItem: {}
    }
  },
  watch: {
    async session_id(newValue) {
      if (newValue != null) {
        await this.mount()
      }
    }
  },
  async mounted() {
    await this.mount()
  },
  methods: {
    async mount() {
      const response = await this.HTTP('GET', '/api/sess/GetInfoSession/?id='+this.session_id, null)
      if (response.success) {
        this.session = response.session
        this.name_value = this.session.name
        this.worker_count_value = this.session.worker_count
        await this.RefreshStatistic()
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
    async AlterField(Field, Value) {
      const response = await this.HTTP('POST', '/api/sess/AlterSession', {
        id: this.session.id,
        field: Field,
        new_value: Value
      })

      if (response.success) {
        this.msg_txt = response.msg_txt 
        this.msg_id = 1
        this.session.name = this.name_value
      } else {
        this.msg_txt = response.msg_txt 
        this.msg_id = 0
      }
      if (Field === 'status') {
        await this.mount()
      }  
    },
    selectProxyPreset(id) {
      this.selectId = id
      this.$refs.pli.forEach(async (element) => {
        if (id == element.id) {
          element.classList.add(selectedClassDay)
          element.classList.add(selectedClassDark)
        } else {
          element.classList.remove(selectedClassDay)
          element.classList.remove(selectedClassDark)
        }
      })
    },
    selectModule(id) {
      this.selectId = id
      this.$refs.mli.forEach(async (element) => {
        if (id == element.id) {
          element.classList.add(selectedClassDay)
          element.classList.add(selectedClassDark)
        } else {
          element.classList.remove(selectedClassDay)
          element.classList.remove(selectedClassDark)
        }
      })
    },
    async ItemSelect(i) {
      this.select_item = i
      if (i === 4) {
        this.insertItem = {}
        this.insertItem.proxy_type = 18
        this.insertItem.timeout = 15000
        return
      }
      
      let url
      switch (i) {
        case 1: url = '/api/list/GetComboLists';   break;
        case 2: url = '/api/prox/GetProxyPresets'; break;
        case 3: url = '/api/modl/GetModules';      break;
      }
      const response = await this.HTTP('GET', url, null)
      if (response.success) {
        switch (i) {
          case 1: this.tree = response.tree;    break;
          case 2: this.tree = response.presets; break;
          case 3: this.tree = response.modules; break;
        }  
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
    },
    async SelectAction(a) {
      if (this.selectId === -1) {
        this.msg_txt = 'Select an item'
        this.msg_id = 0
        return
      }
      await this.AlterField(a, this.selectId)
      await this.mount()
      this.select_item = 0
      this.selectId = -1
    },
    async PullAction(a) {
      const response = await this.HTTP('POST', '/api/sess/ActionSession', {
        id: this.session_id,
        action: a
      })
      if (response.success) {
        this.msg_txt = response.msg_txt
        this.msg_id = 1
      } else {
        this.msg_txt = response.msg_txt
        this.msg_id = 0
      }
    },
    async RefreshStatistic() {
      const response = await this.HTTP('GET', '/api/sess/GetStatistic/?id='+this.session_id, null)
      if (response.success) {
        this.session.finish_time = response.finish_time
        this.session.status = response.status
        this.session.result = response.result
        this.session.statistic = response.statistic
      } else {
        this.msg_txt = response.msg_txt 
        this.msg_id = 0
      } 
    },
    async UploadSessionProxy() {
      this.msg_txt = null 
      this.msg_id = -1
      if (this.insertItem.timeout < 1000 || this.insertItem.timeout == null) {
        this.msg_txt = 'Please specify a timeout'
      }     
      if (this.insertItem.file == null) {
        this.msg_txt = 'Please specify the file with proxy'
      }
      if (this.insertItem.file.name.slice(-4).toUpperCase() !== '.TXT') {
        this.msg_txt = 'File format does not match'
      }
      if (this.msg_txt != null) {
        this.msg_id = 0
        return
      }

      let reader = new FileReader()
      reader.onload  = async () => { 
        this.insertItem.data = btoa(reader.result)
        this.insertItem.session_id = this.session_id
        delete this.insertItem.file        
        const response = await this.HTTP('POST', '/api/sess/UploadProxy', this.insertItem)
        if (response.success) {
          this.insertItem = null
          this.select_item = 0
          this.msg_txt = 'proxy has been loaded'
          this.msg_id = 1
          await this.mount()
        } else {
          this.msg_txt = response.msg_txt
          this.msg_id = 0
        }
      }
      reader.onerror = () => { 
        this.msg_txt = 'Error reading file'
        this.msg_id = 0 
      }
      reader.readAsText(this.insertItem.file)
    },
    parentMethod(action, id) {
      if (action === "select") {
        this.selectId = id
      }
      this.showId = -2
    }
  }
}

</script>
