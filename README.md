# Summary
Golang http web server running by default on port 8080 that is intended for testing OS level signal catching

Image is based on busybox:1.32.1-glibc, is small (~11Mb) because it takes advantage of multi-stage building

docker hub: https://hub.docker.com/r/fabianlee/docker-golang-signal-web

# Example of sending OS level signals

```
  sudo kill -SIGINT <PID>
  sudo kill -SIGUSR1 <PID>
```

# Environment variables

* PORT - listen port, defaults to 8080
* APP_CONTEXT - base context path of app, defaults to '/'


# Makefile targets
* docker-build (builds image)
* docker-run-fg (runs container in foreground, ctrl-C to exit)
* docker-run-bg (runs container in background)
* k8s-apply (applies deployment to kubernetes cluster)
* k8s-delete (removes deployment on kubernetes cluster)


# Pushing new container image using Github Actions

```
newtag=v1.0.0; git tag $newtag && git push origin $newtag
git commit -a -m "new OCI image built by Github Actions $newtag" && git push
```

# Deleting tag

```
tagtodel=v1.0.1
# delete locally
git tag -d $tagtodel
# delete remotely
git push origin :refs/tags/$tagtodel
```


