import { useState } from 'react';
import PropTypes from 'prop-types';

const CommentForm = ({ newsId, onSubmit }) => {
  const [text, setText] = useState('');

  const handleSubmit = (e) => {
    e.preventDefault();
    if (text.trim() === '') {
      alert('Comment text is required');
      return;
    }
    onSubmit({ news_id: newsId, text });
    setText('');
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4 bg-base-200 p-4 rounded-lg shadow-md">
      <textarea
        value={text}
        onChange={(e) => setText(e.target.value)}
        placeholder="Write a comment"
        className="textarea textarea-bordered w-full h-24"
      />
      <button
        type="submit"
        className="btn btn-primary w-full"
      >
        Submit
      </button>
    </form>
  );
};

CommentForm.propTypes = {
  newsId: PropTypes.number.isRequired,
  onSubmit: PropTypes.func.isRequired,
};

export default CommentForm;
