
from . import main
from .. import db
from ..models import Permission, Role, User
from flask import render_template


@main.route('/create/<token>', methods=['GET', 'POST'])
def create(token):
    
    u = User.query().filter_by(self.token = token).first()
    print(u._slack_id)

    return

