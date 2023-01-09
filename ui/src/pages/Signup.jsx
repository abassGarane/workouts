import React, { useState } from 'react'
import { useSignup } from '../hooks/useSignup'

const Signup = () => {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")
  const [email, setEmail] = useState("")

  const { signup, loading, error } = useSignup()

  const createUser = async (e) => {
    e.preventDefault()
    let user = {
      name: username,
      email,
      password,
    }
    await signup(user.name, user.email, user.password)
  }
  return (
    <form className='signup' onSubmit={createUser}>
      <h3>Sign up</h3>
      <label >User Name :</label>
      <input type="text" name="username" value={username} onChange={e => setUsername(e.target.value)} />
      <label >Email :</label>
      <input type="email" name="email" value={email} onChange={e => setEmail(e.target.value)} />
      <label >Password:</label>
      <input type="password" name="password" value={password} onChange={e => setPassword(e.target.value)} />
      <button type="submit" disabled={loading}>Sign up</button>
      {error && <div className="error">{error}</div>}
    </form>
  )
}

export default Signup
