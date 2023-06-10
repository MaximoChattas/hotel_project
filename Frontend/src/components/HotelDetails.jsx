import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

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
          throw new Error("Error fetching hotel details");
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
    <div>
      <h1>{hotel.name}</h1>
      <p>
        Dirección: {hotel.street_name}, {hotel.street_number}
      </p>
      <p>Description: {hotel.description}</p>
      {/* Render your calendar component for reservations here */}
    </div>
  );
};

export default HotelDetails;
