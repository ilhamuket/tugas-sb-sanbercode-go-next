import { useState } from 'react';
import NewsList from '../components/News/NewsList';
import NewsForm from '../components/News/NewsForm';
import DashboardLayout from '../components/Dashboard/DashboardLayout';
import { useNews } from '../context/NewsContext';

const NewsPage = () => {
  const { createNews, updateNews } = useNews();
  const [editingNews, setEditingNews] = useState(null);

  const handleEdit = (news) => {
    setEditingNews(news);
  };

  const handleSubmit = (newsData) => {
    if (editingNews) {
      updateNews(editingNews.id, newsData);
      setEditingNews(null); // Reset editing state after update
    } else {
      createNews(newsData);
    }
  };

  return (
    <DashboardLayout>
      <div className="container mx-auto px-4 py-8">
        <h1 className="text-4xl font-bold mb-6 text-center">News</h1>
        <div className="mb-8">
          <NewsForm
            news={editingNews}
            onSubmit={createNews}
            onUpdate={handleSubmit}
          />
        </div>
        <NewsList onEdit={handleEdit} />
      </div>
    </DashboardLayout>
  );
};

export default NewsPage;
