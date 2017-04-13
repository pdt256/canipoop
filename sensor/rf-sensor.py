import RPi.GPIO as GPIO
from lib_nrf24 import NRF24
import time
import spidev
import pyrebase
import threading

config = {
  "apiKey": os.environ['FIREBASE-API-KEY'],
  "authDomain": "canipoop-4efd0.firebaseapp.com",
  "databaseURL": "https://canipoop-4efd0.firebaseio.com/",
  "storageBucket": "canipoop-4efd0"
}

firebase = pyrebase.initialize_app(config)
auth = firebase.auth()
db = firebase.database()
last_is_open = 3

def door_changed(location, is_open):
    print("door changed")
    db.child(location).set(is_open)


GPIO.setmode(GPIO.BCM)

pipes = [[0xE8, 0xE8, 0xF0, 0xF0, 0XE1], [0xF0, 0xF0, 0xF0, 0xF0, 0xE1]]

radio = NRF24(GPIO, spidev.SpiDev())
radio.begin(0, 17)

radio.setPayloadSize(32)
radio.setChannel(0x76)
radio.setDataRate(NRF24.BR_1MBPS)
radio.setPALevel(NRF24.PA_MAX)

radio.setAutoAck(True)
radio.enableDynamicPayloads()
radio.enableAckPayload()

radio.openReadingPipe(1, pipes[1])
radio.printDetails()
radio.startListening()

while True:

    while not radio.available(0):
        time.sleep(1/100)

    receivedMessage = []
    radio.read(receivedMessage, radio.getDynamicPayloadSize())
    #print("Received: {}".format(receivedMessage))

    #print("Translating our received Message into unicode characters...")
    string = ""

    for n in receivedMessage:
        if (n >= 32 and n <= 126):
            string += chr(n)
    print("Received: {}".format(string))


    if string == "H":
        is_open = 0
    elif string == "L":
        is_open = 1
    else:
        continue

    if is_open != last_is_open:
        last_is_open = is_open
        t = threading.Thread(target=door_changed, args=("office/br2", is_open))
        t.daemon = True
        t.start()
