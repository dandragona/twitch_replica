import React, { Component } from 'react';

class Stream extends Component {
    state = {
        userName: this.props.user_name,
        title: this.props.title,
        viewers: this.props.viewers,
        img: this.props.img
      }
    render() { 
        console.log('props', this.props);
        return (
            <div>
                <span>Title: {this.state.title}</span>
                <p>User: {this.state.userName}</p>
                <p>Viewers: {this.state.viewers}</p>
                <img 
                src={this.state.img}
                alt={this.state.userName + " logo"}/>
            </div>
          );
    }
}
 
export default Stream;