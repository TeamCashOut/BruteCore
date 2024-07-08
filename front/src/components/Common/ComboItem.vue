<template> 
  
  <v-contextmenu v-if="isShowId !== -2" @show="parentMethod('show', node.id)" ref="contextmenu">
    <v-contextmenu-item :auto-hide="false" @click="parentMethod('delete', node.id)">Удалить</v-contextmenu-item>
    <v-contextmenu-item v-if="node.type != 3" :auto-hide="false" @click="parentMethod('insert folder', node.id)">Добавить папку</v-contextmenu-item>
    <v-contextmenu-item v-if="node.type != 3" :auto-hide="false" @click="parentMethod('insert file', node.id)">Добавить комбо-лист</v-contextmenu-item>
  </v-contextmenu>

  <li v-contextmenu:contextmenu @click="parentMethod('select', node.id)" ref="li">
    <a href="#" class="flex items-center p-2 text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700 group">    
      <!-- Angle -->
      <svg v-if="node.type == 2 && !isOpen" @click="isOpen = !isOpen" class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 8 14">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 13 5.7-5.326a.909.909 0 0 0 0-1.348L1 1"/>
      </svg>
      <svg v-if="node.type == 2 && isOpen" @click="isOpen = !isOpen" class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 8">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 5.326 5.7a.909.909 0 0 0 1.348 0L13 1"/>
      </svg>

      <!-- Folder -->
      <svg v-if="node.type == 2 && !isOpen" @click="isOpen = !isOpen" class="w-10 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-3 0 18 18">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M1 5v11a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H1Zm0 0V2a1 1 0 0 1 1-1h5.443a1 1 0 0 1 .8.4l2.7 3.6H1Z"/>
      </svg>
      <svg v-if="node.type == 2 && isOpen" @click="isOpen = !isOpen" class="w-10 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="-3 0 18 18">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.539 17h12.476l4-9H5m-2.461 9a1 1 0 0 1-.914-1.406L5 8m-2.461 9H2a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1h5.443a1 1 0 0 1 .8.4l2.7 3.6H16a1 1 0 0 1 1 1v2H5"/>
      </svg>

      <!-- File -->
      <svg v-if="node.type == 3" class="w-6 h-6 text-gray-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 16 20">
        <path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 1v4a1 1 0 0 1-1 1H1m14-4v16a.97.97 0 0 1-.933 1H1.933A.97.97 0 0 1 1 18V5.828a2 2 0 0 1 .586-1.414l2.828-2.828A2 2 0 0 1 5.828 1h8.239A.97.97 0 0 1 15 2Z"/>
      </svg>
      
      <span class="ms-5">{{ node.name }}</span>
      <span class="ml-auto py-1">{{ node.create_time }}</span>  
    </a>
  </li>

  <div v-if="isOpen" class="ms-10">
    <ComboItem class="space-y-2 font-medium" :node="child" v-for="child in node.child" :key="child" 
    @call-parent-method="parentMethod" :isShowId="isShowId" :isSelectId="isSelectId"/>
  </div>

</template>

<script>

const selectedClassDay = 'bg-gray-100'
const selectedClassDark = 'dark:bg-gray-700'

export default {
  name: "ComboItem",
  props: {
    node: {
      type: Object,
      required: true
    },
    isShowId: {
      type: Number,
      required: true
    },
    isSelectId: {
      type: Number,
      required: true
    }
  },
  watch: {
    isShowId(newValue) {
      if (this.node.id != newValue) {
        this.$refs.contextmenu.hide()
      }
    },
    isSelectId(newValue) {
      if (this.node.id != newValue) {
        this.$refs.li.classList.remove(selectedClassDay)
        this.$refs.li.classList.remove(selectedClassDark)
      }
    }
  },
  data() {
    return {
      isOpen: false,
      isVisible: true,
    }
  },
  methods: {
    parentMethod(action, id) {
      if (action == 'select') {
        this.$refs.li.classList.add(selectedClassDay)
        this.$refs.li.classList.add(selectedClassDark)
      }   
      this.$emit('call-parent-method', action, id)
    }
  }
}

</script>
