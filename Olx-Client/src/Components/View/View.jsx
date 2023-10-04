import React from "react";
import { MyContext } from "../../App";
import { useContext } from "react";

import "./View.css";
import ImageView from "./imgview";
function View({ id }) {
  const value = useContext(MyContext);

  return (
    <div>
      {value.map((item) => {
        if (id === item.product.productid) {
          {
            console.log("fgdjhagdsagdh");
          }
          return (
            <div className="viewParentDiv">
              <ImageView item={item}></ImageView>
              <div className="rightSection">
                <div className="productDetails">
                  <p>&#x20B9; {item.product.price} </p>
                  <span>{item.product.productname}</span>
                  <p>{item.product.productdesc}</p>
                </div>
                <div className="contactDetails">
                  <p>Seller details</p>
                  <p>{item.product.soldby}</p>
                  <p>{item.address.house}</p>
                  <p>{item.address.locality}</p>
                  <p>{item.address.city}</p>
                  <p>{item.address.district}</p>
                  <p>{item.address.state}</p>
                  <p>{item.address.pin}</p>
                </div>
              </div>
            </div>
          );
        }
      })}
    </div>
  );
}
export default View;
