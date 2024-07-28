import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useContext } from 'react';
import { AuthContext } from '../context/AuthContext';
import NewsLayout from '../components/News/NewsLayout';
import NewsCard from '../components/News/NewsCard';
import { useNews } from '../context/NewsContext';

const Home = () => {
  const { user } = useContext(AuthContext); 
  const navigate = useNavigate();
  const { news, loading, error } = useNews();

  // Efek ini tidak diperlukan lagi jika tidak perlu memaksa pengguna login
  useEffect(() => {
    // Jika diperlukan, bisa digunakan untuk logika lain
  }, [user, navigate]);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="w-16 h-16 border-b-4 border-blue-600 rounded-full animate-spin"></div>
      </div>
    );
  }
  if (error) return <div>Error loading news.</div>;

  return (
    <NewsLayout>
      <div className="grid grid-cols-1 gap-8 md:grid-cols-2 lg:grid-cols-3">
        {news.map((newsItem) => (
          <NewsCard key={newsItem.id} news={newsItem} />
        ))}
      </div>
    </NewsLayout>
  );
};

export default Home;
