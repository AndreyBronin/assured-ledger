#!/usr/bin/env bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
NUM_DISCOVERY_NODES=5
NUM_DISCOVERY_NODES=${NUM_DISCOVERY_NODES:-5}
KUBECTL=${KUBECTL:-"kubectl"}
USE_MANIFESTS=${USE_MANIFESTS:-"local"}
INSOLAR_IMAGE=${INSOLAR_IMAGE:-"insolar/assured-ledger:latest"}
ARTIFACTS_DIR=${ARTIFACTS_DIR:-"/tmp/insolar"}
set -x
check_dependencies() {
  command -v docker >/dev/null 2>&1 || {
    echo >&2 "docker is required. Aborting."
    exit 1
  }
  command -v kubectl >/dev/null 2>&1 || {
    echo >&2 "kubectl is required. Aborting."
    exit 1
  }
  check_docker_images
  check_ingress_installation
}

# Delete this after image templating will be done, and images will be in insolar hub
check_docker_images() {
  if [ "$(docker images "$INSOLAR_IMAGE" -q)" = "" ]; then
    echo >&2 "make sure you made 'make docker-build'"
    exit 1
  fi
}

check_ingress_installation() {
  if [ "$($KUBECTL get pods -n ingress-nginx -l app.kubernetes.io/name=ingress-nginx | grep -c Running)" = "0" ]; then
    $KUBECTL apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-0.31.0/deploy/static/provider/cloud/deploy.yaml
  fi
}

delete_ingress() {
  $KUBECTL delete -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-0.31.0/deploy/static/provider/cloud/deploy.yaml
}

run_network() {
  $KUBECTL apply -k "$DIR/$USE_MANIFESTS/" 2>&1 || {
    echo >&2 "kubectl apply failed. Aborting."
    exit 1
  }
}

wait_for_complete_network_state() {
  #set +x
  ready=0
  for try in {0..60}; do
    if [ "$($KUBECTL -n insolar get po bootstrap -o jsonpath="{.status.phase}")" = "Succeeded" ]; then
      ready=1
      break
    fi
    sleep 1s
  done

  if [ $ready = 0 ]; then
    echo "bootstrap was not started"
    $KUBECTL -n insolar describe pods
    exit 1
  fi
  echo "bootstrap completed"

  ready=0
  for try in {0..30}; do
    if [ "$($KUBECTL -n insolar exec -i deploy/pulsewatcher -- bash -c 'pulsewatcher -c /etc/pulsewatcher/pulsewatcher.yaml -s' | grep 'READY' | grep -c -v 'NOT')" = "1" ]; then
      echo "network is ready!"
      ready=1
      break
    else
      echo "network is not ready"
      sleep 2
    fi
  done

  if [ $ready = 0 ]; then
    echo "network was not started"
    $KUBECTL -n insolar describe pods
    exit 1
  fi
  set -x
}

copy_bootstrap_config_to_temp() {
  mkdir -p "$ARTIFACTS_DIR"
  $KUBECTL -n insolar get cm bootstrap-yaml -o jsonpath='{.data.bootstrap\.yaml}' >"$ARTIFACTS_DIR/bootstrap.yaml"
}

check_dependencies
echo "Starting insolar"
run_network
wait_for_complete_network_state
copy_bootstrap_config_to_temp
echo "Insolar started"
set +x
