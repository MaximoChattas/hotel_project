import React from "react";
import Home from "./pages/Home";
import Habitaciones from "./pages/Habitaciones";
import "./index.css";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/habitaciones" element={<Habitaciones />} />
      </Routes>
    </>
  );
}

export default App;
