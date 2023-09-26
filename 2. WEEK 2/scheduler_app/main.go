package main
import (
    "fmt"
    "scheduler_app/scheduler"
)
func main() {
    fmt.Println("Mulai Aplikasi Scheduler")

    // Jadwalkan FibonacciTask
    scheduler.ScheduleTask(scheduler.FibonacciTask)

    // Jadwalkan AkronimTask
    scheduler.ScheduleTask(scheduler.AkronimTask)

    // Jadwalkan ScrabbleScoreTask
    scheduler.ScheduleTask(scheduler.ScrabbleScoreTask)
}
