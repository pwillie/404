# 404 not found
Docker Golang container to respond with http 404 (not found) to every request

Motivation for this service is to run inside a Kubernetes cluster and provide the default 404 container for ingress. 

Container will listen on port 8080 and will respond with http 404 to every request.

Also contains a container status endpoint at /status that will respond with http 200.
