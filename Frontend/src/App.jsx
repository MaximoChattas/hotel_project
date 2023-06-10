import React from "react";
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import HotelList from "./components/HotelList";
import HotelDetails from "./Components/HotelDetails";

const App = () => {
  return (
   <div>
   <BrowserRouter>
   <Routes>
    <Route path="/" exact element={<HotelList />} />
    <Route path="/hotel/:id" element={<HotelDetails />} />
   </Routes>
   </BrowserRouter>
   </div>
  );
};

export default App;
