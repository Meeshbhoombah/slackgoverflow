
from flask import current_app
from slackclient import SlackClient
from ..models import User


class Message(object):

    
    def __init__(self, User):
        self._id = current_app.config['SLACK_CLIENT_ID']
        self.secret = current_app.config['SLACK_CLIENT_SECRET']
        self.u = User

        self.sc = SlackClient(client_id = self._id, client_secret = self.secret)


    def onboard(self):
        token = self.u.generate_registration_token() 

        msg = {
            "text": "Thanks for joining the Architect channel!"
        }

        self.sc.api_call(
            "chat.PostMessage",
            channel = self.u._slack_id,
            text = msg
        )


    def welcome_back(self)
        msg = {
            "text": "Welcome back!"
        }


