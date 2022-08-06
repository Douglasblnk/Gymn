import { resolve } from 'path'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import fg from 'fast-glob'
import Components from 'unplugin-vue-components/vite'
import AutoImport from 'unplugin-auto-import/vite'
import Pages from 'vite-plugin-pages'
import { VitePWA } from 'vite-plugin-pwa'
import Unocss from 'unocss/vite'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': [ resolve(__dirname, './src') ],
    },
  },
  optimizeDeps: {
    include: [
      'vue',
      'vue-router',
      'axios',
    ],
  },
  plugins: [
    vue(),

    Pages({
      extendRoute(route) {
        const { path } = route

        if (path === '/')
          return route

        return {
          ...route,
          meta: { auth: true },
        }
      },
    }),

    AutoImport({
      include: [
        /\.vue\??/, // .vue
      ],
      imports: [
        'vue',
        'vue-router',
        'vue/macros',
      ],
      dirs: fg.sync('src/**/composables', { onlyDirectories: true }),
      vueTemplate: true,
      eslintrc: {
        enabled: true,
      },
      dts: 'src/auto-imports.d.ts',
    }),

    VitePWA({
      manifest: {
        name: 'Gymn',
        icons: [],
        start_url: '/index.html',
        display: 'standalone',
      },
    }),

    Components({
      dirs: [ 'src/**/components' ],
      resolvers: [
        IconsResolver(),
      ],
      dts: 'src/components.d.ts',
    }),

    Icons({
      autoInstall: true,
    }),

    Unocss(),
  ],
  server: {
    port: 3000,
  },
})
