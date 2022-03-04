import { createApp } from 'vue'
import Root from './App.vue'
import './index.css'

const app = createApp(Root)

import utils from './libs/utils'
app.use(utils);

app.mount('#app')
