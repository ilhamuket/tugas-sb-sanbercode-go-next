import { BookProvider } from './BookContext';
import BookTable from './BookTable';
import BookForm from './BookForm';
import './TugasContext.css';

const TugasContext = () => {
    return (
        <BookProvider>
            <div className="tugas-crud-axios">
                <h1>Tugas CRUD Books</h1>
                <div className="button-container">
                    <a href="https://tugas-sb-sanbercode-go-next-2024-xnnms-projects.vercel.app/swagger/index.html" target="_blank" rel="noopener noreferrer" className="swagger-link">Backend Live</a>
                    <a href="https://tugas-sb-sanbercode-go-next-2024-alu5.vercel.app" target="_blank" rel="noopener noreferrer" className="swagger-link">React Live</a>
                </div>
                <BookForm />
                <BookTable />
            </div>
        </BookProvider>
    );
};

export default TugasContext;