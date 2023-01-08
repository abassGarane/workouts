import React, { useEffect } from "react";
import WorkoutDetails from "../components/WorkoutDetails";
import WorkoutForm from "../components/WorkoutForm";
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";

const Home = () => {
  // const [workouts, setWorkouts] = useState(null);
  const { workouts, dispatch } = UseWorkoutsContext()
  useEffect(() => {
    const fetchWorkouts = async () => {
      const res = await fetch("/api");
      const json = await res.json();
      if (res.ok) {
        console.log(json);
        // setWorkouts(json);
        console.log(json)
        dispatch({ type: "SET_WORKOUTS", payload: json })
      }
    };
    fetchWorkouts();
  }, [dispatch]);

  return (
    <div className="home">
      <div>
        {workouts &&
          workouts.map((workout) => (
            <WorkoutDetails key={workout.id} workout={workout} />
          ))}
      </div>
      <WorkoutForm />
    </div>
  );
};

export default Home;
