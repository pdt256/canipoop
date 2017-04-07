import time # so we can use "sleep" to wait between actions
import RPi.GPIO as io # import the GPIO library we just installed but call it "io"
import pyrebase
import threading
import os

config = {
  "apiKey": os.environ['FIREBASE-API-KEY'],
  "authDomain": "canipoop-4efd0.firebaseapp.com",
  "databaseURL": "https://canipoop-4efd0.firebaseio.com/",
  "storageBucket": "canipoop-4efd0"
}

door_pin = 23

firebase = pyrebase.initialize_app(config)
auth = firebase.auth()
db = firebase.database()

io.setmode(io.BCM)
io.setup(door_pin, io.IN, pull_up_down=io.PUD_UP)


def door_changed(location, is_open):
    # print("door changed")
    db.child(location).set(is_open)


def pin_callback(channel):
    is_open = io.input(door_pin)
    t = threading.Thread(target=door_changed, args=("office/br1", is_open))
    t.daemon = True
    t.start()

io.add_event_detect(door_pin, io.BOTH, callback=pin_callback)

while True:
    time.sleep(1e6)
