import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import Games from "./components/games"
import { Link } from 'react-router-dom';
import Streams from "./components/streams"


class App extends React.Component {
   render() {
      return (
      <Router>
      <div>
        <Route exact path="/" component={Home} />
        <Route exact path="/directory" component={Games} />
        <Route path={'/dashboard/:game'} component={Streams}/>
      </div>
    </Router>)
   }
}
export default App;

class Home extends React.Component {
  render() {
     return (
        <div>
           <h1>Home...</h1>
           <Link to="/directory"><button>Browse game directory!</button></Link>
        </div>
     )
  }
}