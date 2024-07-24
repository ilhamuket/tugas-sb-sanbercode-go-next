import { useState } from 'react';
import useAuth from '../../hooks/useAuth';

const Register = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const { handleRegister } = useAuth();

  const onSubmit = async (e) => {
    e.preventDefault();
    try {
      await handleRegister({ username, password, email, role: 'editor' });
    } catch (error) {
      console.error('Registration failed:', error);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <div className="card w-full max-w-sm bg-base-100 shadow-xl">
        <div className="card-body">
          <h2 className="text-2xl font-bold mb-4">Register</h2>
          <form onSubmit={onSubmit}>
            <div className="mb-4">
              <label className="label">
                <span className="label-text">Username</span>
              </label>
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="input input-bordered w-full"
                required
              />
            </div>
            <div className="mb-4">
              <label className="label">
                <span className="label-text">Password</span>
              </label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="input input-bordered w-full"
                required
              />
            </div>
            <div className="mb-4">
              <label className="label">
                <span className="label-text">Email</span>
              </label>
              <input
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="input input-bordered w-full"
                required
              />
            </div>
            {/* Hidden input for role */}
            <input type="hidden" name="role" value="editor" />
            <button
              type="submit"
              className="btn btn-primary w-full"
            >
              Register
            </button>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Register;
