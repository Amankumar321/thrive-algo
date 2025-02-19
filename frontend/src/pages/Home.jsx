import React, { useState, useEffect } from 'react';
import Calendar from '../components/Calendar';
import { getHolidaysAPI, deleteHolidayAPI } from '../api';

const Home = () => {
  const [holidays, setHolidays] = useState([]);

  useEffect(() => {
    fetchHolidays();
  }, []);

  const fetchHolidays = async () => {
    try {
      const response = await getHolidaysAPI();
      setHolidays(response.data ?? []);
    } catch (error) {
      console.error('Error fetching holidays:', error);
    }
  };

  const deleteHoliday = async (id) => {
    if (!window.confirm('Are you sure you want to delete this holiday?')) return;

    try {
      await deleteHolidayAPI(id);
      fetchHolidays();
    } catch (error) {
      console.error('Error deleting holiday:', error);
    }
  };

  const groupHolidaysByMonthYear = () => {
    const groups = [];
  
    holidays.forEach((holiday) => {
      const date = new Date(holiday.date);
      const monthYear = date.toLocaleString('default', { month: 'long', year: 'numeric' });
  
      // Check if the group for the current month-year already exists
      let group = groups.find((g) => g.label === monthYear);
  
      if (!group) {
        // If no group exists, create a new one
        group = { label: monthYear, holidays: [] };
        groups.push(group);
      }
  
      // Add the holiday to the group
      group.holidays.push(holiday);
    });
  
    // Sort each group’s holidays by date
    groups.forEach((group) => {
      group.holidays.sort((a, b) => new Date(a.date) - new Date(b.date));
    });
  
    // Optionally, sort the groups by date (oldest month first)
    groups.sort((a, b) => new Date(a.holidays[0].date) - new Date(b.holidays[0].date));
  
    return groups;
  };
  

  const groupedHolidays = groupHolidaysByMonthYear();

  return (
    <div className="p-2 sm:p-4 w-full ">
      <h1 className="text-2xl font-bold mb-12 mt-2 sm:mt-4 flex justify-center text-crimson tracking-wide">
        🗓️ &nbsp;Holiday Calendar
      </h1>
      <div className="xl:w-2/3 mx-auto flex items-center justify-center">
        <Calendar holidays={holidays} fetchHolidays={fetchHolidays} />
      </div>

      <div className="mt-8 md:w-1/2 xl:w-1/3 mx-auto flex flex-col items-center justify-center">
        <h1 className="text-2xl font-bold mb-12 mt-2 sm:mt-4 flex justify-center text-crimson tracking-wide">
          Holiday List
        </h1>
        {groupedHolidays.map((group) => (
          <div key={group.label} className="w-full mb-6">
            <h3 className="text-lg font-semibold border-b border-gray-300 pb-2 mb-4">{group.label}</h3>
            <ul className="space-y-2 w-full">
              {group.holidays.map((holiday) => (
                <li
                  key={holiday.id}
                  className="w-full flex items-center justify-between bg-gray-50 p-2"
                >
                  <span className="text-gray-600 text-sm break-words">{new Date(holiday.date).toLocaleDateString('en-US', { day: 'numeric' })}</span>
                  <span className="text-gray-800 text-sm break-words">{holiday.name}</span>
                  <button
                    onClick={() => deleteHoliday(holiday.id)}
                    className="px-2 py-0 text-crimson text-lg"
                  >
                    ✕
                  </button>
                </li>
              ))}
            </ul>
          </div>
        ))}
        {Object.keys(groupedHolidays).length === 0 && (
          <p className="text-gray-600 italic">No holidays to display.</p>
        )}
      </div>
    </div>
  );
};

export default Home;
