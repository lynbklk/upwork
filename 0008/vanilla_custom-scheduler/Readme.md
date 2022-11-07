# Custom Scheduler



## Submission Notes

### Done

The finished solution is:

- Written in Go 1.18
- Implementing initiating configuration with k8s cluster and scheduling unbounds pods to a random node.
- The solution comes with a Docker file used for building and pushing the docker file into a private or public registry such as Dockerhub.
- The solution comes with manifests used to deploy the scheduler into the kubernetes cluster
- Simple and concise


### Features included

The golang code contains 1 file with following logic:

- **Initate Scheduler** The first step is to initiate the structure Scheduler with kubernetes client configuration as well as the name of the scheduler.
- **Run Scheduler** This function is used to watch unbound pods and loop this event channel to bind a pod into a random node.

The functions used: 
- **NewScheduler** Function used to initiate a new scheduler with kubernetes client configuration as well as the name of the cluster
- **Run** Function used to watch unbound pods and loop this event channel to bind a pod into a random node.
- **GetRandomNode** Function used to find a random node in the cluster
- **BindPod** Bind pod into a given node.

### Dockerfile

Based on Alpine image this image is used to build the binary of the golang program and to expose it. How to build and push the image?
```
docker build  -t ${YOUR_IMAGE}:${VERSION}  . 
```

To push into the registry:
```
docker push  
```

### Manifests

The deployment for the custom scheduler is using iliasnaamane/scheduler:5.6 image. Image that has been built and pushed into the Docker registry.

- 1: Deployment.yml: Used to deploy the custom scheduler and matching it with a service account.
- 2: Rbac.yml: Service Account used for the deployment to grant him access the same ClusterRole as the default kube scheduler.
- 3: example.yml:  Example of deployment to see if the custom scheduler is working perfectly.


### How to test your project?
- **Apply the Rbac** First step is to apply the ClusterRole & Service Account using ``` kubectl apply -f rbac.yml ``` 
- **Deploy your custom scheduler** To deploy you have to run ``` kubectl apply -f deployment.yml ``` 
- **Deploy example**  Deploy pod example using ``` kubectl apply -f example.yml ``` 
- **Check** The example pod should be scheduled correctly and the custom scheduler should provide logs ``` kubectl logs custom-scheduler-xxxxxx``` 


