import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Navbar from "../NavBar/NavBar";

const HotelList = () => {
  const [hotels, setHotels] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchHotels = async () => {
      try {
        const response = await fetch("http://localhost:8090/hotel");
        if (response.ok) {
          const data = await response.json();
          setHotels(data);
        } else {
          throw new Error("Error fetching hotels");
        }
      } catch (error) {
        setError(error.message);
      }
    };

    fetchHotels();
  }, []);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (hotels.length === 0) {
    return <div>Loading...</div>;
  }

  return (
    <>
      <Navbar />
      <h2>Hoteles</h2>
      <div className="row">
        {hotels.map((hotel) => (
          <div key={hotel.id} className="col-md-4 mb-4">
            <div className="card">
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
