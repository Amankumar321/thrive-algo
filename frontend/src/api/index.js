import axios from 'axios';

axios.defaults.baseURL = 'https://thrive-algo.onrender.com'

export const getHolidaysAPI = async () => { return await axios.get('/api/holidays') };
export const addHolidayAPI =  async (date, holidayName) => { return await axios.post('/api/holidays', { date, name: holidayName }) };
export const deleteHolidayAPI = async (id) => { return await axios.delete(`/api/holidays/${id}`) };