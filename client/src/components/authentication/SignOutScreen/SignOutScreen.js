import React, { useEffect } from 'react';
import { useHistory } from 'react-router-dom';

const SignOutScreen = () => {
  const history = useHistory();
  useEffect(() => {
    localStorage.clear();
    history.replace('/');
  });
  return <div>See ya! 👋</div>;
};

export default SignOutScreen;
