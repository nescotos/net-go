package main

type Handler func(*Client, interface{})

type CheckHandler func(string) (Handler, bool)
