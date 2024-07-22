import useAuth from '../../hooks/useAuth';

const ChangePassword = () => {
  // eslint-disable-next-line no-undef
  const [password, setPassword] = useState('');
  const { handleChangePassword } = useAuth();

  const onSubmit = (e) => {
    e.preventDefault();
    handleChangePassword({ password });
  };

  return (
    <div>
      <h2>Change Password</h2>
      <form onSubmit={onSubmit}>
        <div>
          <label>New Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>
        <button type="submit">Change Password</button>
      </form>
    </div>
  );
};

export default ChangePassword;
