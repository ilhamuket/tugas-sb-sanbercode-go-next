import { createContext, useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import axiosInstance from '../api/axiosConfig';
import { login, register, changePassword } from '../api/auth';

export const AuthContext = createContext();

// eslint-disable-next-line react/prop-types
export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null); // Start with null to indicate unauthenticated
  const [loading, setLoading] = useState(true); // To handle loading state
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    const token = localStorage.getItem('token');
    const savedUser = localStorage.getItem('user');

    if (token) {
      axiosInstance.defaults.headers.common['Authorization'] = `Bearer ${token}`;
      if (savedUser) {
        setUser(JSON.parse(savedUser));
        setLoading(false);
      } else {
        const fetchUserDetails = async () => {
          try {
            const response = await axiosInstance.get('/profile');
            setUser(response.data);
            localStorage.setItem('user', JSON.stringify(response.data));
          } catch (error) {
            console.error('Failed to fetch user details:', error);
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            navigate('/login');
          } finally {
            setLoading(false);
          }
        };
        fetchUserDetails();
      }
    } else {
      setLoading(false);
      if (location.pathname !== '/login' && location.pathname !== '/register') {
        navigate('/login');
      }
    }
  }, [navigate, location.pathname]);

  const handleLogin = async (userData) => {
    try {
      const response = await login(userData);
      const token = response.data.token;
      localStorage.setItem('token', token);
      axiosInstance.defaults.headers['Authorization'] = `Bearer ${token}`;
      const userResponse = await axiosInstance.get('/profile');
      setUser(userResponse.data);
      localStorage.setItem('user', JSON.stringify(userResponse.data));
      navigate('/');
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  const handleRegister = async (userData) => {
    try {
      await register(userData);
      navigate('/login');
    } catch (error) {
      console.error('Registration failed:', error);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setUser(null);
    navigate('/login');
  };

  const handleChangePassword = async (passwordData) => {
    try {
      await changePassword(passwordData);
      handleLogout();
    } catch (error) {
      console.error('Change password failed:', error);
    }
  };

  if (loading) {
    return <div>Loading...</div>; // Show loading state while fetching data
  }

  return (
    <AuthContext.Provider value={{ user, handleLogin, handleRegister, handleLogout, handleChangePassword }}>
      {children}
    </AuthContext.Provider>
  );
};
