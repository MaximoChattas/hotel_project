import React, { useState, useEffect } from 'react';
import jwt_decode from 'jwt-decode';
import Navbar from '../NavBar/NavBar';

function Login() {
  const [loggedIn, setLoggedIn] = useState(false);
  const [loading, setLoading] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [userProfile, setUserProfile] = useState(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setLoggedIn(true);
      const decoded = jwt_decode(token);
      setUserProfile(decoded);
    }
  }, []);

  const handleLogin = async (e) => {
    e.preventDefault();
    setLoading(true);
    setError('');

    try {
      const response = await fetch('http://localhost:8090/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (response.status === 202) {
        const { token, user } = await response.json();
        localStorage.setItem('token', token);
        const decoded = jwt_decode(token);
        setLoggedIn(true);
        setUserProfile(user);
      } else {
        throw new Error('Invalid email or password');
      }
    } catch (error) {
      console.error(error);
      setError('Invalid email or password');
    } finally {
      setLoading(false);
    }
  };

  const handleLogout = () => {
    localStorage.removeItem('token');
    setLoggedIn(false);
    setUserProfile(null);
  };

  return (
    <>
      <Navbar />
      {error && <p>{error}</p>}
      {loggedIn ? (
        <>
            <h3>Profile</h3>
            <p>Name: {userProfile.name}</p>
            <p>Last name: {userProfile.last_name}</p>
            <p>DNI: {userProfile.dni}</p>
            <p>Email: {userProfile.email}</p>
            <p>ID: {userProfile.id}</p>
            <button onClick={handleLogout}>Logout</button>
        </>
      ) : (
        <>
            <h2>Login</h2>
            <form onSubmit={handleLogin}>
            <div>
                <label>Email:</label>
                <input type="email" value={email} disabled={loggedIn} onChange={(e) => setEmail(e.target.value)} />
            </div>
            <div>
                <label>Password:</label>
                <input type="password" value={password} disabled={loggedIn} onChange={(e) => setPassword(e.target.value)} />
            </div>
            <button type="submit" disabled={loading}>
                {loading ? 'Loading...' : 'Login'}
            </button>
            </form>
        </>
      )}
    </>
  );
}

export default Login;
