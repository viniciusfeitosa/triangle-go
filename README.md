# Triangle-GO  [![Build Status](https://travis-ci.org/viniciusfeitosa/triangle-go.svg?branch=master)](https://travis-ci.org/viniciusfeitosa/triangle-go)

This application is to detect the type of triangle.
The app use Go (verion 1.5.2)
The repository have a binary of the application, you can go to Step 3 for use the app
To build the app is necessary install Go and make the group of directories from your GOPATH like the github (github.com/viniciusfeitosa/triangle-go)

## Operation:

1. Start the app with Makefile

2. Input the value of triangle sides like ex: 2,3,4 (The values should be separated by commas)

3. After that you will see the result. Throws errors if you input a wrong value.

## Running:

Inside the directory application...

### Step 1 (Run the tests):

make test

### Step 2 (Building a executable jar):

make build

### Step 3 (Run the application):

make run
