import React from "react";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";

import AddCustomer from './AddCustomer';


import Logout from './Logout';


import Home from './Home';

export default function BasicExample() {
  return (
    <Router>
      <div>
        <ul>
          <li>
            <Link to="/home">Home</Link>
          </li>
          <li>
            <Link to="/addCustomer">Add Customer</Link>
          </li>
          <li>
            <Link to="/logout">Logout</Link>
          </li>
        </ul>

        <hr />

        {}
        <Switch>
          <Route exact path="/home">
            <Home />
          </Route>
          <Route path="/addCustomer">
            <AddCustomer />
          </Route>
          <Route path="/logout">
            <Logout />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

