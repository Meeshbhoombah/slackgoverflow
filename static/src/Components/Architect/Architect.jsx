import React, { Component } from 'react';

class Architect extends Component {

    render() {
        return (
            <div className="home">
                <div className="header-content">
                    <h1 id="h1-1" className="h1-primary">#Hello</h1>
                    <h1 id="h1-2" className="h1-primary">#SLACKOVERFLOW</h1>
                    <h1 id="h1-3" className="h1-primary">#SLACKOVERFLOW</h1>
                    <p className="slogan">"Slack Overflow" x <span aria-label="emoji" role="img">💎</span>... in the PC Slack Workspace</p>
                    <button id="btn-1" className="btn-primary" onClick={this.handleClick}></button>
                    <button id="btn-2" className="btn-primary" onClick={this.handleClick}></button>
                    <button id="btn-3" className="btn-primary" onClick={this.handleClick}>ADD TO SLACK</button>
                
                    <div className="archie-section">
                    <p id="a-2" className="archie-text">Architect</p>
                    <p id="a-3" className="archie-text">Architect</p>
                    {/* <p id="h1-3" className="archie-text">Architect</p> */}
                    </div>

                </div>


            </div>
        );
    };
};

export default Architect;
