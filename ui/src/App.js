import { Route, Routes, HashRouter, Navigate } from "react-router-dom";
import Navbar from "./components/Navbar";
import { UseAuthContext } from "./hooks/useAuthContect";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Signup from "./pages/Signup";

function App() {

  const { user } = UseAuthContext()
  return (
    <div className="App">
      <HashRouter>
        <Navbar />
        <div className="pages">
          <Routes>
            <Route path="/" element={user ? <Home /> : <Navigate to="/login" />} />
            <Route path="/login" element={!user ? <Login /> : <Navigate to="/" />} />
            <Route path="/signup" element={!user ? <Signup /> : <Navigate to="/" />} />
          </Routes>
        </div>
      </HashRouter>
    </div>
  );
}

export default App;
