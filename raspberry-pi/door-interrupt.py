import RPi.GPIO as GPIO
import time
import spidev
import pyrebase
import threading
import os

token = ""

config = {
    "apiKey": token,
    "authDomain": "canipoop-4efd0.firebaseapp.com",
    "databaseURL": "https://canipoop-4efd0.firebaseio.com/",
    "storageBucket": "canipoop-4efd0"
}

firebase = pyrebase.initialize_app(config)
db = firebase.database()
last_is_open = 3


def door_changed(location, is_open):
    print("door changed:", location)
    lastChange = int(time.time())
    db.child(location).update({"isOpen": is_open, "lastChange": lastChange, "lastUpdate": lastChange}, token)


GPIO.setmode(GPIO.BCM)

door_pin = 23
GPIO.setup(door_pin, GPIO.IN, pull_up_down=GPIO.PUD_UP)  # activate input with PullUp


def pin_callback(channel):
    is_open = GPIO.input(door_pin)
    t = threading.Thread(target=door_changed, args=("office1/br1", is_open))
    t.daemon = True
    t.start()


GPIO.add_event_detect(door_pin, GPIO.BOTH, callback=pin_callback)

pin_callback(0)

while True:
    time.sleep(1e6)