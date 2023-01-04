import React from "react";

const WorkoutDetails = ({ workout }) => {
  const handleDelete = async () => {
    const res = await fetch(`api/${workout.id}/`, { method: "DELETE" })
    if (!res.ok) {
      alert("unable to delete workout")
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
