


import { useState } from "react";

import React from "react"


  
  
function Add() {

    

  const [name, setName] = useState("");
  const [mobile, setMobile] = useState("");
  const [address, setAddress] = useState("");
  const [message, setMessage] = useState("");
  


  let handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let res = fetch("http://localhost:8080/create", {
        method: "POST",
        body: JSON.stringify({
          name: name,
          mobile: mobile,
          address: address
        }),
      });
      let resJson = res.json();
      if (res.status === 200) {
        setName("");
        setAddress("");
        setMessage("User created successfully");
      } else {
        setMessage("Some error occured");
      }
    } catch (err) {
      console.log(err);
    }
22  };


    return (
      <div className="container">
        <br /> 
        
        <p>Add New Customer</p>
        <div className="row">
          <div className="container">
		  <form onSubmit={handleSubmit}>
        <div class="form-group">
          <label for="name">Customer name</label>
          <input value={name}
          placeholder="Name"
          onChange={(e) => setName(e.target.value)} class="form-control" tabindex="1" />
        </div>
        
        <div class="form-group">
          <label for="mobile">Customer Mobile：</label>
          <input value={mobile}
          placeholder="Mobile"
          onChange={(e) => setMobile(e.target.value)} class="form-control" tabindex="2" />
        </div>
        
        <div class="form-group">
          <label for="address">Address：</label>
          <input value={address}
          placeholder="Address"
          onChange={(e) => setAddress(e.target.value)} class="form-control" tabindex="3" />
        </div>
        
        <button>Add Customer</button>
        <a href="#" title="don't create the customer">Cancel</a>
      </form>
          </div>
			
			<a href="/add">Add New Customer</a>
        </div>
      </div>
    );
}


  
export default Add;