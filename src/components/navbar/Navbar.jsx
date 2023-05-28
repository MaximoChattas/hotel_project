import React, { useState } from "react";
import "./Navbar.css";
import Log from "../Login/Log";
import Login from "../Login/Login";

function Navbar() {
  const [EstadoLogin, CambiarEstadoLogin] = useState(false);
  return (
    <header>
      <div className="container">
          <h1 >MIRANDA</h1>
          <div className="contenedorBotones">
            <button onClick={() => CambiarEstadoLogin(!EstadoLogin)} className="boton">
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
