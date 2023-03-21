# NGINX example

Builds a simple Nginx containert using the [Source-to-image](https://docs.openshift.com/container-platform/4.12/cicd/builds/build-strategies.html#builds-strategy-s2i-build_build-strategies-docker) build strategy.

## Create a project/namespace

```shell
oc new-project nginx-example
```

*Note:* `deployment.yaml` expects this namespace, i.e. if you want to use a different namespace, change the deployment configuration accordingly.

## Build the container

```shell
oc apply -f examples/nginx/imagestream.yaml
oc apply -f examples/nginx/build.yaml
```

*Note:* The build configuration expects the *source code* in folder `examples/nginx/public`.

## Deploy the container

```shell
oc apply -f examples/nginx/deployment.yaml
```

## Expose the container to the internet

```shell
oc apply -f examples/nginx/service.yaml
oc apply -f examples/nginx/route.yaml
```