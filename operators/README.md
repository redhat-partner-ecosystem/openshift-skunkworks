# Installation

## Install the GitOps and Pipelines operators

Subscribe to the operators:

```shell
oc apply -f operators/openshift-gitops-operator.yaml
oc apply -f operators/openshift-pipeline-operator.yaml
```

Verify that the default GitOps instance is up-and-running:

```shell
oc get pods -n openshift-gitops
```

## Install the Gitea operator

Install the catalog source:

```shell
oc apply -f https://raw.githubusercontent.com/redhat-gpte-devopsautomation/gitea-operator/master/catalog_source.yaml
```

Subscribe to the operators:

```shell
oc apply -f operators/gitea-operator.yaml
oc apply -f operators/gitea-instance.yaml
```

## Install Red
That's all ...


## Usefull commands

```shell
# list installed operators
oc get csv

# list available operators
oc get packagemanifests -n openshift-marketplace

```
