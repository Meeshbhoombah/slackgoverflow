"""Recieve and respond to user texts via Twilio.
"""

import pprint
from flask import request
from flask_restful import Resource, reqparse
from .model import User

class User(Resource):

    def post(self):
        """Create a new user"""
        pprint.pprint(request)

        return {'status': 'OK'}, 200


