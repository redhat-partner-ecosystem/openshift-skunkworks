BUILD_NAMESPACE = buildspace
DEVSPACES_NAMESPACE = devspaces

.PHONY: namespaces
namespaces:
	oc new-project ${DEVSPACES_NAMESPACE}
	oc new-project ${BUILD_NAMESPACE}

.PHONY: install-gitops-operators
install-gitops-operators:
	oc apply -f operators/openshift-pipeline-operator.yaml
	oc apply -f operators/openshift-gitops-operator.yaml
	
.PHONY: install-devspaces-operators
install-devspaces-operators:
	oc apply -f https://raw.githubusercontent.com/redhat-gpte-devopsautomation/gitea-operator/master/catalog_source.yaml
	oc apply -f operators/openshift-devspaces-operator.yaml
	oc apply -f operators/gitea-operator.yaml

.PHONY: config-devspaces
config-devspaces:
	oc apply -f operators/openshift-devspaces-instance.yaml
	oc apply -f operators/gitea-instance.yaml

.PHONY: config-roles
config-roles:
	oc policy add-role-to-user system:image-builder \
		system:serviceaccount:${BUILD_NAMESPACE}:builder \
		--namespace=openshift

.PHONY: install
install: namespaces install-gitops-operators install-devspaces-operators config-devspaces
