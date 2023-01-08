import React, { useState } from 'react'

const Login = () => {
  const [password, setPassword] = useState("")
  const [email, setEmail] = useState("")

  const createUser = async (e) => {
    e.preventDefault()
    let user = {
      email,
      password,
    }
    console.log(user)
    // const res = await fetch("/auth/login", {
    //   method: "POST",
    //   body: JSON.stringify(user),
    // })
    // const data = await res.json()
    // console.log(data)
  }
  return (
    <form className='login' onSubmit={createUser}>
      <h3>Log in</h3>
      <label >Email :</label>
      <input type="email" name="email" value={email} onChange={e => setEmail(e.target.value)} />
      <label >Password:</label>
      <input type="password" name="password" value={password} onChange={e => setPassword(e.target.value)} />
      <button type="submit">Log in</button>
    </form>
  )
}

export default Login
