"""Creates an instance of the Phase-0 server"""


from config import config
from flask import Flask


def create_app(server):
    """Configures the server, db, and routes"""
    app = Flask(__name__)

    app.config.from_object(config[server] or 'default') 
    config[server].init_app(app)

