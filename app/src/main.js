import { createApp } from 'vue'

import 'uno.css'
import 'uno:components.css'
import 'uno:utilities.css'
import '@/styles/index.sass'
import '@/styles/transitions.sass'
import '@/styles/variables.sass'

import App from './App.vue'

import createRouterInstance from '@/routes'

const router = createRouterInstance()
const app = createApp(App)

app.use(router).mount('#app')
