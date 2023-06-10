import React from "react";
import { useState } from "react";
import { Routes, Route } from 'react-router-dom'
import HotelList from "./components/HotelList/HotelList";
import HotelDetails from "./components/HotelDetails/HotelDetails";
import Login from "./components/Login/Login";

export const LoginContext = React.createContext();

const App = () => {

  const [loggedIn, setLoggedIn] = useState(false);

  const handleLogout = () => {
    localStorage.removeItem("token");
    setLoggedIn(false);
  };

  return (
    <div>
      <LoginContext.Provider value={{ loggedIn, handleLogout }}>
        <Routes>
          <Route path="/" exact element={<HotelList />} />
          <Route path="/hotel/:id" element={<HotelDetails />} />
          <Route path="/login" element={<Login />} />
        </Routes>
      </LoginContext.Provider>
    </div>
  );
};

export default App;
