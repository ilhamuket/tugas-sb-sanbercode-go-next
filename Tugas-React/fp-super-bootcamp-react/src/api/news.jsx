import axios from './axiosConfig';

export const getAllNews = async () => {
  const response = await axios.get('/news');
  return response.data;
};

export const getNewsById = async (id) => {
  const response = await axios.get(`/news/${id}`);
  return response.data;
};

export const createNews = async (newsData) => {
  const response = await axios.post('/news', newsData);
  return response.data;
};

export const updateNews = async (id, newsData) => {
  const response = await axios.put(`/news/${id}`, newsData);
  return response.data;
};

export const deleteNews = async (id) => {
  const response = await axios.delete(`/news/${id}`);
  return response.data;
};
