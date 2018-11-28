"""Handles User CRUD via `/user`"""


import pprint
from flask import request
from flask_restful import Resource, reqparse


from .model import User


class User(Resource):

    def post(self):
        """CREATE"""
        return {'status': 'OK'}, 200


