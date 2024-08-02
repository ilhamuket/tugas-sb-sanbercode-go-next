"use client";
import { useEffect } from 'react';
import useNilaiStore from '@/stores/nilai';
import { useRouter } from 'next/router';
import Link from 'next/link';

const NilaiDetail = () => {
  const { getNilaiById, selectedNilai } = useNilaiStore((state) => ({
    getNilaiById: state.getNilaiById,
    selectedNilai: state.selectedNilai,
  }));
  const router = useRouter();
  const { id } = router.query;

  useEffect(() => {
    if (id) {
      getNilaiById(id);
    }
  }, [id, getNilaiById]);

  return (
    <div className="p-6">
      <h1 className="mb-4 text-xl font-semibold">Nilai Detail</h1>
      {selectedNilai ? (
        <div className="space-y-4">
          <p><strong>ID:</strong> {selectedNilai.ID}</p>
          <p><strong>Indeks:</strong> {selectedNilai.Indeks}</p>
          <p><strong>Skor:</strong> {selectedNilai.Skor}</p>
          <p><strong>Mahasiswa ID:</strong> {selectedNilai.MahasiswaID}</p>
          <p><strong>Mata Kuliah ID:</strong> {selectedNilai.MataKuliahID}</p>
          <p><strong>User ID:</strong> {selectedNilai.UsersID}</p>
          <Link href={`/nilai/${selectedNilai.ID}/edit`} className="text-blue-500 hover:underline">
            Edit
          </Link>
        </div>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default NilaiDetail;
