import { useState, useEffect, useContext } from 'react';
import { UserContext } from '../../context/UserContext';
import Swal from 'sweetalert2';

const ProfileForm = () => {
  const { profile, fetchProfile, editProfile } = useContext(UserContext);
  const [bio, setBio] = useState('');
  const [picture, setPicture] = useState('');

  useEffect(() => {
    fetchProfile();
  }, []);

  useEffect(() => {
    if (profile) {
      setBio(profile.profile.bio || '');
      setPicture(profile.profile.picture || '');
    }
  }, [profile]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await editProfile({ bio, picture });
      Swal.fire({
        icon: 'success',
        title: 'Profile updated successfully!',
        text: 'Your profile has been updated.',
      });
    } catch (error) {
      Swal.fire({
        icon: 'error',
        title: 'Failed to update profile',
        text: 'There was an error updating your profile.',
      });
    }
  };

  return (
    <div className="w-full max-w-sm shadow-xl card bg-base-100">
      <div className="card-body">
        <h1 className="mb-4 text-2xl font-bold">Profile</h1>
        <form onSubmit={handleSubmit}>
          <label className="block mb-4">
            <span className="label-text">Bio:</span>
            <textarea
              value={bio}
              onChange={(e) => setBio(e.target.value)}
              className="w-full textarea textarea-bordered"
            />
          </label>
          <label className="block mb-4">
            <span className="label-text">Picture URL:</span>
            <input
              type="text"
              value={picture}
              onChange={(e) => setPicture(e.target.value)}
              className="w-full input input-bordered"
            />
          </label>
          <button type="submit" className="w-full btn btn-primary">
            Update Profile
          </button>
        </form>
      </div>
    </div>
  );
};

export default ProfileForm;
