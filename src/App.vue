<template>
    <div class="antialiased">
        <nav class="relative z-10 bg-gray-800">
            <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
                <div class="flex items-center justify-between h-16">
                    <div class="flex items-center">
                        <div class="flex-shrink-0 rounded-full bg-gray-200 text-gray-800">
                            <svg fill="currentColor" viewBox="0 0 20 20" class="w-6 h-6 m-2">
                                <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z"></path>
                            </svg>
                        </div>
                        <div class="hidden md:block">
                            <div class="ml-10 flex items-baseline">
                                <a href="#map" :class="{'text-white bg-gray-900': activeComponent =='#map', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != '#map'}" class="px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700">Karte</a>
                                <a href="#list" :class="{'text-white bg-gray-900': activeComponent =='#list', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != '#list'}" class="ml-4 px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700">Liste</a>
                            </div>
                        </div>
                    </div>
                    <div class="hidden md:block">
                        <div class="ml-4 flex items-center md:ml-6">
                            <div class="ml-3 relative">
                                <div>
                                    <button v-on:click="menuOpen = !menuOpen" class="max-w-xs flex items-center text-sm p-1 border-2 border-transparent text-gray-400 rounded-full hover:text-white focus:outline-none focus:text-white focus:bg-gray-700">
                                        <svg class="h-6 w-6" stroke="currentColor" fill="none" viewBox="0 0 24 24">
                                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"></path>
                                        </svg>
                                    </button>
                                </div>
                                <div v-show="menuOpen" x-transition:enter="transition ease-out duration-100" x-transition:enter-start="transform opacity-0 scale-95" x-transition:enter-end="transform opacity-100 scale-100" x-transition:leave="transition ease-in duration-75" x-transition:leave-start="transform opacity-100 scale-100" x-transition:leave-end="transform opacity-0 scale-95" class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg">
                                    <div class="py-1 rounded-md bg-white shadow-xs">
                                        <AutoRefresh />
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="-mr-2 flex md:hidden">
                        <button @click="menuOpen = !menuOpen" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:bg-gray-700 focus:text-white">
                            <svg class="h-6 w-6" stroke="currentColor" fill="none" viewBox="0 0 24 24">
                                <path :class="{'hidden': menuOpen, 'inline-flex': !menuOpen }" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                                <path :class="{'hidden': !menuOpen, 'inline-flex': menuOpen }" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                            </svg>
                        </button>
                    </div>
                </div>
            </div>
            <div :class="{'block': menuOpen, 'hidden': !menuOpen}" class="md:hidden">
                <div class="px-2 pt-2 pb-3 sm:px-3">
                    <a href="#map" :class="{'text-white bg-gray-900': activeComponent =='#map', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != '#map'}" class="block px-3 py-2 rounded-md text-base font-medium focus:outline-none focus:text-white focus:bg-gray-700">Karte</a>
                    <a href="#list" :class="{'text-white bg-gray-900': activeComponent =='#list', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != '#list'}" class="mt-1 block px-3 py-2 rounded-md text-base font-medium focus:outline-none focus:text-white focus:bg-gray-700">Liste</a>
                </div>
                <div class="pt-4 pb-3 border-t border-gray-700">
                    <div class="flex items-center px-5">
                        <div class="flex-shrink-0">
                            <svg class="h-6 w-6 text-white" stroke="currentColor" fill="none" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"></path>
                            </svg>
                        </div>
                        <div class="ml-3">
                            <div class="text-base font-medium leading-none text-white">Einstellungen</div>
                        </div>
                    </div>
                    <div class="mt-3 px-2">
                        <AutoRefresh />
                    </div>
                </div>
            </div>
        </nav>
        <main>
            <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
                <STWAData />
                <keep-alive>
                    <Map v-if="activeComponent == '#map'" />
                    <List v-else-if="activeComponent == '#list'" />
                </keep-alive>
            </div>
        </main>
    </div>
</template>

<script>
import AutoRefresh from './components/AutoRefresh'
import Map from './components/Map'
import List from './components/List'
import STWAData from './components/stwa/Data'

export default {
    name: 'stwaApp',
    components: {
        AutoRefresh,
        Map,
        List,
        STWAData,
    },
    data () {
        return {
            menuOpen: false,
            activeComponent: document.location.hash != '' ? document.location.hash : '#map',
        }
    },
    mounted () {
        addEventListener('hashchange', (event) => {
            switch (event.target.document.location.hash) {
                case '#list':
                    this.activeComponent = '#list'
                    break
                case '#map':
                default:
                    this.activeComponent = '#map'
                    break
            }
        })
    }
}
</script>

<style>
@import 'https://rsms.me/inter/inter.css';

@tailwind base;
@tailwind components;
@tailwind utilities;
</style>
