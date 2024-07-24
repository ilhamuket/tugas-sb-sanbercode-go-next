import { useState } from 'react';
import useAuth from '../../hooks/useAuth';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const { handleLogin } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await handleLogin({ username, password });
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="card w-full max-w-sm bg-base-100 shadow-xl">
        <div className="card-body">
          <h1 className="text-2xl font-bold mb-4">Login</h1>
          <form onSubmit={handleSubmit}>
            <label className="block mb-4">
              <span className="label-text">Username:</span>
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="input input-bordered w-full"
              />
            </label>
            <label className="block mb-6">
              <span className="label-text">Password:</span>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="input input-bordered w-full"
              />
            </label>
            <button
              type="submit"
              className="btn btn-primary w-full"
            >
              Login
            </button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
