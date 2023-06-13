import React, { useContext } from "react";
import { Link, NavLink } from "react-router-dom";
import { LoginContext, UserProfileContext } from '../../App';
import'./NavBar.css';

function Navbar() {
  const { loggedIn } = useContext(LoginContext);
  const { userProfile } = useContext(UserProfileContext);

  return (
    <header>
      <div className="container">
        <NavLink className="nav-link" to="/">
          <h1 className="asd">MIRANDA</h1>
        </NavLink>
        <div className="contenedorBotones">
          {loggedIn ? (
            <NavLink className="nav-link" to="/profile">
              <button className="boton">Hola {userProfile.name}</button>
            </NavLink>
          ) : (
            <NavLink className="nav-link" to="/login">
              <button className="boton">Iniciar sesion</button>
            </NavLink>
          )}
        </div>
      </div>
    </header>
  );
}


export default Navbar;
