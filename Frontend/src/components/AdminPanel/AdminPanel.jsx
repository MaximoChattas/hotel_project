import React, { useContext } from "react";
import { useNavigate } from 'react-router-dom';
import { LoginContext, UserProfileContext } from '../../App';

function AdminPanel() {
    const { loggedIn } = useContext(LoginContext);
    const { userProfile } = useContext(UserProfileContext);
    const navigate = useNavigate();

    if (!loggedIn || userProfile.role !== "Admin") {
        return null;
      }

      return (
        <>
            <div>
                <h4>Panel de Administración</h4>
            </div>
            <div>
                <button onClick={() => navigate('/loadhotel')}>Nuevo Hotel</button>
                <button onClick={() => navigate('/admin/reservations/hotel')}>Ver reservas por Hotel</button>
                <button onClick={() => navigate('/admin/reservations/user')}>Ver reservas por usuario</button>
            </div>
        </>
      )

}

export default AdminPanel