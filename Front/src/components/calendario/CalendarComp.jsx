import React, { useState } from "react";
import { DateRange } from "react-date-range";
import { addDays } from "date-fns";
import "react-date-range/dist/styles.css";
import "react-date-range/dist/theme/default.css";

const CalendarComp = () => {
  //date state
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

  const handleReservation = () => {
    // Aquí puedes realizar la lógica para guardar la reserva con las fechas seleccionadas
    console.log("Reserva guardada:", range);
    setReservationSaved(true); // Actualizar el estado a reserva guardada
  };

  //determina la fecha de hoy
  const today = new Date();

  return (
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
      <button onClick={handleReservation} disabled={reservationSaved}>
        {reservationSaved ? "Reserva Guardada" : "Guardar Reserva"}
      </button>
    </div>
  );
};

export default CalendarComp;
