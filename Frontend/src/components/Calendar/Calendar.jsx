import React, { useState, useContext } from "react";
import { LoginContext, UserProfileContext } from '../../App';
import { DateRange } from "react-date-range";
import { addDays, format, differenceInHours } from "date-fns";
import "react-date-range/dist/styles.css";
import "react-date-range/dist/theme/default.css";

const Calendar = ({ hotel_id, hotelRate }) => {
  const { userProfile } = useContext(UserProfileContext);

  const [ reservation, setReservation ] = useState(null)
  
  // date state
  const [range, setRange] = useState([
    {
      startDate: new Date(),
      endDate: new Date(),
      key: "selection",
    },
  ]);

  const [reservationSaved, setReservationSaved] = useState(false);

  const handleSelectDates = (ranges) => {
    setRange([ranges.selection]);
  };

  const calculateAmount = () => {
    const start_date = range[0].startDate;
    const end_date = range[0].endDate;

    // Calculate check-in and checkout dates with the desired times
    const checkInDate = new Date(start_date);
    checkInDate.setHours(15, 0, 0, 0);

    const checkoutDate = new Date(end_date);
    checkoutDate.setHours(11, 0, 0, 0);

    const hours = differenceInHours(checkoutDate, checkInDate);
    const nights = Math.ceil(hours / 24)
    let amount = nights * hotelRate;
    amount = amount.toFixed(2); //Limit to 2 decimals
    amount = parseFloat(amount).toLocaleString(); //Thousand separators

    return amount;
  };

  const handleReservation = async () => {
    const start_date = range[0].startDate;
    const end_date = range[0].endDate;

    // Calculate check-in and checkout dates with the desired times
    const checkInDate = new Date(start_date);
    checkInDate.setHours(15, 0, 0, 0);

    const checkoutDate = new Date(end_date);
    checkoutDate.setHours(11, 0, 0, 0);

    try {
      const reservationData = {
        user_id: parseInt(userProfile.id),
        hotel_id: parseInt(hotel_id),
        start_date: format(checkInDate, "dd-MM-yyy HH:mm"),
        end_date: format(checkoutDate, "dd-MM-yyy HH:mm"),
      };

      const response = await fetch("http://localhost:8090/reserve", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(reservationData),
      });

      if (response.ok) {
        setReservationSaved(true);
        const data = await response.json();
        setReservation(data);
      } else {
        console.log("Error saving reservation.");
      }
    } catch (error) {
      console.log("Error:", error);
    }
  };

  const today = new Date();
  const amount = calculateAmount();

  return (
    <>
      <div className="calendarWrap">
        <DateRange
          ranges={range}
          onChange={handleSelectDates}
          editableDateInputs={true}
          moveRangeOnFirstSelection={false}
          months={2}
          minDate={today}
          direction="horizontal"
          className="calendarElement"
        />
      </div>
      <div>
        <p>Total: ${amount}</p>
        <button onClick={handleReservation} disabled={reservationSaved}>
          {reservationSaved ? "Reserva Guardada" : "Reservar"}
        </button>
      </div>
    </>
  );
};

export default Calendar;
