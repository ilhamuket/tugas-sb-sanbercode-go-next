import { createContext, useState, useEffect } from 'react';
import axios from '../api/axiosConfig';

// Context untuk User
export const UserContext = createContext();

// eslint-disable-next-line react/prop-types
export const UserProvider = ({ children }) => {
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [isLoggedIn, setIsLoggedIn] = useState(false);

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
        setError(error);
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
      setError(error);
    }
  };

  const addUser = async (userData) => {
    try {
      const response = await axios.post('/users', userData);
      setUsers([...users, response.data]);
    } catch (error) {
      setError(error);
    }
  };

  const updateUser = async (id, userData) => {
    try {
      const response = await axios.put(`/users/${id}`, userData);
      setUsers(users.map((user) => (user.id === id ? response.data : user)));
    } catch (error) {
      setError(error);
    }
  };

  const removeUser = async (id) => {
    try {
      await axios.delete(`/users/${id}`);
      setUsers(users.filter((user) => user.id !== id));
    } catch (error) {
      setError(error);
    }
  };

  return (
    <UserContext.Provider
      value={{ users, fetchUserById, addUser, updateUser, removeUser, loading, error }}
    >
      {children}
    </UserContext.Provider>
  );
};
