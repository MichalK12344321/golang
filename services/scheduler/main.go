package main

import(
  "github.com/MichalK12344321/golang/services/scheduler/scheduler"
)

func main() {
  // place for custom getopt
  schedulerService := scheduler.NewService()
  schedulerService.Start()
  // place for return code
}
