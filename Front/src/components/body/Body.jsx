import React from "react";
import "./Body.css";
import hotel1 from "../../assets/img/hotel1.png";
import hotel2 from "../../assets/img/hotel2.png";
import hotel3 from "../../assets/img/hotel3.png";
import { NavLink } from "react-router-dom";

function Body() {
  return (
    <div className="hotel-card">
      <div className="hoteles">
        <NavLink className="nav-link" to="/habitaciones">
          {/* <button className="botonhoteles"> */}
          <img src={hotel1} alt="hotel" className="hotel-image" />
          <h2 className="hotel-title">Hotel Perfecta</h2>
          {/* </button> */}
        </NavLink>
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
        <NavLink className="nav-link" to="/habitaciones">
          {/* <button className="botonhoteles"> */}
          <img src={hotel2} alt="hotel" className="hotel-image" />
          <h2 className="hotel-title">Hotel Don</h2>
          {/* </button> */}
        </NavLink>
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
        <NavLink className="nav-link" to="/habitaciones">
          {/* <button className="botonhoteles"> */}
          <img src={hotel3} alt="hotel" className="hotel-image" />
          <h2 className="hotel-title">Hotel Traicion</h2>
          {/* </button> */}
        </NavLink>
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

// import React from "react";
// import "./Body.css";
// import hotel1 from "../../assets/img/hotel1.png";
// import hotel2 from "../../assets/img/hotel2.png";
// import hotel3 from "../../assets/img/hotel3.png";
// import { NavLink } from "react-router-dom";

// const photos = [
//   { id: 1, src: hotel1, title: "foto 1" },
//   { id: 2, src: hotel2, title: "foto 2" },
//   { id: 3, src: hotel3, title: "foto 3" },
//   { id: 4, src: hotel1, title: "foto 4" },
// ];

// function Body() {
//   return (
//     <div className="gallery-container">
//       {photos.map((photo) => (
//         <div key={photo.id} className="photo-item">
//           <img src={photo.src} alt={photo.title} />
//           <h3>{photo.title}</h3>
//         </div>
//       ))}
//     </div>
//   );
// }

// export default Body;
