
from flask import request, current_app, make_response
from . import listen
from .messages import Message
import hashlib
import hmac
import sys
from time import time
import json
from ..import db
from ..models import User
from slackclient import SlackClient


def verify_signature(timestamp, signature, SIGNING_SECRET):
    # Compare the generated hash and incoming request signature
    if hasattr(hmac, "compare_digest"):
        req = str.encode('v0:' + str(timestamp) + ':') + request.data
        request_hash = hmac.new(str.encode(SIGNING_SECRET), req, hashlib.sha256)
        request_hash = 'v0=' + request_hash.hexdigest()

        # Compare byte strings for Python 2
        if (sys.version_info[0] == 2):
            return hmac.compare_digest(bytes(request_hash), bytes(signature))
        else:
            return hmac.compare_digest(request_hash, signature)

    else:
        req = str.encode('v0:' + str(timestamp) + ':') + request.data
        request_hash = 'v0=' + hmac.new(
            str.encode(self.SIGNING_SECRET),
            req, hashlib.sha256
        ).hexdigest()

        if len(request_hash) != len(signature):
            return False
        
        result = 0
        
        if isinstance(request_hash, bytes) and isinstance(signature, bytes):
            for x, y in zip(request_hash, signature):
                result |= x ^ y
        else:
            for x, y in zip(request_hash, signature):
                result |= ord(x) ^ ord(y)
        
        return result == 0


def member_joined_channel(event_data):
    """Handles `member_joined_channel` event when User joins the channel`"""
    sc = SlackClient(token = current_app.config['SLACK_BOT_TOKEN'])

    try:
        assert event_data['event']['channel'] == 'CEET14B25'
    except (AssertionError):
        return make_response("Not the #devp2p channel.", 404)

    # Get slack_id for User who triggered event
    member_id = event_data['event']['user']

    # Check if preexisting user
    u = User.query.filter_by(_slack_id = member_id).first()
    print(u)

    if u is None:
        # first time
        u = User(slack_id = member_id)

        db.session.add(u)
        db.session.commit()

        token = u.generate_registration_token() 

        msg = {
            "text": "Welcome to Architect, to confirm your account please sign in."
        }
       
        sc.api_call(
            "chat.postMessage",
            channel = self.u._slack_id,
            text = msg
        )

    else:
        u.pong()

        msg = {
            "text": "Welcome back!"
        }

        resp = sc.api_call(
            "chat.postMessage",
            channel = u._slack_id,
            text = msg
        )

        print(resp)

    response = make_response("Success.", 200)
    response.headers['X-Slack-Powered-By'] = 'Architect'
    return response
    

Handle = {
        'member_joined_channel' : member_joined_channel
}


@listen.route('/slack/event', methods=['POST'])
def event():
    """Verify each event's timestamp/signature and handle."""
    timestamp = request.headers.get('X-Slack-Request-Timestamp')
    if abs(time() - int(timestamp)) > 60 * 5:
        current_app.logger.error('Invalid request timestamp.')
        return make_response("You shall not pass.", 403)

    signature = request.headers.get('X-Slack-Signature')
    SIGNING_SECRET = current_app.config['SLACK_SIGNING_SECRET']

    if not verify_signature(timestamp, signature, SIGNING_SECRET):
        current_app.logger.error('Invalid request signature')
        return make_response("You shall not pass.", 403)

    # Parse the request payload into JSON
    event_data = json.loads(request.data.decode('utf-8'))

    # Echo the URL verification challenge code back to Slack
    if "challenge" in event_data:
        return make_response(
            event_data.get("challenge"), 200, {"content_type": "application/json"}
        )

    # Parse the Event payload and emit the event to the event listener
    if "event" in event_data:
        event_type = event_data["event"]["type"]

        handler = Handle[event_type]
        handler(event_data)

        response = make_response("Success.", 200)
        response.headers['X-Slack-Powered-By'] = 'Architect'
        return response



