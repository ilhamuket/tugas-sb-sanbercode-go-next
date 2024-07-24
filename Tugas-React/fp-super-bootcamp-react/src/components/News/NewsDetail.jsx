import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { useNews } from '../../context/NewsContext';
import { useComments } from '../../context/CommentsContext';
import NewsLayout from './NewsLayout';
import CommentList from '../Comments/CommentList';
import CommentForm from '../Comments/CommentForm';

const NewsDetail = () => {
  const { id } = useParams();
  const navigate = useNavigate();
  const { getNewsById, error: newsError } = useNews();
  const { fetchComments, createComment } = useComments();
  const [newsItem, setNewsItem] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        // Fetch news item
        const item = await getNewsById(id);
        setNewsItem(item);
        
        // Fetch comments for the news item
        await fetchComments(id);
      } catch (error) {
        console.error('Failed to fetch news item or comments:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [id, getNewsById, fetchComments]);

  const handleCommentSubmit = async (commentData) => {
    try {
      await createComment(commentData);
      // Fetch comments again to update the list
      await fetchComments(id);
    } catch (error) {
      console.error('Failed to submit comment:', error);
    }
  };

  if (loading) return <p className="text-center">Loading comments...</p>;

  if (newsError) return <div>Error loading news detail: {newsError.message}</div>;
  if (!newsItem) return <div>No news item found.</div>;

  return (
    <NewsLayout>
      <div className="container mx-auto px-4 py-8">
        <div className="bg-white shadow-lg rounded-lg p-6 mb-8">
          <h1 className="text-3xl font-bold mb-4">{newsItem.title}</h1>
          <p className="text-gray-700 mb-4">{newsItem.content}</p>
          <button
            onClick={() => navigate('/')}
            className="btn btn-secondary"
          >
            Back to Home
          </button>
        </div>
        <div className="bg-white shadow-lg rounded-lg p-6">
          <CommentList newsId={newsItem.id} />
          <div className="mt-6"> {/* Jarak atas 6 untuk Tailwind CSS */}
            <CommentForm newsId={newsItem.id} onSubmit={handleCommentSubmit} />
        </div>
        </div>
      </div>
    </NewsLayout>
  );
};

export default NewsDetail;
