#!/bin/sh
curl -X POST localhost:8080/alter --data-binary @$1
