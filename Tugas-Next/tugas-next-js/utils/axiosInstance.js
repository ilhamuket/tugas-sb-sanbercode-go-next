import axios from 'axios';
import useAuthStore from '@/stores/authStore'; 

const axiosInstance = axios.create({
  baseURL: baseUrl, 
});

axiosInstance.interceptors.request.use(
  (config) => {
    const { token } = useAuthStore.getState();
    
    if (config.method !== 'get' && token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default axiosInstance;
