"use client";
import { useState, useEffect, Suspense } from 'react';
import useNilaiStore from '@/stores/nilai';
import useMahasiswaStore from '@/stores/mahasiswa';
import useMataKuliahStore from '@/stores/matakuliah';
import useAuthStore from '@/stores/authStore'; 
import { useRouter } from 'next/navigation';

const CreateNilai = () => {
  const { createNilai } = useNilaiStore((state) => ({
    createNilai: state.createNilai,
  }));
  const { mahasiswaList, fetchMahasiswa } = useMahasiswaStore((state) => ({
    mahasiswaList: state.mahasiswaList,
    fetchMahasiswa: state.fetchMahasiswa,
  }));
  const { mataKuliahList, fetchMataKuliah } = useMataKuliahStore((state) => ({
    mataKuliahList: state.mataKuliahList,
    fetchMataKuliah: state.fetchMataKuliah,
  }));
  const { user } = useAuthStore(); 

  const [nilaiData, setNilaiData] = useState({
    indeks: '',
    skor: '',
    mahasiswa_id: '',
    mata_kuliah_id: '',
    users_id: user?.id || '', 
  });

  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    const fetchData = async () => {
      await fetchMahasiswa();
      await fetchMataKuliah();
      setIsLoading(false);
    };
    fetchData();
  }, [fetchMahasiswa, fetchMataKuliah]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const dataToSend = {
      indeks: nilaiData.indeks,
      skor: parseInt(nilaiData.skor, 10),
      mahasiswa_id: parseInt(nilaiData.mahasiswa_id, 10),
      mata_kuliah_id: parseInt(nilaiData.mata_kuliah_id, 10),
      users_id: user?.id || '', 
    };

    await createNilai(dataToSend);
    router.push('/nilai');
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6">
      <h1 className="mb-4 text-xl font-semibold">Create Nilai</h1>
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
          {mahasiswaList.map((mahasiswa) => {
            if (!mahasiswa.ID) {
              console.error('Mahasiswa item missing ID:', mahasiswa);
              return null; 
            }
            return (
              <option key={mahasiswa.ID} value={mahasiswa.ID}>
                {mahasiswa.Nama}
              </option>
            );
          })}
        </select>

        <select
          name="mata_kuliah_id"
          value={nilaiData.mata_kuliah_id}
          onChange={(e) => setNilaiData({ ...nilaiData, mata_kuliah_id: e.target.value })}
          className="w-full px-4 py-2 text-black border border-gray-300 rounded"
          required
        >
          <option value="" disabled>Select Mata Kuliah</option>
          {mataKuliahList.map((mataKuliah) => {
            if (!mataKuliah.ID) {
              console.error('Mata Kuliah item missing ID:', mataKuliah);
              return null; 
            }
            return (
              <option key={mataKuliah.ID} value={mataKuliah.ID}>
                {mataKuliah.Nama}
              </option>
            );
          })}
        </select>

        {/* Removed users_id input field */}
        <button
          type="submit"
          className="px-4 py-2 text-white bg-blue-500 rounded"
        >
          Create
        </button>
      </form>
    </div>
  );
};

const CreateNilaiPage = () => {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <CreateNilai />
    </Suspense>
  );
};

export default CreateNilaiPage;
