from flask import Blueprint


listen = Blueprint('listen', __name__)


from . import slack

