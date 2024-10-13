# DevOpsifying-GO-Application

## Description,
    DevOpsifying the GO web Appication by containerizing the App, Deploying in the Kubernetes Cluster, automating the CICD Pipeline with GitOps Practice

## pre-requisite,


## Environment Setup,
- (optional) Install Go language for running in local environment
  https://go.dev/doc/install
- Install Docker in your local machine for building the image and running Containers on top of your base or host machine
  ```bash
     docker version
  ```
- create K8s Cluster

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
  
- Create kubernetes Cluster and K8s Deployment, Service, and Ingress for deploying the application in k8s cluster
-  

