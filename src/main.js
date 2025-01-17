import './assets/main.css'

/* import the fontawesome core */
import { library } from '@fortawesome/fontawesome-svg-core'

/* import font awesome icon component */
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

/* import specific icons */
//import { all } from '@awesome.me/kit-ad69ce3fbe/icons'

import {fas} from '@fortawesome/free-solid-svg-icons'
library.add(fas)

import { createApp } from 'vue'
import App from './App.vue'

const app = createApp(App)

app.component('font-awesome-icon', FontAwesomeIcon)

app.mount('#app')


