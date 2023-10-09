# Document keeper - The Guardian of PDFs and PNGs

The Document Keeper is the go to service to have your need of random PDFs and PNGs fulfilled to your hearts contempt.

To ask for a document you just need to access http://localhost:4096/document/{randomNumber} and voila you will get a random document. However, the Keeper is not always very certain where he keeps the forbidden PDF documents so once in a while you might an error message about an unprocessable entity.

## Requirements
If you want to run Document keeper locally, you will need the following tools:

- make
- docker
- minikube

## Local development
Once you have clone the DocumentKeeper repository you will only need to execute a few commands to setup and run it locally.

### Building Document Keeper

To generate a package with all the dependencies, run the following command:

```bash
make build
```

### Building and Running Document Keeper
Once you have all the dependencies working as expected, you will run the Document Keeper container by executing the following procedure:

* Start a minikube cluster: `minikube start --kubernetes-version=v1.23.0 --memory=4g`
* Run `eval $(minikube -p minikube docker-env)` in order to use the docker daemon inside the minikube cluster
* Run `make docker-build ` in order to build the image add it to the minikube cluster

These steps will make the app be deployed to the minikube cluster. After that you can watch the logs by running `kubectl -n default logs <container_name>`

If you want to make requests to the Document keeper container you will need to run the following command to
port-foward requests to it:


## Improvements list

## Feedback

## Maintainers
| name            | slack            | Channel                            |
|-----------------|------------------|------------------------------------|
| Rita Canavarro  | @rita.canavarro  | Cloud Native Computing foundations |


