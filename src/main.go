package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "github.com/pborman/uuid"
)

// can be overriden with -ldflags
var Version = "n/a"
var BuildTime = "n/a"
var BuiltBy = "n/a"

func StartWebServer() {

    // handlers
    http.HandleFunc("/healthz", handleHealth)
    http.HandleFunc("/shutdown", handleShutdown)

    // APP_CONTEXT defaults to root
    appContext := getenv("APP_CONTEXT","/")
    log.Printf("app context: %s", appContext)
    http.HandleFunc(appContext, handleApp)

    port := getenv("PORT","8080")
    log.Printf("Starting web server on port %s", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        panic(err)
    }

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type","application/json")
    fmt.Fprintf(w, "{ \"health\":\"ok\" }" )
}

func handleApp(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type","text/plain")

    // print main hello message
    fmt.Fprintf(w, "Hello, %s\n", "World")

    // 'Host' header is promoted to Request.Host field and removed from Header map
    fmt.Fprintf(w, "Host: %s\n", provideDefault(r.Host,"empty"))
}

// provide default for value
func provideDefault(value,defaultVal string) string {
  if len(value)==0 { 
    return defaultVal
  }
  return value
}
// pull from OS environment variable, provide default
func getenv(key, fallback string) string {
    value := os.Getenv(key)
    if len(value) == 0 {
        return fallback
    }
    return value
}
// invoked by GET request (not SIGNAL)
func handleShutdown(w http.ResponseWriter, r *http.Request) {
    log.Printf("About to abruptly exit")
    os.Exit(0)
}

func recieveSignalLoop(sigc chan os.Signal) {
  for {
    fmt.Println("SIGNAL waiting to receive...")
    s := <-sigc
    fmt.Printf("SIGNAL received!!! ")
    fmt.Println(s)
  }
}

func main() {

    fmt.Println("Version: ",Version)
    fmt.Println("BuildTime: ",BuildTime)
    fmt.Println("BuiltBy: ",BuiltBy)

    uuidWithHyphen := uuid.NewRandom()
    fmt.Println(uuidWithHyphen)

    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc,
        syscall.SIGHUP,
        syscall.SIGTERM,
        syscall.SIGUSR1,
        syscall.SIGUSR2,
        syscall.SIGQUIT)
    go recieveSignalLoop(sigc)

    StartWebServer()
}

