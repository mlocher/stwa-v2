<template>
  <div class="antialiased">
    <nav class="relative z-10 bg-gray-800">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex items-center justify-between h-16">
          <div class="flex items-center">
            <div class="flex-shrink-0 rounded-full bg-gray-200">
              <svg fill="currentColor" viewBox="0 0 20 20" class="w-6 h-6 m-2 text-gray-800">
                <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z"></path>
              </svg>
            </div>
            <div class="hidden md:block">
              <div class="ml-10 flex items-baseline">
                <a href="#map" v-on:click="navigateTo('map')" :class="{'text-white bg-gray-900': activeComponent =='map', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != 'map'}" class="px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700">Karte</a>
                <a href="#list" v-on:click="navigateTo('list')" :class="{'text-white bg-gray-900': activeComponent =='list', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != 'list'}" class="ml-4 px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700">Liste</a>
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
                    <div class="flex justify-between block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100">
                      <span>Auto Refresh?</span>
                      <span class="relative inline-flex items-center justify-center flex-shrink-0 h-5 w-10 cursor-pointer focus:outline-none" role="checkbox" tabindex="0" v-on:click="toggleAutoRefresh()" v-on:keydown.space.prevent="toggleAutoRefresh()" :aria-checked="autoRefresh.toString()">
                        <span aria-hidden="true" :class="{ 'bg-indigo-600': autoRefresh, 'bg-gray-200': !autoRefresh }" class="absolute h-4 w-9 mx-auto rounded-full transition-colors ease-in-out duration-200"></span>
                        <span aria-hidden="true" :class="{ 'translate-x-5': autoRefresh, 'translate-x-0': !autoRefresh, 'shadow-outline border-blue-300': autoRefreshFocused }" class="absolute left-0 inline-block h-5 w-5 border border-gray-200 rounded-full bg-white shadow transform transition-transform ease-in-out duration-200"></span>
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="-mr-2 flex md:hidden">
            <button @click="mobileMenuOpen = !mobileMenuOpen" class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:bg-gray-700 focus:text-white">
              <svg class="h-6 w-6" stroke="currentColor" fill="none" viewBox="0 0 24 24">
                <path :class="{'hidden': mobileMenuOpen, 'inline-flex': !mobileMenuOpen }" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
                <path :class="{'hidden': !mobileMenuOpen, 'inline-flex': mobileMenuOpen }" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>
          </div>
        </div>
      </div>
      <div :class="{'block': mobileMenuOpen, 'hidden': !mobileMenuOpen}" class="md:hidden">
        <div class="px-2 pt-2 pb-3 sm:px-3">
          <a href="#map" v-on:click="navigateTo('map')" :class="{'text-white bg-gray-900': activeComponent =='map', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != 'map'}" class="block px-3 py-2 rounded-md text-base font-medium focus:outline-none focus:text-white focus:bg-gray-700">Karte</a>
          <a href="#list" v-on:click="navigateTo('list')" :class="{'text-white bg-gray-900': activeComponent =='list', 'text-gray-300 hover:text-white hover:bg-gray-700': activeComponent != 'list'}" class="mt-1 block px-3 py-2 rounded-md text-base font-medium focus:outline-none focus:text-white focus:bg-gray-700">Liste</a>
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
            <div class="flex justify-between block px-3 py-2 rounded-md text-base font-medium text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:text-white focus:bg-gray-700">
              <span>Auto Refresh?</span>
              <span class="relative inline-flex items-center justify-center flex-shrink-0 h-5 w-10 cursor-pointer focus:outline-none" role="checkbox" tabindex="0" v-on:click="toggleAutoRefresh()" v-on:keydown.space.prevent="toggleAutoRefresh()" :aria-checked="autoRefresh.toString()">
                <span aria-hidden="true" :class="{ 'bg-indigo-600': autoRefresh, 'bg-gray-200': !autoRefresh }" class="absolute h-4 w-9 mx-auto rounded-full transition-colors ease-in-out duration-200"></span>
                <span aria-hidden="true" :class="{ 'translate-x-5': autoRefresh, 'translate-x-0': !autoRefresh, 'shadow-outline border-blue-300': autoRefreshFocused }" class="absolute left-0 inline-block h-5 w-5 border border-gray-200 rounded-full bg-white shadow transform transition-transform ease-in-out duration-200"></span>
              </span>
            </div>
          </div>
        </div>
      </div>
    </nav>
    <main>
      <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <keep-alive>
          <StationList v-if="activeComponent == 'list'" :stationData="stwa.geoJSON" />
          <Map v-if="activeComponent == 'map'" v-on:map-loaded="mapLoaded()" :stationData="stwa.geoJSON" />
        </keep-alive>
      </div>
    </main>
  </div>
</template>

<script>
import axios from 'axios'
import StationList from './components/StationList.vue'
import Map from './components/Map.vue'

export default {
  name: 'stwaApp',
  components: {
    StationList,
    Map
  },
  data () {
    return {
      menuOpen: false,
      mobileMenuOpen: false,
      activeComponent: 'map',
      focused: false,
      stwa: {
        geoJSON: {}
      },
      autoRefresh: true,
      autoRefreshFocused: false,
      autoRefreshIntervall: 15000
    }
  },
  computed: {
    navigationClasses: function () {
        return {
            'text-white bg-gray-900': this.activeComponent == "map",
            'text-gray-300 hover:text-white hover:bg-gray-700': this.activeComponent == "list",
        }
    }
  },
  methods: {
    navigateTo: function(target) {
      switch (target) {
        case 'map':
          this.activeComponent = 'map'
          break
        case 'list':
          this.activeComponent = 'list'
          break
        default:
          this.activeComponent = 'map'
      }
    },
    toggleAutoRefresh: function () {
      this.autoRefresh = !this.autoRefresh
      if (this.autoRefresh) {
        this.fetchStwaData()
        this.autoRefreshID = setInterval(this.fetchStwaData, this.autoRefreshIntervall)
      } else {
        clearInterval(this.autoRefreshID)
      }
    },
    fetchStwaData: function () {
      this.stwa.isLoading = true
      axios.get('https://stwa.locher.at/api/stwa')
        .then(response => {
          this.stwa.geoJSON = response.data
        })
        .catch(error => {
          this.stwa.isError = true
          this.stwa.error = error
        })
        .finally(() => this.stwa.isLoading = false)
    },
    mapLoaded: function () {
      this.fetchStwaData()
      if (this.autoRefresh) {
        this.autoRefreshID = setInterval(this.fetchStwaData, this.autoRefreshIntervall)
      }
    }
  }
}
</script>

<style>
@import 'https://rsms.me/inter/inter.css';

@tailwind base;
@tailwind components;
@tailwind utilities;
</style>
