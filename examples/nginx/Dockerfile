FROM registry.access.redhat.com/ubi9/nginx-120

# Add application sources to a directory where the assemble script expects them
# and set permissions so that the container runs without root access
ADD --chown=1001:0 public /tmp/src

# Run the assemble script to install the dependencies
RUN /usr/libexec/s2i/assemble

# Run script uses standard ways to run the application
CMD /usr/libexec/s2i/run