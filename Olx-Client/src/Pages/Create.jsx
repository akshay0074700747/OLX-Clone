import React, { Fragment } from "react";
import Header from "../Components/Header/Header";
import Create from "../Components/Create/Create";

const CreatePage = ({ count, setcount }) => {
  return (
    <Fragment>
      <Create count={count} setcount={setcount} />
    </Fragment>
  );
};

export default CreatePage;
