import PropTypes from 'prop-types';
import { createContext, useState, useEffect, useContext } from 'react';
import * as newsAPI from '../api/news';
import Swal from 'sweetalert2'; 
import { useNavigate } from 'react-router-dom'; 


const NewsContext = createContext();

export const useNews = () => useContext(NewsContext);


export const NewsProvider = ({ children }) => {
  const [news, setNews] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate(); 

  useEffect(() => {
    const fetchNews = async () => {
      setLoading(true);
      try {
        const data = await newsAPI.getAllNews();
        setNews(data);
      } catch (error) {
        handleError(error);
      }
      setLoading(false);
    };

    fetchNews();
  }, []);

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

  const createNews = async (newsData) => {
    try {
      const newNews = await newsAPI.createNews(newsData);
      setNews([...news, newNews]);
    } catch (error) {
      handleError(error);
    }
  };

  const updateNews = async (id, newsData) => {
    try {
      const updatedNews = await newsAPI.updateNews(id, newsData);
      setNews(news.map((newsItem) => (newsItem.id === id ? updatedNews : newsItem)));
    } catch (error) {
      handleError(error);
    }
  };

  const deleteNews = async (id) => {
    try {
      await newsAPI.deleteNews(id);
      setNews(news.filter((newsItem) => newsItem.id !== id));
    } catch (error) {
      handleError(error);
    }
  };

  const getNewsById = async (id) => {
    try {
      return await newsAPI.getNewsById(id);
    } catch (error) {
      handleError(error);
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


NewsProvider.propTypes = {
  children: PropTypes.node.isRequired,
};
