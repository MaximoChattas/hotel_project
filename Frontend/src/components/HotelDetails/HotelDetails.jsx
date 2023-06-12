import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Calendar from "../Calendar/Calendar";
import Navbar from "../NavBar/NavBar";

const HotelDetails = () => {
  const { id } = useParams();
  const [hotel, setHotel] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchHotelDetails = async () => {
      try {
        const response = await fetch(`http://localhost:8090/hotel/${id}`);
        if (response.ok) {
          const data = await response.json();
          setHotel(data);
        } else {
          const errorData = await response.json();
          throw new Error(errorData.error);
        }
      } catch (error) {
        setError(error.message);
      }
    };

    fetchHotelDetails();
  }, [id]);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!hotel) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <Navbar />
      <h1>{hotel.name}</h1>
      <p>
        Dirección: {hotel.street_name} {hotel.street_number}
      </p>
      <p>Description: {hotel.description}</p>
      <p>Precio por noche: ${hotel.rate}</p>
      <Calendar hotel_id={id} hotelRate={hotel.rate}/>
    </>
  );
};

export default HotelDetails;
