import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { ThemeProvider } from './ThemeContext/ThemeContext';
import Navbar from './components/Navbar';
import TugasIntroReact from './Tugas-Intro-ReactJS/TugasIntroReact';
import TugasHooks from './Tugas-Hooks/TugasHooks';
import TugasCrudHooks from './Tugas-CRUD-Hooks/TugasCRUDHooks';
import TugasAxios from './TugasAxios/TugasAxios';
import TugasContext from './TugasContext/TugasContext';
import BookForm from './TugasContext/BookForm';
import { BookProvider } from './TugasContext/BookContext'; // Import BookProvider
import './App.css';

const App = () => {
  return (
    <ThemeProvider>
      <BookProvider> {/* Wrap with BookProvider */}
        <Router>
          <Navbar />
          <div className="App">
            <Routes>
              <Route path="/" element={<TugasIntroReact />} />
              <Route path="/hooks" element={<TugasHooks />} />
              <Route path="/crud" element={<TugasCrudHooks />} />
              <Route path="/axios" element={<TugasAxios />} />
              <Route path="/context" element={<TugasContext />} />
              <Route path="/create" element={<BookForm />} />
              <Route path="/edit/:id" element={<BookForm />} />
            </Routes>
          </div>
        </Router>
      </BookProvider>
    </ThemeProvider>
  );
};

export default App;
