FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder

ENV PKG=/go/src/github.com/flacatus/operator-tests/
WORKDIR ${PKG}

# compile test binary
COPY . .
RUN make

FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

COPY --from=builder /go/src/github.com/flacatus/operator-tests/bin/che-operator-test-harness che-operator-test-harness

ENTRYPOINT [ "/che-operator-test-harness" ]