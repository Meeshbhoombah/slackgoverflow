import React, { Component } from 'react';
import axios from 'axios';
import Loading from 'Components/Loading'

class Home extends Component {
    async handleClick() {
        // resp contains slackoverflow Slack OAuth URL
        await axios.get(`/authorize`)
        .then((res) => {
            window.location = res.data;
        })
        .catch((err) => {
              console.error(err);
        });
    }

    render() {
        return(
          <div className="home">
              {window.location.search ? ( 
                <Loading code={window.location.search}></Loading>
              ) : (
                <div className="header-content">
                    <h1 id="h1-1" className="h1-primary">#SLACKOVERFLOW</h1>
                    <h1 id="h1-2" className="h1-primary">#SLACKOVERFLOW</h1>
                    <h1 id="h1-3" className="h1-primary">#SLACKOVERFLOW</h1>
                    <p className="slogan">"Slack Overflow" x <span aria-label="emoji" role="img">💎</span>... in the PC Slack Workspace</p>
                    <button id="btn-1" className="btn-primary" onClick={this.handleClick}></button>
                    <button id="btn-2" className="btn-primary" onClick={this.handleClick}></button>
                    <button id="btn-3" className="btn-primary" onClick={this.handleClick}>ADD TO SLACK</button>
                </div>
              )}
          </div>
        );
    };
};

export default Home;
