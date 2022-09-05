import { autocompleteExtractorAttributify } from '@unocss/preset-attributify'
import { defineConfig } from 'unocss'
import { preflightColors } from './rules/colors'

import theme from './theme'

export default () => {
  const autocompleteExtractors = [
    autocompleteExtractorAttributify,
  ]

  return defineConfig({
    shortcuts: [
      { column: 'flex flex-wrap flex-col h-auto min-h-0 max-h-full' },
      { title: 'text-32px font-bold text-primary' },
      [ /^offset-(.*)$/, ([ , d ]) => {
        // check value is between 1 and 12
        if (d >= 1 && d <= 12) {
          return `m-l-${ (100 / 12) * d }%`
        }
      },
      ],
    ],
    name: 'preset-gymn',
    layers: {
      root: -99,
    },
    theme,
    preflights: [ preflightColors ],
    autocomplete: {
      extractors: autocompleteExtractors,
    },
  })
}
