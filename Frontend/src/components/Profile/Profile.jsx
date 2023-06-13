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
            <div className="descripcion">
                <h3>Profile</h3>
                <p>Name: {userProfile.name}</p>
                <p>Last name: {userProfile.last_name}</p>
                <p>DNI: {userProfile.dni}</p>
                <p>Email: {userProfile.email}</p>
                <p>ID: {userProfile.id}</p>
                <div>
                    {userProfile.role === "Customer" && <button onClick={()=>navigate(reservationURL)}> Mis Reservas </button>}
                    <button onClick={handleLogout}>Cerrar Sesión</button>
                </div>
            </div>
            <div>
                <AdminPanel />
            </div>
        </>
    )
}

export default Profile;
