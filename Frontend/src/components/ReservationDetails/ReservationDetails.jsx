import React, { useEffect, useState, useContext } from "react";
import { useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import { UserProfileContext } from '../../App';
import "./ReservationDetail.css"

import Navbar from "../NavBar/NavBar";

const ReservationDetails = () => {
  const { id } = useParams();
  const [reservation, setReservation] = useState(null);
  const [hotel, setHotel] = useState(null);
  const [error, setError] = useState(null);
  const { userProfile } = useContext(UserProfileContext);

  useEffect(() => {
    const fetchReservationDetails = async () => {
      try {
        const response = await fetch(`http://localhost:8090/reservation/${id}`);
        if (response.ok) {
          const data = await response.json();
          setReservation(data);

            const hotelResponse = await fetch(`http://localhost:8090/hotel/${data.hotel_id}`);
            if (hotelResponse.ok) {
            const hotelData = await hotelResponse.json();
            setHotel(hotelData);
            } else {
            const errorData = await hotelResponse.json();
            throw new Error(errorData.error);
            }

        } else {
          const errorData = await response.json();
          throw new Error(errorData.error);
        }
      } catch (error) {
        setError(error.message);
      }
    };

    fetchReservationDetails();
  }, [id]);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!reservation) {
    return <div>Loading...</div>;
  }

  if (!userProfile || ((reservation.user_id !== userProfile.id) && (userProfile.role === "Customer"))) {
    return (
        <>
            <Navbar />
            <p>No puedes acceder a este sitio.</p>
        </>
    )
   }

  return (
    <>
      <Navbar />
      <div className="Detalles">
        <h2>Reserva {id}</h2>
        {hotel && (
          <>
            <Link to={`/hotel/${hotel.id}`}>
              <h4>{hotel.name}</h4>
              </Link>
            <h6>{hotel.street_name} {hotel.street_number}</h6>
          </>
        )}
        <p>Inicio: {reservation.start_date}</p>
        <p>Fin: {reservation.end_date}</p>
        <p>Costo: ${reservation.amount}</p>
      </div>
      </>
  );
  
  
};

export default ReservationDetails;
