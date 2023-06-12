import React, { useEffect, useState, useContext } from "react";
import { LoginContext, UserProfileContext } from '../../App';
import { useParams } from "react-router-dom";
import { Link } from "react-router-dom";
import Navbar from "../NavBar/NavBar";

const UserReservations = () => {
  const { id } = useParams();
  const [userReservations, setUserReservations] = useState([]);
  const [hotels, setHotels] = useState([]);
  const [error, setError] = useState(null);
  const { userProfile } = useContext(UserProfileContext);
  const { loggedIn } = useContext(LoginContext);

  useEffect(() => {
    const fetchUserReservations = async () => {
      try {
        const response = await fetch(`http://localhost:8090/user/reservations/${id}`);
        if (response.ok) {
          const data = await response.json();
          setUserReservations(data);

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

    fetchUserReservations();
  }, [id]);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!userReservations) {
    return <div>Loading...</div>;
  }

  if (!loggedIn || (userReservations.user_id !== userProfile.id)) {
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
      <h2>Reservas de {userReservations.user_name} {userReservations.user_last_name}</h2>
  
      {userReservations.reservations ? (
        <>
        {hotels.map(hotel => {
          const hotelReservations = userReservations.reservations.filter(
            reservation => reservation.hotel_id === hotel.id
          );
          if (hotelReservations.length > 0) {
            return (
              <div key={hotel.id}>
                <h3>{hotel.name}</h3>
                {hotelReservations.map(reservation => (
                  <div key={reservation.id}>
                    <Link to={`/reservation/${reservation.id}`}>
                      Nº Reserva: {reservation.id}
                    </Link>
                    <p>Inicio: {reservation.start_date}</p>
                    <p>Fin: {reservation.end_date}</p>
                    <p>Costo: {reservation.amount}</p>
                  </div>
                ))}
              </div>
            );
          } else {
            return null;
          }
        })}
      </>
      ) : (
        <p>No tiene ninguna reserva.</p>
      )}
    </>
  );
  
};

export default UserReservations;
