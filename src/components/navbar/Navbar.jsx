import React, { useState } from "react";
import "./Navbar.css";
import Log from "../Login/Log";
import Reservar from "../Login/Reservar";
import Login from "../Login/Login";

function Navbar() {
  const [EstadoReserva, cambiarEstadoReserva] = useState(false);
  const [EstadoLogin, CambiarEstadoLogin] = useState(false);
  return (
    <header>
      <div className="container">
        <div className="main-navbar">
          <h1 className="mir">MIRANDA</h1>
        </div>
        <div className="main-navbar">
          <div className="contenedorBotones">
            <button
              onClick={() => CambiarEstadoLogin(!EstadoLogin)}
              className="boton"
            >
              Iniciar sesion
            </button>
          </div>
          <Log estadolog={EstadoLogin} cambiarestadolog={CambiarEstadoLogin}>
            <Login />
          </Log>
        </div>
      </div>
    </header>
  );
}

export default Navbar;
