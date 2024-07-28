import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../context/UserContext';
import UserTable from '../components/User/UserTable';
import UserForm from '../components/User/UserForm';
import DashboardLayout from '../components/Dashboard/DashboardLayout';

const UserManagementPage = () => {
  const { users, fetchUserById, addUser, updateUser, removeUser, loading, error } = useContext(UserContext);
  const [editingUser, setEditingUser] = useState(null);

  useEffect(() => {
    // Fetch users if needed on component mount
  }, []);

  const handleEdit = async (userId) => {
    const user = await fetchUserById(userId);
    setEditingUser(user);
  };

  const handleSave = async (userData) => {
    if (editingUser) {
      // Update existing user
      await updateUser(editingUser.id, userData);
      setEditingUser(null);
    } else {
      // Add new user
      await addUser(userData);
    }
  };

  const handleDelete = async (id) => {
    await removeUser(id);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="w-16 h-16 border-b-4 border-blue-600 rounded-full animate-spin"></div>
      </div>
    );
  }
  if (error) return <p className="text-center text-red-500">Error: {error.message}</p>;

  return (
    <DashboardLayout>
      <div className="container flex-1 p-4 mx-auto">
        <h1 className="mb-4 text-3xl font-bold">User Management</h1>
        <UserForm user={editingUser} onSave={handleSave} />
        <UserTable users={users} onEdit={handleEdit} onDelete={handleDelete} />
      </div>
    </DashboardLayout>
  );
};

export default UserManagementPage;
