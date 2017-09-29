// WEMOS MINI
#include <ESP8266WiFi.h>
#include <FirebaseArduino.h>

#define FIREBASE_HOST "canipoop-4efd0.firebaseio.com"
#define WIFI_SSID "farm-guest"
#define WIFI_PASSWORD "FarmingRocks!"
#define DOOR_DEBOUNCE_DELAY 75
#define DOOR_PIN 4 // This is D2 on the ESP8266
//#define MY_DEBUG

volatile boolean switchStatus = LOW;
uint8_t lastSwitchStatus = 2;
const String DOOR_NUMBER = "2";

void setup() {
  Serial.begin(9600);
  pinMode(DOOR_PIN, INPUT_PULLUP);
  connectToWifi();
  Firebase.begin(FIREBASE_HOST);
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
  return digitalRead(DOOR_PIN);
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
