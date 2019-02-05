#!/bin/sh

mockery -name=Registry -dir ../plugin -output .
mockery -all -dir ../config -output .
