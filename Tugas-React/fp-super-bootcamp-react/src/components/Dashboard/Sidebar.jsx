import { Link, useLocation } from 'react-router-dom';
import { FaUser, FaNewspaper } from 'react-icons/fa';

const Sidebar = () => {
  const location = useLocation();

  return (
    <div className="flex flex-col w-64 h-auto p-4 text-white bg-gray-800">
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
