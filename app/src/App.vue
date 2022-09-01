<script setup>
const { routeTransition } = useRouteTransition()
const { isShowing, alertText, alertType } = useAlert()
</script>

<template>
  <RouterView v-slot="{ Component }">
    <Suspense>
      <transition
        :name="routeTransition"
        mode="out-in"
      >
        <component :is="Component" />
      </transition>
    </Suspense>
  </RouterView>

  <GAlert
    :is-showing="isShowing"
    :type="alertType"
    :text="alertText"
  />
</template>
