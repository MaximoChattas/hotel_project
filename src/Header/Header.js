import React, {useState} from 'react';
import './Header.css';
import Log from '../Componentes/Log';
import Reservar from '../Componentes/Reservar';
import Login  from '../Componentes/Login';
function Header() {
  const [EstadoReserva, cambiarEstadoReserva]=useState(false);
  const [EstadoLogin, CambiarEstadoLogin]=useState(false);
  return (
    <header>
          <div className="container">

      <div className="main-header">
       <h1 className='mir'>MIRANDA</h1>
      </div>
      <div className="main-header">
        <div className= "contenedorBotones">
          <button onClick={()=>cambiarEstadoReserva(!EstadoReserva)} className="boton">Reservar fecha</button>
        </div>
        <Reservar
          estado={EstadoReserva}
          cambiarestado={cambiarEstadoReserva}
        >
            <h2>Desde: </h2>
            <h2>Hasta: </h2>
          </Reservar>
      </div>
      <div className="main-header">
        <div className= "contenedorBotones">
          <button onClick={()=>CambiarEstadoLogin(!EstadoLogin)} className="boton">Iniciar sesion</button>
        </div>
        <Log
          estadolog={EstadoLogin}
          cambiarestadolog={CambiarEstadoLogin}
        >
          <Login/>
          </Log>
      </div>
    
      </div>
    </header>
  );
}

export default Header;