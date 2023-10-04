import React from "react";
import Header from "../Components/Header/Header";
import View from "../Components/View/View";
import { useParams } from "react-router-dom";

function ViewPost() {
  const { id } = useParams();
  return (
    <div>
      {/* <Header /> */}
      <View id={id} />
    </div>
  );
}

export default ViewPost;
