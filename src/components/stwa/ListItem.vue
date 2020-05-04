<template>
    <div class="flex items-center px-4 py-4 sm:px-6">
        <div class="min-w-0 flex-1 flex items-center">
            <icon v-bind="properties" />
            <div class="min-w-0 flex-1 px-4 md:grid md:grid-cols-3 md:gap-4">
                <div>
                    <div class="text-sm leading-5 font-medium text-gray-800 truncate">
                        {{ name }}
                    </div>
                    <div class="mt-2 flex items-center text-sm leading-5 text-gray-500">
                        ({{ this.$store.getters['stwa/label'](state) }})
                    </div>
                </div>
                <div class="hidden md:block col-span-2">
                    <div>
                        <div class="flex items-top text-sm leading-5 text-gray-500">
                            <svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400"  fill="currentColor" viewBox="0 0 20 20">
                                <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-7-4a1 1 0 11-2 0 1 1 0 012 0zM9 9a1 1 0 000 2v3a1 1 0 001 1h1a1 1 0 100-2v-3a1 1 0 00-1-1H9z" clip-rule="evenodd"></path>
                            </svg>
                            <span class="truncate">{{ description }}</span>
                        </div>
                        <div class="mt-2 flex items-center text-sm leading-5 text-gray-500">
                            <svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" viewBox="0 0 20 20" fill="currentColor">
                                <path fill-rule="evenodd" d="M4.083 9h1.946c.089-1.546.383-2.97.837-4.118A6.004 6.004 0 004.083 9zM10 2a8 8 0 100 16 8 8 0 000-16zm0 2c-.076 0-.232.032-.465.262-.238.234-.497.623-.737 1.182-.389.907-.673 2.142-.766 3.556h3.936c-.093-1.414-.377-2.649-.766-3.556-.24-.56-.5-.948-.737-1.182C10.232 4.032 10.076 4 10 4zm3.971 5c-.089-1.546-.383-2.97-.837-4.118A6.004 6.004 0 0115.917 9h-1.946zm-2.003 2H8.032c.093 1.414.377 2.649.766 3.556.24.56.5.948.737 1.182.233.23.389.262.465.262.076 0 .232-.032.465-.262.238-.234.498-.623.737-1.182.389-.907.673-2.142.766-3.556zm1.166 4.118c.454-1.147.748-2.572.837-4.118h1.946a6.004 6.004 0 01-2.783 4.118zm-6.268 0C6.412 13.97 6.118 12.546 6.03 11H4.083a6.004 6.004 0 002.783 4.118z" clip-rule="evenodd"></path>
                            </svg>
                            <span class="truncate">Lat {{ location[1] }}, Lon {{ location[0] }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div>
            <svg class="h-5 w-5 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd"/>
            </svg>
        </div>
    </div>
</template>

<script>
import icon from './Icon'

export default {
    name: 'list',
    components: {
        icon
    },
    props: ['geometry', 'properties'],
    data: function () {
        return {
            state: this.properties.state,
            name: this.properties.name,
            description: this.properties.description,
            height: this.properties.height,
            location: this.geometry.coordinates
        }
    }
}
</script>

<style scoped>
</style>
