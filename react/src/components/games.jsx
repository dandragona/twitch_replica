import React, { Component } from 'react';
import Game from "./game"
class Games extends Component {
    state = { 
        games: []
     }
    render() {
        this.getGames()
        return (
            <div>
                {this.state.games.map(game => 
                <Game 
                key={game.name}
                name={game.name} 
                viewers={game.viewers} 
                img={game.img} 
                link={game.link} />)}
            </div>
          );
    }

    getGames = () => {
        if (this.state.games.length !== 0) {
            return 
        }
        fetch('http://localhost:8000/directory', {mode: 'cors'})
        .then(response => response.json())
        .then(data => {
        // Prints result from `response.json()` in getRequest
        for (var i = 0; i < data.length; i++) {
            var game = data[i]
            this.setState({games: this.state.games.concat({
                name: game.Name,
                viewers: game.Viewers,
                img: game.Img,
                link: game.Link
            }
            )});
        }
        })
        .catch(error => console.error(error))
    }
}
 
export default Games;