import { useContext, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { BookContext } from './BookContext';
import './BookForm.css';

const BookForm = () => {
    const { form, handleChange, handleSubmit, books, setForm, setEditIndex } = useContext(BookContext);
    const { id } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        if (id) {
            const bookIndex = books.findIndex(book => book.id === parseInt(id));
            if (bookIndex !== -1) {
                const book = books[bookIndex];
                setForm({
                    title: book.title,
                    description: book.description,
                    image_url: book.image_url,
                    release_year: book.release_year,
                    price: book.price,
                    total_page: book.total_page
                });
                setEditIndex(bookIndex);
            }
        } else {
            setForm({
                title: "",
                description: "",
                image_url: "",
                release_year: "",
                price: "",
                total_page: ""
            });
            setEditIndex(null);
        }
    }, [id, books, setForm, setEditIndex]);

    const handleFormSubmit = async (event) => {
        event.preventDefault();
        const success = await handleSubmit(event);
        if (success) {
            navigate('/context');
            setForm({
                title: "",
                description: "",
                image_url: "",
                release_year: "",
                price: "",
                total_page: ""
            });
        }
    };

    return (
        <div className="form-container">
            <form className="form" onSubmit={handleFormSubmit}>
                <label>
                    Title:
                    <input type="text" name="title" value={form.title} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Description:
                    <textarea name="description" value={form.description} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Image URL:
                    <input type="text" name="image_url" value={form.image_url} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Release Year:
                    <input type="number" name="release_year" value={form.release_year} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Price:
                    <input type="number" name="price" value={form.price} onChange={handleChange} required />
                </label>
                <br />
                <label>
                    Total Page:
                    <input type="number" name="total_page" value={form.total_page} onChange={handleChange} required />
                </label>
                <br />
                <button type="submit">Submit</button>
            </form>
        </div>
    );
};

export default BookForm;
