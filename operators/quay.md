# Installation

## Install Red Hat Quay

### Prerequisite: OpenShift Data Foundations

Red Hat Quay uses some Red Hat OpenShift Data Foundation (ODF) APIs. To install the ODF operator, follow the instructions [here](https://access.redhat.com/documentation/en-us/red_hat_openshift_data_foundation/4.12).

**Important:** Create a default `StorageSystem` and wait until it is operational !

### Install the Red Hat Quay operator

Deploy the initial config params for the Quay operator installation:

```shell
oc create secret generic -n openshift-operators --from-file config.yaml=./operators/quay/config.yaml quay-init-config-bundle
```

Subscribe to the Quay and Quay Bridge operators:

```shell
oc apply -f operators/openshift-quay-operator.yaml
```

Deploy the default Quay instance:

```shell
oc create -n openshift-operators -f operators/quay/quay-registry.yaml
```

**Note:** it might take a couple of minutes until all the Quay services are up-and-running. Check `Workload/Pods` in the OpenShift web console before moving on to the next configuration step.

#### Create the Quay admin user

To create a default `quayadmin` user, make a call to Quay's management API:

```shell
# get the quay api endpoint
oc get route quay-registry-quay -n openshift-operators
```

```shell
# export the route URL
export QUAY=quay-registry-.... 
```

```shell
# create the user
curl -X POST -k  "https://$QUAY/api/v1/user/initialize" \
    --header 'Content-Type: application/json' \
    --data '{ "username": "quayadmin", "password":"quaypass123", "email": "quayadmin@example.com", "access_token": true}'

```

**Important:** save the access token somewhere, it is never shown again !


### Configure the Quay Bridge

Log into Quay as the admin user (quayadmin@quaypass123).

Create a `new organization` (gitopshq).

Create a new `application` within the organization (gitopshq-bridge). 

Create an `Access Token` for the application. Give it full rights to the organization.

Create a secret with the above access token:

```shell
oc create secret -n openshift-operators generic quay-integration-secret \
    --from-literal=token=<access_token>
```

Deploy the Quay Bridge instance:

**Important:** update `operators/quay/quay-integration.yaml` with the actual Quay instance's endpoint URL.

```shell
oc apply -n openshift-operators -f operators/quay/quay-integration.yaml
```

See [https://github.com/quay/quay-bridge-operator](https://github.com/quay/quay-bridge-operator) for more details on the operator.


### Deploy the Quay Container Security Operator

```shell
oc apply -f operators/openshift-quay-security-operator.yaml
```

That's all ...

See [https://github.com/quay/container-security-operator](https://github.com/quay/container-security-operator) for more details on the operator.


## Usefull commands

```shell
# list installed operators
oc get csv

# list available operators
oc get packagemanifests -n openshift-marketplace

```
