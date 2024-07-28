import { createContext, useState, useEffect } from 'react';
import axios from '../api/axiosConfig';
import Swal from 'sweetalert2';
import { useNavigate } from 'react-router-dom';

// Context untuk User
export const UserContext = createContext();

// eslint-disable-next-line react/prop-types
export const UserProvider = ({ children }) => {
  const [users, setUsers] = useState([]);
  const [profile, setProfile] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const navigate = useNavigate();

  const handleError = (error) => {
    if (error.response && error.response.status === 401) {
      Swal.fire({
        icon: 'error',
        title: 'Access Denied',
        text: error.response.data.error || 'You do not have permission to access this page (401).',
      }).then(() => {
        navigate('/login');
      });
    } else {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: error.response.data.error || 'An unexpected error occurred.',
      });
    }
    setError(error);
  };

  useEffect(() => {
    const checkLoginStatus = () => {
      const token = localStorage.getItem('token');
      if (token) {
        setIsLoggedIn(true);
      } else {
        setIsLoggedIn(false);
      }
      setLoading(false);
    };

    checkLoginStatus();
  }, []);

  useEffect(() => {
    const fetchUsers = async () => {
      if (!isLoggedIn) return;

      setLoading(true);
      try {
        const response = await axios.get('/users');
        setUsers(response.data);
      } catch (error) {
        handleError(error);
      } finally {
        setLoading(false);
      }
    };

    fetchUsers();
  }, [isLoggedIn]);

  const fetchUserById = async (id) => {
    try {
      const response = await axios.get(`/users/${id}`);
      return response.data;
    } catch (error) {
      handleError(error);
    }
  };

  const fetchProfile = async () => {
    try {
      const response = await axios.get('/profile');
      setProfile(response.data);
    } catch (error) {
      handleError(error);
    }
  };

  const editProfile = async (profileData) => {
    try {
      const response = await axios.put('/profile', profileData);
      setProfile(response.data);
    } catch (error) {
      handleError(error);
    }
  };

  const addUser = async (userData) => {
    try {
      const response = await axios.post('/users', userData);
      setUsers([...users, response.data]);
    } catch (error) {
      handleError(error);
    }
  };

  const updateUser = async (id, userData) => {
    try {
      const response = await axios.put(`/users/${id}`, userData);
      setUsers(users.map((user) => (user.id === id ? response.data : user)));
    } catch (error) {
      handleError(error);
    }
  };

  const removeUser = async (id) => {
    try {
      await axios.delete(`/users/${id}`);
      setUsers(users.filter((user) => user.id !== id));
    } catch (error) {
      handleError(error);
    }
  };

  return (
    <UserContext.Provider
      value={{ users, profile, fetchUserById, fetchProfile, editProfile, addUser, updateUser, removeUser, loading, error }}
    >
      {children}
    </UserContext.Provider>
  );
};
