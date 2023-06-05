import React from "react";
import Hotel from "./Hotel";
import hotels from "../../data/hotels";
import "./HotelList.css";

const HotelList = () => {
  return (
    <div className="hotel-list">
      {hotels.map((hotel) => (
        <Hotel key={hotel.id} hotel={hotel} />
      ))}
    </div>
  );
};

export default HotelList;
