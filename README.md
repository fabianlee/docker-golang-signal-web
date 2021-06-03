# Summary
Golang http web server running by default on port 8080 that is intended for testing OS level signal catching

Image is based on busybox:1.32.1-glibc, is about ~11Mb because it takes advantage of multi-stage building

sending OS level signals examples:
  sudo kill -SIGINT <PID>
  sudo kill -SIGUSR1 <PID>

# Environment variables

* PORT - listen port, defaults to 8080
* APP_CONTEXT - base context path of app, defaults to '/'

# Prerequisites
* make utility (sudo apt-get install make)

# Makefile targets
* docker-build (builds image)
* docker-run-fg (runs container in foreground, ctrl-C to exit)
* docker-run-bg (runs container in background)
* k8s-apply (applies deployment to kubernetes cluster)
* k8s-delete (removes deployment on kubernetes cluster)
