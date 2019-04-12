import React, { Component } from 'react';
import axios from 'axios';

class Loading extends Component {
    parseCode = (search) => {
        let params = new URLSearchParams(search);
        return params.get("code");
    } 

    async integrate() {
        let search = window.location.search;
        console.log(search);

        await axios.post('/integrate', {
            code: this.parseCode(search),
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
