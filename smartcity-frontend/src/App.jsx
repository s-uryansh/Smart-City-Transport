import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import { Box, CssBaseline } from '@mui/material';
import HumanCreate from "./pages/HumanCreate";
import Register from "./pages/Register";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import DashboardVehicleManager from "./pages/DashboardVehicleManager";
import DashboardPaymentManager from "./pages/DashboardPaymentManager";
import DashboardScheduleManager from "./pages/DashboardScheduleManager";
import DashboardIncidentManager from "./pages/DashboardIncidentManager";
import DashboardRouteManager from "./pages/DashboardRouteManager";
import DashboardMaintenanceManager from "./pages/DashboardMaintenanceManager";
import DashboardAccidentManager from "./pages/DashboardAccidentManager";
import DashboardAdmin from "./pages/DashboardAdmin";

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/human-create" />} />
      <Route path="/human-create" element={<HumanCreate />} />
      <Route path="/register" element={<Register />} />
      <Route path="/login" element={<Login />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/admin-dashboard" element={<DashboardAdmin />} />
      <Route path="/vehicle-dashboard" element={<DashboardVehicleManager />} />
      <Route path="/payment-dashboard" element={<DashboardPaymentManager />} />
      <Route path="/schedule-dashboard" element={<DashboardScheduleManager />} />
      <Route path="/route-dashboard" element={<DashboardRouteManager />} />
      <Route path="/maintenance-dashboard" element={<DashboardMaintenanceManager />} />
      <Route path="/incident-dashboard" element={<DashboardIncidentManager />} />
      <Route path="/accident-dashboard" element={<DashboardAccidentManager />} />
    </Routes>
  );
};

export default App;
