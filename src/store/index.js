import Vue from 'vue'
import Vuex from 'vuex'

import autorefresh from './modules/autorefresh'
import stwa from '../components/stwa/state.js'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    autorefresh,
    stwa,
  },
  strict: debug
})
