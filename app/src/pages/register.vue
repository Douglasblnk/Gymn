<script setup>
const register = () => {
  // todo
}
</script>

<template>
  <div
    un-flex="~ col"
    un-h-screen
  >
    <GTitle
      text="Cadastrar."
      small
      un-flex-grow
    />

    <RegisterFields
      un-flex-grow
      @register="register"
    />
  </div>
</template>

<route lang="yaml">
alias: /registrar
meta:
  transition: slide-down
</route>
