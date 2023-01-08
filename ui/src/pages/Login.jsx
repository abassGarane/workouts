import React, { useState } from 'react'
import { useLogin } from '../hooks/useLogin'

const Login = () => {
  const [password, setPassword] = useState("")
  const [email, setEmail] = useState("")
  const { login, error, loading } = useLogin()

  const createUser = async (e) => {
    e.preventDefault()

    await login(email, password)
  }
  return (
    <form className='login' onSubmit={createUser}>
      <h3>Log in</h3>
      <label >Email :</label>
      <input type="email" name="email" value={email} onChange={e => setEmail(e.target.value)} />
      <label >Password:</label>
      <input type="password" name="password" value={password} onChange={e => setPassword(e.target.value)} />
      <button type="submit" disabled={loading}>Log in</button>
      {error && <div className="error">{error}</div>}
    </form>
  )
}

export default Login
