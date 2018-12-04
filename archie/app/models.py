"""Models for the Phase-0 server"""

from datetime import datetime
from flask import current_app, request, url_for
from . import db


class Permission:
    """Manages Permission values for each User."""
    REACT   = 1
    EARN    = 2
    ASK     = 4
    ANSWER  = 8
    VERIFY  = 16


class Role(db.Model):
    """Assigns and manages roles of each User.
   
    Any person can join the channel and REACT (nurture/hinder) to a question,
    however, only those who complete the signup process can ASK/ANSWER 
    questions. Completing this process also unlocks the ability to EARN Drops.

    Each question is sorted into a number of categories dependent on the 
    context of the question. Initally these categories will be defined by the
    Asker of question (via #'s).
    """
    __tablename__ = 'roles'
    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(64), unique = True)
    default = db.Column(db.Boolean, default = False, index = True)
    permissions = db.Column(db.Integer)
    users = db.relationship('User', backref = 'role', lazy = 'dynamic')


    @staticmethod
    def insert_roles():
        roles = {
                'Unverified'    : []
                'User'          : [Permission.ASK, Permission.ANSWER]
                'Student'       : []
                'MVP'           : []
                'Staff'         : []
                'Instructor'    : []
                'TA'            : []
                'RA'            : []
        }

        default_role = 'User'

