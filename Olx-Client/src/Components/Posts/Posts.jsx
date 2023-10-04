import React from "react";

import Heart from "../../assets/Heart";
import "./Post.css";
import Post from "./Post";


function Posts() {
  return (
    <div className="postParentDiv">
      <div className="moreView">
        <div className="heading">
          <span>Quick Menu</span>
          <span>View more</span>
        </div>
        <div className="cards">
          
          <Post></Post>
        </div>
      </div>
      <div className="recommendations">
        <div className="heading">
          <span>Fresh recommendations</span>
        </div>
        <div className="cards">
          <Post></Post>
        </div>
      </div>
    </div>
  );
}

export default Posts;
