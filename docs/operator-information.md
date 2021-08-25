## Introduction
This document helps in understanding the development process of the operator and some links to the concepts on which this operator is based.

### Concepts
Below are the concepts along with the links which will help in understanding the documents:
1. [CRDs](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/)
2. For extending K8S APIs using custom controllers, there are following tools which can help in generating scaffolding code:
   a. [Kubebuilder](https://book.kubebuilder.io/)
   b. [Operator SDK](https://sdk.operatorframework.io/docs/building-operators/golang/quickstart/)
   For building this operator, Operator SDK was made use of.

### Climon controller details
Following are the links which will help in understanding the core functionality of the climon controller
1. The CRD for climon is provided [here](../config/crd/bases/monitoring.soi.dev_monitoringconfigurations.yaml)
2. Controller responsible for reconcilation of the CRD is [here](../controllers/monitoringconfiguration_controller.go)

### Steps for modification 
1. In case of update in the CRD schema, following commands should be executed:
   '''bash
   make manifests
   make generate
   '''
2. Update the controller accordingly.
3. Create a pull request as per [doc](../CONTRIBUTING.md)
