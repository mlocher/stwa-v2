<template>
    <div id="map" class="absolute z-0 top-0 bottom-0 left-0 w-full mt-16"></div>
</template>

<script>
import mapboxgl from 'mapbox-gl'
mapboxgl.accessToken = process.env.VUE_APP_MAPBOX_TOKEN

export default {
    name: 'stwaMap',
    mounted () {
        let map = new mapboxgl.Map({
            container: 'map',
            style: 'mapbox://styles/mapbox/light-v10',
            center: [16.7776, 47.8650],
            zoom: 10
        });

        map.on('load', () => {
            map.fitBounds([[16.6102,47.6294],[16.9419,47.99]]);
            map.addControl(new mapboxgl.NavigationControl());
            map.addControl(new mapboxgl.GeolocateControl({
                positionOptions: {
                    enableHighAccuracy: true
                },
                fitBoundsOptions: {
                    maxZoom: map.getZoom()
                },
                trackUserLocation: true
            }));

            this.$store.map = map
            this.$root.$emit('map-loaded')
        });
    }
}


        // this.popup = new mapboxgl.Popup({
        //     closeButton: false,
        //     closeOnClick: false,
        //     offset: 16
        // });

        // this.map.on('load', () => {
        //     let stwaMarkerLight = new Image(16, 16)
        //     stwaMarkerLight.addEventListener('load', () => {
        //         this.map.addImage('stwaMarkerLight', stwaMarkerLight)
        //     })
        //     stwaMarkerLight.src = markerLight

        //     let stwaMarkerDark = new Image(16, 16)
        //     stwaMarkerDark.addEventListener('load', () => {
        //         this.map.addImage('stwaMarkerDark', stwaMarkerDark)
        //     })
        //     stwaMarkerDark.src = markerDark

        //     let stwaMarkerGreen = new Image(16, 16)
        //     stwaMarkerGreen.addEventListener('load', () => {
        //         this.map.addImage('stwaMarkerGreen', stwaMarkerGreen)
        //     })
        //     stwaMarkerGreen.src = markerGreen

        //     this.map.addSource('stwa', {
        //         type: 'geojson',
        //         cluster: false,
        //         clusterRadius: 10,
        //         clusterMaxZoom: 12,
        //         data: this.stwaData,
        //         attribution: 'STWA data &copy; <a href="https://www.lsz-b.at/">LSZ Burgenland</a>'
        //     });

        //     this.map.addLayer({
        //         'id': 'stwa-status',
        //         'type': 'circle',
        //         'source': 'stwa',
        //         'paint': {
        //             'circle-color': ['case',
        //                 ['==', ['get', 'state'], 'standby'], tailwind.theme.colors.gray['200'],
        //                 ['==', ['get', 'state'], 'advance_warning'], tailwind.theme.colors.yellow['300'],
        //                 ['==', ['get', 'state'], 'gale_warning'], tailwind.theme.colors.red['500'],
        //                 ['==', ['get', 'state'], 'out_of_order'], tailwind.theme.colors.blue['200'],
        //                 tailwind.theme.colors.gray['200'] ],
        //             'circle-radius': 14
        //         }
        //     });

        //     this.map.addLayer({
        //         'id': 'stwa-marker',
        //         'type': 'symbol',
        //         'source': 'stwa',
        //         'layout': {
        //             'icon-image': ['case',
        //                 ['==', ['get', 'state'], 'standby'], 'stwaMarkerGreen',
        //                 ['==', ['get', 'state'], 'advance_warning'], 'stwaMarkerDark',
        //                 ['==', ['get', 'state'], 'gale_warning'], 'stwaMarkerLight',
        //                 ['==', ['get', 'state'], 'out_of_order'], 'stwaMarkerDark',
        //                 'stwaMarkerDark'],
        //             'icon-allow-overlap': true,
        //             //'text-field': ['get', 'name'],
        //             //'text-size': 10,
        //             //'text-variable-anchor': ['left', 'right'],
        //             //'text-radial-offset': 1.25,
        //             //'text-optional': true
        //         },
        //         'paint': {
        //             //'text-color': tailwind.theme.colors.gray['700'],
        //             //'text-halo-blur': 2,
        //         }
        //     });
        // });

        // this.map.on('mouseenter', 'stwa-marker', (e) => {
        //     this.map.getCanvas().style.cursor = 'pointer'

        //     var coordinates = e.features[0].geometry.coordinates.slice();
        //     var name = e.features[0].properties.name;
        //     var description = e.features[0].properties.description;
        //     var stateClasses = null;
        //     var stateLabel = null

        //     switch (e.features[0].properties.state) {
        //         case 'standby':
        //             stateClasses = 'bg-gray-200 text-green-500';
        //             stateLabel = 'Bereitschaft';
        //             break;
        //         case 'advance_warning':
        //             stateClasses = 'bg-yellow-300 text-gray-500';
        //             stateLabel = 'Vorwarnung';
        //             break;
        //         case 'gale_warning':
        //             stateClasses = 'bg-red-500 text-gray-200';
        //             stateLabel = 'Vorwarnung';
        //             break;
        //         case 'out_of_order':
        //             stateClasses = 'bg-blue-200 text-gray-500';
        //             stateLabel = 'Außer Betrieb';
        //             break;
        //         case 'unknown':
        //         default:
        //             stateClasses = 'bg-gray-200 text-gray-500';
        //             stateLabel = 'Unbekannt';
        //             break;
        //     }

        //     // Ensure that if the map is zoomed out such that multiple
        //     // copies of the feature are visible, the popup appears
        //     // over the copy being pointed to.
        //     while (Math.abs(e.lngLat.lng - coordinates[0]) > 180) {
        //         coordinates[0] += e.lngLat.lng > coordinates[0] ? 360 : -360;
        //     }

        //     this.popup
        //         .setLngLat(coordinates)
        //         .setHTML(`<div class="sm:flex sm:items-start">
        //             <div class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full ${stateClasses} sm:mx-0 sm:h-10 sm:w-10">
        //                 <svg fill="currentColor" viewBox="0 0 20 20" class="w-6 h-6">
        //                     <path d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z"></path>
        //                 </svg>
        //             </div>
        //             <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
        //                 <h3 class="text-md font-medium leading-snug text-gray-900">
        //                     STWA ${name} <br />
        //                     <span class="text-xs text-gray-500">${stateLabel}</span>
        //                 </h3>
        //                 <div class="mt-2">
        //                     <p class="text-xs leading-snug text-gray-500">
        //                         ${description}
        //                     </p>
        //                 </div>
        //             </div>
        //             </div>`)
        //         .addTo(this.map);
        // })

        // this.map.on('mouseleave', 'stwa-marker', () => {
        //     this.map.getCanvas().style.cursor = '';
        //     this.popup.remove();
        // })
</script>

<style>
@import "~mapbox-gl/dist/mapbox-gl.css";

.mapboxgl-popup-content {
    @apply shadow-lg rounded-md;
}
</style>
