import { useContext } from 'react';
import { BookContext } from './BookContext';
import './BookForm.css';

const BookForm = () => {
    const { form, handleChange, handleSubmit } = useContext(BookContext);

    return (
        <form className="form" onSubmit={handleSubmit}>
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
    );
};

export default BookForm;
