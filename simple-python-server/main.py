
from time import sleep
from flask import Flask

app = Flask(__name__)


@app.route('/shivam')
def hello_world():
    sleep(2)
    return {"sh": "shi"}


if __name__ == "__main__":
    app.run()