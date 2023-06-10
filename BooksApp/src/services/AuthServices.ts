// AuthServices.ts
const requestConfig = {
  basePath: 'http://localhost:8081'
}

export function login(payload: any) {
  const requestOptions = createRequestOptions(payload)

  return fetch(`${requestConfig.basePath}/users/login`, requestOptions)
    .then((response) => {
      if (!response.ok) {
        throw new Error('Request failed with status ' + response.status)
      }
      return response.json()
    })
    .then((data) => {
      if (data.error) {
        console.log('Error:', data.message)
      } else {
        console.log(data)
      }
      return data
    })
    .catch((error) => {
      console.log('Error:', error)
      throw error
    })
}

function createRequestOptions(payload: any) {
  return {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(payload)
  }
}
