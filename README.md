## GoPushServer
Simple push notification server made with go

## How to build
git clone git@github.com:acidsound/gopushserver.git
cd gopushserver
git submodule init
git submodule update

## Editing configuration
for Android : copy config.example.json to config.json, and set APIKey and serverURL.
for iOS : (todo)

## for Google App Engine
Set application name in app.yaml (default: gopushserver)

## Packages
go-gcm : http://github.com/googollee/go-gcm.git
apns : http://github.com/pranavraja/apns.git