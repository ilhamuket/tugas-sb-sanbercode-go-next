import { useState, useEffect } from 'react';
import PropTypes from 'prop-types';

const NewsForm = ({ news = {}, onSubmit, onUpdate }) => {
  
  const [title, setTitle] = useState(news?.title || '');
  const [content, setContent] = useState(news?.content || '');

  useEffect(() => {
    
    if (news) {
      setTitle(news.title || '');
      setContent(news.content || '');
    }
  }, [news]);

  const resetForm = () => {
    setTitle('');
    setContent('');
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const data = { title, content };

    if (news?.id) {
      onUpdate(data); 
    } else {
      onSubmit(data);
    }

    resetForm(); 
  };

  return (
    <form onSubmit={handleSubmit} className="card bg-base-100 shadow-xl p-6">
      <h2 className="text-2xl font-bold mb-4">{news?.id ? 'Edit News' : 'Submit News'}</h2>
      <div className="mb-4">
        <label className="label">
          <span className="label-text">Title</span>
        </label>
        <input
          type="text"
          placeholder="Enter title"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="input input-bordered w-full"
          required
        />
      </div>
      <div className="mb-4">
        <label className="label">
          <span className="label-text">Content</span>
        </label>
        <textarea
          placeholder="Enter content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="textarea textarea-bordered w-full"
          rows="4"
          required
        />
      </div>
      <button type="submit" className="btn btn-primary w-full">
        {news?.id ? 'Update' : 'Submit'}
      </button>
    </form>
  );
};

NewsForm.propTypes = {
  news: PropTypes.shape({
    id: PropTypes.number,
    title: PropTypes.string,
    content: PropTypes.string,
  }),
  onSubmit: PropTypes.func.isRequired,
  onUpdate: PropTypes.func,
};

export default NewsForm;
