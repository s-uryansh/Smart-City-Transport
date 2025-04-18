import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
  });

  const [isDarkMode, setIsDarkMode] = useState(true);

  useEffect(() => {
    document.body.className = isDarkMode ? 'dark-mode' : 'light-mode';
  }, [isDarkMode]);

  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData(prev => ({
      ...prev,
      [e.target.name]: e.target.value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const { username, password } = formData;

    try {
      const res = await fetch("https://smart-city-transport-301261782088.us-central1.run.app/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password })
      });

      const data = await res.json();
      if (res.ok) {
        localStorage.setItem('token', data.token);
        //============================================//
        //check users role and show them there page

        if(data.Role == "Admin"){
          navigate('/admin-dashboard');
        }
        if(data.Role == "user"){
          navigate('/dashboard');
        }
        if(data.Role == "Vehicle Manager"){
          navigate('/vehicle-dashboard');
        }
        if(data.Role == "Maintenance Manager"){
          navigate('/maintenance-dashboard');
        }
        if(data.Role == "Payment Manager"){
          navigate('/payment-dashboard');
        }
        if(data.Role == "Schedule Manager"){
          navigate('/schedule-dashboard');
        }
        if(data.Role == "Route Manager"){
          navigate('/route-dashboard');
        }
        if(data.Role == "Incident Manager"){
          navigate('/incident-dashboard');
        }
        if(data.Role == "Accident Manager"){
          navigate('/accident-dashboard');
        }

        //============================================//
      } else {
        alert('Login failed: ' + (data?.message || data?.error || 'Unknown error'));
      }
    } catch (error) {
      console.error('Login error:', error);
      alert('Server error');
    }
  };

  return (
    <div className="container mt-5 position-relative">
      <button
        className="btn btn-sm position-absolute top-0 end-0 m-2"
        onClick={() => setIsDarkMode(prev => !prev)}
        style={{ fontSize: '1.4rem' }}
        aria-label="Toggle dark mode"
      >
        {isDarkMode ? '🌙' : '☀️'}
      </button>

      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">Username</label>
          <input
            type="text"
            name="username"
            className="form-control"
            value={formData.username}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label className="form-label">Password</label>
          <input
            type="password"
            name="password"
            className="form-control"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>

        <button type="submit" className="btn btn-primary">Login</button>
      </form>
    </div>
  );
}

export default Login;
