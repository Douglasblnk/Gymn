const { setAlert } = useAlert()

export default () => {
  const watchError = (error) => {
    watch(
      () => error.value,
      err => err && setAlert({
        text: `${err.statusCode} - ${err.error}`,
        type: 'negative',
      }),
    )
  }

  return {
    watchError,
  }
}
