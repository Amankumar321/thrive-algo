# Thrive Algo - Holiday Calendar Application

This is a full-stack holiday calendar application that allows users to view, add, and delete holidays for a 5-year range (2024-2029). The project is divided into two parts: `backend` (built with Go and MongoDB) and `frontend` (built with React and Tailwind CSS).

---

## Table of Contents
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [API Endpoints](#api-endpoints)
- [Screenshots](#screenshots)
- [Contributing](#contributing)
- [License](#license)

---

## Features
- View holidays in a monthly calendar (like Google Calendar).
- Add holidays with a custom name.
- Delete holidays from the calendar.
- Pagination between months with a 5-year range (2024-2029).
- Display a list of holidays grouped by month and year.

---

## Tech Stack
- **Frontend:** React, Tailwind CSS
- **Backend:** Go (Golang), MongoDB
- **Database:** MongoDB
- **State Management:** React Hooks
- **API Testing:** Go’s `httptest`

---


---

## Getting Started

### Backend Setup
1. Ensure you have Go installed. (Version 1.18+ recommended)
2. Install MongoDB and ensure it is running locally on `localhost:27017`.
3. Clone the repository:
   ```bash
   git clone https://github.com/your-username/thrive-algo.git
   cd thrive-algo/backend
4. Create a .env file with the following content:
MONGODB_URI=mongodb://localhost:27017
5. Install Go dependencies:
```bash
go mod tidy
```
6. Run the backend server:
```bash
go run main.go
```
7. The server will be running on http://localhost:8080.

### Frontend Setup
1. Navigate to the frontend folder:
```bash
cd ../frontend
```
2. Install Node.js dependencies:
```bash
npm install
```
3. Start the development server:
```bash
npm start
```
4. Open http://localhost:3000 in your browser to view the application.


### API Endpoints

1. Get All Holidays

Endpoint: GET /api/holidays
Description: Returns a list of all holidays.

2. Add a Holiday

Endpoint: POST /api/holidays
Request Body:
{
  "date": "2024-02-24",
  "name": "New Holiday"
}
Description: Adds a new holiday.

3. Delete a Holiday

Endpoint: DELETE /api/holidays/{id}
Description: Deletes a holiday by its ID.