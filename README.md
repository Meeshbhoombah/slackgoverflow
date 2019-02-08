# Slackoverflow
- Students have trouble with programming bugs
    + *Stack Overflow?* SO is purposefully daunting to post a question on, and
      questions can be answered more quickly/thorougly if someone understands
      the context in which it exists (better to ask at MS if MS curriculum)
    + *TAs?* But they're constantly busy, have to have set hours, and can be 
      difficult to find. 
    + *Instructors?* See above, and sometimes a person may have a question that
      they feel would make them appear as lacking x in front of an intructor
    + *Peers?* But you have to bother them, and they're still learning, so the
      answers they provide may not be based in accurate or thorough 
      understanding
    + *Google?* But sometimes the resources online are not adapted for someone 
      who is trying to build an understanding and assumes pre-requisite 
      knowledge.
- Students have questions on policy regarding/procedure at Make School
- Most people have similar questions (assumption) 
    + For some tutorials, since students are following the same reference
      material, they often run into the same issues
    + TAs/RAs/Staff/Instructors repeat a lot of effort and lose time
- All Students/Intructors/Staff are on Slack
- Slack can be accessed from anywhere/anytime with an Internet connection

`slackoverflow` - a Stack Overflow-esque channel for the Make School Product
College Slack

## Introducing `slackoverflow`...
**Ask questions in the `slackoverflow` channel w/** `/ask`
- Mark a solution/answer by reacting with a ✅
- Questions are encouraged to be thorough:
    + Include steps covered, references, code, etc.
    + Protip: provide more context to your question with a Gist link
- Ask questions anonymously: `/ask anon "[question]"`

**Recieve Staff and Instructor verified answers**
- Staff/Instructors are invited to verify answers (in their domain) by reacting
  with a ✅

**Reply to any question's thread with answers/comments**
- Answers are enouraged to be thorough:
    + Comprehensive, adressing a wide audience, following up, examples, etc.

### Benefits
Moves towards the dissolution of issues stated in intro +:

**Those asking questions...**
- Save time and effort
    + Async discussion - no need to... 
        * Set time to meet
        * Wait for office hours
        * Hope that someone is awake at 4:00 AM when you really need help
- Ask socially discrediting questions
    + Eg:
        * How does someone use a list in python? (as a senior)
        * Why did we have a Ben Carson movie night?

**Those answering questions...**
- Save time and effort...
    + Avoid answering the same questions over-and-over again
    + Make better use of 1-on-1s/office hours
- Practice technical discussion 

## User Stories
[View the Project Here](https://github.com/orgs/archproj/projects/5).

### Drops
- OP of question 
    + is rewarded for... 
        * **facilitating discussion:**
            - _+1 Drop_ for each reply to the thread (except their own)
        * **highlighting content/policy/procedure that is unclear**
            - _+1 Drop_ for each ✅ reaction from Staff/Instructors
- OP of elected answer
    + is rewarded for... 
        * **supporting the original poster of the question:**        
            - _+1 Drop_ for answering the question or Staff/Intructor 
              verification (cannot recieve +1 for both)
        * **providing greater understanding**
            - _+1 Drop_ for each previous answer verified by the question 
              OP/Staff/Instructors
        * **answering quickly**
            - _+1 Drop/(Total number of 20 minute intervals since question 
              posted)_

#### Benefits
- Incentivize asking/answering/commenting on questions in `slackoverflow`
- Incentivize answering questions quickly

### Planned
- Mechanisms to prevent behavior which defies Make School community values/
  hurts `slackoverflow`
    + Spam - clogs the system
    + Duplicated questions
        * Same question filteration system
        * Merit-based ability unlocks
- Notifications
- "Support" by reacting with ❤️ to indicate that the reactor also had the 
  particular question/answer
    + Benefits:
        * Gague which content needs further development to communicate ideas to 
          students

## Getting Started
These instructions will get you a copy of the project up and running on your 
local machine for development and testing purposes. See deployment for notes on 
how to deploy the project on a live system.

### Prerequisites
- Postgres SQL `11`
- `virtualenv` `16.0.0`
- Git `2.17.1`
- Python `3.7.0`

### Installing
Clone this repository.
```
$ git clone [URL] 
```

Navigate to the current directory.
```
$ cd make/archie/
```

Create a `python3` virtual environment with `virtualenv`. For instructions on 
using Docker, view the `README` in the directory above. 
```
$ virtualenv -p python3 venv
```

Install dependencies, located in `requirements.txt`, with `pip`.
```
(venv) $ pip install -r requirements.txt
```

## Running Tests
Archie's test suite is located in `/test`.

Archie leverages Flask's built in integration of the 
[click](https://click.palletsprojects.com/en/7.x/)
command-line interface to automate testing. Visit 
[the Flask documentation](http://flask.pocoo.org/docs/0.12/cli/) 
for more information.

Simply run the `test` command.
```
$ flask test
test_app_exists (test_deploy.ServerRunningTestCase) ... 
...
```

## Deploying
Assuming Postgres is installed, start the server in the background. It is
recommended to use the Desktop application, but the same can be accomplished
via the command-line.
```
$ pg_ctl -D /usr/local/var/postgres start
```
To later stop the server:
```
$ pg_ctl -D /usr/local/var/postgres stop
```

Activate the virtual environment and set the **required** environment variables.
```
. venv/bin/activate
(venv) bash-3.2$: source .env
```

Set up the database using the `click` built-in for Flask. Do `flask --help` to
view all commands. First run the shell then create the database using the
injected variables.
```
$ flask shell
Python 3.7.0 (default, Jun 29 2018, 20:14:27) 
...
>>> db.create_all()
```

Then migrate and configure User Roles. 
```
$ flask deploy
```

Now you can run the server with `flask run`.
```
$ flask run
 * Serving Flask app "archie.py" (lazy loading)
...
```

