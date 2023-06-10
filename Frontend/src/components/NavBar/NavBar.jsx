import React, { useContext } from "react";
import { Link, NavLink } from "react-router-dom";
import Login from "../Login/Login"
import Style from "../NavBar/NavBar.css"

function Navbar() {

  return (
    <header>
      <div className="container">
        <NavLink className="nav-link" to="/">
          <h1 className="asd">MIRANDA</h1>
        </NavLink>
        <div className="contenedorBotones">
        <NavLink className="nav-link" to="/login">
              <button className="boton">Iniciar sesion</button>
        </NavLink>
        </div>
      </div>
    </header>
  );
}

export default Navbar;
