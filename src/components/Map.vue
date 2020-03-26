<template>
    <div id="map" class="absolute z-0 top-0 bottom-0 left-0 w-full mt-16"></div>
</template>

<script>
import mapboxgl from 'mapbox-gl'
mapboxgl.accessToken = 'pk.eyJ1IjoibWxvY2hlciIsImEiOiJjazVmcG40dnowdGgyM2xtY3RxY2k3dm5oIn0.d13YAqiMPFPRtYBK7hJLSQ'

import resolveConfig from 'tailwindcss/resolveConfig'
import tailwindConfig from '../../tailwind.config'
const tailwind = resolveConfig(tailwindConfig)

import markerLight from '../assets/img/marker-stwa-light.svg'
import markerDark from '../assets/img/marker-stwa-dark.svg'
import markerGreen from '../assets/img/marker-stwa-green.svg'

export default {
    name: 'stwaMap',
    props: ['stationData'],
    data () {
        return {
            map: null,
            stwaMarker: null
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

            let stwaMarkerLight = new Image(16, 16)
            stwaMarkerLight.addEventListener('load', () => {
                this.map.addImage('stwaMarkerLight', stwaMarkerLight)
            })
            stwaMarkerLight.src = markerLight

            let stwaMarkerDark = new Image(16, 16)
            stwaMarkerDark.addEventListener('load', () => {
                this.map.addImage('stwaMarkerDark', stwaMarkerDark)
            })
            stwaMarkerDark.src = markerDark

            let stwaMarkerGreen = new Image(16, 16)
            stwaMarkerGreen.addEventListener('load', () => {
                this.map.addImage('stwaMarkerGreen', stwaMarkerGreen)
            })
            stwaMarkerGreen.src = markerGreen

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
                'paint': {
                    'circle-color': ['case',
                        ['==', ['get', 'state'], 'standby'], tailwind.theme.colors.gray['200'],
                        ['==', ['get', 'state'], 'advance_warning'], tailwind.theme.colors.yellow['300'],
                        ['==', ['get', 'state'], 'gale_warning'], tailwind.theme.colors.red['500'],
                        ['==', ['get', 'state'], 'out_of_order'], tailwind.theme.colors.blue['200'],
                        tailwind.theme.colors.gray['200'] ],
                    'circle-radius': 10
                }
            });

            this.map.addLayer({
                'id': 'stwa-marker',
                'type': 'symbol',
                'source': 'stwa',
                'layout': {
                    'icon-image': ['case', 
                        ['==', ['get', 'state'], 'standby'], 'stwaMarkerGreen',
                        ['==', ['get', 'state'], 'advance_warning'], 'stwaMarkerLight',
                        ['==', ['get', 'state'], 'gale_warning'], 'stwaMarkerLight',
                        ['==', ['get', 'state'], 'out_of_order'], 'stwaMarkerDark',
                        'stwaMarkerDark'],
                    'icon-allow-overlap': true,
                    'text-field': ['get', 'name'],
                    'text-size': 10,
                    'text-variable-anchor': ['left', 'right'],
                    'text-radial-offset': 1.25,
                    'text-optional': true
                },
                'paint': {
                    'text-color': tailwind.theme.colors.gray['700'],
                    'text-halo-blur': 2,
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
