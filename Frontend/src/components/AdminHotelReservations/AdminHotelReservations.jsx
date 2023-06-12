import React, { useEffect, useState, useContext } from "react";
import { LoginContext, UserProfileContext } from '../../App';
import { Link } from "react-router-dom";
import Navbar from "../NavBar/NavBar";

const AdminHotelReservations = () => {
  const [hotelReservations, setHotelReservations] = useState({ reservations: [] });
  const [hotels, setHotels] = useState([]);
  const [error, setError] = useState(null);
  const { userProfile } = useContext(UserProfileContext);
  const { loggedIn } = useContext(LoginContext);

  useEffect(() => {
    const fetchHotelReservations = async () => {
      try {
        const response = await fetch(`http://localhost:8090/reservation`);
        if (response.ok) {
          const data = await response.json();
          setHotelReservations({ reservations: data });

          const hotelResponse = await fetch(`http://localhost:8090/hotel`);
          if (hotelResponse.ok) {
            const hotelData = await hotelResponse.json();
            setHotels(hotelData);
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

    fetchHotelReservations();
  }, []);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!hotelReservations) {
    return <div>Loading...</div>;
  }

  if (!loggedIn || userProfile.role !== "Admin") {
    return (
      <>
        <Navbar />
        <p>No puedes acceder a este sitio.</p>
      </>
    );
  }

  return (
    <>
      <Navbar />
      <h2>Reservas</h2>

      {hotels.map(hotel => {
        const filteredReservations = hotelReservations.reservations || [];
        const hotelReservationsFiltered = filteredReservations.filter(
          reservation => reservation.hotel_id === hotel.id
        );
        if (hotelReservationsFiltered.length > 0) {
          return (
            <div key={hotel.id}>
              <h3>{hotel.name}</h3>
              {hotelReservationsFiltered.map(reservation => (
                <div key={reservation.id}>
                  <Link to={`/reservation/${reservation.id}`}>
                    Nº Reserva: {reservation.id}
                  </Link>
                  <p>Inicio: {reservation.start_date}</p>
                  <p>Fin: {reservation.end_date}</p>
                  <p>Costo: {reservation.amount}</p>
                  <p>Usuario Nº: {reservation.user_id}</p>
                </div>
              ))}
            </div>
          );
        } else {
          return (
            <div key={hotel.id}>
              <h3>{hotel.name}</h3>
              <p>Sin reservas.</p>
            </div>
          );
        }
      })}
    </>
  );
};

export default AdminHotelReservations;
