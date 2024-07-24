import { Link, useLocation } from 'react-router-dom';
import { FaUser, FaNewspaper } from 'react-icons/fa';

const Sidebar = () => {
  const location = useLocation();

  return (
    <div className="w-64 bg-gray-800 text-white h-auto p-4 flex flex-col">
      <h2 className="text-2xl mb-6">Dashboard</h2>
      <ul className="flex-1">
        <li className={`mb-1 ${location.pathname === '/users' ? 'bg-gray-600' : ''}`}>
          <Link to="/users" className="flex items-center p-2 text-white hover:bg-gray-600">
            <FaUser className="mr-2" /> User Management
          </Link>
        </li>
        <li className={`mb-1 ${location.pathname === '/news' ? 'bg-gray-600' : ''}`}>
          <Link to="/news" className="flex items-center p-2 text-white hover:bg-gray-600">
            <FaNewspaper className="mr-2" /> News
          </Link>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
