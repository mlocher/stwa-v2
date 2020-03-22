<template>
    <div id="map" class="absolute z-0 top-0 bottom-0 left-0 w-full mt-16">
        <div id="stwaMarker" class="hidden rounded-full w-6 h-6 bg-yellow-200 text-gray-500">
            <svg class="h-4 w-4 m-1" viewBox="0 0 20 20" fill="currentColor">
                <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z"></path>
            </svg>
        </div>
    </div>
</template>

<script>
import mapboxgl from 'mapbox-gl'
mapboxgl.accessToken = 'pk.eyJ1IjoibWxvY2hlciIsImEiOiJjazVmcG40dnowdGgyM2xtY3RxY2k3dm5oIn0.d13YAqiMPFPRtYBK7hJLSQ'

export default {
    name: 'stwaMap',
    props: ['stationData'],
    data () {
        return {
            map: null
        }
    },
    watch: {
        stationData: function(data) {
            if (this.map.getSource('stwa')) {
                this.map.getSource('stwa').setData(data)
            }
        }
    },
    mounted () {
        this.map = new mapboxgl.Map({
            container: 'map',
            style: 'mapbox://styles/mapbox/light-v10',
            center: [16.7776, 47.8650],
            zoom: 10
        });

        this.map.on('load', () => {
            this.map.addControl(new mapboxgl.NavigationControl());
            this.map.addControl(new mapboxgl.GeolocateControl({
                positionOptions: {
                    enableHighAccuracy: true
                },
                fitBoundsOptions: {
                    maxZoom: this.map.getZoom()
                },
                trackUserLocation: true
            }));
            this.map.fitBounds([[16.6102,47.6294],[16.9419,47.99]]);

            this.map.loadImage('/img/stwa.png', (error, image) => {
                if (error) throw error;
                this.map.addImage('stwa', image);
            });

            this.map.addSource('stwa', {
                type: 'geojson',
                cluster: false,
                clusterRadius: 10,
                clusterMaxZoom: 12,
                data: null,
                attribution: 'STWA data &copy; <a href="https://www.lsz-b.at/">LSZ Burgenland</a>'
            });

            this.map.addLayer({
                'id': 'stwa-status',
                'type': 'circle',
                'source': 'stwa',
                'layout': {
                    'circle-sort-key': ['case',
                        ['==', ['get', 'state'], 'gale_warning'], 8,
                        ['==', ['get', 'state'], 'advance_warning'], 4,
                        ['==', ['get', 'state'], 'standby'], 2,
                        ['==', ['get', 'state'], 'out_of_order'], 1,
                        0 ],
                },
                'paint': {
                    'circle-color': ['case',
                        ['==', ['get', 'state'], 'standby'], '#0e9f6e',
                        ['==', ['get', 'state'], 'advance_warning'], '#faca15',
                        ['==', ['get', 'state'], 'gale_warning'], '#f05252',
                        ['==', ['get', 'state'], 'out_of_order'], '#c3ddfd',
                        '#e5e7eb' ],
                    'circle-radius': 10
                }
            });

            this.map.addLayer({
                'id': 'stwa-marker',
                'type': 'symbol',
                'source': 'stwa',
                'layout': {
                    'symbol-sort-key': ['case',
                        ['==', ['get', 'state'], 'gale_warning'], 0,
                        ['==', ['get', 'state'], 'advance_warning'], 1,
                        ['==', ['get', 'state'], 'standby'], 2,
                        ['==', ['get', 'state'], 'out_of_order'], 4,
                        8 ],
                    'icon-image': 'stwa',
                    //'icon-image': ['coalesce', ['image', 'stwa'], ['image', 'monument-11']],
                    'icon-allow-overlap': true,
                    'text-field': ['get', 'name'],
                    'text-font': ['Open Sans Semibold'],
                    'text-size': 9,
                    'text-variable-anchor': ['left', 'right'],
                    'text-justify': 'auto',
                    'text-radial-offset': 1,
                    'text-optional': true
                },
                'paint': {
                    'text-color': '#6b7280',
                    'text-halo-blur': 2,
                    'icon-color': ['case',
                        ['==', ['get', 'state'], 'standby'], '#0e9f6e',
                        ['==', ['get', 'state'], 'advance_warning'], '#faca15',
                        ['==', ['get', 'state'], 'gale_warning'], '#f05252',
                        ['==', ['get', 'state'], 'out_of_order'], '#c3ddfd',
                        '#e5e7eb' ],
                }
            });

            this.$emit('map-loaded')
        });
    }
}
</script>

<style scoped>
@import "~mapbox-gl/dist/mapbox-gl.css";
</style>
