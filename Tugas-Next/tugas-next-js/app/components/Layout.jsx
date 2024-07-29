import Navbar from './Navbar';

const Layout = ({ children }) => {
  return (
    <div className="flex flex-col min-h-screen">
      <Navbar />
      <main className="flex-grow w-full p-0 mx-0">
        {children}
      </main>
      <footer className="p-4 text-center text-white bg-white">
        &copy; 2024 Your Company
      </footer>
    </div>
  );
};

export default Layout;
