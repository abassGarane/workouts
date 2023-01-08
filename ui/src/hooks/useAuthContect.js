import { useContext } from "react"
import { AuthContext } from "../context/auth_context"

export const UseAuthContext = () => {
  const ctx = useContext(AuthContext)
  if (!ctx) {
    throw Error("Use auth context can only be used inside a WorkoutContextProvider")
  }
  return ctx
}
