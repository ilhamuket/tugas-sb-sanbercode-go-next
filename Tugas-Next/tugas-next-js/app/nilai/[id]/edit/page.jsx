"use client";
import { useEffect, useState } from 'react';
import useNilaiStore from '@/stores/nilai';
import useMahasiswaStore from '@/stores/mahasiswa';
import useMataKuliahStore from '@/stores/matakuliah';
import useAuthStore from '@/stores/authStore'; 
import { useParams, useRouter } from 'next/navigation';

const EditNilai = () => {
  const { getNilaiById, updateNilai, selectedNilai } = useNilaiStore((state) => ({
    getNilaiById: state.getNilaiById,
    updateNilai: state.updateNilai,
    selectedNilai: state.selectedNilai,
  }));
  const { mahasiswaList, fetchMahasiswa } = useMahasiswaStore((state) => ({
    mahasiswaList: state.mahasiswaList,
    fetchMahasiswa: state.fetchMahasiswa,
  }));
  const { mataKuliahList, fetchMataKuliah } = useMataKuliahStore((state) => ({
    mataKuliahList: state.mataKuliahList,
    fetchMataKuliah: state.fetchMataKuliah,
  }));
  const { user } = useAuthStore((state) => ({ user: state.user })); 

  const [nilaiData, setNilaiData] = useState({
    indeks: '',
    skor: '',
    mahasiswa_id: '',
    mata_kuliah_id: '',
    users_id: '', 
  });

  const [isLoading, setIsLoading] = useState(true);
  const { id } = useParams();
  const router = useRouter();

  useEffect(() => {
    const fetchData = async () => {
      await fetchMahasiswa();
      await fetchMataKuliah();
      if (id) {
        getNilaiById(id);
      }
    };
    fetchData();
  }, [id, fetchMahasiswa, fetchMataKuliah, getNilaiById]);

  useEffect(() => {
    if (selectedNilai) {
      setNilaiData({
        indeks: selectedNilai.Indeks,
        skor: selectedNilai.Skor,
        mahasiswa_id: selectedNilai.MahasiswaID,
        mata_kuliah_id: selectedNilai.MataKuliahID,
        users_id: user?.ID || '', 
      });
      setIsLoading(false);
    }
  }, [selectedNilai, user]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (id) {
      const dataToSend = {
        indeks: nilaiData.indeks,
        skor: parseInt(nilaiData.skor, 10),
        mahasiswa_id: parseInt(nilaiData.mahasiswa_id, 10),
        mata_kuliah_id: parseInt(nilaiData.mata_kuliah_id, 10),
        users_id: parseInt(nilaiData.users_id, 10), 
      };
      await updateNilai(id, dataToSend);
      router.push('/nilai');
    }
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6">
      <h1 className="mb-4 text-xl font-semibold">Edit Nilai</h1>
      <form onSubmit={handleSubmit} className="space-y-4">
        <select
          name="indeks"
          value={nilaiData.indeks}
          onChange={(e) => setNilaiData({ ...nilaiData, indeks: e.target.value })}
          className="w-full px-4 py-2 text-black border border-gray-300 rounded"
          required
        >
          <option value="" disabled>Select Indeks</option>
          <option value="A">A</option>
          <option value="B">B</option>
          <option value="C">C</option>
          <option value="D">D</option>
          <option value="E">E</option>
        </select>
        <input
          type="number"
          name="skor"
          value={nilaiData.skor}
          onChange={(e) => setNilaiData({ ...nilaiData, skor: e.target.value })}
          placeholder="Skor"
          className="w-full px-4 py-2 border border-gray-300 rounded"
          required
        />
        <select
          name="mahasiswa_id"
          value={nilaiData.mahasiswa_id}
          onChange={(e) => setNilaiData({ ...nilaiData, mahasiswa_id: e.target.value })}
          className="w-full px-4 py-2 text-black border border-gray-300 rounded"
          required
        >
          <option value="" disabled>Select Mahasiswa</option>
          {mahasiswaList.map((mahasiswa) => (
            <option key={mahasiswa.ID} value={mahasiswa.ID}>
              {mahasiswa.Nama}
            </option>
          ))}
        </select>

        <select
          name="mata_kuliah_id"
          value={nilaiData.mata_kuliah_id}
          onChange={(e) => setNilaiData({ ...nilaiData, mata_kuliah_id: e.target.value })}
          className="w-full px-4 py-2 text-black border border-gray-300 rounded"
          required
        >
          <option value="" disabled>Select Mata Kuliah</option>
          {mataKuliahList.map((mataKuliah) => (
            <option key={mataKuliah.ID} value={mataKuliah.ID}>
              {mataKuliah.Nama}
            </option>
          ))}
        </select>

        {/* Hapus input users_id */}
        <button
          type="submit"
          className="px-4 py-2 text-white bg-blue-500 rounded"
        >
          Update
        </button>
      </form>
    </div>
  );
};

export default EditNilai;
