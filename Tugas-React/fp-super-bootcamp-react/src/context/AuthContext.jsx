import { createContext, useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import axiosInstance from '../api/axiosConfig';
import { login, register, changePassword } from '../api/auth';
import Swal from 'sweetalert2';

export const AuthContext = createContext();

// eslint-disable-next-line react/prop-types
export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null); 
  const [loading, setLoading] = useState(true); 
  const navigate = useNavigate();
  const location = useLocation();

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
    console.error('Error:', error);
  };

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
            handleError(error);
            localStorage.removeItem('token');
            localStorage.removeItem('user');
            if (location.pathname !== '/login' && location.pathname !== '/register') {
              navigate('/login');
            }
          } finally {
            setLoading(false);
          }
        };
        fetchUserDetails();
      }
    } else {
      setLoading(false);
      if (location.pathname !== '/login' && location.pathname !== '/register' && location.pathname !== '/' && !location.pathname.startsWith('/news/')) {
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
      handleError(error);
    }
  };

  const handleRegister = async (userData) => {
    try {
      await register(userData);
      navigate('/login');
    } catch (error) {
      handleError(error);
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
      Swal.fire({
        icon: 'success',
        title: 'Password Changed',
        text: 'Your password has been successfully changed.',
      }).then(() => {
        handleLogout();
      });
    } catch (error) {
      handleError(error);
    }
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="w-16 h-16 border-b-4 border-blue-600 rounded-full animate-spin"></div>
      </div>
    );
  }

  return (
    <AuthContext.Provider value={{ user, handleLogin, handleRegister, handleLogout, handleChangePassword }}>
      {children}
    </AuthContext.Provider>
  );
};
