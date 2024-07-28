import { useNews } from '../../context/NewsContext';
import NewsItem from './NewsItem';

// eslint-disable-next-line react/prop-types
const NewsList = ({ onEdit }) => {
  const { news, loading, error } = useNews();

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="w-16 h-16 border-b-4 border-blue-600 rounded-full animate-spin"></div>
      </div>
    );
  }
  if (error) return <p className="text-center text-red-500">Error: {error.message}</p>;

  return (
    <div className="space-y-4">
      {news.map((item) => (
        <div key={item.id} className="p-6 shadow-xl card bg-base-100">
          <NewsItem news={item} onEdit={onEdit} />
        </div>
      ))}
    </div>
  );
};

export default NewsList;
