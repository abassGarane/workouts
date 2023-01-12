import { useState } from "react"
import { UseAuthContext } from "./useAuthContect"

export const useSignup = () => {
  const [error, setError] = useState(null)
  const [loading, setLoading] = useState(false)
  const { dispatch } = UseAuthContext()
  const signup = async (name, email, password) => {
    setLoading(true)
    setError(null)
    const res = await fetch("/auth/signup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        name,
        email,
        password
      })
    })
    const json = await res.json()

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
    signup, loading, error
  }
}
