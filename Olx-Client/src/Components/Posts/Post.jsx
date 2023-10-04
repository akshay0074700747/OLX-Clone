import { useContext } from "react";
import Heart from "../../assets/Heart";
import "./Post.css";
import { MyContext } from "../../App";
import { Link } from "react-router-dom";

export default function Post() {
  const value = useContext(MyContext);
  return (
    <div>
      {value.map((item) => {
        // Assuming item.images[0].image contains raw PNG image data
        const base64ImageData = `data:image/png;base64,${item.images[0].image}`;
        return (
          <Link to={`/details/${item.product.productid}`}>
            <div className="card" key={item.id}>
              <div className="favorite">
                <Heart></Heart>
              </div>
              <div className="image">
                <img src={base64ImageData} alt="" />
              </div>
              <div className="content">
                <p className="rate">&#x20B9; {item.product.price}</p>
                <span className="kilometer">{item.address.city}</span>
                <p className="name">{item.product.productname}</p>
              </div>
              <div className="date">
                <span>{item.product.soldby}</span>
              </div>
            </div>
          </Link>
        );
      })}
    </div>
  );
}
