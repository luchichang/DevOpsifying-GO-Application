# DevOpsifying-GO-Application

## Description,
    DevOpsifying the GO web Appication by containerizing the App, Deploying in the Kubernetes Cluster, automating the CICD Pipeline with GitOps Practice

## Objective, 
- Containerize the application and store the image artifactory in the docker hub repository
- Create k8s deployment, service, ingress in k8s cluster
- Create __Ingress Controller__
- Create __EKS__ Cluster
- Create __Helm__ Chart.
- automate the Workflow with __CICD Pipeline__

## pre-requisite,
- MultiStage Docker build
- Distroless Image
- K8s Object ( Deployment, Replica set, pods)
- k8s service (ClusterIp, NodePort, LoadBalancer)
- k8s Ingress & Ingress Controller
- Helm & Helm Chart


## Environment Setup,
- (optional) Install Go language for running in local environment
  https://go.dev/doc/install
- Install Docker in your local machine for building the image and running Containers on top of your base or host machine
  ```bash
     docker version
  ```
- install __Kubectl__ by following the instructions in below link
  https://docs.aws.amazon.com/eks/latest/userguide/install-kubectl.html
- install __eksctl__ by following the instructions in below link
  https://eksctl.io/installation/
- install __awscli__ by following the instructions in below link
  https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html
- install the nginx Ingress Controller
  https://kubernetes.github.io/ingress-nginx/deploy/
- Install __HelmCLI__ by following the instructions from the below URL
  https://helm.sh/docs/intro/install/

## Steps,
- First build & run the Application locally. so, that DevOps Engineer can get the nuances of how the web app is working.
  - run the Unit test to check that application is working as expected and assure the will not break
    ```bash
       go test
    ```
    (NOTE: if you get **OK** followed by Module Name then the application is running perfectly ok) 
  - install the necessary dependencies for the application
    ```bash
      go mod download
    ```
  - run the go app locally
    ```bash
      go run main.go
    ```
    ![image](https://github.com/user-attachments/assets/91eae8f3-6b63-4d81-bdb0-e58af381ddf4)

    after successful execution of this step web app is now accessible in the local host ***http://localhost:8080***
  - build the binary file
    ```bash
      go build -o main .
    ```
  - Execute/Run the build file
    
    ```bash
      ./main
    ```
- Containerize the GO app
  - Building the image from the Docker file
    ```bash
      docker build -t go-app .
    ```
             [or]
    pull the image from repository
    ```bash
       docker pull dinesht0006/go-app
    ```
  - running the container from the build image
    ```bash
      docker run -itd -p 8080:8080 go-app
    ```
    Hurrah 🥳 now go web app is accessible from the <public ip>:8080
    
- Create kubernetes Cluster and K8s Deployment, Service, and Ingress for deploying the application in k8s cluster
  - create __AWS EKS__ cluster
     ```bash
       eksctl create cluster --name prod-cluster --region ap-south-1
    ```
         similarly for deleting the EKS cluster execute the below command
    ```bash
       eksctl delete cluster --name <cluster-name> --region <created-region>
    ```
  - ensure EKS cluster is created
    ```bash
       eksctl get cluster
    ```
  - create object by applying the deployment.yaml file to the kubectl
    ```bash
       kubectl apply -f deployment.yaml
    ```
  - create service by applying the service.yaml file to the kubectl
    ```bash
       kubectl apply -f service.yaml
    ```
  - check the pod is running and try to access it in the nodeport
    ```bash
       kubectl get pod
    ```
        for viewing the service 
    ```bash
       kubectl get svc
    ```
  - create Ingress by applying the ingress.yaml file to the kubectl
    ```bash
       kubectl apply -f ingress.yaml
    ```
  - ensure ingress is created
    ```bash
       kubectl get ing
    ```
  - (pre requisite) inorder to make ingress in effect first ingress controller need to be created
   
  - ensure ingress controlled is mapped to the ingress
    ```bash
       kubectl edit ing <ingress-podname> -n <namespace>
    ```
  - get the Ingress controller Load balancer IP address
    ```bash
      nslookup <loadbalancer address>
    ``` 
NOTE: Ingress controller maps the ingress using the __ingressClassName__ attribute
  - create custom domain name and map it to the __/etc/hosts__ file
     ```bash
        <ingress Controllers load balancer IP>  <Custom Domain Name>
     ```
Note: Custom Domain Name mapped to the __/etc/hosts__ file is accessible only locally

NOTE: Pods can be debuged within the cluster 
```bash
  kubectl run -i --tty --rm debug --image=busybox -- sh  
```
     
-  Create Helm Charts
  - create the Helm Chart
    ```bash
       helm create <chart name>
    ```
    
  - create the release name and run the helm chart
    ```bash
       helm install <release-name> <chart path>
    ``` 
  - uninstall the release
    ```bash
       helm uninstall <release-name>
    ```
-  CI Pipeline Setup

