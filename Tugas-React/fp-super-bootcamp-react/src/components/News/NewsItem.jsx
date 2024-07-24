import { useState, useEffect } from 'react';
import PropTypes from 'prop-types';
import { useNews } from '../../context/NewsContext';
import { useComments } from '../../context/CommentsContext';
import CommentForm from '../Comments/CommentForm';
import CommentList from '../Comments/CommentList';

const NewsItem = ({ news }) => {
  const { id, title, content } = news;
  const { deleteNews } = useNews();
  const { fetchComments, createComment } = useComments();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      if (id) {
        try {
          await fetchComments(id);
        } catch (error) {
          console.error("Failed to fetch comments", error);
        } finally {
          setLoading(false);
        }
      }
    };

    fetchData();
  }, [id, fetchComments]);

  const handleDelete = () => {
    deleteNews(id);
  };

  const handleCreateComment = (commentData) => {
    createComment(commentData);
  };

  if (loading) return <p className="text-center">Loading comments...</p>;

  return (
    <div>
      <h3 className="text-xl font-semibold mb-2">{title}</h3>
      <p className="mb-4">{content}</p>
      <button
        onClick={handleDelete}
        className="absolute top-8 right-8 btn btn-error btn-sm"
      >
        Delete
      </button>

      {/* Comment List */}
      <div className="mb-6"> {/* Add margin-bottom here */}
        <CommentList newsId={id} />
      </div>

      {/* Comment Form */}
      <div> {/* Optional wrapper if you want to add margin-top */}
        <CommentForm newsId={id} onSubmit={handleCreateComment} />
      </div>
    </div>
  );
};

// PropTypes validation
NewsItem.propTypes = {
  news: PropTypes.shape({
    id: PropTypes.number.isRequired,
    title: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
  }).isRequired,
};

export default NewsItem;
