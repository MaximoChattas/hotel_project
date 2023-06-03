import React, { useState } from "react";
import "./Navbar.css";
import { Link, NavLink } from "react-router-dom";

function Navbar() {
  return (
    <header>
      <div className="container">
        <NavLink className="nav-link" to="/">
          <h1 className="asd">MIRANDA</h1>
        </NavLink>
        <div className="contenedorBotones">
        <NavLink className="nav-link" to="/Iniciar sesion">
            Iniciar sesion
          </NavLink>
        </div>
      </div>
    </header>
  );
}

export default Navbar;
