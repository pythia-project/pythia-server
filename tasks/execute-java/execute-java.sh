#!/bin/sh
PATH=/usr/lib/jvm/java-8-openjdk-i386/bin:$PATH
mkdir /tmp/work
cp /task/Pre.java /tmp/work/Pre.java
cd /tmp/work
javac Pre.java
java Pre
