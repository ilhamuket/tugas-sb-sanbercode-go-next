import { useState } from 'react';
import PropTypes from 'prop-types';

const NewsForm = ({ news = {}, onSubmit }) => {
  const [title, setTitle] = useState(news.title || '');
  const [content, setContent] = useState(news.content || '');

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit({ title, content });
  };

  return (
    <form onSubmit={handleSubmit} className="card bg-base-100 shadow-xl p-6">
      <h2 className="text-2xl font-bold mb-4">Submit News</h2>
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
        Submit
      </button>
    </form>
  );
};

NewsForm.propTypes = {
  news: PropTypes.shape({
    title: PropTypes.string,
    content: PropTypes.string,
  }),
  onSubmit: PropTypes.func.isRequired,
};

export default NewsForm;
