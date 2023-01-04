import React, { useState } from "react";
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";

const WorkoutForm = () => {
  const [type, setType] = useState("");
  const [reps, setReps] = useState("");
  const [load, setLoad] = useState("");
  const [error, setError] = useState(null);

  const { dispatch } = UseWorkoutsContext()
  const handleSubmit = async (e) => {
    e.preventDefault();
    let workout = {
      type,
      reps,
      load,
    };
    console.log(JSON.stringify(workout))
    const res = await fetch("http://localhost:4000/api/", {
      method: "POST",
      body: JSON.stringify(workout),
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (!res.ok) {
      const errMsg = await res.text()
      setError(errMsg)
    }
    const json = await res.json();
    if (!res.ok) {
      setError("Unable to create document");
      console.log(res.body)
    }
    if (res.ok) {
      setError(null);
      dispatch({ type: "CREATE_WORKOUT", payload: json })
      console.log("New workout created", json);
    }
    setType("");
    setLoad("");
    setReps("");
  };

  return (
    <form className="create" onSubmit={handleSubmit}>
      <h3>Add a New Workout</h3>
      <label>Exercise Type :</label>
      <input
        type="text"
        onChange={(e) => setType(e.target.value)}
        value={type}
        required
      />
      <label>Load (in Kgs):</label>
      <input
        type="number"
        onChange={(e) => setLoad(e.target.value)}
        value={load}
        required
      />
      <label>Reps : </label>
      <input
        type="number"
        onChange={(e) => setReps(e.target.value)}
        value={reps}
        required
      />
      <button type="submit">Add Workout</button>
      {error && <div className="error">{error}</div>}
    </form>
  );
};

export default WorkoutForm;
