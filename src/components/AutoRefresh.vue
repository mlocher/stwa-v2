<template>
    <div>
        <div class="flex block justify-between items-baseline px-3 md:px-4 py-2 text-base md:text-sm text-gray-400 md:text-gray-700 rounded-md md:rounded-none font-medium md:font-normal hover:text-white hover:bg-gray-700 md:hover:text-gray-700 md:hover:bg-gray-100 focus:outline-none focus:text-white focus:bg-gray-700 md:focus:text-gray-700 md:focus:bg-gray-100">
            <div>Auto Refresh?</div>
            <div>
                <span class="relative inline-flex items-center justify-center flex-shrink-0 h-5 w-10 cursor-pointer focus:outline-none" role="checkbox" tabindex="0" v-on:click="toggle()" v-on:keydown.space.prevent="toggle()" :aria-checked="enabled.toString">
                    <span aria-hidden="true" :class="{ 'bg-indigo-600': enabled, 'bg-gray-200': !enabled }" class="absolute h-4 w-9 mx-auto rounded-full transition-colors ease-in-out duration-200"></span>
                    <span aria-hidden="true" :class="{ 'translate-x-5': enabled, 'translate-x-0': !enabled }" class="absolute left-0 inline-block h-5 w-5 border border-gray-200 rounded-full bg-white shadow transform transition-transform ease-in-out duration-200"></span>
                </span>
            </div>
        </div>
        <div class="flex block justify-between items-baseline px-3 md:px-4 py-2 text-base md:text-sm text-gray-400 md:text-gray-700 rounded-md md:rounded-none font-medium md:font-normal hover:text-white hover:bg-gray-700 md:hover:text-gray-700 md:hover:bg-gray-100 focus:outline-none focus:text-white focus:bg-gray-700 md:focus:text-gray-700 md:focus:bg-gray-100">
            <label for="interval">Intervall (Sek)</label>
            <div class="relative w-16 rounded-md shadow-sm">
                <input id="interval" v-model.number.lazy="interval" type="number" class="form-input block w-full sm:text-sm sm:leading-5 text-gray-700" />
            </div>
        </div>
    </div>
</template>

<script>
export default {
    computed: {
        enabled () {
            return this.$store.getters['autorefresh/enabled']
        },
        interval: {
            get: function () {
                return this.$store.getters['autorefresh/interval']
            },
            set: function (newInterval) {
                this.$store.dispatch('autorefresh/setInterval', { interval: newInterval })
            }
        }
    },
    methods: {
        toggle: function () {
            this.$store.dispatch('autorefresh/toggle')
        }
    }
}
</script>

<style scoped>
</style>
