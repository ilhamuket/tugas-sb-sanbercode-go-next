import { Link } from 'react-router-dom';
import { useContext } from 'react';
import { BookContext } from './BookContext';
import './BookTable.css';

const BookTable = () => {
    const { books, handleDelete } = useContext(BookContext);

    return (
        <div>
            <div className="btn-container">
                <Link to="/create" className="btn-link">Create</Link>
            </div>
            <table className="table">
                <thead>
                    <tr>
                        <th>No</th>
                        <th>Title</th>
                        <th>Description</th>
                        <th>Image URL</th>
                        <th>Release Year</th>
                        <th>Price</th>
                        <th>Total Page</th>
                        <th>Thickness</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {books.map((book, index) => (
                        <tr key={book.id}>
                            <td>{index + 1}</td>
                            <td>{book.title}</td>
                            <td>{book.description}</td>
                            <td>
                                <a href={book.image_url} target="_blank" rel="noopener noreferrer">
                                    {book.image_url}
                                </a>
                            </td>
                            <td>{book.release_year}</td>
                            <td>{book.price}</td>
                            <td>{book.total_page}</td>
                            <td>{book.thickness}</td>
                            <td>
                                <Link to={`/edit/${book.id}`} className="btn-edit">Edit</Link>
                                <button className="btn-delete" onClick={() => handleDelete(index)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default BookTable;
