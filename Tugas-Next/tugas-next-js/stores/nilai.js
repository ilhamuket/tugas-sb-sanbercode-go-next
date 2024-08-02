import create from 'zustand';
import { persist } from 'zustand/middleware';
import axios from 'axios';
import { baseUrl } from "@/utils/constant";
import useAuthStore from '@/stores/authStore';  

const useNilaiStore = create(persist(
  (set) => ({
    nilaiList: [],
    selectedNilai: null,
    error: null,

    fetchNilai: async () => {
      const { token } = useAuthStore.getState();  
      try {
        const response = await axios.get(`${baseUrl}/nilai`, {
          headers: { Authorization: `Bearer ${token}` }  
        });
        set({ nilaiList: response.data, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    getNilaiById: async (id) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.get(`${baseUrl}/nilai/${id}`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set({ selectedNilai: response.data, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    createNilai: async (nilaiData) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.post(`${baseUrl}/nilai`, nilaiData, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          nilaiList: [...state.nilaiList, response.data],
          selectedNilai: response.data,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    updateNilai: async (id, nilaiData) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.patch(`${baseUrl}/nilai/${id}`, nilaiData, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          nilaiList: state.nilaiList.map((nilai) =>
            nilai.ID === id ? response.data : nilai
          ),
          selectedNilai: response.data,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    deleteNilai: async (id) => {
      const { token } = useAuthStore.getState();
      try {
        await axios.delete(`${baseUrl}/nilai/${id}`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          nilaiList: state.nilaiList.filter((nilai) => nilai.ID !== id),
          selectedNilai: null,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    clearError: () => set({ error: null }),
  }),
  {
    name: 'nilai-storage',
    getStorage: () => localStorage,
  }
));

export default useNilaiStore;
