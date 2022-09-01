const routeTransition = ref('fade')

export default () => {
  const setRouteTransition = (transition) => {
    routeTransition.value = transition || 'fade'
  }

  return {
    setRouteTransition,
    routeTransition,
  }
}
