import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import { NewsProvider } from './context/NewsContext';
import { CommentsProvider } from './context/CommentsContext';
import { UserProvider } from './context/UserContext'; 
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ChangePasswordPage from './pages/ChangePasswordPage';
import Home from './pages/Home';
import NewsPage from './pages/NewsPage'; 
import NewsDetail from './components/News/NewsDetail'; 
import UserManagementPage from './pages/UserManagementPage'; 
import ProtectedRoute from './components/ProtectedRoute';
import ProfilePage from './pages/ProfilePage';

const App = () => {
  return (
    <Router>
      <AuthProvider>
        <NewsProvider>
          <CommentsProvider>
            <UserProvider> {/* Wrap UserProvider around the other providers */}
              <Routes>
                <Route path="/login" element={<LoginPage />} />
                <Route path="/register" element={<RegisterPage />} />
                <Route path="/change-password" element={<ChangePasswordPage />} />
                <Route path="/"  element={<Home />}/>
                <Route path="/news" element={<ProtectedRoute element={<NewsPage />} />} />
                <Route path="/profile" element={<ProtectedRoute element={<ProfilePage />} />} />
                <Route path="/news/:id" element={<NewsDetail />}  />
                <Route path="/users" element={<ProtectedRoute element={<UserManagementPage />} />} /> {/* Add route for UserManagementPage */}
              </Routes>
            </UserProvider>
          </CommentsProvider>
        </NewsProvider>
      </AuthProvider>
    </Router>
  );
};

export default App;
