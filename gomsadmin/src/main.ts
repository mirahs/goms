import { createApp } from 'vue'

import { createPinia } from 'pinia'
import piniaPluginPersist from 'pinia-plugin-persist'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import '@/assets/css/global.css'

import App from './App.vue'

import router from './router'


const pinia = createPinia()
pinia.use(piniaPluginPersist)

createApp(App).use(pinia).use(ElementPlus).use(router).mount('#app')
