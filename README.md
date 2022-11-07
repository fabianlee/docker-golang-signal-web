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


# Github Actions will construct OCI image for semantic tag

```
# show latest tags
git tag
git describe --tags

# get latest semantic version tag, construct patch+1
semantic_version=$(git describe --tags | grep -Po '^v[0-9]*.[0-9]*.[0-9]*')
major_minor=$(echo "$semantic_version" | cut -d'.' -f1-2)
patch=$(echo "$semantic_version" | cut -d'.' -f3)
((patch++))

# push new semantic version tag
newtag="${major_minor}.${patch}"
echo "old version: $semantic_version new_version: ${newtag}"
git commit -a -m "changes for new tag $newtag"
git tag $newtag && git push origin $newtag
```

# Deleting tag

```
tagtodel=v1.0.1
# delete locally
git tag -d $tagtodel
# delete remotely
git push origin :refs/tags/$tagtodel
```


