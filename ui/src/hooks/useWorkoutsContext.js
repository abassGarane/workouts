import { useContext } from "react"
import { WorkoutContext } from "../context/workout_context"

export const UseWorkoutsContext = () => {
  const ctx = useContext(WorkoutContext)
  if (!ctx) {
    throw Error("Use workouts context can only be used inside a WorkoutContextProvider")
  }
  return ctx
}
