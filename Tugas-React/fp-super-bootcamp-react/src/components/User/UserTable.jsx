import PropTypes from 'prop-types';

const UserTable = ({ users, onEdit, onDelete }) => {
  return (
    <div className="overflow-x-auto">
      <table className="table-auto w-full border-collapse border border-gray-300">
        <thead>
          <tr className="bg-gray-200">
            <th className="border border-gray-300 px-4 py-2">ID</th>
            <th className="border border-gray-300 px-4 py-2">Username</th>
            <th className="border border-gray-300 px-4 py-2">Email</th>
            <th className="border border-gray-300 px-4 py-2">Role</th>
            <th className="border border-gray-300 px-4 py-2">Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id} className="hover:bg-gray-100">
              <td className="border border-gray-300 px-4 py-2">{user.id}</td>
              <td className="border border-gray-300 px-4 py-2">{user.username || 'N/A'}</td>
              <td className="border border-gray-300 px-4 py-2">{user.email || 'N/A'}</td>
              <td className="border border-gray-300 px-4 py-2">
                {user.roles?.length > 0 ? user.roles[0].name : 'N/A'}
              </td>
              <td className="border border-gray-300 px-4 py-2">
                <button 
                  onClick={() => onEdit(user.id)} 
                  className="btn btn-primary mr-2 px-4 py-2 bg-blue-500 text-white rounded"
                >
                  Edit
                </button>
                <button 
                  onClick={() => onDelete(user.id)} 
                  className="btn btn-error px-4 py-2 bg-red-500 text-white rounded"
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

UserTable.propTypes = {
  users: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.number.isRequired,
      username: PropTypes.string.isRequired,
      email: PropTypes.string.isRequired,
      roles: PropTypes.arrayOf(
        PropTypes.shape({
          name: PropTypes.string.isRequired,
        })
      ).isRequired,
    })
  ).isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};

export default UserTable;
