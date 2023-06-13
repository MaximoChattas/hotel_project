import React, { useContext } from "react";
import { useNavigate } from 'react-router-dom';
import Navbar from "../NavBar/NavBar";
import { LoginContext, UserProfileContext } from '../../App';
import AdminPanel from "../AdminPanel/AdminPanel";
import "./Profile.css"

function Profile() {
    const { loggedIn, setLoggedIn } = useContext(LoginContext);
    const { userProfile, setUserProfile } = useContext(UserProfileContext);
    const navigate = useNavigate();

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('userProfile');
        setLoggedIn(false);
        setUserProfile(null);
        navigate('/');
    };

    if (!loggedIn) {
        return (
            <>
                <Navbar />
                <div className="descripcion">
                    <p>No puedes acceder a este sitio.</p>
                </div>
            </>
        )
      }

    const reservationURL = "/user/reservations/"+userProfile.id;

    return (
        <>
            <Navbar />
            <div className="descripcion">
                <h3>Profile</h3>
                <p>Name: {userProfile.name}</p>
                <p>Last name: {userProfile.last_name}</p>
                <p>DNI: {userProfile.dni}</p>
                <p>Email: {userProfile.email}</p>
                <p>ID: {userProfile.id}</p>
                <div>
                    {userProfile.role === "Customer" && <button onClick={()=>navigate(reservationURL)}> Mis Reservas </button>}
                    {userProfile.role === "Customer" && <button onClick={()=>navigate("/user/reservations/range")}> Reservas por Rango </button>}
                    <button onClick={handleLogout}>Cerrar Sesión</button>
                </div>
                <AdminPanel />
            </div>
        </>
    )
}

export default Profile;
