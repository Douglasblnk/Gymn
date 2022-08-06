module.exports = {
  extends: [
    '@antfu',
    '.eslintrc-auto-import.json',
  ],
  rules: {
    'array-bracket-spacing': [ 'error', 'always' ],
    'object-curly-spacing': [ 'error', 'always' ],
    'computed-property-spacing': [ 'error', 'always' ],
    'vue/no-multiple-template-root': 'off',
    'template-curly-spacing': [ 'error', 'always' ],
    'curly': [ 'error', 'multi' ],
    'comma-dangle': [ 'error', {
      arrays: 'always-multiline',
      objects: 'always-multiline',
      imports: 'only-multiline',
      exports: 'always-multiline',
      functions: 'always-multiline',
    } ],
    'max-len': 'off',
    'vue/max-attributes-per-line': [ 'error', {
      singleline: {
        max: 1,
      },
      multiline: {
        max: 1,
      },
    } ],
  },
  globals: {
    defineProps: 'readonly',
    defineEmits: 'readonly',
    defineExpose: 'readonly',
    withDefaults: 'readonly',
  },
}
