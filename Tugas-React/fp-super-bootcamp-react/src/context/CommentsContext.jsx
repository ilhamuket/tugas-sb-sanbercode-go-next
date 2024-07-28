import { createContext, useState, useContext, useCallback } from 'react';
import PropTypes from 'prop-types';
import * as commentsAPI from '../api/comments';
import Swal from 'sweetalert2';
import { useNavigate } from 'react-router-dom';

const CommentsContext = createContext();

export const useComments = () => useContext(CommentsContext);

export const CommentsProvider = ({ children }) => {
  const [commentsByNewsId, setCommentsByNewsId] = useState({});
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleError = (error) => {
    if (error.response && error.response.status === 401) {
      Swal.fire({
        icon: 'error',
        title: 'Access Denied',
        text: error.response.data.error || 'You do not have permission to access this page (401).',
      }).then(() => {
        navigate('/login');
      });
    } else {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: error.response.data.error || 'An unexpected error occurred.',
      });
    }
    setError(error);
  };

  const fetchComments = useCallback(async (newsId) => {
    setLoading(true);
    try {
      const data = await commentsAPI.getCommentsByNews(newsId);
      setCommentsByNewsId((prev) => ({
        ...prev,
        [newsId]: data,
      }));
    } catch (error) {
      handleError(error);
    } finally {
      setLoading(false);
    }
  }, []);

  const createComment = useCallback(async (commentData) => {
    try {
      if (!commentData.news_id || !commentData.text) {
        throw new Error('Missing required fields');
      }

      const newComment = await commentsAPI.createComment(commentData);
      setCommentsByNewsId((prev) => {
        if (!prev[commentData.news_id]) return prev;
        return {
          ...prev,
          [commentData.news_id]: [...prev[commentData.news_id], newComment],
        };
      });
    } catch (error) {
      handleError(error);
    }
  }, []);

  const updateComment = useCallback(async (id, commentData) => {
    try {
      const updatedComment = await commentsAPI.updateComment(id, commentData);
      setCommentsByNewsId((prev) => {
        const newsComments = prev[updatedComment.news_id] || [];
        return {
          ...prev,
          [updatedComment.news_id]: newsComments.map((comment) =>
            comment.id === id ? updatedComment : comment
          ),
        };
      });
    } catch (error) {
      handleError(error);
    }
  }, []);

  const deleteComment = useCallback(async (id) => {
    try {
      await commentsAPI.deleteComment(id);
      setCommentsByNewsId((prev) => {
        const newsId = Object.keys(prev).find((newsId) =>
          prev[newsId].some((comment) => comment.id === id)
        );
        if (!newsId) return prev;
        return {
          ...prev,
          [newsId]: prev[newsId].filter((comment) => comment.id !== id),
        };
      });
    } catch (error) {
      handleError(error);
    }
  }, []);

  return (
    <CommentsContext.Provider
      value={{
        commentsByNewsId,
        loading,
        error,
        fetchComments,
        createComment,
        updateComment,
        deleteComment,
      }}
    >
      {children}
    </CommentsContext.Provider>
  );
};

// PropTypes validation
CommentsProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
