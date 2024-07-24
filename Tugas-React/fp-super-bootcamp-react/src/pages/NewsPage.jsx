import NewsList from '../components/News/NewsList';
import NewsForm from '../components/News/NewsForm';
import { useNews } from '../context/NewsContext';

const NewsPage = () => {
  const { createNews } = useNews();

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-4xl font-bold mb-6 text-center">News</h1>
      <div className="mb-8">
        <NewsForm onSubmit={createNews} />
      </div>
      <NewsList />
    </div>
  );
};

export default NewsPage;
