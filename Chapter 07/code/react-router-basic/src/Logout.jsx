import React from "react";
  
class Logout extends React.Component {
  constructor(props) {
    super(props);

    this.logout = this.logout.bind(this);
	

  }
  
  logout() {
    localStorage.removeItem("id_token");
    localStorage.removeItem("access_token");
    localStorage.removeItem("profile");
    //location.reload();
  }
  
  render() {
  
  return (
    <div>
      <h1>Logout</h1>
	  
        <span className="pull-right">
          <a onClick={this.logout}>Log out</a>
        </span>
    </div>
  );
 }
};
  
export default Logout;