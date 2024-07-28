import ProfileForm from '../components/User/Profileform';
import { UserProvider } from '../context/UserContext';
import DashboardLayout from '../components/Dashboard/DashboardLayout';

const ProfilePage = () => {
  return (
    <DashboardLayout>
    <UserProvider>
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <ProfileForm />
      </div>
    </UserProvider>
    </DashboardLayout>
  );
};

export default ProfilePage;
