import { Link } from "react-router-dom";
import React from "react";


  
  
class Home extends React.Component {
  
  
  constructor(props) {
    super(props);
    this.state = {
      jokes: []
    };

    this.serverRequest = this.serverRequest.bind(this);

  }

 

  serverRequest() {
	  
	  
	  fetch('http://localhost:8080/customers').then(response => {
	        return response.json();
	      }).then(json => {
			  console.log("json",json)
	          this.setState({ jokes: json});
	      });

  }

  componentDidMount() {
    this.serverRequest();
  }

  render() {
    return (
      <div className="container">
        <br /> 
        
        <h2>CRM</h2>
		<br />
        <p>List of Customers</p>
        <div className="row">
          <div className="container">
		  <table border="1">
	       <tr>
		   <th>Name</th><th> Mobile</th> <th>Address</th>
		    </tr>
            {this.state.jokes.map((item, index) => (
				<tr border="1">
				<td border="1">
				{item.name} </td>
				<td border="1">{item.mobile}</td>
				<td border="1"> {item.address}</td>
				</tr>
              
            ))}
			<br />
			</table>
          </div>
			
			<a href="/add">Add New Customer</a>
        </div>
      </div>
    );
  }
}

export default Home;