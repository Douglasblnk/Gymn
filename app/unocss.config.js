import { defineConfig } from 'unocss'
import presetWind from '@unocss/preset-wind'
import presetAttributify from '@unocss/preset-attributify'
import presetWebFonts from '@unocss/preset-web-fonts'
import transformerDirective from '@unocss/transformer-directives'
import presetMini from '@unocss/preset-mini'
import presetIcons from '@unocss/preset-icons'
import presetCore from './src/unocss/index'

export default defineConfig({
  presets: [
    presetMini(),
    presetWind(),
    presetCore(),
    presetAttributify(),
    presetIcons({
      extraProperties: {
        display: 'inline-block',
      },
      warn: true,
    }),
    presetWebFonts({
      provider: 'google',
      fonts: {
        nunito: [
          {
            name: 'Nunito',
            weights: [ '400', '700', '800', '900' ],
            italic: true,
          },
        ],
      },
    }),
  ],
  transformers: [
    transformerDirective(),
  ],
})
