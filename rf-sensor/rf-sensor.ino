#include<RF24.h>
#include<LowPower.h>

const int switchInputPin = 2;
const int powerSamplePin = 5;
boolean switchStatus;
boolean lastSwitchStatus;

// ce, csn pins
RF24 radio(9, 10);

void setup() {
  pinMode(switchInputPin, INPUT);
  pinMode(powerSamplePin, OUTPUT);
  digitalWrite(powerSamplePin, LOW);

  switchStatus = digitalRead(switchInputPin);
  
  radio.begin();
  radio.setPALevel(RF24_PA_MAX);
  radio.setChannel(0x76);
  radio.openWritingPipe(0xF0F0F0F0E1LL);
  radio.enableDynamicPayloads();

  getSwitchStatus();
  sendSwitchStatus();
}

void loop() {
  getSwitchStatus();

  if (switchStatus != lastSwitchStatus) {
    sendSwitchStatus();
  }

  LowPower.powerDown(SLEEP_1S, ADC_OFF, BOD_OFF);
}

boolean getSwitchStatus() {
  digitalWrite(powerSamplePin, HIGH);
  switchStatus = digitalRead(switchInputPin);
  digitalWrite(powerSamplePin, LOW);
}

void sendSwitchStatus() {
  boolean sendWasSuccessful;
  
  radio.powerUp();
  if (switchStatus == HIGH) {
    sendWasSuccessful = radio.write("H", 1);
  } else {
    sendWasSuccessful = radio.write("L", 1);
  }
  radio.powerDown();

  if (sendWasSuccessful) {
    lastSwitchStatus = switchStatus;
  }
}

