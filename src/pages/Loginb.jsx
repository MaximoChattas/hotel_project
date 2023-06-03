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
    </>
  );
}

export default Login;
