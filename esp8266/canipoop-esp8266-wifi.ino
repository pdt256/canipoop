// WEMOS MINI

#include <ESP8266WiFi.h>
#include <FirebaseArduino.h>

#define FIREBASE_HOST "canipoop-4efd0.firebaseio.com"
#define FIREBASE_AUTH ""
#define WIFI_SSID ""
#define WIFI_PASSWORD ""
#define DOOR_DEBOUNCE_DELAY 75
#define interruptPin 4 // This is D2 on the ESP8266
//#define MY_DEBUG

volatile boolean switchStatus = LOW;
uint8_t lastSwitchStatus = 2;
const String DOOR_NUMBER = "2";

void setup() {
  Serial.begin(9600);
  pinMode(interruptPin, INPUT_PULLUP);
  connectToWifi();
  Firebase.begin(FIREBASE_HOST, FIREBASE_AUTH);
}

void loop() {
  uint8_t doorStatus = getSwitchStatus();
    
  if (doorStatus != lastSwitchStatus) {
    #ifdef MY_DEBUG
      Serial.print("Door Value: ");
      Serial.println(doorStatus);
    #endif
    statusChanged();
  }
}

void connectToWifi() {
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);
  #ifdef MY_DEBUG
  Serial.print("connecting");
  #endif
  while (WiFi.status() != WL_CONNECTED) {
    #ifdef MY_DEBUG
    Serial.print(".");
    #endif
    delay(500);
  }
  #ifdef MY_DEBUG
  Serial.println();
  Serial.print("connected: ");
  Serial.println(WiFi.localIP());
  #endif
}

boolean getSwitchStatus() {
  return digitalRead(interruptPin);
}

void statusChanged() {
  delay(DOOR_DEBOUNCE_DELAY);
  uint8_t doorStatus = getSwitchStatus();

  if (doorStatus != lastSwitchStatus) {
    if (doorStatus == HIGH) {
      setDoorStatus(0);
    } else {
      setDoorStatus(1);
    }
 
    lastSwitchStatus = doorStatus;   
  }
}

void setDoorStatus(int doorStatusInt) {
   Firebase.setInt("office1/br" + DOOR_NUMBER + "/isOpen", doorStatusInt);
   #ifdef MY_DEBUG
   if (Firebase.failed()) {
       Serial.print("setDoorOpen Failed");
       Serial.println(Firebase.error());  
       return;
   }
   #endif
}
