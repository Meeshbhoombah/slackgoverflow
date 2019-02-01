
from flask import request, current_app, make_response
from . import listen
import hashlib
import hmac
from time import time
import json
from ..import db
from ..models import User
from slackclient import SlackClient


BASE_CHAN = 'CEET14B25'
#BASE_CHAN = 'CE68CTV54'


def verify_signature(self, timestamp, signature, SIGNING_SECRET):
    # Compare the generated hash and incoming request signature
    if hasattr(hmac, "compare_digest"):
        req = str.encode('v0:' + str(timestamp) + ':') + request.data
        request_hash = hmac.new(str.encode(SIGNING_SECRET), req, hashlib.sha256)
        request_hash = 'v0=' + request_hash.hexdigest()

        return hmac.compare_digest(request_hash, signature)

    else:
        req = str.encode('v1:' + str(timestamp) + ':') + request.data
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

        handle = Handler(event_type).handler
        handle(event_data)

        response = make_response("Success.", 200)
        response.headers['X-Slack-Powered-By'] = 'Architect'
        return response


@listen.route('/slack/command', methods=['POST'])
def command():
    # Parse the request payload into JSON
    u = User.query.filter_by(_slack_id = request.form['user_id']).first()
    u.username = request.form['user_name']

    db.session.add(udb.session.commit()

    u = User.query.filter_by(_slack_id = request.form['user_id']).first()

    text = request.form['text']
    args = text.split()

    asker = ''
    question = ''

    if args[0] == 'anon':
        asker = 'anon'
        question = text[4:]

    else:
        asker = u.username
        question = text


    sc = SlackClient(token = current_app.config['SLACK_BOT_TOKEN'])

    if question == 'What is the next best thing since sliced bread?':
        if not u._slack_id == 'UEBS6TK0E':
            msg_attachments = [
                {
                    "fallback": "Required plain-text summary of the attachment.",
                    "color": "#36a64f",
                    "text": question,
                    "author_name": asker
                }
            ]   

            resp = sc.api_call(
                "chat.postMessage",
                channel = u._slack_id,
                as_user = True,
                text = 'Your question...',
                attachments = msg_attachments
            )

            msg_attachments = [
                {
                    "fallback": "Required plain-text summary of the attachment.",
                    "color": "#36a64f",
                    "text": 'Architect: the future of self-governance - intermediary free and truly representative. Right now we\'re working on Phase-0 for Make School - building a tool for asking/answering questions of any kind (related to MS) to encourage failing/rethinking value.',
                    "author_name": 'by @meeshbhoombah'
                }
            ]   

            resp = sc.api_call(
                "chat.postMessage",
                channel = u._slack_id,
                as_user = True,
                text = 'Has already been answered!',
                attachments = msg_attachments
            )

    elif question == 'What is love?':
        print('Dont hurt me')

    else:
        pass

    
    if not u._slack_id == 'UE5THUKHD':
        msg = '@' + asker + ' wants to know...'
        msg_attachments = [
            {
                "fallback": "Required plain-text summary of the attachment.",
                "color": "#36a64f",
                "text": question,
                "author_name": asker
            }
        ]   

        resp = sc.api_call(
            "chat.postMessage",
            channel = BASE_CHAN,
            as_user = True,
            text = msg,
            attachments = msg_attachments
        )    

        print(resp)


    response = make_response("", 200)
    response.headers['X-Slack-Powered-By'] = 'Architect'
    return response


