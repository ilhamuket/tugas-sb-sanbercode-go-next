import { useContext, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';

const Home = () => {
  const { user, handleLogout } = useContext(AuthContext);
  const navigate = useNavigate();

  useEffect(() => {
    if (user === null) {
      navigate('/login');
    }
  }, [user, navigate]);

  const handleLogoutClick = () => {
    handleLogout(); // Memanggil fungsi logout dari AuthContext
    navigate('/login'); // Redirect ke halaman login setelah logout
  };

  return (
    <div>
      <h1>Welcome to the Home Page</h1>
      {user ? (
        <>
          <p>Welcome back, {user.bio || 'User'}!</p>
          <button onClick={handleLogoutClick}>Logout</button>
        </>
      ) : (
        <p>Loading...</p>
      )}
      {/* Content of the home page */}
    </div>
  );
};

export default Home;
