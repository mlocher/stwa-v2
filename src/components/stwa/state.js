import axios from 'axios'

// TODO: missing an error state for the XHR request

export default {
    namespaced: true,
    state: {
        data: null,
        loading: false,
        lastRefresh: null,
        labels: {
            standby: 'Bereitschaft',
            advance_warning: 'Vorwarnung',
            gale_warning: 'Sturmwarnung',
            out_of_order: 'Außer Betrieb',
            unknown: 'Unbekannt'
        }
    },
    getters: {
        loading: state => { return state.loading },
        data: state => { return state.data },
        label: state => (stwaState) => { return state.labels[stwaState] || state.labels['unknown'] }
    },
    mutations: {
        setData (state, { data }) {
            state.lastRefresh = Date.now()
            state.data = data
        },
        startLoading (state) {
            state.loading = true
        },
        endLoading (state) {
            state.loading = false
        }
    },
    actions: {
        fetch (context) {
            context.commit('startLoading')
            axios.get(process.env.VUE_APP_API_URL)
                .then(response => {
                    context.commit('setData', { data: response.data })
                })
                .catch(error => {
                    console.log(error)
                })
                .finally(() => {
                    context.commit('endLoading')
                })
        }
    }
}
