<script>
import mapboxgl from 'mapbox-gl'

import resolveConfig from 'tailwindcss/resolveConfig'
import tailwindConfig from '../../../tailwind.config'
const tailwind = resolveConfig(tailwindConfig)

import Vue from 'vue'
import iconComponent from './Icon'
var Icon = Vue.extend(iconComponent)

export default {
    name: 'stwaData',
    computed: {
        stwaData () {
            return this.$store.getters['stwa/data']
        }
    },
    watch: {
        stwaData: function(data) {
            if (this.$store.map && this.$store.map.getSource('stwa')) {
                this.$store.map.getSource('stwa').setData(data)
            }
        }
    },
    methods: {
        initLayer: function () {
            let map = this.$store.map

            map.addSource('stwa', {
                type: 'geojson',
                cluster: false,
                data: this.stwaData,
                attribution: 'STWA data <a href="https://www.lsz-b.at/" target="_blank">&copy; LSZ Burgenland</a>'
            })

            map.addLayer({
                id: 'stwaStatus',
                type: 'circle',
                source: 'stwa',
                paint: {
                    'circle-color': ['case',
                        ['==', ['get', 'state'], 'standby'], tailwind.theme.colors.green[200],
                        ['==', ['get', 'state'], 'advance_warning'], tailwind.theme.colors.yellow[300],
                        ['==', ['get', 'state'], 'gale_warning'], tailwind.theme.colors.red[500],
                        ['==', ['get', 'state'], 'out_of_order'], tailwind.theme.colors.blue[200],
                        tailwind.theme.colors.gray[200]
                    ],
                    'circle-radius': 11,
                }
            })

            map.addLayer({
                id: 'stwaMarker',
                type: 'symbol',
                source: 'stwa',
                layout: {
                    'icon-image': 'stwa-15',
                    'icon-allow-overlap': true,
                    'icon-ignore-placement': true,
                    'text-field': ['get', 'name'],
                    'text-size': 12,
                    'text-variable-anchor': ['left', 'right'],
                    'text-radial-offset': 1,
                    'text-optional': true
                    },
                paint: {
                    'text-color': tailwind.theme.colors.gray[700],
                }
            })
        },
        initPopups: function () {
            let map = this.$store.map
            let popup = new mapboxgl.Popup({
                closeButton: false,
                closeOnClick: false,
                offset: 16
            })

            map.on('mouseenter', 'stwaMarker', (event) => {
                map.getCanvas().style.cursor = 'pointer'

                let coordinates = event.features[0].geometry.coordinates.slice()
                let { name, description, state } = event.features[0].properties
                let stateLabel = this.$store.getters['stwa/label'](state)

                // Ensure that if the map is zoomed out such that multiple
                // copies of the feature are visible, the popup appears over
                // the copy being pointed to.
                while (Math.abs(event.lngLat.lng - coordinates[0]) > 180) {
                    coordinates[0] += event.lngLat.lng > coordinates[0] ? 360 : -360;
                }

                var icon = new Icon({
                    propsData: { state: state }
                })
                icon.$mount()

                popup.setLngLat(coordinates)
                    .setHTML(`
                        <div class="sm:flex sm:items-start">
                            <div id="stwaIconContainer" class="mx-auto flex items-center justify-center">
                            </div>
                            <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                                <h3 class="text-md font-medium leading-snug text-gray-900">
                                    STWA ${name} <br />
                                    <span class="text-xs text-gray-500">${stateLabel}</span>
                                </h3>
                                <div class="mt-2">
                                    <p class="text-xs leading-snug text-gray-500">
                                        ${description}
                                    </p>
                                </div>
                            </div>
                        </div>`)
                    .addTo(map)

                    document.getElementById('stwaIconContainer').append(icon.$el)
            })

            map.on('mouseleave', 'stwaMarker', () => {
                map.getCanvas().style.cursor = ''
                popup.remove()
            })
        },
    },
    mounted () {
        this.$store.commit('autorefresh/addCallbackAction', { callback: 'stwa/fetch' })
        this.$store.dispatch('stwa/fetch')

        this.$root.$on('map-loaded', () => {
            let map = this.$store.map

            map.loadImage('/img/marker/stwa-15.png', (error, image) => {
                if (error) throw error
                map.addImage('stwa-15', image)

                this.initLayer()
                this.initPopups()
            })
        })
    },
    render () {
        return null
    }
}
</script>
