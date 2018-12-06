"""Models for the Phase-0 server"""

from datetime import datetime
from flask import current_app, request, url_for
from . import db


class Permission:
    REACT   = 1
    EARN    = 2
    ASK     = 4
    ANSWER  = 8
    VERIFY  = 16


class Role(db.Model):
    
    __tablename__ = 'roles'

    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(64), unique = True)
    default = db.Column(db.Boolean, default = False, index = True)
    permissions = db.Column(db.Integer)
    users = db.relationship('User', backref = 'role', lazy = 'dynamic')


    def __init__(self, **kwargs):
        super(Role, self).__init__(**kwargs)
        if self.permissions is None:
            self.permissions = 0


    @staticmethod
    def insert_roles():
        roles = {
                'Unregistered'  : [Permission.REACT]
                'User'          : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT]

                'Student'       : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT]
                'MVP'           : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT, 
                                    Permission.VERIFY]

                'TA'            : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT,
                                    Permission.VERIFY]
                'RA'            : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT
                                    Permission.VERIFY]

                'Staff'         : [Permission.ASK, Permission.ANSWER,
                                    Permission.REACT, Permission.VERIFY]
                'Instructor'    : [Permission.ASK, Permission.ANSWER,
                                    Permission.REACT, Permission.VERIFY]
        }

        default_role = 'Unregistered'

        for r in roles:
            role = Role.query.filter_by(name = r).first()

            if role is None:
                role = Role(name = r)
            
            role.reset_permissions()

            for perm in role[r]:
                role.add_permission(perm)

            role.default = (role.name == default_role)
            db.session.add(role)

        db.session.commit()


    def add_permission(self, perm):
        if not self.has_permission(perm):
            self.permissions += perm


    def remove_permission(self, perm):
        if self.has_permission(perm):
            self.permissions -= perm


    def reset_permissions(self):
        self.permissions = 0


    def has_permission(self, perm):
        return self.permissions & perm == perm


    def __repr__(self):
        return '<Role %r>' % self.name


class User(db.Model):

