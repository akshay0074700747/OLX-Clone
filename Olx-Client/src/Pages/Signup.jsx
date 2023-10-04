import React from 'react';

import Signup from '../Components/Signup/Signup';

function SignupPage({error,seterror,count,setcount}) {
  return (
    <div>
      <Signup error={error} seterror={seterror} count={count} setcount={setcount}/>
    </div>
  );
}

export default SignupPage;