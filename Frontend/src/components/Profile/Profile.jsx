import React, { useContext } from "react";
import { useNavigate } from 'react-router-dom';
import Navbar from "../NavBar/NavBar";
import { LoginContext, UserProfileContext } from '../../App';

function Profile() {
    const { loggedIn, setLoggedIn } = useContext(LoginContext);
    const { userProfile, setUserProfile } = useContext(UserProfileContext);
    const navigate = useNavigate();
    const reservationURL = "/user/reservations/"+userProfile.id;

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('userProfile');
        setLoggedIn(false);
        setUserProfile(null);
        navigate('/');
    };

    if (!userProfile) {
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
            <div>
                <h3>Mi Perfil</h3>
                <p>Nombre: {userProfile.name}</p>
                <p>Apellido: {userProfile.last_name}</p>
                <p>DNI: {userProfile.dni}</p>
                <p>Email: {userProfile.email}</p>
                <p>Nº de usuario: {userProfile.id}</p>
            </div>
            <div>
                <button onClick={()=>navigate(reservationURL)}> Mis Reservas </button>
                <button onClick={handleLogout}>Cerrar Sesión</button>
            </div>
            
        </>
    )
}

export default Profile;
