import { createApp } from 'vue'
import { defineRule } from 'vee-validate'
import { email, required } from '@vee-validate/rules'

import App from './App.vue'
import createRouterInstance from '@/routes'

import 'uno.css'
import 'uno:components.css'
import 'uno:utilities.css'
import '@/styles/index.sass'
import '@/styles/transitions.sass'
import '@/styles/variables.sass'

defineRule('required', required)
defineRule('email', email)

const router = createRouterInstance()
const app = createApp(App)

app.use(router).mount('#app')
