import { UseAuthContext } from "./useAuthContect"

export const useLogout = () => {
  const { dispatch } = UseAuthContext()
  const logout = () => {
    localStorage.removeItem("user")
    dispatch({ type: "LOGOUT" })
  }
  return { logout }
}
