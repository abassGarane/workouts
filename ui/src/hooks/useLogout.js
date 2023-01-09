import { UseAuthContext } from "./useAuthContect"
import { UseWorkoutsContext } from "./useWorkoutsContext"

export const useLogout = () => {
  const { dispatch } = UseAuthContext()
  const { dispatch: wkoutsDispatch } = UseWorkoutsContext()
  const logout = () => {
    localStorage.removeItem("user")
    dispatch({ type: "LOGOUT" })
    wkoutsDispatch({ type: "SET_WORKOUTS", payload: null })
  }
  return { logout }
}
