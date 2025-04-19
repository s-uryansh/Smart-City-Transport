import React, { useState,useEffect } from 'react';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import { Dropdown } from 'react-bootstrap';
import axios from 'axios';
import Snackbar from '@mui/material/Snackbar';
import MuiAlert from '@mui/material/Alert';
function Dashboard() {
//===========================Profile===========================
  const [showModal, setShowModal] = useState(false);
  const [profileData, setProfileData] = useState(null);
  const [formData, setFormData] = useState({ FName: '', LName: '', DOB: '', Age: '', V_ID: '' });
  const getProfile = async () => {
    try {
      const token = localStorage.getItem("token");
      const res = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/humans/', {
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (!res.ok) throw new Error(`Failed to fetch: ${res.status}`);

      const data = await res.json();
      setProfileData(data);
      setFormData({
        FName: data.fname || '',
        LName: data.lname || '',
        DOB: data.dob ? data.dob.split('T')[0] : '',
        Age: data.age || '',
        V_ID: data.v_id || ''
      });
    } catch (err) {
      console.error(err);
      showSnackbar("Error fetching profile", 'error');
    }
  };
  const updateProfile = async () => {
    try {
      const token = localStorage.getItem("token");
      const res = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/humans/', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        },
        body: JSON.stringify({
          ...formData,
          Age: parseInt(formData.Age),
          V_ID: parseInt(formData.V_ID),
        }),
      });

      if (res.ok) {
        showSnackbar("Profile updated successfully!");
        getProfile();
      } else {
        const text = await res.text();
        if (text.includes("foreign key constraint fails")) {
          setErrorMessage("Can't delete: MySQL foreign key rule.");
          setShowErrorModal(true);
        } else {
          showSnackbar("Error updating profile", 'error');
        }
      }
    } catch (err) {
      console.error(err);
      showSnackbar("server error", 'error');
    }
  };
  const deleteProfile = async () => {
    const confirmDelete = window.confirm('Are you sure you want to delete your profile?');
    if (!confirmDelete) return;

    try {
      const token = localStorage.getItem("token");
      const res = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/humans/', {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });

      if (res.ok) {
        showSnackbar("profile deleted successfully");
        setProfileData(null);
        setShowModal(false);
      } else {
        const text = await res.text();
          setErrorMessage("Can't delete: MySQL foreign key rule bro.");
          setShowErrorModal(true);
        //   alert('Failed to delete profile');
        
      }
    } catch (err) {
      console.error(err);
      alert('Server error');
    }
  };
//===========================Errors===========================
  const [showErrorModal, setShowErrorModal] = useState(false);
  const [errorMessage, setErrorMessage] = useState('');
//===========================Dark Mode===========================
  const [isDarkMode, setIsDarkMode] = useState(true);
  useEffect(() => {
    const savedTheme = localStorage.getItem('darkMode');
    if (savedTheme) setIsDarkMode(savedTheme === 'true');
  }, []);
  useEffect(() => {
    document.body.className = isDarkMode ? 'dark-mode' : 'light-mode';
    localStorage.setItem('darkMode', isDarkMode);
  }, [isDarkMode]);
//===========================Vehicle===========================
const [newVehicleId, setNewVehicleId] = useState('');  
const [showVehiclesModal, setShowVehiclesModal] = useState(false);
const [vehiclesData, setVehiclesData] = useState([]);
const [showVehicleUpdateModal, setShowVehicleUpdateModal] = useState(false);
const [scheduleVehicleId, setScheduleVehicleId] = useState(null);
const [selectedVehicleId, setSelectedVehicleId] = useState(null);
const [updateFormData, setUpdateFormData] = useState({
  current_location: '',
  status: '',
});
const toggleModal = () => setShowModal(prev => !prev);
const bookVehicle = async (vehicleId) => {
  try {
    const token = localStorage.getItem("token");
    const res = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/vehicles/${vehicleId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify({
        status: "Booked"
      }),
    });

    if (res.ok) {
      showSnackbar("Vehicle booked successfully!");
      toggleVehiclesModal(); // Refresh list after booking
    } else {
      showSnackbar("Failed to book vehicle.", "error");
    }
  } catch (err) {
    console.error(err);
    showSnackbar("Server error while booking vehicle.", "error");
  }
};
const toggleVehiclesModal = async () => {
  if (showVehiclesModal) {
    setShowVehiclesModal(false); // If already open, just close it
    return;
  }

  try {
    const token = localStorage.getItem("token");
    const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/vehicles/all", {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) throw new Error("Failed to fetch vehicles");

    const data = await res.json();
    // //console.log("Fetched vehicles:", data);
    const normalizedData = Array.isArray(data) ? data : [data];
    setVehiclesData(normalizedData);
    setShowVehiclesModal(true);
  } catch (err) {
    console.error(err);
    showSnackbar("Error fetching vehicles", 'error');
  }
};
//===========================Incident===========================
  const [incidentData, setIncidentData] = useState(null);
  const [showIncidentModal, setShowIncidentModal] = useState(false);
  const [editingIncidentId, setEditingIncidentId] = useState(null);
  const [editedDescription, setEditedDescription] = useState('');  
  const [newIncidentId, setNewIncidentId] = useState('');
  const [newDescription, setNewDescription] = useState('');
  const [showCreateForm, setShowCreateForm] = useState(false);
  const updateIncidentDescription = async (incidentId, updatedDescription) => {
    try {
      const token = localStorage.getItem('token');
  
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/incident/${incidentId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({ description: updatedDescription }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to update incident');
      }
  
      const data = await response.json();
      setEditingIncidentId(null);
      fetchIncidentData();
      showSnackbar('Incident updated successfully!');
    } catch (error) {
      console.error('Error updating incident:', error);
      showSnackbar('Error updating incident.', 'error');
    }
  };
  const deleteIncident = async (incidentId) => {
    try {
      const token = localStorage.getItem('token');
  
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/incident/${incidentId}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to delete incident');
      }
  
      // //console.log('Incident deleted successfully');
  
      // Optional: close the modal and refresh the incident list
      setShowIncidentModal(false);
  
      // Optionally remove it from local state if you're storing a list
      setIncidentData((prev) => prev?.filter((i) => i.incident_id !== incidentId));
      showSnackbar('Incident deleted successfully!');
    } catch (error) {
      console.error('Error deleting incident:', error);
      showSnackbar('Error deleting incident.', 'error');
    }
  };
  const createIncident = async () => {
    try {
      const token = localStorage.getItem('token');
      // //console.log(newIncidentId , newDescription)
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/incident/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          INCIDENT_ID: parseInt(newIncidentId),
          Description: newDescription,
        }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to create incident');
      }
  
      const data = await response.json();
      // //console.log('Incident created:', data);
      // setShowIncidentModal(false);
  
      // Optionally update local list if your backend returns full incident
      setIncidentData((prev) => [...prev, data]);
  
      // Clear the form and hide it
      setNewIncidentId('');
      setNewDescription('');
      setShowCreateForm(false);
      fetchIncidentData();
      showSnackbar('Incident created successfully!');
    } catch (error) {
      console.error('Error creating incident:', error);
      showSnackbar('Error creating incident.', 'error');
    }
  };
  const fetchIncidentData = async () => {    
    try {
      const token = localStorage.getItem("token");
      const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/incident/", {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
        }
      });
  
      if (!res.ok) throw new Error("Failed to fetch incident");
  
      const data = await res.json();
      setIncidentData(Array.isArray(data) ? data : [data]);
      setShowIncidentModal(true);
    } catch (err) {
      console.error(err);
      showSnackbar("Error fetching incident", 'error');
    }
  };
//===========================Snackbar===========================
  const [openSnackbar, setOpenSnackbar] = useState(false);
  const [snackbarMessage, setSnackbarMessage] = useState('');
  const [snackbarSeverity, setSnackbarSeverity] = useState('success');
  const showSnackbar = (message, severity = 'success') => {
    setSnackbarMessage(message);
    setSnackbarSeverity(severity);
    setOpenSnackbar(true);
  };
  const Alert = React.forwardRef(function Alert(props, ref) {
    return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
  });
//===========================Maintenance===========================
  const [newMaintenanceId, setNewMaintenanceId] = useState('');
  const [maintenanceData, setMaintenanceData] = useState([]); // State for maintenance data
  const [showMaintenanceModal, setShowMaintenanceModal] = useState(false); // State to control maintenance modal visibility
  const [editingMaintenanceId, setEditingMaintenanceId] = useState(null);
  const [editIssueReported, setEditIssueReported] = useState('');
  const [editRepairStatus, setEditRepairStatus] = useState('Pending');
  const [editVehicleId, setEditVehicleId] = useState('');
  const [isCreating, setIsCreating] = useState(false);
  const [selectedMaintenance, setSelectedMaintenance] = useState(null);
  const updateMaintenance = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance/${editingMaintenanceId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          maintenance_id: parseInt(editingMaintenanceId),
          v_id: parseInt(editVehicleId),
          issue_reported: editIssueReported,
          repair_status: editRepairStatus,
        }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to update maintenance record');
      }
  
      // Fetch updated maintenance data
      await fetchMaintenanceData(); // Refresh the maintenance data
      showSnackbar('Maintenance record updated successfully!');
      setEditingMaintenanceId(null); // Clear editing state
      setEditVehicleId('');
      setEditIssueReported('');
      setEditRepairStatus('Pending');
    } catch (error) {
      console.error('Error updating maintenance record:', error);
      showSnackbar('Error updating maintenance record.', 'error');
    }
};
const createMaintenance = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          maintenance_id: parseInt(newMaintenanceId),
          v_id: parseInt(newVehicleId),
          issue_reported: newIssueReported,
          repair_status: newRepairStatus,
        }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to create maintenance record');
      }
  
      await fetchMaintenanceData(); // Refresh the maintenance data
      showSnackbar('Maintenance record created successfully!');
  
      // ✅ Hide the form after successful creation
      setIsCreating(false); 
  
      // ✅ Clear the form fields
      setNewMaintenanceId('');
      setNewVehicleId('');
      setNewIssueReported('');
      setNewRepairStatus('Pending');
    } catch (error) {
      console.error('Error creating maintenance record:', error);
      showSnackbar('Error creating maintenance record.', 'error');
    }
};
const deleteMaintenance = async (maintenanceId) => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance/${maintenanceId}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to delete maintenance record');
      }
  
      // Fetch updated maintenance data
      await fetchMaintenanceData(); // Refresh the maintenance data
      showSnackbar('Maintenance record deleted successfully!');
    } catch (error) {
      console.error('Error deleting maintenance record:', error);
      showSnackbar('Error deleting maintenance record.', 'error');
    }
};
const fetchMaintenanceData = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance/all', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to fetch maintenance data');
      }

      const data = await response.json();
      setMaintenanceData(data); // Store the fetched data
      setShowMaintenanceModal(true); // Show the modal
      showSnackbar('Maintenance data retrieved!');
    } catch (error) {
      console.error('Error fetching maintenance data:', error);
      showSnackbar('Error fetching maintenance data.', 'error');
    }
};
const fetchSingleMaintenance = async (maintenanceId) => {
    try {
      const token = localStorage.getItem('token');
      const res = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance/${maintenanceId}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`
        }
      });
      if (!res.ok) throw new Error('Failed to fetch maintenance record');
  
      const data = await res.json();
      ////console.log(data)
      // Assuming the response has keys: maintenance_id, v_id, issue_reported, repair_status
      setEditingMaintenanceId(data.maintenance_id);
      setEditVehicleId(data.v_id);
      setEditIssueReported(data.issue_reported);
      setEditRepairStatus(data.repair_status);
    } catch (err) {
      console.error(err);
      showSnackbar("Error fetching maintenance record.", "error");
    }
};
//===========================Maintenance History===========================
  const [maintenanceHistoryData, setMaintenanceHistoryData] = useState([]);
  const [showMaintenanceHistoryModal, setShowMaintenanceHistoryModal] = useState(false);
  const [newIssueReported, setNewIssueReported] = useState('');
  const [newRepairStatus, setNewRepairStatus] = useState('Pending');
  const [showCreateMaintenanceForm, setShowCreateMaintenanceForm] = useState(false);
  const fetchMaintenanceHistory = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance-history/all', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to fetch maintenance history');
      }

      const data = await response.json();
      setMaintenanceHistoryData(data); // Store the fetched data
      setShowMaintenanceHistoryModal(true); // Show the modal
      showSnackbar('Maintenance history retrieved!');
    } catch (error) {
      console.error('Error fetching maintenance history:', error);
      showSnackbar('Error fetching maintenance history.', 'error');
    }
};
const createMaintenanceHistory = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance-history/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          m_id: parseInt(newMaintenanceId),
          v_id: parseInt(newVehicleId),
        }),
      });

      if (!response.ok) {
        throw new Error('Failed to create maintenance history');
      }

      // Fetch updated maintenance history data
      await fetchMaintenanceHistory(); // Refresh the maintenance history data
      showSnackbar('Maintenance history created successfully!');
      setNewMaintenanceId('');
      setNewVehicleId('');
    } catch (error) {
      console.error('Error creating maintenance history:', error);
      showSnackbar('Error creating maintenance history.', 'error');
    }
};
const deleteMaintenanceHistory = async (m_id, v_id) => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/maintenance-history/${m_id}/${v_id}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to delete maintenance history');
      }

      // Fetch updated maintenance history data
      await fetchMaintenanceHistory(); // Refresh the maintenance history data
      showSnackbar('Maintenance history deleted successfully!');
    } catch (error) {
      console.error('Error deleting maintenance history:', error);
      showSnackbar('Error deleting maintenance history.', 'error');
    }
};
//===========================Accident History===========================
  const [accidentHistoryData, setAccidentHistoryData] = useState([]);
  const [showAccidentHistoryModal, setShowAccidentHistoryModal] = useState(false);
  // const [editVehicleId, setEditVehicleId] = useState('');
  // const [editIssueReported, setEditIssueReported] = useState('');
  // const [editRepairStatus, setEditRepairStatus] = useState('Pending');
  const fetchAccidentHistory = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/accident-history/', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error('Failed to fetch accident history');
      }

      const data = await response.json();
      setAccidentHistoryData(data); // Store the fetched data
      setShowAccidentHistoryModal(true); // Show the modal
      showSnackbar('Accident history retrieved!');
    } catch (error) {
      console.error('Error fetching accident history:', error);
      showSnackbar('Error fetching accident history.', 'error');
    }
  };
//===========================Operates on===========================
  const [operatesOnData, setOperatesOnData] = useState([]);
  const [showOperatesOnModal, setShowOperatesOnModal] = useState(false);
  const [newOperatesOnVehicleId, setNewOperatesOnVehicleId] = useState('');
  const [newOperatesOnStationId, setNewOperatesOnStationId] = useState('');
  const fetchOperatesOn = async () => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/operates_on/all', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to fetch operates on records');
    }

    const data = await response.json();
    setOperatesOnData(data); // Store the fetched data
    setShowOperatesOnModal(true); // Show the modal
    showSnackbar('Operates on records retrieved!');
  } catch (error) {
    console.error('Error fetching operates on records:', error);
    showSnackbar('Error fetching operates on records.', 'error');
  }
};
const createOperatesOn = async () => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/operates_on/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        v_id: parseInt(newOperatesOnVehicleId),
        s_id: parseInt(newOperatesOnStationId),
      }),
    });

    if (!response.ok) {
      throw new Error('Failed to create operates on record');
    }

    // Fetch updated operates on data
    await fetchOperatesOn(); // Refresh the operates on data
    showSnackbar('Operates on record created successfully!');
    setNewOperatesOnVehicleId('');
    setNewOperatesOnStationId('');
  } catch (error) {
    console.error('Error creating operates on record:', error);
    showSnackbar('Error creating operates on record.', 'error');
  }
};
const deleteOperatesOn = async (v_id, s_id) => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app /operates_on/${v_id}/${s_id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to delete operates on record');
    }

    // Fetch updated operates on data
    await fetchOperatesOn(); // Refresh the operates on data
    showSnackbar('Operates on record deleted successfully!');
  } catch (error) {
    console.error('Error deleting operates on record:', error);
    showSnackbar('Error deleting operates on record.', 'error');
  }
};
//===========================Payment===========================
  const [paymentsData, setPaymentsData] = useState([]);
  const [showPaymentsModal, setShowPaymentsModal] = useState(false);
  const [newPaymentId, setNewPaymentId] = useState('');
  const [newAmount, setNewAmount] = useState('');
  const [newMethod, setNewMethod] = useState('');
  const fetchPayments = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/payments/', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to fetch payments');
      }
  
      const data = await response.json();
  
      // If no data returned (null or empty array), redirect to dashboard
      if (!data || data.length === 0) {
        showSnackbar('No payment found.', 'info');
        window.location.href = '/dashboard';
        return;
      }
  
      setPaymentsData(data); // Store the fetched data
      setShowPaymentsModal(true); // Show the modal
      showSnackbar('Payments retrieved!');
    } catch (error) {
      console.error('Error fetching payments:', error);
      showSnackbar('Error fetching payments.', 'error');
    }
  };
  const createPayment = async () => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/payments/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        payment_id: parseInt(newPaymentId),
        amount: parseFloat(newAmount),
        method: newMethod,
      }),
    });

    if (!response.ok) {
      throw new Error('Failed to create payment');
    }

    // Fetch updated payments data
    await fetchPayments(); // Refresh the payments data
    showSnackbar('Payment created successfully!');
    setNewPaymentId('');
    setNewAmount('');
    setNewMethod('');
  } catch (error) {
    console.error('Error creating payment:', error);
    showSnackbar('Error creating payment.', 'error');
  }
};
const updatePayment = async () => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/payments/', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
      body: JSON.stringify({
        payment_id: parseInt(newPaymentId),
        amount: parseFloat(newAmount),
        method: newMethod,
      }),
    });

    if (!response.ok) {
      throw new Error('Failed to update payment');
    }

    // Fetch updated payments data
    await fetchPayments(); // Refresh the payments data
    showSnackbar('Payment updated successfully!');
  } catch (error) {
    console.error('Error updating payment:', error);
    showSnackbar('Error updating payment.', 'error');
  }
};
const deletePayment = async (payment_id) => {
  try {
    const token = localStorage.getItem('token');
    const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/payments/${payment_id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error('Failed to delete payment');
    }

    // Fetch updated payments data
    await fetchPayments(); // Refresh the payments data
    showSnackbar('Payment deleted successfully!');
  } catch (error) {
    console.error('Error deleting payment:', error);
    showSnackbar('Error deleting payment.', 'error');
  }
};
//===========================Performs Maintenance===========================
  const [performsMaintenanceData, setPerformsMaintenanceData] = useState([]);
  const [showPerformsMaintenanceModal, setShowPerformsMaintenanceModal] = useState(false);
  const [newStaffId, setNewStaffId] = useState('');
  const fetchPerformsMaintenance = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/performs-maintenance/all', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to fetch performs maintenance records');
      }
  
      const data = await response.json();
      setPerformsMaintenanceData(data); // Store the fetched data
      setShowPerformsMaintenanceModal(true); // Show the modal
      showSnackbar('Performs maintenance records retrieved!');
    } catch (error) {
      console.error('Error fetching performs maintenance records:', error);
      showSnackbar('Error fetching performs maintenance records.', 'error');
    }
  };
  const createPerformsMaintenance = async () => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/performs-maintenance/', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          m_id: parseInt(newMaintenanceId),
          staff_id: parseInt(newStaffId),
        }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to create performs maintenance record');
      }
  
      // Fetch updated performs maintenance data
      await fetchPerformsMaintenance(); // Refresh the performs maintenance data
      showSnackbar('Performs maintenance record created successfully!');
      setNewMaintenanceId('');
      setNewStaffId('');
    } catch (error) {
      console.error('Error creating performs maintenance record:', error);
      showSnackbar('Error creating performs maintenance record.', 'error');
    }
  };
  const deletePerformsMaintenance = async (m_id, staff_id) => {
    try {
      const token = localStorage.getItem('token');
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/performs-maintenance/${m_id}/${staff_id}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to delete performs maintenance record');
      }
  
      // Fetch updated performs maintenance data
      await fetchPerformsMaintenance(); // Refresh the performs maintenance data
      showSnackbar('Performs maintenance record deleted successfully!');
    } catch (error) {
      console.error('Error deleting performs maintenance record:', error);
      showSnackbar('Error deleting performs maintenance record.', 'error');
    }
  };
//===========================Route===========================
  const [showRouteModal, setShowRouteModal] = useState(false);
  const [routeOperation, setRouteOperation] = useState('');
  const [routeId, setRouteId] = useState('');
  const [journeyTime, setJourneyTime] = useState('');
  const [newRouteFollowedId, setNewRouteFollowedId] = useState('');
  const [startPoint, setStartPoint] = useState('');
  const [endPoint, setEndPoint] = useState('');
  const [distance, setDistance] = useState('');
  const [routeData, setRouteData] = useState(null);
  const createRoute = async () => {
    try {
      const token = localStorage.getItem("token");
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/routes/${routeId}`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          JOURNEY_TIME: journeyTime,
          START_POINT: startPoint,
          END_POINT: endPoint,
          DISTANCE: parseFloat(distance),
        }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to create route');
      }
  
      showSnackbar('Route created successfully!');
      setShowRouteModal(false);
      resetRouteFields();
    } catch (error) {
      console.error('Error creating route:', error);
      showSnackbar('Error creating route.', 'error');
    }
  };
  const viewRoute = async () => {
    try {
      const token = localStorage.getItem("token");
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/routes/${routeId}`, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });
      console.log('Journey Time:', routeData.journey_time); 
      if (!response.ok) {
        throw new Error('Failed to fetch route');
      }
  
      const data = await response.json();
      console.log('Fetched route:', data); // Log the response
      setRouteData(data); // Store the fetched route data
      showSnackbar('Route fetched successfully!');
    } catch (error) {
      console.error('Error fetching route:', error);
      showSnackbar('Error fetching route.', 'error');
    }
  };
  const editRoute = async () => {
    try {
      const token = localStorage.getItem("token");
  
      // Log the journey time for debugging
      //console.log('Journey Time:', journeyTime);
  
  
      // Get the current date
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/routes/${routeId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify({
          JOURNEY_TIME: journeyTime,
          START_POINT: startPoint,
          END_POINT: endPoint,
          DISTANCE: parseFloat(distance),
        }),
      });
  
      if (!response.ok) {
        throw new Error('Failed to update route');
      }
  
      showSnackbar('Route updated successfully!');
      setShowRouteModal(false);
      resetRouteFields(); // Reset fields after successful update
    } catch (error) {
      console.error('Error updating route:', error);
      showSnackbar('Error updating route.', 'error');
    }
  };
  const deleteRoute = async () => {
    try {
      const token = localStorage.getItem("token");
      const response = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/routes/${routeId}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });
  
      if (!response.ok) {
        throw new Error('Failed to delete route');
      }
  
      showSnackbar('Route deleted successfully!');
      setShowRouteModal(false);
      // Reset fields
      resetRouteFields();
    } catch (error) {
      console.error('Error deleting route:', error);
      showSnackbar('Error deleting route.', 'error');
    }
  }; 
//===========================Route Followed===========================
  const [routeFollowedData, setRouteFollowedData] = useState([]);
  const [showRouteFollowedModal, setShowRouteFollowedModal] = useState(false);
  const [scheduleFollowedData, setScheduleFollowedData] = useState([]);
  const createRouteFollowed = async () => {
    try {
      const token = localStorage.getItem("token");
  
      const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/route-followed/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ r_id: parseInt(newRouteFollowedId) }),
      });
  
      if (!res.ok) throw new Error("Failed to create route followed");
  
      await fetchRouteFollowed(); // Refresh list
      setNewRouteFollowedId('');
      showSnackbar("Route followed created!");
    } catch (err) {
      console.error(err);
      showSnackbar("Error creating route followed", "error");
    }
  };
  const deleteRouteFollowed = async (routeId) => {
    try {
      const token = localStorage.getItem("token");
  
      const res = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/route-followed/${routeId}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
  
      if (!res.ok) throw new Error("Failed to delete route followed");
  
      await fetchRouteFollowed(); // Refresh list
      showSnackbar("Route followed deleted!");
    } catch (err) {
      console.error(err);
      showSnackbar("Error deleting route followed", "error");
    }
  };
  
  const fetchRouteFollowed = async () => {
    try {
      const token = localStorage.getItem("token");
  
      const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/route-followed/", {
        headers: { Authorization: `Bearer ${token}` },
      });
  
      if (!res.ok) throw new Error("Failed to fetch route followed data");
  
      const data = await res.json();
  
      // Expecting data.route_id to be an array
      setRouteFollowedData(data.route_id || []);
      setShowRouteFollowedModal(true);
      showSnackbar("Route followed data fetched!");
    } catch (err) {
      console.error(err);
      showSnackbar("Error fetching route followed data", "error");
    }
  };  
  const resetRouteFields = () => {
    setRouteId('');
    setJourneyTime('');
    setStartPoint('');
    setEndPoint('');
    setDistance('');
  };  
//===========================Schedule Followed===========================
  const [showScheduleFollowedModal, setShowScheduleFollowedModal] = useState(false);
  const [newScheduleRId, setNewScheduleRId] = useState('');
  const fetchScheduleFollowed = async () => {
    try {
      const token = localStorage.getItem("token");
  
      const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/schedule-followed/", {
        headers: { Authorization: `Bearer ${token}` },
      });
  
      if (!res.ok) throw new Error("Failed to fetch schedule followed");
  
      const data = await res.json();
      setScheduleFollowedData(data["Schedule-Followed ID"] || []);
      setScheduleVehicleId(data["Vehicle ID"]);
      setShowScheduleFollowedModal(true);
      showSnackbar("Schedule followed data fetched!");
    } catch (err) {
      console.error(err);
      showSnackbar("Error fetching schedule followed", "error");
    }
  };
  const createScheduleFollowed = async () => {
    try {
      const token = localStorage.getItem("token");
  
      const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/schedule-followed/", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          s_id: parseInt(newScheduleSId),
          r_id: parseInt(newScheduleRId),
        }),
      });
  
      if (!res.ok) throw new Error("Failed to create schedule followed");
  
      await fetchScheduleFollowed();
      setNewScheduleSId('');
      setNewScheduleRId('');
      showSnackbar("Schedule followed created!");
    } catch (err) {
      console.error(err);
      showSnackbar("Error creating schedule followed", "error");
    }
  };
//===========================Schedule===========================
  const [scheduleData, setScheduleData] = useState([]);
  const [showScheduleModal, setShowScheduleModal] = useState(false);
  const [newScheduleId, setNewScheduleId] = useState('');
  const [newScheduleVId, setNewScheduleVId] = useState('');
  const [newDepartureTime, setNewDepartureTime] = useState('');
  const [newArrivalTime, setNewArrivalTime] = useState('');
  const [newScheduleSId, setNewScheduleSId] = useState('');
  const fetchSchedule = async () => {
    try {
      const token = localStorage.getItem("token");
      const res = await fetch("https://smartcity-backend-try1-301261782088.asia-south1.run.app/schedule/", {
        headers: { Authorization: `Bearer ${token}` },
      });
      
      if (!res.ok) throw new Error("Failed to fetch schedule");
      
      const data = await res.json();
      console.log(data)
      const normalized = Array.isArray(data) ? data : [data];
      setScheduleData(normalized);
      setShowScheduleModal(true);
      showSnackbar("Schedule data fetched!");
    } catch (err) {
      console.error(err);
      showSnackbar("Error fetching schedule", "error");
    }
  };
  const deleteSchedule = async (scheduleId) => {
    try {
      const token = localStorage.getItem("token");
      const res = await fetch(`https://smartcity-backend-try1-301261782088.asia-south1.run.app/schedule/${scheduleId}`, {
        method: "DELETE",
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
  
      if (!res.ok) throw new Error("Failed to delete schedule");
  
      showSnackbar("Schedule deleted successfully!");
      fetchSchedule();
    } catch (err) {
      console.error(err);
      showSnackbar("Error deleting schedule", "error");
    }
  };
  const createOrUpdateSchedule = async () => {
    const body = {
      schedule_id: parseInt(newScheduleId),
      r_id: parseInt(newScheduleRId),
      v_id: parseInt(newScheduleVId),
      departure_time: formatTimeToISOString(newDepartureTime),
      arrival_time: formatTimeToISOString(newArrivalTime),
    };
    
    console.log("Sending payload:", body);
  
    try {
      const token = localStorage.getItem("token");
  
      const isUpdate = scheduleData.some(s => s.schedule_id === body.schedule_id);
  
      const url = isUpdate
        ? `https://smartcity-backend-try1-301261782088.asia-south1.run.app/schedule/${body.schedule_id}`
        : `https://smartcity-backend-try1-301261782088.asia-south1.run.app/schedule/`;
  
      const method = isUpdate ? "put" : "post";
  
      const res = await axios[method](url, body, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
  
      showSnackbar(isUpdate ? "Schedule updated!" : "Schedule created!");
      fetchSchedule();
      resetForm();
    } catch (err) {
      console.error("Error saving schedule", err);
    }
  };
//===========================Helpers===========================
const resetForm = () => {
  setNewScheduleId('');
  setNewScheduleRId('');
  setNewScheduleVId('');
  setNewDepartureTime('');
  setNewArrivalTime('');
};
function formatTimeToISOString(timeStr) {
  // Ensure timeStr is in "HH:mm:ss" format
  if (/^\d{2}:\d{2}$/.test(timeStr)) {
    timeStr += ":00";
  }

  return `1970-01-01T${timeStr}Z`;
}
const theme = createTheme({
  palette: {
    mode: isDarkMode ? 'dark' : 'light',
  },
});
//==========================Logout===============================
const logout = async () => {
  try {
    const token = localStorage.getItem('token');
    await fetch('https://smartcity-backend-try1-301261782088.asia-south1.run.app/auth/logout', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${token}`,
      },
    });
  } catch (error) {
    console.error('Logout failed:', error);
  }

  localStorage.removeItem('token');
  showSnackbar('Logged out successfully!', 'info');
  window.location.href = '/human-create'; // or use navigate('/login') with React Router
};

//===============================================================
return (
  <div className="container mt-5">
  <div className="max-w-7xl mx-auto px-6 py-16">
  <h1 className="text-4xl font-semibold text-center text-gray-800 dark:text-white mb-12">
      Smart City Transport System
    </h1>
    <h2>Welcome Maintenance Manager</h2>

    <div className="flex justify-center mb-10 gap-4">
      <button
        className="bg-white/30 backdrop-blur-md px-4 py-2 rounded-lg shadow hover:shadow-lg transition duration-300 dark:bg-white/10 dark:text-white"
        onClick={() => {
          toggleModal();
          getProfile();
        }}
      >
        Profile
      </button>
      <button
        className="bg-white/30 backdrop-blur-md px-4 py-2 rounded-lg shadow hover:shadow-lg transition duration-300 dark:bg-white/10 dark:text-white"
        onClick={toggleVehiclesModal}
      >
        Vehicles
      </button>
    </div>
        <div style={{ position: "fixed", bottom: "20px", right: "20px", zIndex: 999 }}>
  { <Dropdown>
  <Dropdown.Toggle variant="success" id="dropdown-basic">
    ⚙️
      <i className="fas fa-cog"></i>
    </Dropdown.Toggle>

    <Dropdown.Menu>
      {/* <Dropdown.Item onClick={fetchAccidentHistory}>Accident History</Dropdown.Item> */}
      <Dropdown.Item onClick={fetchMaintenanceHistory}>Maintenance History</Dropdown.Item>
      <Dropdown.Item onClick={fetchOperatesOn}>Operates On</Dropdown.Item>
      <Dropdown.Item onClick={fetchPerformsMaintenance}>Performs Maintenance</Dropdown.Item>
      {/* <Dropdown.Item onClick={fetchScheduleFollowed}>Schedule Followed</Dropdown.Item> */}
      {/* <Dropdown.Item onClick={fetchRouteFollowed}>Route Followed</Dropdown.Item> */}
    </Dropdown.Menu>
  </Dropdown> }
</div>

      </div>
{showScheduleFollowedModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Schedule Followed</h5>
          <button type="button" className="btn-close" onClick={() => setShowScheduleFollowedModal(false)}></button>
        </div>
        <div className="modal-body">
          <p><strong>Vehicle ID:</strong> {scheduleVehicleId}</p>
          {scheduleFollowedData.length > 0 ? (
            <ul className="list-group mb-3">
            {scheduleFollowedData.map((entry, index) => (
              <li key={index} className="list-group-item">
                Route ID: {entry.r_id}, Schedule ID: {entry.s_id}
              </li>
            ))}
          </ul>
          ) : (
            <p>No schedule followed records found.</p>
          )}

          <h6 className="mt-3">Add New Schedule Followed</h6>
          <div className="input-group mb-2">
            <input
              type="number"
              className="form-control"
              placeholder="Route ID"
              value={newScheduleRId}
              onChange={(e) => setNewScheduleRId(e.target.value)}
            />
            <input
              type="number"
              className="form-control"
              placeholder="Schedule ID"
              value={newScheduleSId}
              onChange={(e) => setNewScheduleSId(e.target.value)}
            />
            <button className="btn btn-success" onClick={createScheduleFollowed}>
              Create
            </button>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowScheduleFollowedModal(false)}>
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
)}

{showRouteFollowedModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Route Followed</h5>
          <button type="button" className="btn-close" onClick={() => setShowRouteFollowedModal(false)}></button>
        </div>
        <div className="modal-body">
          {routeFollowedData.length > 0 ? (
            <ul className="list-group mb-3">
              {routeFollowedData.map((id, index) => (
                <li key={index} className="list-group-item d-flex justify-content-between align-items-center">
                  Route ID: {id}
                  <button className="btn btn-danger btn-sm" onClick={() => deleteRouteFollowed(id)}>
                    Delete
                  </button>
                </li>
              ))}
            </ul>
          ) : (
            <p>No route followed data available.</p>
          )}

          <div className="mt-3">
            <h6>Add New Route Followed</h6>
            <div className="input-group">
              <input
                type="number"
                className="form-control"
                placeholder="Enter Route ID"
                value={newRouteFollowedId}
                onChange={(e) => setNewRouteFollowedId(e.target.value)}
              />
              <button className="btn btn-success" onClick={createRouteFollowed}>
                Create
              </button>
            </div>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowRouteFollowedModal(false)}>
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
)}


{showPerformsMaintenanceModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Performs Maintenance Records</h5>
          <button type="button" className="btn-close" onClick={() => setShowPerformsMaintenanceModal(false)}></button>
        </div>
        <div className="modal-body">
          {performsMaintenanceData.length > 0 ? (
            <div className="table-responsive">
              <table className="table table-striped table-bordered">
                <thead>
                  <tr>
                    <th>Maintenance ID</th>
                    <th>Staff ID</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {performsMaintenanceData.map((record, index) => (
                    <tr key={index}>
                      <td>{record.m_id}</td>
                      <td>{record.staff_id}</td>
                      <td>
                        <button
                          className="btn btn-danger btn-sm"
                          onClick={() => deletePerformsMaintenance(record.m_id, record.staff_id)}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p>No performs maintenance records available.</p>
          )}
          <div className="mt-3">
            <h6>Create New Performs Maintenance Record</h6>
            <div className="mb-2">
              <input
                type="number"
                className="form-control"
                placeholder="Maintenance ID"
                value={newMaintenanceId}
                onChange={(e) => setNewMaintenanceId(e.target.value)}
              />
            </div>
            <div className="mb-2">
              <input
                type="number"
                className="form-control"
                placeholder="Staff ID"
                value={newStaffId}
                onChange={(e) => setNewStaffId(e.target.value)}
              />
            </div>
            <button className="btn btn-success" onClick={createPerformsMaintenance}>
              Create
            </button>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowPerformsMaintenanceModal(false)}>Close</button>
        </div>
      </div>
    </div>
  </div>
)}
{showOperatesOnModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Operates On Records</h5>
          <button type="button" className="btn-close" onClick={() => setShowOperatesOnModal(false)}></button>
        </div>
        <div className="modal-body">
          {operatesOnData.length > 0 ? (
            <div className="table-responsive">
              <table className="table table-striped table-bordered">
                <thead>
                  <tr>
                    <th>Vehicle ID</th>
                    <th>Station ID</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {operatesOnData.map((record, index) => (
                    <tr key={index}>
                      <td>{record.v_id}</td>
                      <td>{record.s_id}</td>
                      <td>
                        <button
                          className="btn btn-danger btn-sm"
                          onClick={() => deleteOperatesOn(record.v_id, record.s_id)}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p>No operates on records available.</p>
          )}
          <div className="mt-3">
            <h6>Create New Operates On Record</h6>
            <div className="mb-2">
              <input
                type="number"
                className="form-control"
                placeholder="Vehicle ID"
                value={newOperatesOnVehicleId}
                onChange={(e) => setNewOperatesOnVehicleId(e.target.value)}
              />
            </div>
            <div className="mb-2">
              <input
                type="number"
                className="form-control"
                placeholder="Station ID"
                value={newOperatesOnStationId}
                onChange={(e) => setNewOperatesOnStationId(e.target.value)}
              />
            </div>
            <button className="btn btn-success" onClick={createOperatesOn}>
              Create
            </button>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowOperatesOnModal(false)}>Close</button>
        </div>
      </div>
    </div>
  </div>
)}
{showMaintenanceHistoryModal && (
        <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
          <div className="modal-dialog" role="document">
            <div className="modal-content p-3">
              <div className="modal-header">
                <h5 className="modal-title">Maintenance History</h5>
                <button type="button" className="btn-close" onClick={() => setShowMaintenanceHistoryModal(false)}></button>
              </div>
              <div className="modal-body">
                {maintenanceHistoryData.length > 0 ? (
                  <div className="table-responsive">
                    <table className="table table-striped table-bordered">
                      <thead>
                        <tr>
                          <th>M_ID</th>
                          <th>V_ID</th>
                          <th>Actions</th>
                        </tr>
                      </thead>
                      <tbody>
                        {maintenanceHistoryData.map((item, index) => (
                          <tr key={index}>
                            <td>{item.m_id}</td>
                            <td>{item.v_id}</td>
                            <td>
                              <button
                                className="btn btn-danger btn-sm"
                                onClick={() => deleteMaintenanceHistory(item.m_id, item.v_id)}
                              >
                                Delete
                              </button>
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                ) : (
                  <p>No maintenance history available.</p>
                )}
                <div className="mt-3">
                  <h6>Create New Maintenance History</h6>
                  <div className="mb-2">
                    <input
                      type="number"
                      className="form-control"
                      placeholder="Maintenance ID"
                      value={newMaintenanceId}
                      onChange={(e) => setNewMaintenanceId(e.target.value)}
                    />
                  </div>
                  <div className="mb-2">
                    <input
                      type="number"
                      className="form-control"
                      placeholder="Vehicle ID"
                      value={newVehicleId}
                      onChange={(e) => setNewVehicleId(e.target.value)}
                    />
                  </div>
                  <button className="btn btn-success" onClick={createMaintenanceHistory}>
                    Create
                  </button>
                </div>
              </div>
              <div className="modal-footer">
                <button className="btn btn-secondary" onClick={() => setShowMaintenanceHistoryModal(false)}>Close</button>
              </div>
            </div>
          </div>
        </div>
)}
{showAccidentHistoryModal && (
        <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
          <div className="modal-dialog" role="document">
            <div className="modal-content p-3">
              <div className="modal-header">
                <h5 className="modal-title">Accident History</h5>
                <button type="button" className="btn-close" onClick={() => setShowAccidentHistoryModal(false)}></button>
              </div>
              <div className="modal-body">
                {accidentHistoryData.length > 0 ? (
                  <div className="table-responsive">
                    <table className="table table-striped table-bordered">
                      <thead>
                        <tr>
                          <th>Vehicle ID</th>
                          <th>Incident ID</th>
                        </tr>
                      </thead>
                      <tbody>
                        {accidentHistoryData.map((item, index) => (
                          <tr key={index}>
                            <td>{item.v_id}</td>
                            <td>{item.i_id}</td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                ) : (
                  <p>No accident history available.</p>
                )}
              </div>
              <div className="modal-footer">
                <button className="btn btn-secondary" onClick={() => setShowAccidentHistoryModal(false)}>Close</button>
              </div>
            </div>
          </div>
        </div>
)}
<div className="position-absolute top-0 end-0 p-3 d-flex gap-2 align-items-start">
        {/* Dark Mode Button */}
        <button
          className={`btn btn-sm ${isDarkMode ? 'btn-light' : 'btn-dark'}`}
          onClick={() => setIsDarkMode(!isDarkMode)}
        >
          {isDarkMode ? '☀️' : '🌙'}
        </button>

        {/* Three Dots Dropdown */}
        <div className="dropdown">
          <button
            className="btn btn-dark"
            type="button"
            id="dropdownMenuButton"
            data-bs-toggle="dropdown"
            aria-expanded="false"
          >
            &#8942;
          </button>
          <ul
            className="dropdown-menu dropdown-menu-end"
            aria-labelledby="dropdownMenuButton"
          >
            <li>
              <button className="dropdown-item" onClick={fetchIncidentData}>
                Incident
              </button>
            </li>
            <li>
        <button className="dropdown-item" onClick={fetchMaintenanceData}>
          Maintenance
        </button>
      </li>
            <li>
        <button className="dropdown-item" onClick={fetchPayments}>
          Payments
        </button>
      </li>
            <li>
            <button className="dropdown-item" onClick={fetchSchedule}>
  Schedule
</button>

      </li>
            <li>
            <button className="dropdown-item" onClick={() => {
  setShowRouteModal(true);
}}>
  Manage Route
</button>
      </li>
          </ul>
        </div>
</div>
{showScheduleModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog modal-lg" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Schedules</h5>
          <button type="button" className="btn-close" onClick={() => setShowScheduleModal(false)}></button>
        </div>
        <div className="modal-body">
          {scheduleData.length > 0 ? (
            <table className="table table-bordered">
              <thead>
                <tr>
                  <th>Schedule ID</th>
                  <th>Route ID</th>
                  <th>Vehicle ID</th>
                  <th>Departure</th>
                  <th>Arrival</th>
                  <th>Actions</th>

                </tr>
              </thead>
              <tbody>
  {scheduleData.map((item, index) => (
    
    <tr key={index}>
      <td>{item.schedule_id}</td>
      <td>{item.r_id}</td>
      <td>{item.v_id}</td>
      <td>{new Date(item.departure_time).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}</td>
      <td>{new Date(item.arrival_time).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}</td>
      <td>
        <button
          className="btn btn-warning btn-sm me-1"
          onClick={() => {
            setNewScheduleId(item.schedule_id);
            setNewScheduleRId(item.r_id);
            setNewScheduleVId(item.v_id);
            setNewDepartureTime(item.departure_time.split('T')[1].slice(0, 5));
            setNewArrivalTime(item.arrival_time.split('T')[1].slice(0, 5));
          }}
        >
          
          
          Edit
        </button>
        <button
          className="btn btn-danger btn-sm"
          onClick={() => deleteSchedule(item.schedule_id)}
        >
          Delete
        </button>
      </td>
    </tr>
  ))}
</tbody>


            </table>
          ) : (
            <p>No schedule data available.</p>
          )}

          <h6 className="mt-3">Add New Schedule</h6>
          <div className="row g-2">
            <div className="col">
              <input
                type="number"
                className="form-control"
                placeholder="Schedule ID"
                value={newScheduleId}
                onChange={(e) => setNewScheduleId(e.target.value)}
              />
            </div>
            <div className="col">
              <input
                type="number"
                className="form-control"
                placeholder="Route ID"
                value={newScheduleRId}
                onChange={(e) => setNewScheduleRId(e.target.value)}
              />
            </div>
            <div className="col">
              <input
                type="number"
                className="form-control"
                placeholder="Vehicle ID"
                value={newScheduleVId}
                onChange={(e) => setNewScheduleVId(e.target.value)}
              />
            </div>
            <div className="col">
              <input
                type="time"
                className="form-control"
                placeholder="Departure Time"
                value={newDepartureTime}
                onChange={(e) => setNewDepartureTime(e.target.value)}
              />
            </div>
            <div className="col">
              <input
                type="time"
                className="form-control"
                placeholder="Arrival Time"
                value={newArrivalTime}
                onChange={(e) => setNewArrivalTime(e.target.value)}
              />
            </div>
            <div className="col-12 mt-2">
            <button className="btn btn-success" onClick={createOrUpdateSchedule}>
  {scheduleData.some(s => s.schedule_id === parseInt(newScheduleId)) ? "Update" : "Create"}
</button>

            </div>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowScheduleModal(false)}>
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
)}
{showRouteModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Manage Route</h5>
          <button type="button" className="btn-close" onClick={() => setShowRouteModal(false)}></button>
        </div>
        <div className="modal-body">
          <div className="mb-3">
            <label className="form-label">Select Action</label>
            <select
              className="form-select"
              value={routeOperation}
              onChange={(e) => setRouteOperation(e.target.value)}
            >
              <option value="CREATE">Create</option>
              <option value="VIEW">View</option>
              <option value="EDIT">Edit</option>
              <option value="DELETE">Delete</option>
            </select>
          </div>
          <div className="mb-3">
            <label className="form-label">Route ID</label>
            <input
              type="text"
              className="form-control"
              value={routeId}
              onChange={(e) => setRouteId(e.target.value)}
            />
          </div>
          {routeOperation === 'EDIT'&& (
            <>
              <div className="mb-3">
                <label className="form-label">Journey Time</label>
                <input
                  type="text"
                  className="form-control"
                  value={journeyTime}
                  onChange={(e) => setJourneyTime(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Start Point</label>
                <input
                  type="text"
                  className="form-control"
                  value={startPoint}
                  onChange={(e) => setStartPoint(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">End Point</label>
                <input
                  type="text"
                  className="form-control"
                  value={endPoint}
                  onChange={(e) => setEndPoint(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Distance</label>
                <input
                  type="number"
                  className="form-control"
                  value={distance}
                  onChange={(e) => setDistance(e.target.value)}
                />
              </div>
            </>
          )}
          {routeOperation === 'CREATE'&& (
            <>
              <div className="mb-3">
                <label className="form-label">Journey Time</label>
                <input
                  type="text"
                  className="form-control"
                  value={journeyTime}
                  onChange={(e) => setJourneyTime(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Start Point</label>
                <input
                  type="text"
                  className="form-control"
                  value={startPoint}
                  onChange={(e) => setStartPoint(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">End Point</label>
                <input
                  type="text"
                  className="form-control"
                  value={endPoint}
                  onChange={(e) => setEndPoint(e.target.value)}
                />
              </div>
              <div className="mb-3">
                <label className="form-label">Distance</label>
                <input
                  type="number"
                  className="form-control"
                  value={distance}
                  onChange={(e) => setDistance(e.target.value)}
                />
              </div>
            </>
          )}
{routeOperation === 'VIEW' && routeData && (
  <div className="mt-3">
    <h6>Route Details</h6>
    <p><strong>Route ID:</strong> {routeData.r_id}</p>
    <p><strong>Journey Time:</strong> {routeData.journey_time} Hours</p>
    <p><strong>Start Point:</strong> {routeData.start_point}</p>
    <p><strong>End Point:</strong> {routeData.end_point}</p>
    <p><strong>Distance:</strong> {routeData.distance} km</p>
  </div>
)}
          <div className="mb-3">
            <button className="btn btn-primary" onClick={() => {
              if (routeOperation === 'CREATE') createRoute();
              else if (routeOperation === 'VIEW') viewRoute();
              else if (routeOperation === 'EDIT') editRoute();
              else if (routeOperation === 'DELETE') deleteRoute();
            }}>
              {routeOperation}
            </button>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowRouteModal(false)}>Close</button>
        </div>
      </div>
    </div>
  </div>
)}
{showPaymentsModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Payments</h5>
          <button type="button" className="btn-close" onClick={() => setShowPaymentsModal(false)}></button>
        </div>
        <div className="modal-body">
          {paymentsData.length > 0 ? (
            <div className="table-responsive">
              <table className="table table-striped table-bordered">
                <thead>
                  <tr>
                    <th>Payment ID</th>
                    <th>Passenger ID</th>
                    <th>Amount</th>
                    <th>Method</th>
                    <th>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {paymentsData.map((payment, index) => (
                    <tr key={index}>
                      <td>{payment.payment_id}</td>
                      <td>{payment.passenger_id}</td>
                      <td>{payment.amount}</td>
                      <td>{payment.method}</td>
                      <td>
                        <button
                          className="btn btn-danger btn-sm"
                          onClick={() => deletePayment(payment.payment_id)}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p>No payment records available.</p>
          )}
          <div className="mt-3">
            <h6>Create New Payment</h6>
            <div className="mb-2">
              <input
                type="number"
                className="form-control"
                placeholder="Payment ID"
                value={newPaymentId}
                onChange={(e) => setNewPaymentId(e.target.value)}
              />
            </div>
            <div className="mb-2">
              <input
                type="number"
                className="form-control"
                placeholder="Amount"
                value={newAmount}
                onChange={(e) => setNewAmount(e.target.value)}
              />
            </div>
            <div className="mb-2">
                <select
                  className="form-control"
                  value={newMethod}
                  onChange={(e) => setNewMethod(e.target.value)}
                  required
                >
                  <option value="">Select Payment Method</option>
                  <option value="Cash">Cash</option>
                  <option value="Credit Card">Credit Card</option>
                  <option value="Debit Card">Debit Card</option>
                  <option value="Mobile Payment">Mobile Payment</option>
                  <option value="Bank Transfer">Bank Transfer</option>
                  <option value="UPI">UPI</option>
                  <option value="Crypto">Crypto</option>
                </select>
              </div>
            <button className="btn btn-success" onClick={createPayment}>
              Create
            </button>
          </div>
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowPaymentsModal(false)}>Close</button>
        </div>
      </div>
    </div>
  </div>
)}
{showMaintenanceModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Maintenance Records</h5>
          <button type="button" className="btn-close" onClick={() => setShowMaintenanceModal(false)}></button>
        </div>
        <div className="modal-body">
          {Array.isArray(maintenanceData) && maintenanceData.length > 0 ?(
            <div className="table-responsive">
              <table className="table table-striped table-bordered">
                <thead>
                  <tr>
                    <th style={{ width: '20%' }}>Maintenance ID</th>
                    <th style={{ width: '20%' }}>Vehicle ID</th>
                    <th style={{ width: '40%' }}>Issue Reported</th>
                    <th style={{ width: '10%' }}>Repair Status</th>
                    <th style={{ width: '10%' }}>Actions</th>
                  </tr>
                </thead>
                <tbody>
                  {maintenanceData.map((item, index) => (
                    <tr key={index}>
                      <td>{item.maintenance_id}</td>
                      <td>{item.v_id}</td>
                      <td>{item.issue_reported}</td>
                      <td>{item.repair_status}</td>
                      <td>
                        <button
                          className="btn btn-warning btn-sm me-1"
                          onClick={() => {
                            setEditingMaintenanceId(item.maintenance_id);
                            setEditVehicleId(item.v_id);
                            setEditIssueReported(item.issue_reported);
                            setEditRepairStatus(item.repair_status);
                          }}
                        >
                          Edit
                        </button>
                        <button
                          className="btn btn-danger btn-sm"
                          onClick={() => deleteMaintenance(item.maintenance_id)}
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p>No maintenance records available.</p>
          )}
          {editingMaintenanceId && (
            <div className="mt-3">
              <h6>Edit Maintenance Record</h6>
              <div className="mb-2">
                <input
                  type="number"
                  className="form-control"
                  placeholder="Maintenance ID"
                  value={editingMaintenanceId}
                  readOnly // Make it read-only since it's being edited
                />
              </div>
              <div className="mb-2">
                <input
                  type="number"
                  className="form-control"
                  placeholder="Vehicle ID"
                  value={editVehicleId}
                  readOnly
/>
              </div>
              <div className="mb-2">
                <input
                  type="text"
                  className="form-control"
                  placeholder="Issue Reported"
                  value={editIssueReported}
                  onChange={(e) => setEditIssueReported(e.target.value)}
                />
              </div>
              <div className="mb-2">
                <select
                  className="form-control"
                  value={editRepairStatus}
                  onChange={(e) => setEditRepairStatus(e.target.value)}
                >
                  <option value="Pending">Pending</option>
                  <option value="In Progress">In Progress</option>
                  <option value="Completed">Completed</option>
                </select>
              </div>
              <button className="btn btn-success" onClick={updateMaintenance}>
                Update
              </button>
              <button className="btn btn-secondary ms-2" onClick={() => setEditingMaintenanceId(null)}>
                Cancel
              </button>
            </div>
          )}
        </div>
        <div className="modal-footer">
          <button className="btn btn-secondary" onClick={() => setShowMaintenanceModal(false)}>Close</button>
        </div>
      </div>
    </div>
  </div>
)}

{/* Incident Modal */}
{showIncidentModal && incidentData && (

  <div className="modal d-block" tabIndex="-1">
    <div className="modal-dialog modal-lg">
      <div className="modal-content">
        <div className="modal-header">
          <h5 className="modal-title">Incident Reports</h5>
          <button type="button" className="btn-close" onClick={() => setShowIncidentModal(false)}></button>
        </div>
        
        <div className="modal-body">
          {incidentData.length > 0 ? (
            <div className="table-responsive">
              <table className="table table-striped table-bordered">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Description</th>
                    <th>Time</th>
                    <th>Status</th>
                    <th>Vehicle ID</th>
                  </tr>
                </thead>
                <tbody>
                  {incidentData.map((incident, index) => (
                    <tr key={index}>
                      <td className="d-flex align-items-center">
                        {incident.incident_id}
                        <button
                          className="btn btn-danger btn-sm ms-2"
                          onClick={() => deleteIncident(incident.incident_id)}
                        >
                          Delete
                        </button>
                      </td>
                      <td>
                          {editingIncidentId === incident.incident_id ? (
                            <>
                              <input
                                type="text"
                                className="form-control form-control-sm"
                                value={editedDescription}
                                onChange={(e) => setEditedDescription(e.target.value)}
                              />
                              <div className="mt-1">
                                <button
                                  className="btn btn-success btn-sm me-1"
                                  onClick={() => updateIncidentDescription(incident.incident_id, editedDescription)}
                                >
                                  Save
                                </button>
                                <button
                                  className="btn btn-secondary btn-sm"
                                  onClick={() => setEditingIncidentId(null)}
                                >
                                  Cancel
                                </button>
                              </div>
                            </>
                          ) : (
                            <>
                              {incident.description}
                              <button
                                className="btn btn-warning btn-sm ms-2"
                                onClick={() => {
                                  setEditingIncidentId(incident.incident_id);
                                  setEditedDescription(incident.description);
                                }}
                              >
                                Edit
                              </button>
                            </>
                          )}
                        </td>

                      <td>
                        {new Date(incident.report_time_date).toLocaleString('en-IN', {
                          dateStyle: 'medium',
                          timeStyle: 'short',
                        })}
                      </td>
                      {/* <td>{incident.status}</td> */}
                      <td>{incident.v_id}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <p>No incidents to show.</p>
          )}
        </div>
        {showCreateForm && (
                            <div className="p-3 border-top">
                              <h6>Create New Incident</h6>
                              <div className="mb-2">
                                <input
                                  type="number"
                                  className="form-control"
                                  placeholder="Incident ID"
                                  value={newIncidentId}
                                  onChange={(e) => setNewIncidentId(e.target.value)}
                                />
                              </div>
                              <div className="mb-2">
                                <input
                                  type="text"
                                  className="form-control"
                                  placeholder="Description"
                                  value={newDescription}
                                  onChange={(e) => setNewDescription(e.target.value)}
                                />
                              </div>
                              <button className="btn btn-success" onClick={createIncident}>
                                Submit
                              </button>
                            </div>
                          )}
                <div className="modal-footer">
                  <button type="button" className="btn btn-secondary" onClick={() => setShowIncidentModal(false)}>
                    Close
                  </button>
                  <button
            type="button"
            className="btn btn-primary"
            onClick={() => setShowCreateForm(!showCreateForm)}
          >
            {showCreateForm ? 'Cancel' : 'Create Incident'}
    </button>
        </div>
      </div>
    </div>
  </div>
)}

      {showModal && (
        <div className="modal d-block" tabIndex="-1" role="dialog" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
          <div className="modal-dialog" role="document">
            <div className="modal-content p-3">
              <div className="modal-header d-flex justify-content-between align-items-center">
                {/* <button className="btn btn-sm btn-light px-1 py-1 me-1" onClick={toggleModal}>
                  ⬅ Back
                </button> */}
                <h5 className="modal-title mb-0">👤 Profile</h5>
                {/* <button type="button" className="btn-close" onClick={toggleModal}>
                  ⬅ Back
                </button> */}
              </div>
              <div className="modal-body">
                {profileData ? (
                  <>
                    <p><strong>Name:</strong> {profileData.fname} {profileData.lname}</p>
                    <p><strong>DOB:</strong> {profileData.dob.split('T')[0]}</p>
                    {/* <p><strong>Age:</strong>{profileData.Age}</p> */}
                    <hr />
                    <h6 className="mb-3">📝 Edit Profile</h6>

                    <label className="form-label">First Name</label>
                    <input
                      type="text"
                      placeholder="First Name"
                      className="form-control mb-2"
                      value={formData.FName}
                      onChange={(e) => setFormData({ ...formData, FName: e.target.value })}
                    />

                    <label className="form-label">Last Name</label>
                    <input
                      type="text"
                      placeholder="Last Name"
                      className="form-control mb-2"
                      value={formData.LName}
                      onChange={(e) => setFormData({ ...formData, LName: e.target.value })}
                    />

                    <label className="form-label">Date of Birth</label>
                    <input
                      type="date"
                      className="form-control mb-2"
                      value={formData.DOB}
                      onChange={(e) => setFormData({ ...formData, DOB: e.target.value })}
                    />

                    <label className="form-label">Age</label>
                    <input
                      type="number"
                      placeholder="Age"
                      className="form-control mb-2"
                      value={formData.Age}
                      onChange={(e) => setFormData({ ...formData, Age: e.target.value })}
                    />

                    <label className="form-label">Vehicle ID</label>
                    <input
                      type="number"
                      placeholder="Vehicle ID"
                      className="form-control mb-3"
                      value={formData.V_ID}
                      onChange={(e) => setFormData({ ...formData, V_ID: e.target.value })}
                    />
                  </>
                ) : <p>Loading profile...</p>}
              </div>
              <div className="modal-footer">
                <button className="btn btn-success" onClick={updateProfile}>✏️ Save Changes</button>
                <button className="btn btn-danger" onClick={deleteProfile}>🗑️ Delete Profile</button>
                <button className="btn btn-secondary" onClick={toggleModal}>Close</button>
              </div>
            </div>
          </div>
        </div>
      )}
{showVehiclesModal && (
      <div className="modal d-block" tabIndex="-1" role="dialog" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
        <div className="modal-dialog" role="document">
          <div className="modal-content p-3">
            <div className="modal-header d-flex justify-content-between align-items-center">
              <h5 className="modal-title mb-0">🚗 Vehicles</h5>
              {/* <button type="button" className="btn-close" onClick={() => setShowVehiclesModal(false)}>⬅ Back</button> */}
            </div>
            <div className="modal-body">
            {vehiclesData.length > 0 ? (
              <ul className="list-group">
                {vehiclesData
                  .filter(vehicle => vehicle.status === "Available")
                  .map((vehicle, index) => (
                    <li key={index} className="list-group-item d-flex justify-content-between align-items-start">
                      <div>
                        <strong>ID:</strong> {vehicle.vehicle_id} <br />
                        <strong>Current Location:</strong> {vehicle.current_location}
                      </div>
                      <button
                        className="btn btn-sm btn-success"
                        onClick={() => bookVehicle(vehicle.vehicle_id)}
                      >
                        ✅ Select
                      </button>
                    </li>
                  ))}
              </ul>
            ) : (
              <p>No available vehicles found.</p>
            )}
            </div>
            <div className="modal-footer">
              <button className="btn btn-secondary" onClick={() => setShowVehiclesModal(false)}>Close</button>
            </div>
          </div>
        </div>
      </div>
)}
{showVehicleUpdateModal && (
  <div className="modal d-block" tabIndex="-1" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
    <div className="modal-dialog" role="document">
      <div className="modal-content p-3">
        <div className="modal-header">
          <h5 className="modal-title">Update Vehicle</h5>
          <button type="button" className="btn-close" onClick={() => setShowVehicleUpdateModal(false)}></button>
        </div>
        <div className="modal-body">
          <label className="form-label">Current Location</label>
          <input
            type="text"
            className="form-control mb-2"
            value={updateFormData.current_location}
            onChange={(e) => setUpdateFormData({ ...updateFormData, current_location: e.target.value })}
          />

<label>Status</label>
<select
  className="form-control"
  value={updateFormData.status}
  onChange={(e) =>
    setUpdateFormData({ ...updateFormData, status: e.target.value })
  }
>
<option value="Available">Available</option>
<option value="Booked">Booked</option>

</select>

        </div>
        <div className="modal-footer">
          <button className="btn btn-primary" onClick={updateVehicle}>Save</button>
          <button className="btn btn-secondary" onClick={() => setShowVehicleUpdateModal(false)}>Cancel</button>
        </div>
      </div>
    </div>
  </div>
)}
<div className="position-absolute bottom-0 start-0 p-2">
  <button className="btn btn-outline-danger" onClick={logout}>
    Logout
  </button>
</div>


<Snackbar
  open={openSnackbar}
  autoHideDuration={3000}
  onClose={() => setOpenSnackbar(false)}
  anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
>
  <Alert
    onClose={() => setOpenSnackbar(false)}
    severity={snackbarSeverity}
    sx={{ width: '100%' }}
  >
    {snackbarMessage}
  </Alert>
</Snackbar>
{showErrorModal && (
        <div className="modal d-block" tabIndex="-1" role="dialog" style={{ backgroundColor: 'rgba(0,0,0,0.5)' }}>
          <div className="modal-dialog" role="document">
            <div className="modal-content p-3">
              <div className="modal-header">
                <h5 className="modal-title">⚠️ Operation Failed</h5>
                <button type="button" className="btn-close" onClick={() => setShowErrorModal(false)}></button>
              </div>
              <div className="modal-body">
                <p>{errorMessage}</p>
              </div>
              <div className="modal-footer">
                <button className="btn btn-secondary" onClick={() => setShowErrorModal(false)}>Close</button>
              </div>
            </div>
          </div>
        </div>
      )}

    </div>

  );
}

export default Dashboard;
