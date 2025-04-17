import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

function HumanCreate() {
  const [isDarkMode, setIsDarkMode] = useState(true);

  useEffect(() => {
    document.body.className = isDarkMode ? 'dark-mode' : 'light-mode';
  }, [isDarkMode]);

  const [formData, setFormData] = useState({
    ID_NO: '',
    FName: '',
    LName: '',
    DOB: '',
    V_ID: '',
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    setFormData(prev => ({
      ...prev,
      [e.target.name]: e.target.value
    }));
  };

  const calculateAge = (dob) => {
    const birthDate = new Date(dob);
    const today = new Date();
    let age = today.getFullYear() - birthDate.getFullYear();
    const m = today.getMonth() - birthDate.getMonth();
    if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
      age--;
    }
    return age;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const age = calculateAge(formData.DOB);

    const payload = {
      ...formData,
      ID_NO: parseInt(formData.ID_NO),
      Age: age,
      V_ID: parseInt(formData.V_ID)
    };

    try {
      const res = await fetch('http://localhost:8080/humans/', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload)
      });

      if (res.ok) {
        alert('Human created successfully!');
        navigate('/register');
      } else {
        const errData = await res.json();
        alert('Error creating human: ' + errData.message);
      }
    } catch (err) {
      console.error(err);
      alert('Server error');
    }
  };

  return (
    <div className={`container mt-5 p-4 rounded shadow ${isDarkMode ? 'bg-dark text-white' : 'bg-light'}`}>
      {/* Dark Mode Toggle Icon */}
      <div className="text-end mb-3">
        <button
          onClick={() => setIsDarkMode(!isDarkMode)}
          className="btn btn-sm btn-outline-light"
          title="Toggle Dark Mode"
        >
          {isDarkMode ? '🌙' : '☀️'}
        </button>
      </div>

      <h2 className="mb-4">Create Human</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label className="form-label">ID Number</label>
          <input
            type="number"
            name="ID_NO"
            className="form-control"
            value={formData.ID_NO}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label className="form-label">First Name</label>
          <input
            type="text"
            name="FName"
            className="form-control"
            value={formData.FName}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label className="form-label">Last Name</label>
          <input
            type="text"
            name="LName"
            className="form-control"
            value={formData.LName}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label className="form-label">Date of Birth</label>
          <input
            type="date"
            name="DOB"
            className="form-control"
            value={formData.DOB}
            onChange={handleChange}
            required
          />
        </div>

        <div className="mb-3">
          <label className="form-label">Vehicle ID (V_ID)</label>
          <input
            type="number"
            name="V_ID"
            className="form-control"
            value={formData.V_ID}
            onChange={handleChange}
            required
          />
        </div>

        <div className="d-flex justify-content-between">
  <button type="submit" className="btn btn-primary">
    Create
  </button>

  <button
    type="button"
    className="btn btn-outline-info"
    onClick={() => navigate('/login')}
  >
    Login
  </button>

  <button
    type="button"
    className="btn btn-outline-info"
    onClick={() => navigate('/register')}
  >
    User Register
  </button>
</div>

      </form>
    </div>
  );
}

export default HumanCreate;
