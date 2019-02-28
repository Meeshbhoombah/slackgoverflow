import React, { Component } from 'react';
import '../../Styles/home.scss';
import axios from 'axios';
class Home extends Component {

    handleClick() {
        // axios GET /authorize
        axios.get(`/authorize`).then((res) => {
            console.log("res:", res); 
        });
    }

    render() {
        return (
            <div className="home">
            <h1 className="h1-primary">#SLACKOVERFLOW</h1>
            <button className="btn-primary" onClick={this.handleClick}>ADD TO SLACK</button>
            </div>
        );
    };
};

export default Home;
