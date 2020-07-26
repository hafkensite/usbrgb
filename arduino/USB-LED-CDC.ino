#include <DigiCDC.h>
#include <Adafruit_NeoPixel.h>

#define LED_PIN   1
#define LED_COUNT 2
Adafruit_NeoPixel strip(LED_COUNT, LED_PIN, NEO_GRB + NEO_KHZ800);

void setup() {
  strip.begin();
  strip.setPixelColor(0, 0, 128, 0);
  strip.setPixelColor(1, 0, 0, 128);
  strip.show();
  SerialUSB.begin();
  SerialUSB.delay(100);
  strip.setPixelColor(0, 0, 0, 0);
  strip.setPixelColor(1, 0, 0, 0);
  strip.show();

}

void loop() {
  int avail = SerialUSB.available();
  if (avail == 3) {
    char r = SerialUSB.read();
    char g = SerialUSB.read();
    char b = SerialUSB.read();
    strip.setPixelColor(0, r, g, b);
    strip.setPixelColor(1, r, g, b);
    strip.show();
  } else if (avail == 6) {
    char r = SerialUSB.read();
    char g = SerialUSB.read();
    char b = SerialUSB.read();
    strip.setPixelColor(0, r, g, b);
    r = SerialUSB.read();
    g = SerialUSB.read();
    b = SerialUSB.read();
    strip.setPixelColor(1, r, g, b);
    strip.show();
  } else if (avail != 0) {
    strip.setPixelColor(0, 255, 0, 0);
    strip.show();
    while (avail > 0) {
      SerialUSB.read();
      avail--;
    }
    strip.setPixelColor(0, 0, 0, 0);
    strip.show();
  }

  SerialUSB.delay(100);               // keep usb alive // can alos use SerialUSB.refresh();
}
