import { useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import Register from '../components/Auth/Register';

const RegisterPage = () => {
  const { user } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (user) {
      // Jika pengguna sudah login, arahkan ke halaman utama
      navigate('/');
    }
  }, [user, navigate]);

  return (
    <div>
      {/* Hanya tampilkan halaman login jika pengguna belum login */}
      {!user && <Register />}
    </div>
  );
};

export default RegisterPage;
