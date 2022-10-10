# About

It's a project to learn kubernetes

### How ?

We've created some micro-services to deploy them on Kubernetes using minikube

### How to deploy ?

Firstly, start minikube :
`minikube start`

After this, deploy the kustomisation file :
`kubectl deploy -k .`

And then, expose the front services on your local machine :
`minikube service kubernetes-front --url`

### Who ?

Remy LEFEBVRE, Thomas NEVIANI, Marwen JAIEM-BLOT, Tristan GUALINI, Alexis BESSAGUET, Pierre Fourny 