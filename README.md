# XTuples
A data-driven job synchronization service inspired by the tuple space model.

# XTuples
XTuples is a data-driven job synchronization service inspired by the tuple space model. 

## Features

- **Agent Image**: Deploys agents for monitoring and managing tasks.
- **Jobs**: Python-based package for background job processing.
- **Services**: Set of modular services, each with its own configuration.
- **Controller Image**: Manages orchestration tasks.
- **xtuple Helmchart**: Helm Chart for KeyDB uses ClusterIP.


## Getting Started

### Prerequisites

- Minikube or Docker registry running at `minikube:5000`
- Kubernetes cluster setup (local or cloud-based)
- Docker
- Python and virtualenv for jobs

### Setup

1. Clone the repository:
    ```bash
    git clone https://github.com/ditrit/XTuples.git
    cd XTuples
    ```

2. Make sure you are running uild Docker images:
    Run the `build_images.sh` script to build and push images:
    ```bash
    ./build_images.sh
    ```

3. Deploy to Kubernetes:
    Follow the instructions for deploying to your Kubernetes cluster using Helm charts (see below).

## Usage

### Deploying KeyDB with Helm

## TODO!
- XTuples uses Helm charts for managing Kubernetes deployments.
- Customize the chart values and deploy using the following steps.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

