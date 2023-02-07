import { createApp } from 'vue'
import { Quasar, Notify, Dialog } from 'quasar'

// Import icon libraries
import '@quasar/extras/material-icons/material-icons.css'

// Import Quasar css
import 'quasar/src/css/index.sass'
import App from './App.vue'
import i18n from './locales'
const app = createApp(App)
app.use(Quasar, {
    plugins: { Notify, Dialog },
    config: {
        notify: { /* look at QuasarConfOptions from the API card */ }
    } // import Quasar plugins and add here
})
app.use(i18n)
// Assumes you have a <div id="app"></div> in your index.html
app.mount('#app')
