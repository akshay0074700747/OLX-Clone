import React from 'react';
import Login from '../Components/Login/Login';

function LoginPage({error,seterror,count,setcount}) {
  return (
    <div>
      <Login error={error} seterror={seterror} count={count} setcount={setcount}/>
    </div>
  );
}

export default LoginPage;