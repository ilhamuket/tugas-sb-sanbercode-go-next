import PropTypes from 'prop-types';
import { useContext } from 'react';
import { AuthContext } from '../../context/AuthContext';
import { useNavigate } from 'react-router-dom';
import { FaUser } from 'react-icons/fa'; 

const NewsLayout = ({ children }) => {
  const { user, handleLogout } = useContext(AuthContext);
  const navigate = useNavigate();

  const handleLogoutClick = () => {
    handleLogout();
    navigate('/login');
  };

  return (
    <div className="flex flex-col min-h-screen bg-gray-100">
      <header className="bg-blue-600 text-white py-4">
        <div className="container mx-auto px-4 flex items-center justify-between">
          <h1 className="text-3xl font-bold">News App</h1>
          {user ? (
            <div className="flex items-center">
              <span className="text-lg mr-4">
                <FaUser className="inline-block mr-2" /> Welcome, {user.username || 'User'}
              </span>
              <button
                onClick={handleLogoutClick}
                className="btn btn-primary"
              >
                Logout
              </button>
            </div>
          ) : (
            <button
              onClick={() => navigate('/login')}
              className="btn btn-primary"
            >
              Login
            </button>
          )}
        </div>
      </header>
      <main className="flex-1 container mx-auto px-4 py-8">
        {children}
      </main>
      <footer className="bg-blue-600 text-white py-4">
        <div className="container mx-auto px-4 text-center">
          <p>&copy; 2024 News App. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
};

NewsLayout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default NewsLayout;
