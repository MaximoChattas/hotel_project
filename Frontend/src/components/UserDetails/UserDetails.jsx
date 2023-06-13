import React, { useEffect, useState, useContext } from "react";
import { LoginContext, UserProfileContext } from '../../App';
import { useParams } from "react-router-dom";
import Calendar from "../Calendar/Calendar";
import Navbar from "../NavBar/NavBar";

const UserDetails = () => {
  const { id } = useParams();
  const [user, setUser] = useState(null);
  const [error, setError] = useState(null);
  const { loggedIn } = useContext(LoginContext);
  const { userProfile } = useContext(UserProfileContext);

  useEffect(() => {
    const fetchUserDetails = async () => {
      try {
        const response = await fetch(`http://localhost:8090/user/${id}`);
        if (response.ok) {
          const data = await response.json();
          setUser(data);
        } else {
          const errorData = await response.json();
          throw new Error(errorData.error);
        }
      } catch (error) {
        setError(error.message);
      }
    };

    fetchUserDetails();
  }, [id]);

  if (error) {
    return <div>Error: {error}</div>;
  }

  if (!user) {
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
      <div className="UserDetalle">
        <h1>Perfil de Usuario</h1>
        <h2>{user.name} {user.last_name}</h2>
        <p>Número de usuario: {user.id}</p>
        <p>Email: {user.email}</p>
      </div>
      </>
  );
};

export default UserDetails;