import React, {Component} from "react";
import { Link } from 'react-router-dom';
import Streams from "./streams";

class Game extends Component {
    state = {  
        name: this.props.name,
        viewers: this.props.viewers,
        img: this.props.img,
        link: this.props.link,
    }
    render() { 
        console.log('props', this.props);
        return (
            <div>
                <span>Game: {this.state.name}</span>
                <p>Viewers: {this.state.viewers}</p>
                <img 
                src={this.state.img}
                alt={this.state.name + " logo"}/>
                <Link to={"/directory/" + this.state.name}><button>View streams for {this.state.name}.</button></Link>
            </div>
          );
    }
}
 
export default Game;