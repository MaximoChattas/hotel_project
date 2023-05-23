import React from 'react';
import './Body.css'; 
import hotel from '../assets/img/hotel1.png';

function Body() {
  return (
    <div className="hotel-card">
      <div className="hoteles">
        <img src={hotel} alt="hotel" className="hotel-image"/>
        <div className="stars">
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
          <span className="star">&#9733;</span>
        </div>
        <h2 className="hotel-title">Hotel Perfecta</h2>
        <h3 className='Precio'>$19.367 por dia</h3>
     </div>
      
      <div className="hoteles">
        <img src={hotel} alt="hotel" className="hotel-image"/>
      <div className="stars">
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
       </div>
       <h2 className="hotel-title">Hotel Don</h2> 
       <h3 className='Precio'>$19.367 por dia</h3>
      </div>


      <div className="hoteles">
        <img src={hotel} alt="hotel" className="hotel-image"/>
        <div className="stars">
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
        <span className="star">&#9733;</span>
      </div>
      <h2 className="hotel-title">Hotel Traicion</h2>
      <h3 className='Precio'>$19.367 por dia</h3>
      </div>
        
        
    </div>
    
  );
}

export default Body;