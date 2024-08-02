import create from 'zustand';
import axios from 'axios';
import { baseUrl } from '@/utils/constant';

const useNilaiStore = create((set) => ({
  nilai: [],
  selectedNilai: null,
  error: null,

  fetchNilai: async () => {
    try {
      const response = await axios.get(`${baseUrl}/nilai`);
      set({ nilai: response.data, error: null });
    } catch (error) {
      set({ error: error.response?.data?.error || 'Failed to fetch Nilai' });
    }
  },

  fetchNilaiById: async (id) => {
    try {
      const response = await axios.get(`${baseUrl}/nilai/${id}`);
      set({ selectedNilai: response.data, error: null });
    } catch (error) {
      set({ error: error.response?.data?.error || 'Failed to fetch Nilai' });
    }
  },

  createNilai: async (newNilai) => {
    try {
      const response = await axios.post(`${baseUrl}/nilai`, newNilai);
      set((state) => ({
        nilai: [...state.nilai, response.data],
        error: null,
      }));
    } catch (error) {
      set({ error: error.response?.data?.error || 'Failed to create Nilai' });
    }
  },

  updateNilai: async (id, updatedNilai) => {
    try {
      const response = await axios.patch(`${baseUrl}/nilai/${id}`, updatedNilai);
      set((state) => ({
        nilai: state.nilai.map((item) => (item.id === id ? response.data : item)),
        selectedNilai: response.data,
        error: null,
      }));
    } catch (error) {
      set({ error: error.response?.data?.error || 'Failed to update Nilai' });
    }
  },

  deleteNilai: async (id) => {
    try {
      await axios.delete(`${baseUrl}/nilai/${id}`);
      set((state) => ({
        nilai: state.nilai.filter((item) => item.id !== id),
        selectedNilai: null,
        error: null,
      }));
    } catch (error) {
      set({ error: error.response?.data?.error || 'Failed to delete Nilai' });
    }
  },

  clearError: () => set({ error: null }),
}));

export default useNilaiStore;
