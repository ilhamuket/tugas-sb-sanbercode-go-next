import { useNews } from '../../context/NewsContext';
import NewsItem from './NewsItem';

const NewsList = () => {
  const { news, loading, error } = useNews();

  if (loading) return <p className="text-center">Loading...</p>;
  if (error) return <p className="text-center text-red-500">Error: {error.message}</p>;

  return (
    <div className="space-y-4">
      {news.map((item) => (
        <div key={item.id} className="card bg-base-100 shadow-xl p-6">
          <NewsItem news={item} />
        </div>
      ))}
    </div>
  );
};

export default NewsList;
