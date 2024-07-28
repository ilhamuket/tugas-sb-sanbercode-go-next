import axios from './axiosConfig';

export const getComment = async (id) => {
  const response = await axios.get(`/comments/${id}`);
  return response.data;
};

export const getCommentsByNews = async (newsId) => {
  const response = await axios.get(`/news/comments/${newsId}`);
  return response.data;
};

export const createComment = async (commentData) => {
  const response = await axios.post('/comments', commentData);
  return response.data;
};

export const updateComment = async (id, commentData) => {
  const response = await axios.put(`/comments/${id}`, commentData);
  return response.data;
};

export const deleteComment = async (id) => {
  await axios.delete(`/comments/${id}`);
};
