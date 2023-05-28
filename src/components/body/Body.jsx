import React from "react";
import "./Body.css";
import hotel1 from "../../assets/img/hotel1.png";
import hotel2 from "../../assets/img/hotel2.png";
import hotel3 from "../../assets/img/hotel3.png";

function Body() {
  return (
    <div className="hotel-card">
      <div className="hoteles">
        <button className="botonhoteles">
          <img src={hotel1} alt="hotel" className="hotel-image" />
          <h2 className="hotel-title">Hotel Perfecta</h2>
        </button>
        <div className="stars">
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
        </div>
        <h3 className="Precio">$19.369 por dia</h3>
      </div>
      
      <div className="hoteles">
        <button className="botonhoteles">
        <img src={hotel2} alt="hotel" className="hotel-image" />
        <h2 className="hotel-title">Hotel Don</h2>
        </button>
        <div className="stars">
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
        </div>
        <h3 className="Precio">$19.369 por dia</h3>
      </div>

      <div className="hoteles">
        <button className="botonhoteles">
        <img src={hotel3} alt="hotel" className="hotel-image" />
        <h2 className="hotel-title">Hotel Traicion</h2>
        </button>
        <div className="stars">
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
        </div>
        <h3 className="Precio">$19.369 por dia</h3>
      </div>
    </div>
  );
}

export default Body;
