//fix any later
export function login(payload: any) {
  const requestOptions = {
    method: 'POST',
    body: JSON.stringify(payload)
    // headers: {
    //   'Content-Type': 'application/json'
    // }
  }

  return fetch('http://localhost:8081/users/login', requestOptions)
    .then((response) => response.json())
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
