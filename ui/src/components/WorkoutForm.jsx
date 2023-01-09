import React, { useState } from "react";
import { UseAuthContext } from "../hooks/useAuthContect";
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";

const WorkoutForm = () => {
  const [type, setType] = useState("");
  const [reps, setReps] = useState("");
  const [load, setLoad] = useState("");
  const [error, setError] = useState(null);

  const { user } = UseAuthContext()
  const { dispatch } = UseWorkoutsContext()
  const handleSubmit = async (e) => {
    e.preventDefault();
    if (!user) {
      setError("Must be logged in!!!")
      return
    }
    let workout = {
      type,
      reps,
      load,
    };
    const res = await fetch("http://localhost:4000/api/", {
      method: "POST",
      body: JSON.stringify(workout),
      headers: {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${user.token}`
      },
    });
    if (!res.ok) {
      const errMsg = await res.text()
      setError(errMsg)
    }
    const json = await res.json();
    if (!res.ok) {
      setError("Unable to create document");
    }
    if (res.ok) {
      setError(null);
      dispatch({ type: "CREATE_WORKOUT", payload: json })
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
        min={0}
      />
      <label>Reps : </label>
      <input
        type="number"
        onChange={(e) => setReps(e.target.value)}
        value={reps}
        required
        min={0}
      />
      <button type="submit">Add Workout</button>
      {error && <div className="error">{error}</div>}
    </form>
  );
};

export default WorkoutForm;
