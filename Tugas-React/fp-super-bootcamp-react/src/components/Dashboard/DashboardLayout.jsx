import Sidebar from './Sidebar';

// eslint-disable-next-line react/prop-types
const DashboardLayout = ({ children }) => {
  return (
    <div className="flex h-auto min-h-screen">
      <Sidebar />
      <div className="flex-1 p-6">
        {children}
      </div>
    </div>
  );
};

export default DashboardLayout;
