import PropTypes from 'prop-types';
import { createContext, useState, useEffect, useContext } from 'react';
import * as newsAPI from '../api/news';

// Context and Hook
const NewsContext = createContext();

export const useNews = () => useContext(NewsContext);

// Provider Component
export const NewsProvider = ({ children }) => {
  const [news, setNews] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchNews = async () => {
      setLoading(true);
      try {
        const data = await newsAPI.getAllNews();
        setNews(data);
      } catch (error) {
        setError(error);
      }
      setLoading(false);
    };

    fetchNews();
  }, []);

  const createNews = async (newsData) => {
    try {
      const newNews = await newsAPI.createNews(newsData);
      setNews([...news, newNews]);
    } catch (error) {
      setError(error);
    }
  };

  const updateNews = async (id, newsData) => {
    try {
      const updatedNews = await newsAPI.updateNews(id, newsData);
      setNews(news.map((newsItem) => (newsItem.id === id ? updatedNews : newsItem)));
    } catch (error) {
      setError(error);
    }
  };

  const deleteNews = async (id) => {
    try {
      await newsAPI.deleteNews(id);
      setNews(news.filter((newsItem) => newsItem.id !== id));
    } catch (error) {
      setError(error);
    }
  };

  const getNewsById = async (id) => {
    try {
      return await newsAPI.getNewsById(id);
    } catch (error) {
      setError(error);
      throw error;
    }
  };

  return (
    <NewsContext.Provider
      value={{ news, loading, error, createNews, updateNews, deleteNews, getNewsById }}
    >
      {children}
    </NewsContext.Provider>
  );
};

// PropTypes
NewsProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
