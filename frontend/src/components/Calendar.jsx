import React, { useState } from 'react';
import { addHolidayAPI } from '../api';

const Calendar = ({ holidays, fetchHolidays }) => {
  const [month, setMonth] = useState(new Date().getMonth());
  const [year, setYear] = useState(new Date().getFullYear());

  const addHoliday = async (date) => {
    const holidayName = prompt('Enter holiday name:');
    if (!holidayName) return;

    try {
      await addHolidayAPI(date, holidayName);
      fetchHolidays();
    } catch (error) {
      console.error('Error adding holiday:', error);
    }
  };

  const daysInMonth = new Date(year, month + 1, 0).getDate();
  const firstDayOfMonth = new Date(year, month, 1).getDay();

  const renderCalendar = () => {
    const calendarCells = [];

    // Empty cells for days before the first day
    for (let i = 0; i < firstDayOfMonth; i++) {
      calendarCells.push(<div key={`empty-${i}`} className="border border-gray-200 h-20"></div>);
    }

    // Calendar cells for each day of the month
    for (let day = 1; day <= daysInMonth; day++) {
      const dateStr = `${year}-${String(month + 1).padStart(2, '0')}-${String(day).padStart(2, '0')}`;
      const todayDateStr = new Date().toISOString().split('T')[0]

      calendarCells.push(
        <div key={dateStr} className={`border ${dateStr === todayDateStr ? 'border-crimson' : 'border-gray-200'} h-20 relative group bg-white`}>
          <span className="absolute top-2 left-2 text-sm font-bold">{day}</span>
          <button
            onClick={() => addHoliday(dateStr)}
            className="hidden group-hover:block absolute text-lg bottom-2 right-2 text-crimson border border-red-200 bg-red-50 px-2 rounded"
          >
            +
          </button>
        </div>
      );
    }

    return calendarCells;
  };

  return (
    <div className="w-full p-2 sm:p-4 border border-solid border-gray-200 bg-gray-50">
      <div className="flex justify-center items-center mb-6">
        <button
          onClick={() => {
            setMonth((prev) => (prev === 0 ? 11 : prev - 1));
            if (month === 0 && year > 2024) setYear((prev) => prev - 1);
          }}
          className={`bg-red-50 py-1 px-2 rounded border border-red-100 ${year === 2024 && month === 0 ? 'hidden' : ''}`}
        >
          &#10094;&nbsp;&nbsp;Prev
        </button>
        <div className="w-48 flex justify-center">
          <h2 className="text-xl font-semibold">{new Date(year, month).toLocaleString('default', { month: 'long', year: 'numeric' })}</h2>
        </div>
        <button
          onClick={() => {
            setMonth((prev) => (prev === 11 ? 0 : prev + 1));
            if (month === 11 && year < 2025) setYear((prev) => prev + 1);
          }}c
          className={`bg-red-50 py-1 px-2 rounded border border-red-100 ${year === 2029 && month === 11 ? 'hidden' : ''}`}
        >
          Next&nbsp;&nbsp;&#10095;
        </button>
      </div>

      <div className="grid grid-cols-7 gap-2">
        {['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'].map((day) => (
          <div key={day} className="font-bold text-center border-b-2 pb-2 text-crimson">
            {day}
          </div>
        ))}
        {renderCalendar()}
      </div>
    </div>
  );
};

export default Calendar;