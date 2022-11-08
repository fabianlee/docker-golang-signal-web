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
* docker-test-fg (runs container in foreground, ctrl-C to exit)
* docker-run-bg (runs container in background)
* k8s-apply (applies deployment to kubernetes cluster)
* k8s-delete (removes deployment on kubernetes cluster)

# Create Github release

Must have [Go Lang](https://fabianlee.org/2022/10/29/golang-installing-the-go-programming-language-on-ubuntu-22-04/) compiler and [Github CLI](https://fabianlee.org/2022/04/21/github-cli-tool-for-repository-operations/) installed on local host as prerequisite.

Github Actions will automatically build OCI image based on this new semantic version.

```
./create_new_gh_release.sh
```


# Deleting tag

```
todel=v1.0.1

# delete local tag, then remote
git tag -d $todel && git push origin :refs/tags/$todel
```

# Deleting release

```
todel=v1.0.1

# delete release and remote tag
gh release delete $todel --cleanup-tag -y

# delete local tag
git tag -d $todel
```
