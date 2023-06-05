import React, { useEffect } from "react";
import "./Hotel.css";
import { NavLink } from "react-router-dom";

const Hotel = ({ hotel }) => {
  // useEffect(() => {
  //   fetch("/hotel")
  //     .then((response) => response.json())
  //     .then((data) => {
  //       hotel(data);
  //     })
  //     .catch((error) => {
  //       console.error("Error:", error);
  //     });
  // }, []);

  return (
    <div className="hotelASD">
      <NavLink to={hotel.link}>
        <img src={hotel.image} alt={hotel.name} />
      </NavLink>
      <h3>{hotel.name}</h3>
      <p>Rating: {hotel.rating}</p>
      <p>${hotel.price}</p>
    </div>
  );
};

export default Hotel;
