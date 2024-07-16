import { useState, useEffect } from 'react';
import axios from 'axios';
import { baseUrl } from '../utils/constants';
import './TugasAxios.css';

const TugasAxios = () => {
    const [books, setBooks] = useState([]);
    const [form, setForm] = useState({
        title: "",
        description: "",
        image_url: "",
        release_year: "",
        price: "",
        total_page: "",
    });

    const [editIndex, setEditIndex] = useState(null);

    useEffect(() => {
        fetchData();
    }, []);

    const fetchData = async () => {
        try {
            const response = await axios.get(`${baseUrl}/books`);
            console.log('Response data:', response.data);
            
            const booksData = Array.isArray(response.data) ? response.data : [];
            setBooks(booksData);
        } catch (error) {
            console.error('Error fetching books:', error);
        }
    };

    const handleChange = (event) => {
        setForm({
            ...form,
            [event.target.name]: event.target.value
        });
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        
        const releaseYearInt = parseInt(form.release_year);
        const totalPageInt = parseInt(form.total_page);
        const priceFloat = parseFloat(form.price);
        console.log('Release Year:', releaseYearInt);
        console.log('Total Page:', totalPageInt);

        if (form.title && form.description && form.image_url && !isNaN(releaseYearInt) && !isNaN(priceFloat) && !isNaN(totalPageInt)) {
            // Validasi Image URL
            if (!isValidURL(form.image_url)) {
                alert("Format URL gambar tidak valid");
                return;
            }

            // Validasi Release Year
            if (releaseYearInt < 1980 || releaseYearInt > 2021) {
                alert("Tahun rilis harus berada di antara 1980 dan 2021");
                return;
            }

            try {
                if (editIndex === null) {
                    console.log('Saving book:', form);
                    await axios.post(`${baseUrl}/books`, {
                        ...form,
                        release_year: releaseYearInt, 
                        total_page: totalPageInt,
                    });
                } else {
                    await axios.patch(`${baseUrl}/books/${books[editIndex].id}`, {
                        ...form,
                        release_year: releaseYearInt, 
                        total_page: totalPageInt,
                    });
                    setEditIndex(null);
                }
                fetchData(); 
                setForm({
                    title: "",
                    description: "",
                    image_url: "",
                    release_year: "",
                    price: "",
                    total_page: ""
                });
            } catch (error) {
                console.error('Error saving book:', error);
            }
        } else {
            alert("All fields are required, and numeric fields must contain valid numbers.");
        }
    };

    const handleEdit = (index) => {
        const book = books[index];
        setForm({
            title: book.title,
            description: book.description,
            image_url: book.image_url,
            release_year: book.release_year,
            price: book.price,
            total_page: book.total_page
        });
        setEditIndex(index);
    };

    const handleDelete = async (index) => {
        console.log('Deleting book:', books[index]);
        try {
            const bookIdToDelete = books[index].id; 
            await axios.delete(`${baseUrl}/books/${bookIdToDelete}`);
            fetchData(); 
        } catch (error) {
            console.error('Error deleting book:', error);
        }
    };    

    const isValidURL = (url) => {
        return /^https?:\/\//.test(url);
    };

    return (
        <div className="tugas-crud-axios">
            <h1>Tugas CRUD Books</h1>
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
                        <tr key={index}>
                            <td>{index + 1}</td>
                            <td>{book.title}</td>
                            <td>{book.description}</td>
                            <td>{book.image_url}</td>
                            <td>{book.release_year}</td>
                            <td>{book.price}</td>
                            <td>{book.total_page}</td>
                            <td>{book.thickness}</td>
                            <td>
                                <button className="btn-edit" onClick={() => handleEdit(index)}>Edit</button>
                                <button className="btn-delete" onClick={() => handleDelete(index)}>Delete</button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default TugasAxios;
