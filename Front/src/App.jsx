import React from "react";
import Home from "./pages/Home";
import Loginb from "./pages/Loginb";
import Habitaciones from "./pages/Habitaciones/Habitaciones";
import { Route, Routes } from "react-router-dom";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/habitaciones" element={<Habitaciones />} />
        <Route path="/Iniciar sesion" element={<Loginb />} />
      </Routes>
    </>
  );
}

export default App;
