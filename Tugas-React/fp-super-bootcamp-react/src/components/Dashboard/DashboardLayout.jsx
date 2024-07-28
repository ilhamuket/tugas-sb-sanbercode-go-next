import { useContext, useState } from 'react';
import PropTypes from 'prop-types';
import { AuthContext } from '../../context/AuthContext';
import Sidebar from './Sidebar';
import ChangePassword from '../Auth/ChangePassword';

const DashboardLayout = ({ children }) => {
  const { user, handleLogout } = useContext(AuthContext);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleModalOpen = () => setIsModalOpen(true);
  const handleModalClose = () => setIsModalOpen(false);

  return (
    <div className="flex flex-col min-h-screen">
      <nav className="shadow-lg bg-base-100">
        <div className="container flex items-center justify-between p-4 mx-auto">
          <a href="/" className="text-xl font-bold">Home</a>
          <div className="flex items-center space-x-4">
            <span>{user?.username}</span>
            <div className="dropdown dropdown-end">
              <label tabIndex={0} className="btn btn-ghost btn-circle avatar">
                <div className="w-10 rounded-full">
                  <img src="https://static.vecteezy.com/system/resources/thumbnails/000/439/863/small/Basic_Ui__28186_29.jpg" alt="avatar" />
                </div>
              </label>
              <ul tabIndex={0} className="p-2 mt-3 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                <li>
                  <button onClick={handleModalOpen}>Change Password</button>
                </li>
                <li>
                  <button onClick={handleLogout}>Logout</button>
                </li>
              </ul>
            </div>
          </div>
        </div>
      </nav>
      <div className="flex flex-1 h-auto">
        <Sidebar />
        <div className="flex-1 p-6">
          {children}
        </div>
      </div>

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

DashboardLayout.propTypes = {
  children: PropTypes.node.isRequired,
};

export default DashboardLayout;
