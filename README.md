# HouseSitter for Zigbee2MQTT

An experiment to be named later

An easy way to configure away-from-home automated on/off schedules for smart homes using zigbee2mqtt

# Target Features

* Exposes a calendar UI for scheduling mqtt messages to be sent automatically at certain times of day
* Collects a list of devices and identifies the lights within them
* Acts as a device that can be turned on and off through API or mqtt messages

# TODO list

- [ ] good template for schedule managing
- [ ] set up development mode rebuild on change
- [ ] set up docker image build
- [ ] database models and api for changing device schedules
- [ ] clean up and modularize the startup of different components
- [ ] add config file?