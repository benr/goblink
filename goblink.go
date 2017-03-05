package main

import (
  "fmt"
  "time"
  "os"
  import "github.com/stianeikeland/go-rpio"
)

func main() {

  fmt.Println("Starting blinker...");
  err := rpio.Open();
  if (err !+ nil){
    fmt.Println(err);
    os.Exit(1);
  }

  defer rpio.Close();
  
  pin := rpio.Pin(14); // GPIO14, Physical Pin 8 (Use GND on Pin 6)
  pin.Output();
  pin.PullDown();

  for {
    time.Sleep(1000 * time.Millisecond); 
    pin.Toggle();
  } 

}
