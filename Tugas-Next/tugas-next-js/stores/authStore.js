import create from 'zustand';
import { persist } from 'zustand/middleware';
import axios from 'axios';
import { baseUrl } from "@/utils/constant";

const useAuthStore = create(persist(
  (set) => ({
    user: null,
    token: null,
    error: null,

    login: async (email, password) => {
      try {
        const response = await axios.post(`${baseUrl}/login`, {
          email,
          password
        });

        if (response.status !== 200) {
          set({ error: 'Login failed' });
          throw new Error('Login failed');
        }

        const data = response.data;
        set({ user: data.user, token: data.token, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    register: async (name, email, password) => {
      try {
        const response = await axios.post(`${baseUrl}/register`, {
          name,
          email,
          password
        });

        if (response.status !== 201) {
          set({ error: 'Registration failed' });
          throw new Error('Registration failed');
        }

        const data = response.data;
        set({ user: data.user, token: data.token, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    logout: () => set({ user: null, token: null }),

    clearError: () => set({ error: null }),
  }),
  {
    name: 'auth-storage', 
    getStorage: () => localStorage, 
  }
));

export default useAuthStore;
