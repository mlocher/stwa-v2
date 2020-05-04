import Vue from 'vue'
import Vuex from 'vuex'

import App from './App'
import store from './store/'

import './registerServiceWorker'

Vue.use(Vuex)
Vue.config.productionTip = false

new Vue({
    el: '#app',
    store,
    render: h => h(App),
})
