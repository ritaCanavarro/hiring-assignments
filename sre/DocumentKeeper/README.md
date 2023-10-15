# Document keeper — The Guardian of PDFs and PNGs

The Document Keeper is the go-to service to have your need of random PDFs and PNGs fulfilled to your hearts' contempt.

To ask for a document you just need to access http://34.118.53.100:4096/document/{randomNumber} and voilà you will get a random document. However, the Keeper is not always very certain where he keeps the forbidden PDF documents so once in a while you might an error message about an unprocessable entity.

## Access the Services running in GKE
- DocumentKeeper: http://34.118.53.100:4096/
- Prometheus UI: http://34.118.53.140:9090/
- Blackbox Exporter: http://34.116.192.232:9115/

## Get DocumentKeeper Metrics
Make a request to http://34.118.53.100:4096/metrics

## Requirements
If you want to run Document keeper and Dummy PDF or PNG locally with Kubernetes, you will need the following tools:

- make
- docker
- minikube

Or if you just want to test the images of the services use the provided Docker Compose with `docker compose up -d` 

## Local development
Once you have cloned the DocumentKeeper repository you will only need to execute a few commands to set up and run it locally. Additionally, you also have to set up its source of truth, the dummy-pdf-or-png service for which a How-to is also provided below.

### Building Dummy PDF or PNG

To generate a package with all the dependencies, run the following command:

```bash
make build
```

### Run Dummy PDF or PNG

To run Dummy PDF or PNG use the following command:

```bash
make run
```

### Building and Running Dummy PDF or PNG
Once you have all the dependencies working as expected, you will run the Dummy PDF or PNG container by executing the following procedure:

* Start a minikube cluster: 
    - Ubuntu: `minikube start --kubernetes-version=v1.23.0 --memory=4g` 
    - Windows: `minikube start --kubernetes-version=v1.23.0 --memory=4g --driver=hyperv`
* Run the commands below in order to use the docker daemon inside the minikube cluster:
    - Ubuntu: `eval $(minikube docker-env)`
    - Windows: `minikube -p minikube docker-env --shell powershell | Invoke-Expression` 
* Run `make build` in order to build the image add it to the minikube cluster
* Go to the folder .\k8scharts\templatesWithValues, which were generated via `helm template . -f values.yaml > templates.yaml` and run the following commands:

`kubectl apply -f .\deployment.yaml`
`kubectl apply -f .\ingress.yaml`
`kubectl apply -f .\service.yaml`

These steps will make the app be deployed to the minikube cluster. After that you can watch the logs by running `kubectl logs <container_name>` or the events by doing `kubectl events <container_name>`.

If you want to make requests to the Dummy PDF or PNG container you will need to run the following command to
port-forward requests to it:

`kubectl get services`
`kubectl port-forward service/dummypdforpng 3000:3000`


### Building Document Keeper

To generate a package with all the dependencies, run the following command:

```bash
make build
```

### Run Document Keeper

To run Document Keeper use the following command:

```bash
make run
```

### Building and Running Document Keeper
Once you have all the dependencies working as expected, you will run the Document Keeper container by executing the following procedure:

* Start a minikube cluster: 
    - Ubuntu: `minikube start --kubernetes-version=v1.23.0 --memory=4g` 
    - Windows: `minikube start --kubernetes-version=v1.23.0 --memory=4g --driver=hyperv`
* Run the commands below in order to use the docker daemon inside the minikube cluster:
    - Ubuntu: `eval $(minikube docker-env)`
    - Windows: `minikube -p minikube docker-env --shell powershell | Invoke-Expression` 
* Run `make build ` in order to build the image add it to the minikube cluster
* Go to the folder .\k8scharts\templatesWithValues, which were generated via `helm template . -f values.yaml > templates.yaml` and run the following commands:

`kubectl apply -f .\deployment.yaml`
`kubectl apply -f .\ingress.yaml`
`kubectl apply -f .\service.yaml`

These steps will make the app be deployed to the minikube cluster. After that you can watch the logs by running `kubectl logs <container_name>` or the events by doing `kubectl events <container_name>`.

If you want to make requests to the Document Keeper container you will need to run the following command to
port-forward requests to it:

`kubectl get services`
`kubectl port-forward service/documentkeeper 4096:4096`

### Check metrics in Prometheus UI
Add the helm repository and install Prometheus with helm chart:

`helm repo add prometheus-community https://prometheus-community.github.io/helm-charts`
`helm repo update`

Go to the directory ./DocumentKeeper/k8scharts/templatesWithValues
`helm install prom prometheus-community/kube-prometheus-stack --values prometheus.yaml`

Apply the serviceMonitor.yaml:
`kubectl apply -f .\serviceMonitor.yaml`

Open the UI in your local browser:
`kubectl get services`
`kubectl port-forward svc/prom-kube-prometheus-stack-prometheus 9090:9090`

### Health check Probing
Go to the directory ./DocumentKeeper/k8scharts/templatesWithValues and install Blackbox Exporter:
`helm install bexp prometheus-community/prometheus-blackbox-exporter --values blackbox.yaml`

Apply the probe.yaml:
`kubectl apply -f .\probe.yaml`

Open the UI in your local browser:
`kubectl get services`
`kubectl port-forward svc/bexp-prometheus-blackbox-exporter 9115:9115`

Make a call via browser to:
`http://localhost:9115/probe?target=http://172.17.0.4:4096/-/ready&module=http_2xx&debug=true`

Go back to the UI and you will be able to see the result of the probing.

## Considerations
In a first iteration, I had chosen to do the step 1 and 3 of the hiring assignment. For step 1, I have experience with microservices and APIs and even though I am still recent to GO (I only know and worked on-and-off with it for a year and I have never done an API with it) I wanted to develop the service in this language, so I could learn more about it, while trying my best to ensure clean code practices — e.g I learned about Gorilla mux for HTTP routing and HttpTest for mocking HTTP requests. 

As for step 3, I have worked with Make, Dockerfile and Helm charts (more with the last one) and I knew I wanted to have that, at least to demonstrate the skills I learned (and am constantly learning) about Docker and Helm. Additionally, I have never worked with CI/CD and GCP as a developer/maintainer, but I wanted to challenge myself and show to the team that I am not scared of a challenge and that I will always try my best to learn and put what I am learning into practice. For CI/CD after some research I decided to use GitHub actions and GCP (because I know it is the provider the team works with and why not learn more about it and try something new?).

As an additional step, I decided to provide the necessary configuration to set up a Prometheus and a Blackbox exporter, so we can query the metrics via Prometheus UI and perform Health checks to the Document keeper service, respectively.

After having done all the aforementioned, because I was even more motivated due to the fact that I had managed to put the CI/CD pipeline working and the services running in GKE I decided to just go for it and attempt the step 2.
For step 2, I used Terraform and the material mentioned in the section "Learning resources" to set up and automate the creation/management of the infrastructure of Document keeper. This step still needs some improvements as mentioned in the "Improvements list" section, and it would be where I would focus more if I had more time for the assignment.

## Improvements list
In this section, I will be listing the improvements I would like to perform on this assignment:

1) Define a way to install the helm charts of both Document keeper and Dummy PDF or PNG via helm install and then have a control flow that would allow it to install the charts in a first run and then update them on subsequent runs. Additionally, also apply the control flow to the Prometheus and Blackbox charts.

2) Find a way to install the VPA CRDs in order to be able to deploy the defined VPA's.

3) Fix the permission errors when trying to pull the service images in the new GKE cluster, created via Terraform.

4) Define a release strategy for the images with version management.

## Feedback
Feedback on improvement points, tips to implement the improvement list or just new ideas that can make this service better, more reliable, secure and performative are always welcome and will be discussed and iterated upon. Therefore, feel free to reach out to me either on the CNCF Slack channel or via LinkedIn. :) 

## Maintainers
| name            | Slack            | Channel                            |
|-----------------|------------------|------------------------------------|
| Rita Canavarro  | @rita.canavarro  | Cloud Native Computing foundations |


## Learning resources

Terraform in 15 min - https://www.youtube.com/watch?v=l5k1ai_GBDE

GitHub Actions Tutorial - Basic Concepts and CI/CD Pipeline with Docker - https://www.youtube.com/watch?v=R8_veQiYBjI

GitHub Actions to GCP https://docs.github.com/en/actions/deployment/deploying-to-your-cloud-provider/deploying-to-google-kubernetes-engine

Blackbox exporter https://medium.com/cloud-native-daily/blackbox-exporter-to-probe-or-not-to-probe-57a7a495534b

GCP Terraform tutorial https://developer.hashicorp.com/terraform/tutorials/gcp-get-started/google-cloud-platform-change

GCP Terraform registry https://registry.terraform.io/providers/hashicorp/google/latest

CPU - https://home.robusta.dev/blog/stop-using-cpu-limits