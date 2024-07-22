import axiosInstance from './axiosConfig';

export const register = (userData) => {
  return axiosInstance.post('/register', userData);
};

export const login = (userData) => {
  return axiosInstance.post('/login', userData);
};

export const changePassword = (passwordData) => {
  return axiosInstance.put('/change-password', passwordData);
};
