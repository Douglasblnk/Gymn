<script setup>
const props = defineProps({
  small: Boolean,
  text: {
    type: String,
    default: '',
  },
})

const customText = computed(() => {
  const lastWord = props.text.charAt(props.text.length - 1)

  return {
    text: props.text.replace(lastWord, ''),
    lastWord,
  }
})
</script>

<template>
  <div
    un-flex
    un-justify-center
    un-items-center
    un-py-2xl
  >
    <div
      un-font="black !italic"
      un-text-white
      :class="small ? 'text-8' : 'text-16'"
    >
      {{ customText.text }}<span un-text-secondary>{{ customText.lastWord }}</span>
    </div>
  </div>
</template>
