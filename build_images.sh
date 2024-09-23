#!/bin/bash
registry="minikube:5000"

# 1. Build base images

# 1.1 Build controller image
pushd controller
    pushd web
    rm -rf dist
    npm install
    npm run build
  popd
  rmdir -f go-http 
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
  docker build -t ${registry}/xtuples_controller:latest -f Dockerfile.xtuples_controller .
  docker push ${registry}/xtuples_controller:latest

popd

# 1.2 Build base agent image
pushd agent
  go generate
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
  docker build -t xtuples_agent:latest -f Dockerfile.xtuples_agent .
popd

# 1.3 Build jobs python package and image
rmdir -rf jobs/bin jobs/lib jobs/pysrc/dist
virtualenv jobs/
  pushd jobs
    source bin/activate
    pip install -q build
    pushd pysrc/
      pushd xtuples_jobs
         echo "Lancement GENRPC"
        ./gengrpc
        echo "GENRPC DONE"
      popd
      python -m build
    popd
    docker build -t xtuples_jobs:latest -f Dockerfile.xtuples_jobs .
  popd
deactivate

# Build exemple service images and push service images to minikube registry
pushd services

  for service in $(ls -d */)
  do 
    pushd $service
      # Get service name from config.yaml
      service_name=$(awk -F: '/^\s*name\s*:/ {print $2; exit}' config.yaml | xargs)
      echo $service_name
      if [ -z $service_name ]; then
        echo "Service name not found in config.yaml"
        exit 1
      fi

      # Write DockerFiles for jobs and agent if it does not exist
      jobs_id=xtuples_${service_name}_jobs
      agent_id=xtuples_${service_name}_agent

      docker_job_file_model="../Dockerfile.xtuples_jobs"
      docker_agent_file_model="../Dockerfile.xtuples_agent"

      docker_job_file="Dockerfile.${jobs_id}"
      docker_agent_file="Dockerfile.${agent_id}"

      if [  ! -f ${docker_job_file} ]; then
        cp ${docker_job_file_model} ${docker_job_file}
        # sed -i "s/<SERVICE_NAME>/${service_name}/g" ${docker_job_file}
      fi
      if [  ! -f docker_agent_file ]; then
        cp ${docker_agent_file_model} ${docker_agent_file}
      fi

      # Build and push service images
      rm -rf bin lib pysrc/dist
      docker build -t ${registry}/${jobs_id}:latest  -f ${docker_job_file}   .
      docker build -t ${registry}/${agent_id}:latest -f ${docker_agent_file} .
      docker push ${registry}/${jobs_id}:latest
      docker push ${registry}/${agent_id}:latest
    popd
  done
popd
