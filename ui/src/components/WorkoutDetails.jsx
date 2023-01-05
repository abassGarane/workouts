import React from "react";
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";

const WorkoutDetails = ({ workout }) => {

  const { dispatch } = UseWorkoutsContext()
  const handleDelete = async () => {
    const res = await fetch(`api/${workout.id}`, { method: "DELETE" })
    if (!res.ok) {
      console.log(res)
      alert("unable to delete workout")
    }
    if (res.ok) {
      dispatch({ type: "DELETE_WORKOUT", payload: workout })
    }
  }
  return (
    <div className="workout_details">
      <h4>{workout.type}</h4>
      <p>
        <strong>Load (Kg): </strong>
        {workout.load}
      </p>
      <p>
        <strong>Reps :</strong>
        {workout.reps}
      </p>
      <p>{workout.created_at}</p>
      <span onClick={handleDelete}>delete</span>
    </div>
  );
};

export default WorkoutDetails;
