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
      <div className="container px-4 py-8 mx-auto">
        <h1 className="mb-6 text-4xl font-bold text-center">News</h1>
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
