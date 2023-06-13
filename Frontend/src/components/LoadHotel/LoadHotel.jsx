import React, { useContext, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { LoginContext, UserProfileContext } from '../../App';
import Navbar from '../NavBar/NavBar';
import "./LoadHotel.css"


function LoadHotel() {

  const [name, setName] = useState('');
  const [street_name, setStreet_name] = useState('');
  const [street_number, setStreet_number] = useState('');
  const [room_amount, setRoom_amount] = useState('');
  const [rate, setRate] = useState('');
  const [description, setDescription] = useState('');

  const { loggedIn } = useContext(LoginContext)
  const { userProfile } = useContext(UserProfileContext);

  const [error, setError] = useState('');
  const navigate = useNavigate();
  
  const handleLoadHotel = async (e) => {
    e.preventDefault();
    setError('');
  
    try {

      if(!name || !street_name || !street_number || !room_amount || !rate)
      {
        throw new Error('Complete todos los campos requeridos')
      }

      const response = await fetch('http://localhost:8090/hotel', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            name,
            street_name,
            street_number: parseInt(street_number),
            room_amount: parseInt(room_amount),
            rate: parseFloat(rate),
            description }),
      });
  
      if (response.status === 201) {

        navigate('/');
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

  if (!loggedIn || (userProfile.role !== "Admin")) {
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
        <div className="contenedorLoad">
            <h2>Cargar Hotel</h2>
            <form onSubmit={handleLoadHotel}>
            <div>
                <label>Nombre:</label>
                <input
                  type="text"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </div>
              <div>
                <label>Calle:</label>
                <input
                  type="text"
                  value={street_name}
                  onChange={(e) => setStreet_name(e.target.value)}
                />
              </div>
              <div>
                <label>Altura:</label>
                <input
                  type="number"
                  pattern="[0-9]*"
                  value={street_number}
                  onChange={(e) => setStreet_number(e.target.value)}
                />
              </div>
              <div>
                <label>Habitaciones:</label>
                <input
                  type="number"
                  pattern="[0-9]*"
                  value={room_amount}
                  onChange={(e) => setRoom_amount(e.target.value)}
                />
              </div>
              <div>
                <label>Tarifa:</label>
                <input
                  type="number"
                  pattern="[0-9]*"
                  value={rate}
                  onChange={(e) => setRate(e.target.value)}
                />
              </div>
              <div>
                <label>Descripción:</label>
                <div>
                    <textarea
                    value={description}
                    onChange={(e) => setDescription(e.target.value)}
                    placeholder="Ingrese la descripción"
                    maxLength={1000}
                    rows={4}
                    cols={50}
                    />
                    <div>
                    Characters disponibles: {1000 - description.length} / {1000}
                    </div>
                </div>
              </div>
              {error && <p className="error-message">{error}</p>}
              <button type="submit">Cargar Hotel</button>
            </form>
          </div>
      </>
  );
}

export default LoadHotel;
