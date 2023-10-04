import React, { useRef, useState } from "react";

import Logo from "../../olx-logo.png";
import "./Signup.css";
import { Link } from "react-router-dom";
import Cookies from "js-cookie";

export default function Signup({error,seterror,count,setcount}) {
  const passref = useRef(null);
  const cnfrmpassref = useRef(null);
  const [invalidmasg, setinvalidmsg] = useState({
    show: false,
    msg: "",
  });
  const [username, setusername] = useState("");
  const [email, setemail] = useState("");
  const [mobile, setmobile] = useState("");
  const [password, setpassword] = useState("");
  const [confirmpass, setconfirmpass] = useState("");

  const onsubmit = (e) => {
    e.preventDefault();
    if (password != confirmpass) {
      setinvalidmsg({
        ...invalidmasg,
        [show]: true,
        [msg]: "Both the passwords must be the same...",
      });
      return;
    }
    const formdata = {
      username: username,
      email: email,
      mobile: mobile,
      password: password,
    };
    //converts the form data into url encoded format
    const urlencoded = new URLSearchParams(formdata).toString();

    fetch("http://localhost:8080/signup", {
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
        console.log("Signup successful:", data["token"]);
        // Cookies.set("jwtToken", data["token"], {
        //   path: "/",
        //   // secure: true,
        //   sameSite: "Strict",
        //   httpOnly: true,
        // });
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

        seterror(!error)
        setcount(count + 1)
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  };

  return (
    <div>
      <div className="signupParentDiv">
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
          <label htmlFor="fname">Email</label>
          <br />
          <input
            className="input"
            type="email"
            id="fname"
            name="email"
            value={email}
            onChange={(e) => setemail(e.target.value)}
            required
          />
          <br />
          <label htmlFor="lname">Mobile</label>
          <br />
          <input
            className="input"
            type="text"
            id="lname"
            name="mobile"
            value={mobile}
            onChange={(e) => setmobile(e.target.value)}
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
            ref={passref}
            onChange={(e) => {
              setpassword(e.target.value);
              if (password != confirmpass) {
                passref.current.style.color = "red";
                cnfrmpassref.current.style.color = "red";
              }
            }}
            required
          />
          <br />
          <label htmlFor="lname">Confirm Password</label>
          <br />
          <input
            className="input"
            type="password"
            id="lname"
            name="confirm-password"
            value={confirmpass}
            ref={cnfrmpassref}
            onChange={(e) => {
              setconfirmpass(e.target.value);
              if (password != confirmpass) {
                passref.current.style.color = "red";
                cnfrmpassref.current.style.color = "red";
              }
            }}
            required
          />
          <br />
          <br />
          <button type="submit">Signup</button>
        </form>
        {invalidmasg.show ? <p color="Red">{invalidmasg.msg}</p> : ""}
        <Link to="/login">Login</Link>
      </div>
    </div>
  );
}
