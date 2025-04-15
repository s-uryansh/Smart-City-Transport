# Smart City Transport System

> **Course Project:** CSD317 – Introduction to Database Systems  
> **Developed by:** [Suryansh Rohil](https://github.com/s-uryansh) && [Arnab Mandal](https://github.com/eros483)

## 📌 Overview

The **Smart City Transport System** is a full-stack web application designed to streamline urban transportation management. Developed as part of the CSD317 course, this project integrates a robust backend with an intuitive frontend to facilitate efficient transport operations within a smart city framework.

## 🏗️ Project Structure

```
Smart-City-Transport-System/
├── Backend/               # Go-based backend services
├── Internals/             # Data and Schema
├── Project Reports/       # Documentation and project reports
├── smartcity-frontend/    # React.js frontend application
├── package-lock.json      # Dependency lock file
└── ...
```

### 🔧 Backend

- **Language:** Go (Golang)
- **Functionality:** Handles API endpoints, business logic, and database interactions.
- **Structure:** Modular design with clear separation of concerns, enhancing maintainability and scalability.

### 🌐 Frontend

- **Framework:** React.js
- **Features:** Provides a user-friendly interface for interaction with the transport system, including real-time updates and responsive design.

### 📄 Reports

- **Location:** `Project Reports/`
- **Contents:** Comprehensive reports detailing system design, database schema, and implementation strategies.

## 🚀 Getting Started

### Prerequisites

- **Node.js** (for frontend development)
- **Go** (for backend development)
- **MongoDB** or another compatible database system

### Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/s-uryansh/Smart-City-Transport-System.git
   cd Smart-City-Transport-System
   ```

2. **Setup Backend:**

   ```bash
   cd Backend
   go mod tidy
   go run cmd/app/main.go
   ```

3. **Setup Frontend:**

   ```bash
   cd smartcity-frontend
   npm install
   npm run dev
   ```
4. Create Database `smartcity`
```MySQL
Tables:
+----------------------+
| Tables_in_smartcity  |
+----------------------+
| accident_history     |
| human                |
| incident             |
| maintenance          |
| maintenance_history  |
| operates_on          |
| payment              |
| performs_maintenance |
| route                |
| route_followed       |
| schedule             |
| schedule_followed    |
| users                |
| vehicle              |
+----------------------+
```
`Schema`: [All the Tables](Database/Tables)

4. **Access the Application:**

   Open your browser and navigate to `http://localhost:5713` to interact with the Smart City Transport System.

## 📚 Features

- **Real-Time Tracking:** Monitor transport vehicles in real-time.
- **User Management:** Secure authentication and user role management.
- **Data Analytics:** Visualize transport data for informed decision-making.
- **Responsive Design:** Accessible across various devices and screen sizes.

## 🛠️ Technologies Used

- **Frontend:** React.js, HTML5, CSS3
- **Backend:** Go (Golang)
- **Database:** MySQL
