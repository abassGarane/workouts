import TimeAgo from "javascript-time-ago";
import en from 'javascript-time-ago/locale/en'
import React from "react";
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";
import DeleteIcon from '@mui/icons-material/Delete';

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
  TimeAgo.addDefaultLocale(en)
  let timeAgo = new TimeAgo("en-US")
  let formatedDate = timeAgo.format(new Date(workout.created_at))
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
      <p>{formatedDate}</p>
      <span onClick={handleDelete}><DeleteIcon /></span>
    </div>
  );
};

export default WorkoutDetails;
