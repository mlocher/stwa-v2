<template>
    <div v-if="stationData" class="bg-white shadow overflow-hidden sm:rounded-md">
        <ul v-for="station in stationData.features" :key="station.properties.name">
            <li class="block hover:bg-gray-50 focus:outline-none focus:bg-gray-50 transition duration-150 ease-in-out">
                <StationDetails :name="station.properties.name" :location="station.geometry.coordinates" :state="station.properties.state" />
            </li>
        </ul>
    </div>
</template>

<script>
import StationDetails from './StationDetails.vue'

export default {
    components: {
        StationDetails
    },
    props: ['stationData']
}
</script>

<style scoped>
</style>
