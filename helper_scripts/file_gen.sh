#!/bin/bash

blocks=$1
OFN=$2
BS=$3
dd if=/dev/urandom iflag=fullblock of=$OFN bs=$BS count=$blocks
