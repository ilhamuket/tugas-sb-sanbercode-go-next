import { BookProvider } from './BookContext';
import BookTable from './BookTable';
import BookForm from './BookForm';
import './TugasContext.css';

const TugasContext = () => {
    return (
        <BookProvider>
            <div className="tugas-crud-axios">
                <h1>Tugas CRUD Books</h1>
                <div className="swagger-link-container">
                    <a href="https://tugas-sb-sanbercode-go-next-2024-beige.vercel.app/swagger/index.html" target="_blank" rel="noopener noreferrer" className="swagger-link">Swagger Backend</a>
                </div>
                <BookForm />
                <BookTable />
            </div>
        </BookProvider>
    );
};

export default TugasContext;