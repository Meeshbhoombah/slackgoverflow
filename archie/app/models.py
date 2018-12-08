
from datetime import datetime
from flask import current_app, request, url_for
from werkzeug import generate_password_hash, check_password_hash
from . import db
from re import match
from itsdangerous import TimedJSONWebSignatureSerializer as Serializer
from itsdangerous import SignatureExpired, BadSignature


ONE_DAY = 86400 # in seconds


class Permission:
    REACT   = 1
    EARN    = 2
    ASK     = 4
    ANSWER  = 8
    VERIFY  = 16


class Role(db.Model):
    
    __tablename__ = 'roles'

    id = db.Column(db.Integer, primary_key = True)
    name = db.Column(db.String(64), unique = True, index = True)
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
                'Unregistered'  : [Permission.REACT],
                'User'          : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT],

                'Student'       : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT],
                'MVP'           : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT, 
                                    Permission.VERIFY],

                'TA'            : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT,
                                    Permission.VERIFY],
                'RA'            : [Permission.ASK, Permission.ANSWER, 
                                    Permission.EARN, Permission.REACT,
                                    Permission.VERIFY],

                'Staff'         : [Permission.ASK, Permission.ANSWER,
                                    Permission.REACT, Permission.VERIFY],
                'Instructor'    : [Permission.ASK, Permission.ANSWER,
                                    Permission.REACT, Permission.VERIFY],
        }

        default_role = 'Unregistered'

        for r in roles:
            role = Role.query.filter_by(name = r).first()

            if role is None:
                role = Role(name = r)
            
            role.reset_permissions()

            for perm in roles[r]:
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

    __tablename__ = 'users'

    id = db.Column(db.Integer, primary_key = True)
    role_id = db.Column(db.Integer, db.ForeignKey('roles.id'))
    _slack_id = db.Column(db.String(9), unique = True, index = True, nullable = False)

    username = db.Column(db.String(128), unique = True, index = True) 
    hashword = db.Column(db.String(128))
    drops = db.Column(db.Integer)

    last_seen = db.Column(db.DateTime(), default = datetime.utcnow)
    created_on = db.Column(db.DateTime(), default = datetime.utcnow)
    updated_on = db.Column(db.DateTime(), default = datetime.utcnow)

    questions = db.relationship('Question', backref = 'author', lazy = 'dynamic')
    answers = db.relationship('Answer', backref = 'author', lazy = 'dynamic') 
    comments = db.relationship('Comment', backref = 'author', lazy = 'dynamic')


    def __init__(self, **kwargs):
        super(User, self).__init__(**kwargs)
        
        if self.role is None:
            self.role = Role.query.filter_by(default = True).first()


    @property
    def password(self):
        raise AttributeError('Password is not a readable attribute.')


    @password.setter
    def password(self, password):
        self.hashword = generate_password_hash(password)


    @property
    def slack_id(self):
        return self._slack_id


    @slack_id.setter
    def slack_id(self, slack_id):
        """`slack_id` must begin w/ 'U' or 'W', & must be 11 chars in length"""
        if match(r'^[UW][A-Z0-9]{8}$', slack_id):
            self._slack_id = slack_id
        else:
            raise AttributeError('Does not match `slack_id` pattern: ^[UW][A-Z0-9]{8}$')


    def verify_password(self, password):
        return check_password_hash(self.hashword, password)


    def generate_registration_token(self, expiration = ONE_DAY):
        s = Serializer(current_app.config['SECRET_KEY'], expiration)

        token = s.dumps({
            'id': self.id,
            'slack_id': self._slack_id
        }).decode('utf-8')
        return token


    def confirm_registration(self, token):
        s = Serializer(current_app.config['SECRET_KEY'])

        try:
            data = s.loads(token.encode('utf-8'))
        except (BadSignature, SignatureExpired):
            return False

        if data.get('id') is not self.id \
            and data.get('slack_id') is not self._slack_id:
                return False

        role = Role.query.filter_by(name = 'User').first()
        self.role = role

        db.session.commit()
        return True


    def can(self, perm):
        return self.role is not None and self.role.has_permission(perm)


class Question(db.Model):

    __tablename__ = 'questions'

    id = db.Column(db.Integer, primary_key = True)
    body = db.Column(db.Text)
    nourishes = db.Column(db.Integer)
    starves = db.Column(db.Integer)
    
    author_id = db.Column(db.Integer, db.ForeignKey('users.id'))
    answer = db.relationship('Answer', backref = 'question', uselist = False)
    comments = db.relationship('Comment', backref = 'question', lazy = 'dynamic')


class Answer(db.Model):

    __tablename__ = 'answers'

    id = db.Column(db.Integer, primary_key = True)
    body = db.Column(db.Text)
    nourishes = db.Column(db.Integer)
    starves = db.Column(db.Integer)
    
    author_id = db.Column(db.Integer, db.ForeignKey('users.id'))
    question_id = db.Column(db.Integer, db.ForeignKey('questions.id'))


class Comment(db.Model):

    __tablename__ = 'comments'

    id = db.Column(db.Integer, primary_key = True)
    body = db.Column(db.Text)
    nourishes = db.Column(db.Integer)
    starves = db.Column(db.Integer)
    
    author_id = db.Column(db.Integer, db.ForeignKey('users.id'))
    question_id = db.Column(db.Integer, db.ForeignKey('questions.id'))


