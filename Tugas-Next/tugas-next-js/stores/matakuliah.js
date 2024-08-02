import create from 'zustand';
import { persist } from 'zustand/middleware';
import axios from 'axios';
import { baseUrl } from "@/utils/constant";
import useAuthStore from '@/stores/authStore';  

const useMataKuliahStore = create(persist(
  (set) => ({
    mataKuliahList: [],
    selectedMataKuliah: null,
    error: null,

    fetchMataKuliah: async () => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.get(`${baseUrl}/mata-kuliah`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set({ mataKuliahList: response.data, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    getMataKuliahById: async (id) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.get(`${baseUrl}/mata-kuliah/${id}`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set({ selectedMataKuliah: response.data, error: null });
      } catch (err) {
        set({ error: err.message });
      }
    },

    createMataKuliah: async (mataKuliahData) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.post(`${baseUrl}/mata-kuliah`, mataKuliahData, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          mataKuliahList: [...state.mataKuliahList, response.data],
          selectedMataKuliah: response.data,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    updateMataKuliah: async (id, mataKuliahData) => {
      const { token } = useAuthStore.getState();
      try {
        const response = await axios.patch(`${baseUrl}/mata-kuliah/${id}`, mataKuliahData, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          mataKuliahList: state.mataKuliahList.map((mataKuliah) =>
            mataKuliah.ID === id ? response.data : mataKuliah
          ),
          selectedMataKuliah: response.data,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    deleteMataKuliah: async (id) => {
      const { token } = useAuthStore.getState();
      try {
        await axios.delete(`${baseUrl}/mata-kuliah/${id}`, {
          headers: { Authorization: `Bearer ${token}` }
        });
        set((state) => ({
          mataKuliahList: state.mataKuliahList.filter((mataKuliah) => mataKuliah.ID !== id),
          selectedMataKuliah: null,
          error: null,
        }));
      } catch (err) {
        set({ error: err.message });
      }
    },

    clearError: () => set({ error: null }),
  }),
  {
    name: 'mata-kuliah-storage',
    getStorage: () => localStorage,
  }
));

export default useMataKuliahStore;
