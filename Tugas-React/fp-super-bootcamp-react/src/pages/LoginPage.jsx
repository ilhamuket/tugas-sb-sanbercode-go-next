import { useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import Login from '../components/Auth/Login';

const LoginPage = () => {
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
      {!user && <Login />}
    </div>
  );
};

export default LoginPage;
