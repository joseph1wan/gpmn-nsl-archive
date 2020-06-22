export default () => {
  const authorizationToken = localStorage.getItem('authorizationToken');
  const isAuthenticated = !!authorizationToken;

  return [isAuthenticated, authorizationToken];
};
