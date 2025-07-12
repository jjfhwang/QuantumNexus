# QuantumNexus: A Distributed Quantum Circuit Simulator

QuantumNexus is a high-performance, distributed quantum circuit simulator written in Go. Designed for researchers and developers pushing the boundaries of quantum computation, QuantumNexus enables the simulation of complex quantum algorithms at scale, exceeding the limitations of single-node simulators. By leveraging a distributed architecture, it distributes the computational workload across multiple machines, drastically reducing simulation time and allowing for exploration of larger, more intricate quantum circuits.

This project aims to provide a robust and flexible platform for simulating quantum circuits of varying sizes and complexities. It offers a comprehensive set of quantum gates and measurement operations, allowing users to define and execute custom quantum algorithms. The distributed nature of QuantumNexus allows researchers to tackle problems that were previously computationally infeasible, opening new avenues for exploration in quantum algorithm design, quantum error correction, and quantum materials simulation. Furthermore, QuantumNexus aims to be easily extensible, allowing developers to integrate new simulation techniques, quantum hardware models, and performance optimization strategies.

QuantumNexus isn't just about raw performance; it also emphasizes user-friendliness and developer experience. The Go programming language provides inherent concurrency and efficiency, while the project's architecture promotes modularity and maintainability. We strive to provide clear documentation, intuitive APIs, and comprehensive examples to facilitate the adoption and use of QuantumNexus by both experienced quantum computing researchers and newcomers to the field. This project aims to democratize access to high-performance quantum circuit simulation, empowering a wider community to contribute to the advancement of quantum technology.

Key Features:

*   **Distributed Simulation Engine:** Employs a message-passing architecture (using gRPC) to distribute the quantum state vector across multiple compute nodes. This allows for simulations to scale beyond the memory limitations of a single machine. The state vector distribution strategy uses a block-cyclic distribution, allowing efficient data access and communication.
*   **Quantum Gate Library:** Includes a comprehensive set of standard quantum gates, such as Hadamard, Pauli-X, Pauli-Y, Pauli-Z, CNOT, and parameterized rotation gates (Rx, Ry, Rz). These gates are implemented using optimized matrix multiplication routines, leveraging Go's built-in support for complex numbers.
*   **Custom Gate Definition:** Provides an API for defining and incorporating custom quantum gates into simulations. This allows users to explore novel quantum operations and tailor the simulator to their specific research needs. The API requires users to specify the matrix representation of the gate in the computational basis.
*   **Measurement and Collapse:** Supports projective measurements on individual qubits or subsets of qubits. The simulator accurately simulates the quantum state collapse resulting from measurement, updating the state vector according to the measurement outcome. Measurement outcomes are returned as probabilities for each possible outcome.
*   **Noise Modeling:** Integrates basic noise models, including depolarization and bit-flip errors, to simulate the effects of imperfections in quantum hardware. Users can configure the error rates for each noise model, allowing for realistic simulations of noisy intermediate-scale quantum (NISQ) devices.
*   **Circuit Definition Language:** Utilizes a simple, human-readable language (likely YAML or JSON) for defining quantum circuits. This language allows users to specify the sequence of gates and measurements to be applied to the quantum state. The circuit definition is parsed and compiled into an internal representation suitable for distributed execution.
*   **Checkpointing and Restart:** Offers the ability to checkpoint the simulation state at regular intervals. This allows users to restart simulations from a previous point in time, which is crucial for long-running simulations or when debugging complex circuits. Checkpoints are stored on disk in a binary format.

Technology Stack:

*   **Go:** The primary programming language, chosen for its performance, concurrency support, and strong standard library.
*   **gRPC:** Used for inter-process communication (IPC) between the master node and worker nodes in the distributed simulation engine. This provides a robust and efficient mechanism for distributing the computational workload.
*   **Protocol Buffers:** Employed for defining the message formats used in gRPC communication, ensuring efficient serialization and deserialization of data.
*   **YAML/JSON:** Utilized for defining quantum circuits in a human-readable format.
*   **BLAS/LAPACK (Optional):** Potentially integrated for optimized linear algebra routines, further improving the performance of the quantum gate implementations. Can be implemented using a Go BLAS/LAPACK library, or by wrapping a C implementation.

Installation:

1.  Ensure you have Go installed (version 1.18 or later) and properly configured. Verify this by running `go version` in your terminal.
2.  Clone the QuantumNexus repository from GitHub:
    `git clone https://github.com/jjfhwang/QuantumNexus.git`
3.  Navigate to the project directory:
    `cd QuantumNexus`
4.  Build the project:
    `go build ./cmd/quantumnexus-master`
    `go build ./cmd/quantumnexus-worker`
5.  (Optional) Build the example client:
    `go build ./cmd/quantumnexus-client`

Configuration:

The following environment variables can be configured to customize the behavior of QuantumNexus:

*   `QN_MASTER_ADDRESS`: The address (host:port) of the master node. Default: `localhost:50051`.
*   `QN_WORKER_ADDRESS`: The address (host:port) of a worker node. Default: `localhost:50052`.
*   `QN_NUM_WORKERS`: The number of worker nodes to use in the distributed simulation. Default: `1`.
*   `QN_MAX_QUBITS`: The maximum number of qubits that can be simulated. Default: `20`. Increasing this value will require more memory.

These environment variables can be set before running the master and worker nodes. For example:

`export QN_MASTER_ADDRESS="192.168.1.100:50051"`

Usage:

1.  Start the master node:
    `./quantumnexus-master`
2.  Start the worker nodes. Each worker node needs to know the address of the master node:
    `./quantumnexus-worker --master-address="localhost:50051"`
3.  (Optional) Run the example client to submit a quantum circuit for simulation:
    `./quantumnexus-client --circuit="path/to/circuit.yaml"`

The `circuit.yaml` file should contain a description of the quantum circuit in the YAML format. The format should define the number of qubits, the initial state, and a sequence of quantum gates and measurements. Refer to the documentation for the specific syntax and available gate definitions.

Example Circuit YAML:

# num_qubits: 2
# initial_state: [0, 0]
# gates:
#   - type: H
#     target: 0
#   - type: CNOT
#     control: 0
#     target: 1
#   - type: MEASURE
#     target: [0, 1]

Contributing:

We welcome contributions to QuantumNexus! Please follow these guidelines:

*   Fork the repository and create a branch for your feature or bug fix.
*   Write clear and concise commit messages.
*   Include unit tests for your code.
*   Submit a pull request with a detailed description of your changes.

License:

This project is licensed under the MIT License. See the [LICENSE](https://github.com/jjfhwang/QuantumNexus/blob/main/LICENSE) file for details.

Acknowledgements:

We would like to acknowledge the contributions of the open-source community to the development of quantum computing tools and libraries, which have inspired and informed the design of QuantumNexus.