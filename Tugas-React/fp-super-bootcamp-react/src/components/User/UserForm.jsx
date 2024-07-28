import PropTypes from 'prop-types';
import { useState, useEffect } from 'react';
import { FaEye, FaEyeSlash } from 'react-icons/fa';

const UserForm = ({ user = null, onSave }) => {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    role: 'admin',
    password: '',
  });

  const [showPassword, setShowPassword] = useState(false);

  useEffect(() => {
    if (user) {
      setFormData({
        username: user.username || '',
        email: user.email || '',
        role: user.roles?.length > 0 ? user.roles[0].name : 'admin',
        password: '',
      });
    }
  }, [user]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await onSave(formData);

      setFormData({
        username: '',
        email: '',
        role: 'admin',
        password: '',
      });
    } catch (error) {
      console.error("Error saving user:", error);
    }
  };

  const togglePasswordVisibility = () => {
    setShowPassword(!showPassword);
  };

  return (
    <form onSubmit={handleSubmit} className="p-4 mb-4 bg-white border rounded-lg shadow-sm">
      <h2 className="mb-4 text-xl font-semibold">{user ? 'Edit User' : 'Add User'}</h2>
      <div className="mb-2">
        <label className="block mb-1 text-sm font-medium">Username</label>
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="Enter username"
          className="w-full input input-bordered"
        />
      </div>
      <div className="mb-2">
        <label className="block mb-1 text-sm font-medium">Email</label>
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          placeholder="Enter email"
          className="w-full input input-bordered"
        />
      </div>
      {!user && (
        <div className="relative mb-2">
          <label className="block mb-1 text-sm font-medium">Password</label>
          <input
            type={showPassword ? 'text' : 'password'}
            name="password"
            value={formData.password}
            onChange={handleChange}
            placeholder="Enter password"
            className="w-full pr-10 input input-bordered"
          />
          <button
            type="button"
            onClick={togglePasswordVisibility}
            className="absolute inset-y-0 right-0 flex items-center pr-3 text-gray-700"
          >
            {showPassword ? <FaEyeSlash /> : <FaEye />}
          </button>
        </div>
      )}
      <div className="mb-2">
        <label className="block mb-1 text-sm font-medium">Role</label>
        <select
          name="role"
          value={formData.role}
          onChange={handleChange}
          className="w-full input input-bordered"
        >
          <option value="admin">Admin</option>
          <option value="editor">Editor</option>
        </select>
      </div>
      <button type="submit" className="mt-4 btn btn-primary">{user ? 'Update' : 'Add'}</button>
    </form>
  );
};

UserForm.propTypes = {
  user: PropTypes.shape({
    username: PropTypes.string,
    email: PropTypes.string,
    roles: PropTypes.arrayOf(
      PropTypes.shape({
        name: PropTypes.string.isRequired,
      })
    ),
  }),
  onSave: PropTypes.func.isRequired,
};

export default UserForm;
