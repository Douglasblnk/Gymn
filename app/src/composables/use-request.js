import Axios from 'axios'

const API_URL = import.meta.env.VITE_API_URL

export default () => {
  const requestLoading = ref(false)
  const requestError = ref(null)
  const requestResult = ref(null)

  const clearRequest = () => {
    requestLoading.value = false
    requestError.value = null
    requestResult.value = null
  }

  const handleError = (err) => {
    const { data } = err.response

    return {
      error: data.Error,
      statusCode: data.StatusCode,
    }
  }

  const executeRequest = async ({ method = 'POST', data, headers, params, uri }) => {
    try {
      requestLoading.value = true

      const response = await Axios({
        method,
        url: `${API_URL}/${uri}`,
        data,
        headers,
        params,
      })

      requestResult.value = response.data
    }
    catch (err) {
      const { setAlert } = useAlert()

      setAlert({ text: JSON.stringify(err) })
      requestError.value = handleError(err)
    }
    finally {
      requestLoading.value = false
    }
  }

  return {
    requestLoading,
    requestError,
    requestResult,
    clearRequest,
    executeRequest,
  }
}
