"use client";
import { useEffect, useState } from 'react';
import useNilaiStore from '@/stores/nilai';
import Link from 'next/link';

const NilaiList = () => {
  const { nilaiList, fetchNilai, deleteNilai } = useNilaiStore((state) => ({
    nilaiList: state.nilaiList,
    fetchNilai: state.fetchNilai,
    deleteNilai: state.deleteNilai,
  }));
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      await fetchNilai();
      setLoading(false);
    };
    fetchData();
  }, [fetchNilai]);

  const handleDelete = async (id) => {
    if (confirm('Are you sure you want to delete this item?')) {
      await deleteNilai(id);
    }
  };

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className="p-6">
      <h1 className="mb-4 text-2xl font-semibold">Nilai List</h1>
      <Link href="/nilai/create" className="inline-block px-4 py-2 mb-4 text-white bg-blue-500 rounded hover:bg-blue-600">
        Add New Nilai
      </Link>
      <table className="min-w-full bg-white border border-gray-200 rounded-lg shadow-md">
        <thead className="bg-gray-100">
          <tr>
            <th className="px-6 py-3 text-left text-gray-600 border-b">No</th>
            <th className="px-6 py-3 text-left text-gray-600 border-b">Indeks</th>
            <th className="px-6 py-3 text-left text-gray-600 border-b">Skor</th>
            <th className="px-6 py-3 text-left text-gray-600 border-b">Mahasiswa</th>
            <th className="px-6 py-3 text-left text-gray-600 border-b">Mata Kuliah</th>
            <th className="px-6 py-3 text-left text-gray-600 border-b">User</th>
            <th className="px-6 py-3 text-left text-gray-600 border-b">Actions</th>
          </tr>
        </thead>
        <tbody>
          {nilaiList.map((nilai, index) => (
            <tr key={nilai.ID}>
              <td className="px-6 py-4 text-gray-800 border-b">{index + 1}</td>
              <td className="px-6 py-4 text-gray-800 border-b">{nilai.Indeks}</td>
              <td className="px-6 py-4 text-gray-800 border-b">{nilai.Skor}</td>
              <td className="px-6 py-4 text-gray-800 border-b">{nilai.Mahasiswa?.Nama || 'N/A'}</td>
              <td className="px-6 py-4 text-gray-800 border-b">{nilai.MataKuliah?.Nama || 'N/A'}</td>
              <td className="px-6 py-4 text-gray-800 border-b">{nilai.User?.name || 'N/A'}</td>
              <td className="px-6 py-4 border-b">
                <Link href={`/nilai/${nilai.ID}/edit`} className="mr-4 text-blue-500 hover:underline">
                  Edit
                </Link>
                <button
                  onClick={() => handleDelete(nilai.ID)}
                  className="text-red-500 hover:underline"
                >
                  Delete
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default NilaiList;
