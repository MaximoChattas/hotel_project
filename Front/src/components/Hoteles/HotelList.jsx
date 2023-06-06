// import React, { useEffect, useState } from "react";
// import Hotel from "./Hotel";
// import hotels from "../../data/hotels";
// import "./HotelList.css";

// const HotelList = () => {
//   const [hotel, setHotel] = useState([]);

//   useEffect(() => {
//     const fetchHotels = async () => {
//       try {
//         const response = await fetch("http://localhost:8090/hotels");
//         setHotel(response.data);
//       } catch (error) {
//         console.log(error);
//       }
//     };

//     fetchHotels();
//   }, []);

//   return (
//     <div className="hotel-list">
//       {hotels.map((hotel) => (
//         <Hotel key={hotel.id} hotel={hotel} />
//       ))}
//     </div>
//   );
// };

// export default HotelList;

import React, { useEffect, useState } from "react";

const HotelList = () => {
  const [hotels, setHotels] = useState([]);

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
        console.log(error);
      }
    };

    fetchHotels();
  }, []);

  return (
    <div>
      <h1>Listado de Hoteles</h1>
      {hotels.map((hotel) => (
        <div key={hotel.id}>
          <h3>{hotel.name}</h3>
          <p>
            Dirección: {hotel.street_name}, {hotel.street_number}
          </p>
          <p>Descripcion: {hotel.description}</p>
          <p>${hotel.rate}</p>
        </div>
      ))}
    </div>
  );
};

export default HotelList;
