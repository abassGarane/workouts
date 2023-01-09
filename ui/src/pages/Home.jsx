import React, { useEffect } from "react";
import WorkoutDetails from "../components/WorkoutDetails";
import WorkoutForm from "../components/WorkoutForm";
import { UseAuthContext } from "../hooks/useAuthContect";
import { UseWorkoutsContext } from "../hooks/useWorkoutsContext";

const Home = () => {
  // const [workouts, setWorkouts] = useState(null);
  const { workouts, dispatch } = UseWorkoutsContext()
  const { user } = UseAuthContext()
  useEffect(() => {
    const fetchWorkouts = async () => {
      const res = await fetch("/api", {
        headers: {
          "Authorization": `Bearer ${user.token}`
        }
      });
      const json = await res.json();
      if (res.ok) {
        dispatch({ type: "SET_WORKOUTS", payload: json })
      }
    };
    if (user) {
      fetchWorkouts();
    }
  }, [dispatch, user]);

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
