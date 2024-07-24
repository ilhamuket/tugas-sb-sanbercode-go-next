import PropTypes from 'prop-types';
import { useState, useEffect } from 'react';

const UserForm = ({ user = null, onSave }) => {
  const [formData, setFormData] = useState({
    username: '',
    email: '',
    role: 'admin', // Default role
    password: '', // Menambahkan password
  });

  useEffect(() => {
    if (user) {
      setFormData({
        username: user.username || '',
        email: user.email || '',
        role: user.roles?.length > 0 ? user.roles[0].name : 'admin', // Default ke 'admin' jika tidak ada role
        password: '', // Kosongkan password saat edit
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
      // Reset form data after successful save
      setFormData({
        username: '',
        email: '',
        role: 'admin', // Reset role ke default
        password: '', // Kosongkan password
      });
    } catch (error) {
      console.error("Error saving user:", error);
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mb-4 p-4 border rounded-lg shadow-sm bg-white">
      <h2 className="text-xl font-semibold mb-4">{user ? 'Edit User' : 'Add User'}</h2>
      <div className="mb-2">
        <label className="block text-sm font-medium mb-1">Username</label>
        <input
          type="text"
          name="username"
          value={formData.username}
          onChange={handleChange}
          placeholder="Enter username"
          className="input input-bordered w-full"
        />
      </div>
      <div className="mb-2">
        <label className="block text-sm font-medium mb-1">Email</label>
        <input
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
          placeholder="Enter email"
          className="input input-bordered w-full"
        />
      </div>
      {!user && (
        <div className="mb-2">
          <label className="block text-sm font-medium mb-1">Password</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            placeholder="Enter password"
            className="input input-bordered w-full"
          />
        </div>
      )}
      <div className="mb-2">
        <label className="block text-sm font-medium mb-1">Role</label>
        <select
          name="role"
          value={formData.role}
          onChange={handleChange}
          className="input input-bordered w-full"
        >
          <option value="admin">Admin</option>
          <option value="editor">Editor</option>
        </select>
      </div>
      <button type="submit" className="btn btn-primary mt-4">{user ? 'Update' : 'Add'}</button>
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
