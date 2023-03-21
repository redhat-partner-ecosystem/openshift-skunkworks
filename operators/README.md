# Installation

## Install the Red Hat GitOps and Red Hat Pipelines operator

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

Subscribe to the operator and create a new instance:

```shell
oc new-project devspaces

oc apply -f operators/gitea-operator.yaml
oc apply -f operators/gitea-instance.yaml
```

## Install the Red Hat Dev Spaces operator

Subscribe to the operator and create a new instance:

```shell
oc new-project devspaces

oc apply -f operators/openshift-devspaces-operator.yaml
oc apply -f operators/openshift-devspaces-instance.yaml
```

That's all ...
