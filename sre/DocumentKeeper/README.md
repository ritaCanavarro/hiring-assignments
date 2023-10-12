# Document keeper - The Guardian of PDFs and PNGs

The Document Keeper is the go to service to have your need of random PDFs and PNGs fulfilled to your hearts contempt.

To ask for a document you just need to access http://localhost:4096/document/{randomNumber} and voila you will get a random document. However, the Keeper is not always very certain where he keeps the forbidden PDF documents so once in a while you might an error message about an unprocessable entity.

## Requirements
If you want to run Document keeper and Dummy PDF or PNG locally, you will need the following tools:

- make
- docker
- minikube

## Local development
Once you have clone the DocumentKeeper repository you will only need to execute a few commands to setup and run it locally. Additionally, you also have to setup its source of truth, the dummy-pdf-or-png service for which an How to is also provided below.

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

### Building and Running Dummy PDF or PNG  - TBR
Once you have all the dependencies working as expected, you will run the Dummy PDF or PNGcontainer by executing the following procedure:

* Start a minikube cluster: `minikube start --kubernetes-version=v1.23.0 --memory=4g` or `minikube start --kubernetes-version=v1.23.0 --memory=4g --driver=hyperv`
* Run `eval $(minikube docker-env)` or `minikube -p minikube docker-env --shell powershell | Invoke-Expression` in order to use the docker daemon inside the minikube cluster
* Run `make build` in order to build the image add it to the minikube cluster
* Go to the folder .\k8scharts\templatesWithValues, which were generated via `helm template . -f values.yaml > templates.yaml` and run the follwoing commands:

`kubectl apply -f .\deployment.yaml`
`kubectl apply -f .\ingress.yaml`
`kubectl apply -f .\service.yaml`
`kubectl apply -f .\vpa.yaml` <- for VPA you must have the CRDs installed

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

### Building and Running Document Keeper - TBR
Once you have all the dependencies working as expected, you will run the Document Keeper container by executing the following procedure:

* Start a minikube cluster: `minikube start --kubernetes-version=v1.23.0 --memory=4g` or `minikube start --kubernetes-version=v1.23.0 --memory=4g --driver=hyperv`
* Run `eval $(minikube docker-env)` or `minikube -p minikube docker-env --shell powershell | Invoke-Expression` in order to use the docker daemon inside the minikube cluster
* Run `make build ` in order to build the image add it to the minikube cluster
* Go to the folder .\k8scharts\templatesWithValues, which were generated via `helm template . -f values.yaml > templates.yaml` and run the follwoing commands:

`kubectl apply -f .\deployment.yaml`
`kubectl apply -f .\ingress.yaml`
`kubectl apply -f .\service.yaml`

These steps will make the app be deployed to the minikube cluster. After that you can watch the logs by running `kubectl logs <container_name>` or the events by doing `kubectl events <container_name>`.

If you want to make requests to the Document Keeper container you will need to run the following command to
port-forward requests to it:

`kubectl get services`
`kubectl port-forward service/documentkeeper 4096:4096`

### Check metrics in Prometheus UI
Add the helm repository and install prometheus with helm chart:

`helm repo add prometheus-community https://prometheus-community.github.io/helm-charts`
`helm repo update`

Go to the directory ./DocumentKeeper/k8scharts/templatesWithValues
`helm install prom prometheus-community/kube-prometheus-stack --values prometheus.yaml`

Apply the serviceMonitor.yaml:
`kubectl apply -f .\serviceMonitor.yaml`

Open the UI in your local browser:
`kubectl get services`
`kubectl port-forward svc/prom-kube-prometheus-stack-prometheus 9090:9090`

### Healtcheck Probing
Go to the directory ./DocumentKeeper/k8scharts/templatesWithValues
`helm install bexp prometheus-community/prometheus-blackbox-exporter --values blackbox.yaml`

Apply the probe.yaml:
`kubectl apply -f .\probe.yaml`

Open the UI in your local browser:
`kubectl get services`
`kubectl port-forward svc/bexp-prometheus-blackbox-exporter 9115:9115`

Make a call via browser to:
`http://localhost:9115/probe?target=http://172.17.0.4:4096/-/ready&module=http_2xx&debug=true`

Go back to the UI and you will be able to see the result of the Probing.

## Considerations - TBD
I have chosen to do the step 1 and 3 of the hiring assignment. For step 1, I have experience with micro services and APIs and even though I am still recent to GO (I only know and worked on-and-off with it for a year and I have never done an API with it) I wanted to develop the service in this language so I could learn more about it, while trying my best to ensure Clean code practices - e.g I learned about Gorilla mux for HTTP routing and HttpTest for mocking HTTP requests. 

As for step 3, I have worked with Make, Dockerfile and Helm charts (more with the last one) and I knew I wanted to have that at least to show a bit of the skills I learned (and am constantly learning) about Docker and Helm. Additionally, I have never worked with CI/CD and GCP as a developer/maintainer but I wanted to challenge myself and show to the team that I am not scared of a challenge and that I will always try my best to learn and put what I am learning into practice. 

As an additionall step I decided to provide the necessary configuration to setup a Prometheus and a Blackbox exporter so we can query the metrics via
Prometheus UI and perform Healthchecks to the Document keeper service, respectively.

## Improvements list
Use Infrastructure as Code, in this case Terraform, to provide the resources needed to run the Document Keeper and see its metrics and logs.

## Feedback
Feedback on improvement points, tips to implement the improvement list or just new ideas that can make this service better, more reliable, secure and performatic are always welcome and will be discussed and iterated. Therefore, feel free to reach out to me either on the CNCF slack channel or via Linkedin. :) 

## Maintainers
| name            | slack            | Channel                            |
|-----------------|------------------|------------------------------------|
| Rita Canavarro  | @rita.canavarro  | Cloud Native Computing foundations |


## Learning resources

Terraform in 15 min - https://www.youtube.com/watch?v=l5k1ai_GBDE
HashiCorp Certified: Terraform Associate 2023 - https://www.udemy.com/course/terraform-beginner-to-advanced/?ranMID=39197&ranEAID=JVFxdTr9V80&ranSiteID=JVFxdTr9V80-jGrjSNVTdCh1rZdy0o78iQ&LSNPUBID=JVFxdTr9V80&utm_source=aff-campaign&utm_medium=udemyads
GitHub Actions Tutorial - Basic Concepts and CI/CD Pipeline with Docker - https://www.youtube.com/watch?v=R8_veQiYBjI
Github Actions to GCP https://docs.github.com/en/actions/deployment/deploying-to-your-cloud-provider/deploying-to-google-kubernetes-engine
Blackbox exporter https://medium.com/cloud-native-daily/blackbox-exporter-to-probe-or-not-to-probe-57a7a495534b