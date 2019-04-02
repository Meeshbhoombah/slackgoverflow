import React, { Component } from 'react';
import axios from 'axios';

class Loading extends Component {
    parseUrl(url) => {
        console.log(typeof url) 

        return url
    } 

    async integrate() {
        await axios.post('/integrate', {
            code: this.parseUrl(window.location.href);,
        })
        .then((response) => {
            console.log(response);
        })
        .catch((err) => {
              console.error(err);
        });   
    }

    componentDidMount() {
        this.integrate();
    };

    render() {
        return(
            <div>
                <div className="sk-cube-grid">
                    <div className="sk-cube sk-cube1"></div>
                    <div className="sk-cube sk-cube2"></div>
                    <div className="sk-cube sk-cube3"></div>
                    <div className="sk-cube sk-cube4"></div>
                    <div className="sk-cube sk-cube5"></div>
                    <div className="sk-cube sk-cube6"></div>
                    <div className="sk-cube sk-cube7"></div>
                    <div className="sk-cube sk-cube8"></div>
                    <div className="sk-cube sk-cube9"></div>
                </div>
            </div>
        );
    };
};

export default Loading; 
