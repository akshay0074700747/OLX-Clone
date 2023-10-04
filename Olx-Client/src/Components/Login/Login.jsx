import React from "react";

import Logo from "../../olx-logo.png";
import "./Login.css";
import { Link } from "react-router-dom";
import { useState } from "react";
import Cookies from "js-cookie";

function Login({ error, seterror, count, setcount }) {
  const [username, setusername] = useState("");
  const [password, setpassword] = useState("");

  const onsubmit = (e) => {
    e.preventDefault();
    const formdata = {
      username: username,
      password: password,
    };
    //converts the form data into url encoded format
    const urlencoded = new URLSearchParams(formdata).toString();

    fetch("http://localhost:8080/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      body: urlencoded,
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        console.log("Login successful:", data["token"]);
        Cookies.set("jwtToken", data["token"], {
          expires: 7,
          path: "/",
          sameSite: "Strict",
        });
        Cookies.set("username", formdata.username, {
          expires: 7,
          path: "/",
          sameSite: "Strict",
        });

        seterror(!error);
        setcount(count + 1);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  };

  return (
    <div>
      <div className="loginParentDiv">
        <img width="200px" height="200px" src={Logo}></img>
        <form onSubmit={onsubmit}>
          <label htmlFor="fname">Username</label>
          <br />
          <input
            className="input"
            type="text"
            id="fname"
            name="username"
            value={username}
            onChange={(e) => setusername(e.target.value)}
            required
          />
          <br />
          <label htmlFor="lname">Password</label>
          <br />
          <input
            className="input"
            type="password"
            id="lname"
            name="password"
            value={password}
            onChange={(e) => setpassword(e.target.value)}
            required
          />
          <br />
          <br />
          <button type="submit">Login</button>
        </form>
        <Link to="/signup">Signup</Link>
      </div>
    </div>
  );
}

export default Login;
