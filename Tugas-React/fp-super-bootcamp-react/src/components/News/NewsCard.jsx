import PropTypes from 'prop-types';

const NewsCard = ({ news }) => {
  // Menggunakan Lorem Picsum untuk gambar acak
  const imageUrl = `https://picsum.photos/400/200?random=${news.id}`;

  return (
    <div className="overflow-hidden bg-white rounded-lg shadow-md">
      <img
        src={imageUrl}
        alt={news.title}
        className="object-cover w-full h-48"
      />
      <div className="p-6">
        <h2 className="mb-4 text-2xl font-bold">{news.title}</h2>
        <p className="mb-4 text-gray-700">{news.content}</p>
        <a
          href={`/news/${news.id}`}
          className="text-blue-500 hover:underline"
        >
          Read more
        </a>
      </div>
    </div>
  );
};

NewsCard.propTypes = {
  news: PropTypes.shape({
    id: PropTypes.number.isRequired,
    title: PropTypes.string.isRequired,
    content: PropTypes.string.isRequired,
    imageUrl: PropTypes.string,
  }).isRequired,
};

export default NewsCard;
