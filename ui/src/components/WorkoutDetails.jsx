import TimeAgo from "javascript-time-ago";
import React from "react";

import en from 'javascript-time-ago/locale/en'
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";
import DeleteIcon from '@mui/icons-material/Delete';
import { UseAuthContext } from "../hooks/useAuthContect";
const WorkoutDetails = ({ workout }) => {
  const { user } = UseAuthContext()
  const { dispatch } = UseWorkoutsContext()
  const handleDelete = async () => {
    if (!user) {
      return
    }
    const res = await fetch(`api/${workout.id}`,
      {
        method: "DELETE",
        headers: {
          "Authorization": `Bearer ${user.token}`
        }
      })

    if (res.ok) {
      dispatch({ type: "DELETE_WORKOUT", payload: workout })
    }
  }

  TimeAgo.addLocale(en)
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
