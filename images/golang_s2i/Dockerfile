FROM registry.access.redhat.com/ubi8/s2i-base:1-429

ENV NAME=golang \
    GO_MAJOR_VERSION=1 \
    GO_MINOR_VERSION=18 \
    GO_PATCH_VERSION=10 \
    CONTAINER_NAME="rhel8/go-toolset" \
    DATA_ROOT=/opt/app-root/data \
    APP_FILE="main.go" \
    GOOS=linux \
    GOARCH=amd64 \
    PORT=8080

# Define the VERSION environment variable in a separate step, so we can
# re-use the GO_MAJOR_VERSION, GO_MINOR_VERSION and GO_PATCH_VERSION variables. 
ENV VERSION=$GO_MAJOR_VERSION.$GO_MINOR_VERSION.$GO_PATCH_VERSION \
    SUMMARY="Platform for building and running Go Applications" \
    DESCRIPTION="Go Toolset available as a container is a base platform for \
building and running various Go applications and frameworks. \
Go is an easy to learn, powerful, statically typed language in the C/C++ \
tradition with garbage collection, concurrent programming support, and memory safety features."

LABEL summary="$SUMMARY" \
      description="$DESCRIPTION" \
      io.k8s.description="$DESCRIPTION" \
      io.k8s.display-name="Go $VERSION" \
      io.openshift.tags="builder,golang,golang${GO_MAJOR_VERSION}${GO_MINOR_VERSION},rh-golang${GO_MAJOR_VERSION}${GO_MINOR_VERSION},go" \
      name="$CONTAINER_NAME" \
      version="$VERSION"

RUN mkdir -p ${DATA_ROOT}
RUN chown -R 1001:0 ${DATA_ROOT} && \
    fix-permissions ${DATA_ROOT} -P
    
RUN INSTALL_PKGS="go-toolset" && \
    yum install -y --setopt=tsflags=nodocs $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    yum clean all -y

# Copy the custom S2I scripts to $STI_SCRIPTS_PATH.
COPY ./s2i/bin/ $STI_SCRIPTS_PATH

RUN chown -R 1001:0 $STI_SCRIPTS_PATH && chown -R 1001:0 $APP_ROOT

USER 1001
EXPOSE 8080

# Set the default CMD to print the usage of the language image.
CMD $STI_SCRIPTS_PATH/usage