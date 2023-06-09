import React, { useContext, useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Navbar from "../NavBar/NavBar";
import { LoginContext } from '../../App';
import "./HotelList.css"

const HotelList = () => {
  const [hotels, setHotels] = useState([]);
  const [error, setError] = useState(null);
  const { loggedIn } = useContext(LoginContext)

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const response = await fetch("http://localhost:8090/hotel");
        if (response.ok) {
          const data = await response.json();
          setHotels(data);
        } else {
          const data = await response.json();
          const errorMessage = data.error || 'Error';
          throw new Error(errorMessage);
        }
      } catch (error) {
        console.error(error);
        setError(error.message);
      }
    };

    fetchHotels();
  }, []);

  if (error) {
    return (
        <>
          <Navbar />
          <div className="fullscreen">Error: {error}</div>
        </>
    );
  }

  if (!hotels) {
    return (
      <>
        <Navbar />
        <h2>Hoteles</h2>
        <p className="fullscreen">No hay hoteles disponibles</p>
      </>
    );
  }

  return (
    <>
      <Navbar />
      <h2>Hoteles</h2>
      <div className="row">
        {hotels.map((hotel) => (
          <div key={hotel.id} className="col-md-4 mb-4">
            <div className="card">
                {hotel.images &&
                    <img className="card-img-top" alt={`Image for ${hotel.name}`} src={`http://localhost:8090/image/${hotel.images[0].id}`}/>}
              <div className="card-body">
                <h5 className="card-title">
                    <Link to={`/hotel/${hotel.id}`}>
                        {hotel.name}
                    </Link>
                </h5>
                <p className="card-text">
                  Dirección: {hotel.street_name} {hotel.street_number}
                </p>
                <p className="card-text">${hotel.rate}</p>
              </div>
            </div>
          </div>
        ))}
      </div>
    </>
  );
};

export default HotelList;
