import { useState } from "react"
import { UseAuthContext } from "./useAuthContect"

export const useLogin = () => {
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(false)
  const { dispatch } = UseAuthContext()
  const login = async (email, password) => {
    setLoading(true)
    setError(null)
    const res = await fetch("/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email,
        password
      })
    })
    const json = await res.json()
    console.log(json)
    if (!res.ok) {
      setLoading(false)
      setError(json.message)
    }
    if (res.ok) {
      console.log(json)
      localStorage.setItem("user", JSON.stringify(json))
      dispatch({ type: "LOGIN", payload: json })
      setLoading(false)
    }
  }
  return {
    login, loading, error
  }
}
