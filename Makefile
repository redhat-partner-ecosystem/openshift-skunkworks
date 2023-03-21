BUILD_NAMESPACE = buildspace
DEVSPACES_NAMESPACE = devspaces

.PHONY: namespaces
namespaces:
	oc new-project ${DEVSPACES_NAMESPACE}
	oc new-project ${BUILD_NAMESPACE}

.PHONY: gitops-operators
gitops-operators:
	oc apply -f operators/openshift-pipeline-operator.yaml
	oc apply -f operators/openshift-gitops-operator.yaml
	

.PHONY: devspaces-operators
gitops-operators:
	oc apply -f operators/openshift-devspaces-operator.yaml
	oc apply -f operators/openshift-devspaces-instance.yaml
	oc apply -f operators/gitea-operator.yaml
	oc apply -f operators/gitea-instance.yaml

.PHONY: install
install: install_tasks install_pipelines
	oc policy add-role-to-user system:image-builder \
		system:serviceaccount:${BUILD_NAMESPACE}:builder \
		--namespace=openshift
	oc apply -f image-builder/builder.yaml
	oc apply -f pipelines/config/pipelines.yaml
	oc apply -f pipelines/config/rolebindings.yaml

