import { createContext, useState, useContext, useCallback } from 'react';
import PropTypes from 'prop-types'; 
import * as commentsAPI from '../api/comments';

const CommentsContext = createContext();

export const useComments = () => useContext(CommentsContext);

export const CommentsProvider = ({ children }) => {
  const [commentsByNewsId, setCommentsByNewsId] = useState({});
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchComments = useCallback(async (newsId) => {
    setLoading(true);
    try {
      const data = await commentsAPI.getCommentsByNews(newsId);
      setCommentsByNewsId((prev) => ({
        ...prev,
        [newsId]: data
      }));
    } catch (error) {
      setError(error);
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
          [commentData.news_id]: [...prev[commentData.news_id], newComment]
        };
      });
    } catch (error) {
      setError(error);
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
          )
        };
      });
    } catch (error) {
      setError(error);
    }
  }, []);

  const deleteComment = useCallback(async (id) => {
    try {
      await commentsAPI.deleteComment(id);
      setCommentsByNewsId((prev) => {
        const newsId = Object.keys(prev).find((id) =>
          prev[id].some((comment) => comment.id === id)
        );
        if (!newsId) return prev;
        return {
          ...prev,
          [newsId]: prev[newsId].filter((comment) => comment.id !== id)
        };
      });
    } catch (error) {
      setError(error);
    }
  }, []);

  return (
    <CommentsContext.Provider value={{ commentsByNewsId, loading, error, fetchComments, createComment, updateComment, deleteComment }}>
      {children}
    </CommentsContext.Provider>
  );
};

// PropTypes validation
CommentsProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
