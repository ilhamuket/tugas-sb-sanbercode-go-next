import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext';
import { NewsProvider } from './context/NewsContext';
import { CommentsProvider } from './context/CommentsContext';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ChangePasswordPage from './pages/ChangePasswordPage';
import Home from './pages/Home';
import NewsPage from './pages/NewsPage';
import ProtectedRoute from './components/ProtectedRoute';

const App = () => {
  return (
    <Router>
      <AuthProvider>
      <NewsProvider>
      <CommentsProvider>
        <Routes>
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/change-password" element={<ChangePasswordPage />} />
          <Route path="/" element={<ProtectedRoute element={<Home />} />} />
          <Route path="/news" element={<ProtectedRoute element={<NewsPage />} />} />
        </Routes>
        </CommentsProvider>
        </NewsProvider>
      </AuthProvider>
    </Router>
  );
};

export default App;
