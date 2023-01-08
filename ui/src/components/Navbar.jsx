import React from 'react'
import { Link } from 'react-router-dom'
import { UseAuthContext } from '../hooks/useAuthContect'
import { useLogout } from '../hooks/useLogout'

const Navbar = () => {
  const { user } = UseAuthContext()
  const { logout } = useLogout()
  const handleLogout = (e) => {
    e.preventDefault()
    logout()
  }
  return (
    <header>
      <div className="container">
        <Link to="/">
          <h1>Muscles Mania</h1>
        </Link>
        <nav>
          {user && (
            <div>
              <button type="button" onClick={handleLogout}>Logout</button>
              <span>{user.user.email}</span>
            </div>
          )}
          {!user && (
            <div>
              <Link to="/login">Login</Link>
              <Link to="/signup">Signup</Link>
            </div>
          )}
        </nav>
      </div>
    </header>
  )
}

export default Navbar
