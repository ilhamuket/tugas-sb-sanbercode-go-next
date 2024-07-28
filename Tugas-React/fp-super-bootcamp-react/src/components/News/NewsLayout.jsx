import PropTypes from 'prop-types';
import { useContext, useState } from 'react';
import { AuthContext } from '../../context/AuthContext';
import { useNavigate } from 'react-router-dom';
import { FaUser } from 'react-icons/fa';
import ChangePassword from '../Auth/ChangePassword';

const NewsLayout = ({ children }) => {
  const { user, handleLogout } = useContext(AuthContext);
  const navigate = useNavigate();
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleLogoutClick = () => {
    handleLogout();
    navigate('/login');
  };

  const handleModalOpen = () => setIsModalOpen(true);
  const handleModalClose = () => setIsModalOpen(false);

  return (
    <div className="flex flex-col min-h-screen bg-gray-100">
      <header className="py-4 text-white bg-blue-600">
        <div className="container flex items-center justify-between px-4 mx-auto">
          <h1 className="text-3xl font-bold">News App</h1>
          {user ? (
            <div className="relative flex items-center">
              <details className="text-black dropdown">
                <summary className="m-1 btn btn-primary">
                  <FaUser className="inline-block mr-2" /> Welcome, {user.username || 'User'}
                </summary>
                <ul className="menu dropdown-content bg-base-100 rounded-box z-[1] w-52 p-2 shadow">
                  <li>
                    <button onClick={() => navigate('/news')}>
                      Dashboard
                    </button>
                  </li>
                  <li>
                    <button onClick={handleModalOpen}>
                      Change Password
                    </button>
                  </li>
                  <li>
                    <button onClick={handleLogoutClick}>
                      Logout
                    </button>
                  </li>
                </ul>
              </details>
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
      <main className="container flex-1 px-4 py-8 mx-auto">
        {children}
      </main>
      <footer className="py-4 text-white bg-blue-600">
        <div className="container px-4 mx-auto text-center">
          <p>&copy; 2024 News App. All rights reserved.</p>
        </div>
      </footer>

      {isModalOpen && (
        <div className="modal modal-open">
          <div className="modal-box">
            <h3 className="text-lg font-bold">Change Password</h3>
            <ChangePassword />
            <div className="modal-action">
              <button className="btn" onClick={handleModalClose}>Close</button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

NewsLayout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default NewsLayout;
