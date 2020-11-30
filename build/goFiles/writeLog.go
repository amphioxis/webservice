package main

import ("fmt")

func writeLog(l Log) (int) {

  fmt.Println("Request number", l.numberOfRequests, ":")
  fmt.Println("")
  fmt.Println("Time of request:", l.date)
  fmt.Println("HTTP-Status:", l.status)
  fmt.Println("Request:")
  fmt.Println(l.request)
  return l.numberOfRequests
}
