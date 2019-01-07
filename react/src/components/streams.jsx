import React, { Component } from 'react';
import Stream from "./stream"

class Streams extends Component {
    state = { 
        game: this.props.match.params.game,
        streams: []
     }
    render() {
        this.getStreams()
        console.log('props', this.props);

        return (
            <div>
                {this.state.streams.map(stream => 
                <Stream 
                key={stream.userName}
                userName={stream.userName} 
                viewers={stream.viewers} 
                img={stream.img}
                title={stream.title}/>)}
            </div>
          );
    }

    getStreams = () => {
        console.log(this.state.game)
        if (this.state.streams.length !== 0) {
            return 
        }
        var url = "http://localhost:8000/directory/";
        url += this.state.game
        fetch(url, {mode: 'cors'})
        .then(response => response.json())
        .then(data => {
        // Prints result from `response.json()` in getRequest
        for (var i = 0; i < data.length; i++) {
            var stream = data[i]
            this.setState({game: this.state.game, streams: this.state.streams.concat({
                userName: stream.UserName,
                viewers: stream.Viewers,
                img: stream.Img,
                title: stream.Title
            }
            )});
        }
        })
        .catch(error => console.error(error))
    }
}

export default Streams;