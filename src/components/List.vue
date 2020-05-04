<template>
    <div v-if="stwaData && Object.keys(stwaData).length != 0" class="bg-white shadow overflow-hidden sm:rounded-md">
        <ul v-for="station in stwaData.features" :key="station.properties.name" class="divide-y divide-gray-200">
            <li class="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50 transition duration-150 ease-in-out">
                <stwaListItem v-bind="station" />
            </li>
        </ul>
    </div>
    <div v-else class="fixed bottom-0 inset-x-0 px-4 pb-6 sm:inset-0 sm:p-0 sm:flex sm:items-center sm:justify-center">
        <div class="bg-white rounded-lg px-4 pt-5 pb-4 overflow-hidden shadow-xl transform transition-all sm:max-w-sm sm:w-full sm:p-6">
            <div>
                <div class="mx-auto flex items-center justify-center rounded-full h-12 w-12 bg-red-100">
                    <svg class="h-8 w-8 text-red-600"  fill="currentColor" viewBox="0 0 20 20">
                    <path fill-rule="evenodd" d="M8.257 3.099c.765-1.36 2.722-1.36 3.486 0l5.58 9.92c.75 1.334-.213 2.98-1.742 2.98H4.42c-1.53 0-2.493-1.646-1.743-2.98l5.58-9.92zM11 13a1 1 0 11-2 0 1 1 0 012 0zm-1-8a1 1 0 00-1 1v3a1 1 0 002 0V6a1 1 0 00-1-1z" clip-rule="evenodd"></path>
                    </svg>
                </div>
                <div class="mt-3 text-center sm:mt-5">
                    <h3 class="text-lg leading-6 font-medium text-gray-900">
                        Keine Information verfügbar
                    </h3>
                    <div class="mt-2">
                        <p class="text-sm leading-5 text-gray-500">
                            Wir konnten noch keine Information bezüglich der Sturmwarnanlagen laden.
                        </p>
                    </div>
                </div>
            </div>
            <div class="mt-5 sm:mt-6">
                <span class="flex w-full rounded-md shadow-sm">
                    <button @click="this.$store.dispatch('autorefresh/refresh')" type="button" class="inline-flex justify-center w-full rounded-md border border-transparent px-4 py-2 bg-gray-600 text-base leading-6 font-medium text-white shadow-sm hover:bg-gray-500 focus:outline-none focus:border-gray-700 focus:shadow-outline-gray transition ease-in-out duration-150 sm:text-sm sm:leading-5">
                        Daten laden...
                    </button>
                </span>
            </div>
        </div>
    </div>
</template>

<script>
import stwaListItem from './stwa/ListItem'

export default {
    name: 'stwaList',
    components: {
        stwaListItem
    },
    computed: {
        stwaData () {
            return this.$store.getters['stwa/data']
        }
    },
}
</script>

<style scoped>
</style>
