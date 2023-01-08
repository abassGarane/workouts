
import { createContext, useReducer } from "react";
export const AuthContext = createContext()
export const authReducder = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      return {
        user: action.payload
      }
    case "LOGOUT":
      return {
        user: null
      }
    default:
      return {
        state
      }

  }
}
export const AuthContextProvider = ({ children }) => {
  const [state, dispatch] = useReducer(authReducder, {
    user: null
  })
  console.log("Auth state", state)
  return (
    <AuthContext.Provider value={{ ...state, dispatch }}>
      {children}
    </AuthContext.Provider>
  )

}
