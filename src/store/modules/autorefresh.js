export default {
    namespaced: true,
    state: {
        id: null,
        enabled: null,
        interval: 15,
        callbacks: []
    },
    getters: {
        enabled: state => { return state.enabled == null ? false : state.enabled },
        interval: state => { return state.interval },
        callbacks: state => { return state.callbacks }
    },
    mutations: {
        setInterval (state, { interval }) {
            state.interval = interval
        },
        addCallbackAction (state, { callback }) {
            if (!state.callbacks.includes(callback)) {
                state.callbacks.push(callback)
            }
        },
        enable (state, { id }) {
            state.enabled = true
            state.id = id
        },
        disable (state) {
            clearInterval(state.id)
            state.id = null
            state.enabled = false
        }
    },
    actions: {
        toggle (context) {
            context.getters.enabled ? context.dispatch('disable') : context.dispatch('enable')
        },
        setInterval (context, { interval }) {
            context.commit('setInterval', { interval: interval })
            if (context.getters.enabled) {
                let oldIntervalID = context.state.id
                context.dispatch('enable')
                clearInterval(oldIntervalID)
            }
        },
        enable (context) {
            context.dispatch('refresh')
            let intervalID = setInterval(() => {
                context.dispatch('refresh')
            }, context.getters.interval * 1000)
            context.commit('enable', { id: intervalID })
        },
        disable (context) {
            context.commit('disable')
        },
        refresh (context) {
            context.getters.callbacks.forEach((callback) => {
                context.dispatch(callback, null, { root: true })
            })
        }
    }
}
