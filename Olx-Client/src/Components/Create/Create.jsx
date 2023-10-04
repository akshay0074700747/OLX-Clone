import React, { Fragment, useState } from "react";
import "./Create.css";
import Cookies from "js-cookie";

const Create = ({ count, setcount }) => {
  const [files, setfiles] = useState([]);
  const [prodname, setprodname] = useState("");
  const [proddesc, setproddesc] = useState("");
  const [prodprice, setprodprice] = useState("");
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [addressadded, setaddressadded] = useState(false);
  const [housename, sethousename] = useState("");
  const [locality, setlocality] = useState("");
  const [city, setcity] = useState("");
  const [district, setdistrict] = useState("");
  const [state, setstate] = useState("");
  const [pin, setpin] = useState("");

  const onmodalsubmit = (e) => {
    e.preventDefault();

    const formdata = {
      house: housename,
      locality: locality,
      city: city,
      district: district,
      state: state,
      pin: pin,
    };
    console.log(formdata.house);
    //converts the form data into url encoded format
    const urlencoded = new URLSearchParams(formdata).toString();

    fetch("http://localhost:8080/address", {
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
        console.log(data);
      })
      .catch((error) => {
        console.error("Error:", error);
      });

    setIsModalOpen(false);
    setaddressadded(true);
  };

  const addAddress = () => {
    setIsModalOpen(true);
  };

  const closePopup = () => {
    setIsModalOpen(false);
  };

  const onsubmit = (e) => {
    console.log(
      "heeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeyyyyyyyyyyyyyyyyyyyyyyy"
    );
    e.preventDefault();
    if (!addressadded) {
      return;
    }
    const formData = new FormData();

    // Append form fields and files to the FormData object
    formData.append("productname", prodname);
    formData.append("productdesc", proddesc);
    formData.append("price", prodprice);

    // Append each selected file
    for (let i = 0; i < files.length; i++) {
      formData.append("images", files[i]);
    }

    const username = Cookies.get("username");

    console.log(username);

    fetch(`http://localhost:8080/sell/${username}`, {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        console.log(data);
        setcount(count + 1);
      })
      .catch((error) => {
        console.error("Error:", error);
      });
    closePopup();
  };

  return (
    <Fragment>
      <div className="centerDiv">
        <form>
          <label htmlFor="fname">Product Name</label>
          <br />
          <input
            className="input"
            type="text"
            id="fname"
            name="productname"
            value={prodname}
            onChange={(e) => setprodname(e.target.value)}
            required
          />
          <br />
          <label htmlFor="fname">Product Description</label>
          <br />
          <input
            className="input"
            type="text"
            id="fname"
            name="productdesc"
            value={proddesc}
            onChange={(e) => setproddesc(e.target.value)}
            required
          />
          <br />
          <label htmlFor="fname">Price</label>
          <br />
          <input
            className="input"
            type="text"
            id="fname"
            name="price"
            value={prodprice}
            onChange={(e) => setprodprice(e.target.value)}
            required
          />
          <br />
        </form>
        <br />
        <img alt="Posts" width="200px" height="200px" src=""></img>
        <form encType="multipart/form-data">
          <br />
          <input
            type="file"
            name="images"
            multiple={true}
            onChange={(e) => setfiles([...files, ...e.target.files])}
          />
          <br />
        </form>
        <br />
        <button onClick={addAddress}>Add Address</button>
        <button onClick={(e) => onsubmit(e)}>Upload and Submit</button>
        <br />
        {isModalOpen && (
          <div className="modal">
            <div className="modal-content">
              <span className="close" onClick={closePopup}>
                &times;
              </span>
              <h2>Add Address</h2>
              <form onSubmit={onmodalsubmit}>
                <label htmlFor="street">House Name:</label>
                <input
                  type="text"
                  id="street"
                  name="house"
                  placeholder="House name"
                  value={housename}
                  onChange={(e) => sethousename(e.target.value)}
                  required
                />
                <br />
                <label htmlFor="street">Locality:</label>
                <input
                  type="text"
                  id="street"
                  name="locality"
                  placeholder="Locality"
                  value={locality}
                  onChange={(e) => setlocality(e.target.value)}
                  required
                />
                <br />
                <label htmlFor="city">City:</label>
                <input
                  type="text"
                  id="city"
                  name="city"
                  placeholder="City"
                  value={city}
                  onChange={(e) => setcity(e.target.value)}
                  required
                />
                <br />
                <label htmlFor="city">District:</label>
                <input
                  type="text"
                  id="city"
                  name="district"
                  placeholder="District"
                  value={district}
                  onChange={(e) => setdistrict(e.target.value)}
                  required
                />
                <br />
                <label htmlFor="state">State:</label>
                <input
                  type="text"
                  id="state"
                  name="state"
                  placeholder="State"
                  value={state}
                  onChange={(e) => setstate(e.target.value)}
                  required
                />
                <br />

                <label htmlFor="zipcode">PIN Code:</label>
                <input
                  type="text"
                  id="zipcode"
                  name="pin"
                  placeholder="PIN Code"
                  value={pin}
                  onChange={(e) => setpin(e.target.value)}
                  required
                />
                <br />

                <button type="submit">Submit Address</button>
              </form>
            </div>
          </div>
        )}
      </div>
    </Fragment>
  );
};

export default Create;
