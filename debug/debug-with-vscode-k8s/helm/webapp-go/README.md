# Helm Chart for webapp-go

Build the Docker image:

```bash
# eval $(minikube docker-env)

docker build -t webapp-go .
```

Install the Helm Chart in a Kubernetes environment (e.g., Minikube):

```bash
# helm lint helm/webapp-go

helm package helm/webapp-go

# helm install --dry-run --debug --name webapp-go webapp-go-0.1.0.tgz

helm install --name webapp-go webapp-go-0.1.0.tgz

# helm del --purge webapp-go
```

Log output should indicate when the Delve server is running:

```bash
$ kubectl logs webapp-go-webapp-go

API server listening at: [::]:2345
```

Use port forwarding to access the webapp and Delve port locally:

```bash
kubectl port-forward service/webapp-go-webapp-go 8080 2345
```

The application will not work until you connect a Delve debugger. This project
includes an example `/.vscode/launch.json` file that can be used with Visual Studio Code.
Within Visual Studio Code, launch debugging via the "Remote debug" launch profile.

**NOTE:** You will need to define breakpoints before connecting the debugger.

Visit the webapp:

[http://localhost:8080/edit/test](http://localhost:8080/edit/test)