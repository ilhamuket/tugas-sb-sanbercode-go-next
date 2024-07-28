import axios from './axiosConfig';

// Mendapatkan semua pengguna
export const getAllUsers = async () => {
  const response = await axios.get('/users');
  return response.data;
};

// Mendapatkan pengguna berdasarkan ID
export const getUserById = async (id) => {
  const response = await axios.get(`/users/${id}`);
  return response.data;
};

// Membuat pengguna baru
export const createUser = async (userData) => {
  const response = await axios.post('/users', userData);
  return response.data;
};

// Memperbarui data pengguna berdasarkan ID
export const updateUser = async (id, userData) => {
  const response = await axios.put(`/users/${id}`, userData);
  return response.data;
};

// Menghapus pengguna berdasarkan ID
export const deleteUser = async (id) => {
  const response = await axios.delete(`/users/${id}`);
  return response.data;
};

// Mendapatkan profil pengguna yang sedang login
export const getProfile = async () => {
  const response = await axios.get('/profile');
  return response.data;
};

// Memperbarui profil pengguna yang sedang login
export const updateProfile = async (profileData) => {
  const response = await axios.put('/profile', profileData);
  return response.data;
};
