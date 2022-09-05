import {
  defineConfig,
  presetAttributify,
  presetIcons,
  presetUno,
  presetWebFonts,
  transformerDirectives,
  transformerVariantGroup,
} from 'unocss'

import presetCore from './src/unocss'

export default defineConfig({
  presets: [
    presetUno(),
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
    transformerVariantGroup(),
    transformerDirectives(),
  ],
})
