import React, { useState, useEffect } from 'react'

const App = () => {
  const [msg, setMsg] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/api/greet')
    .then(res => res.json())
    .then(data => {setMsg(data.message)}) // fetches the message from the Go backend. // Here there was an issue, I mistakenly used data.greeting(function name) instead of data.message (key name in the JSON response). Which I found and fixed immediately.
  }, []);
  return (
    <>
      <h1>React Go</h1>
      <h2>{msg}</h2>
    </>  
  )
}

export default App