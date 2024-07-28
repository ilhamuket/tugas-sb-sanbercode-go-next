import { useNavigate } from 'react-router-dom';
import ChangePassword from '../components/Auth/ChangePassword';

const ChangePasswordPage = () => {
  const navigate = useNavigate();

  return (
    <div className="relative">
      <button
        onClick={() => navigate('/')}
        className="absolute top-4 left-4 btn btn-primary"
      >
        Back to Home
      </button>
      <ChangePassword />
    </div>
  );
};

export default ChangePasswordPage;
