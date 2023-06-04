import React from "react";
import "./Reservar.css";

const Reservar = ({ children, estado, cambiarestado }) => {
  return (
    <>
      {estado && (
        <div className="Overlay">
          <div className="ContenedorReserva">
            <button onClick={() => cambiarestado(false)} className="Cerrar">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="16"
                height="16"
                fill="currentColor"
                class="bi bi-x"
                viewBox="0 0 16 16"
              >
                <path d="M4.646 4.646a.5.5 0 0 1 .708 0L8 7.293l2.646-2.647a.5.5 0 0 1 .708.708L8.707 8l2.647 2.646a.5.5 0 0 1-.708.708L8 8.707l-2.646 2.647a.5.5 0 0 1-.708-.708L7.293 8 4.646 5.354a.5.5 0 0 1 0-.708z" />
              </svg>
            </button>
            {children}
          </div>
        </div>
      )}
    </>
  );
};

export default Reservar;


/* La fecha de reserva debemos colocarla andentro de la pagina donde vamos a reservar el hotel
   <div className="main-navbar">
          <div className="contenedorBotones">
            <button
              onClick={() => cambiarEstadoReserva(!EstadoReserva)}
              className="boton"
            >
              Reservar fecha
            </button>
          </div>
          <Reservar estado={EstadoReserva} cambiarestado={cambiarEstadoReserva}>
            <h2>Desde: </h2>
            <h2>Hasta: </h2>
          </Reservar>
        </div>
     

*/