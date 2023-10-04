import React, { createContext, useLayoutEffect, useState } from "react";
import { Route, Routes, useNavigate } from "react-router-dom";
import "./App.css";

/**
 * ?  =====Import Components=====
 */
import Home from "./Pages/Home";
import SignupPage from "./Pages/Signup";
import LoginPage from "./Pages/Login";
import Cookies from "js-cookie";
import CreatePage from "./Pages/Create";
import ViewPost from "./Pages/ViewPost";

export const MyContext = React.createContext();

function App() {
  const navigate = useNavigate();
  const apiurl = "http://localhost:8080";
  const [data, setdata] = useState([]);
  const [error, seterror] = useState(false);
  const [count, setcount] = useState(0);

  useLayoutEffect(() => {
    setcount(count + 1);
  }, []);

  useLayoutEffect(() => {
    console.log("heereeeeee");
    fetch(apiurl, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${Cookies.get("jwtToken")}`,
      },
    })
      .then((responce) => {
        if (!responce.ok) {
          seterror(!error);
          return;
        }
        return responce.json();
      })
      .then((dataa) => {
        setdata(dataa);
        // console.log(dataa);

        if (!error) {
          navigate("/");
        }
      });
  }, [count]);

  return (
    <MyContext.Provider value={data}>
      {console.log(data)}
      <div>
        <Routes>
          <Route
            exact
            path="/"
            element={
              error ? (
                <LoginPage
                  error={error}
                  seterror={seterror}
                  count={count}
                  setcount={setcount}
                ></LoginPage>
              ) : (
                <Home></Home>
              )
            }
          ></Route>
          <Route path="/details/:id" element={<ViewPost></ViewPost>}></Route>
          <Route
            path="/signup"
            element={
              <SignupPage
                error={error}
                seterror={seterror}
                count={count}
                setcount={setcount}
              ></SignupPage>
            }
          ></Route>
          <Route
            path="/sell"
            element={
              <CreatePage count={count} setcount={setcount}></CreatePage>
            }
          ></Route>
          <Route
            path="/login"
            element={
              <LoginPage
                error={error}
                seterror={seterror}
                count={count}
                setcount={setcount}
              ></LoginPage>
            }
          ></Route>
        </Routes>
      </div>
    </MyContext.Provider>
  );
}

export default App;
