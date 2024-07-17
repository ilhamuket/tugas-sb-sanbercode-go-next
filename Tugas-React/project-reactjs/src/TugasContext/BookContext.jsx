import { createContext, useState, useEffect } from 'react';
import axios from 'axios';
import { baseUrl } from '../utils/constants';

export const BookContext = createContext();

// eslint-disable-next-line react/prop-types
export const BookProvider = ({ children }) => {
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

        if (form.title && form.description && form.image_url && !isNaN(releaseYearInt) && !isNaN(priceFloat) && !isNaN(totalPageInt)) {
            if (!isValidURL(form.image_url)) {
                alert("Format URL gambar tidak valid");
                return;
            }

            if (releaseYearInt < 1980 || releaseYearInt > 2021) {
                alert("Tahun rilis harus berada di antara 1980 dan 2021");
                return;
            }

            try {
                if (editIndex === null) {
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
        <BookContext.Provider value={{
            books,
            form,
            editIndex,
            handleChange,
            handleSubmit,
            handleEdit,
            handleDelete,
        }}>
            {children}
        </BookContext.Provider>
    );
};