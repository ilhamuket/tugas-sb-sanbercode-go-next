import create from 'zustand';
import { persist } from 'zustand/middleware';
import axios from 'axios';
import { baseUrl } from "@/utils/constant";
import useAuthStore from '@/stores/authStore';  

const useMahasiswaStore = create(persist(
  (set) => ({
    mahasiswaList: [],
    selectedMahasiswa: null,
    error: null,

    fetchMahasiswa: async () => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.get(`${baseUrl}/mahasiswa`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set({ mahasiswaList: response.data, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    getMahasiswaById: async (id) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.get(`${baseUrl}/mahasiswa/${id}`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set({ selectedMahasiswa: response.data, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    createMahasiswa: async (mahasiswaData) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.post(`${baseUrl}/mahasiswa`, mahasiswaData, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          mahasiswaList: [...state.mahasiswaList, response.data],
          selectedMahasiswa: response.data,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    updateMahasiswa: async (id, mahasiswaData) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.patch(`${baseUrl}/mahasiswa/${id}`, mahasiswaData, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          mahasiswaList: state.mahasiswaList.map((mahasiswa) =>
            mahasiswa.ID === id ? response.data : mahasiswa
          ),
          selectedMahasiswa: response.data,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    deleteMahasiswa: async (id) => {
      const { token } = useAuthStore.getState();
      try {
        await axios.delete(`${baseUrl}/mahasiswa/${id}`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          mahasiswaList: state.mahasiswaList.filter((mahasiswa) => mahasiswa.ID !== id),
          selectedMahasiswa: null,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    clearError: () => set({ error: null }),
  }),
  {
    name: 'mahasiswa-storage',
    getStorage: () => localStorage,
  }
));

export default useMahasiswaStore;
