import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import { Box, CssBaseline } from '@mui/material';
import HumanCreate from "./pages/HumanCreate";
import Register from "./pages/Register";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";

const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Navigate to="/human-create" />} />
      <Route path="/human-create" element={<HumanCreate />} />
      <Route path="/register" element={<Register />} />
      <Route path="/login" element={<Login />} />
      <Route path="/dashboard" element={<Dashboard />} />
    </Routes>
  );
};

export default App;
