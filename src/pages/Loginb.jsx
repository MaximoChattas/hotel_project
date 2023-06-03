import React, { useState } from "react";
import "./Loginb.css";
import { Link, NavLink } from "react-router-dom";


function Login() {

  // React States
  var database;
  const [errorMessages, setErrorMessages] = useState({});
  const [isSubmitted, setIsSubmitted] = useState(false);
  fetch("http://localhost:5173//item.json")
    .then((res) => res.json())
    .then(
      (result) => {
        database = result;
      },
      (error) => {}
    );

  const errors = {
    uname: "invalid username",
    pass: "invalid password",
  };

  const handleSubmit = (event) => {
    //Prevent page reload
    event.preventDefault();

    var { uname, pass } = document.forms[0];

    // Find user login info
    const userData = database.find((user) => user.username === uname.value);

    // Compare user info
    if (userData) {
      if (userData.password !== pass.value) {
        // Invalid password
        setErrorMessages({ name: "pass", message: errors.pass });
      } else {
        setIsSubmitted(true);
      }
    } else {
      // Username not found
      setErrorMessages({ name: "uname", message: errors.uname });
    }
  };

  // Generate JSX code for error message
  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

  // JSX code for login form
  const renderForm = (
    
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="title">Sign In</div>
        <div className="Username">
          <label>Username </label>
          <input type="text" name="uname" required />
          {renderErrorMessage("uname")}
        </div>
        <div className="Password">
          <label>Password </label>
          <input type="password" name="pass" required />
          {renderErrorMessage("pass")}
        </div>
        <div className="enviar">
          <input type="submit" />
        </div>
      </form>
    </div>
  );

  return (
    <>
    <header>
      <div className="container">
        <NavLink className="nav-link" to="/">
          <h1 className="asd">MIRANDA</h1>
        </NavLink>
      </div>
    </header>
    <div className="login-form">
        {isSubmitted ? <div>User is successfully logged in</div> : renderForm}
        </div>
    
    <footer>
      
        <>
        <NavLink className="nav-link" to="/">
        <svg xmlns="http://www.w3.org/2000/svg" width="50" height="50" fill="currentColor" class="bi bi-box-arrow-left" viewBox="0 0 16 16">
          <path fill-rule="evenodd" d="M6 12.5a.5.5 0 0 0 .5.5h8a.5.5 0 0 0 .5-.5v-9a.5.5 0 0 0-.5-.5h-8a.5.5 0 0 0-.5.5v2a.5.5 0 0 1-1 0v-2A1.5 1.5 0 0 1 6.5 2h8A1.5 1.5 0 0 1 16 3.5v9a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 5 12.5v-2a.5.5 0 0 1 1 0v2z"/>
          <path fill-rule="evenodd" d="M.146 8.354a.5.5 0 0 1 0-.708l3-3a.5.5 0 1 1 .708.708L1.707 7.5H10.5a.5.5 0 0 1 0 1H1.707l2.147 2.146a.5.5 0 0 1-.708.708l-3-3z"/>
          </svg>
        </NavLink>
        <h3>volver al menu</h3>
      </>
      
    </footer>
    </>
  );
}

export default Login;
