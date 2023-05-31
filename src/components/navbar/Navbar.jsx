import React, { useState } from "react";
import "./Navbar.css";
import Log from "../Login/Log";
import Login from "../Login/Login";
import { Link, NavLink } from "react-router-dom";

function Navbar() {
  const [EstadoLogin, CambiarEstadoLogin] = useState(false);
  return (
    <header>
      <div className="container">
        <NavLink className="nav-link" to="/">
          <h1 className="asd">MIRANDA</h1>
        </NavLink>
        <div className="contenedorBotones">
          <button
            onClick={() => CambiarEstadoLogin(!EstadoLogin)}
            className="boton"
          >
            Iniciar sesion
          </button>
          <Log estadolog={EstadoLogin} cambiarestadolog={CambiarEstadoLogin}>
            <Login />
          </Log>
        </div>
      </div>
    </header>
  );
}

export default Navbar;
