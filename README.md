# SwarmMQ - Intelligent Message Routing System

## Overview

SwarmMQ is an intelligent, hierarchical mesh streaming platform that provides high-throughput, low-latency message delivery with adaptive routing capabilities. The MVP implementation focuses on a **single-cluster message router** that demonstrates core distributed systems concepts and routing intelligence.

## MVP Capabilities

The SwarmMQ MVP delivers a working single-cluster system with:

- **High-Performance Message Routing**: Handle 1,000+ messages/second with sub-50ms latency
- **Intelligent Path Selection**: Static hash-based routing with foundation for performance optimization
- **Reliable Delivery**: Exactly-once message delivery with encryption and integrity verification
- **Fault Tolerance**: Graceful handling of single node failures
- **Custom Binary Protocol**: Optimized TCP-based client communication
- **Producer/Consumer Model**: Pull-based consumption with acknowledgments

## Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  Producer       │    │   SwarmMQ        │    │   Consumer      │
│  Client         │───▶│   Cluster        │───▶│   Client        │
│                 │    │  (5-7 nodes)     │    │                 │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

### Core Components

- **Hybrid Nodes**: Each node handles producer ingestion, message routing, and consumer serving
- **Custom Binary Protocol**: Efficient TCP-based communication with length-prefixed framing
- **Hash-Based Routing**: Consistent message distribution using SHA-256 destination hashing
- **In-Memory Storage**: Fast message queues with event sourcing for durability
- **Goroutine Architecture**: One goroutine per client connection for optimal concurrency

## Technology Stack

- **Language**: Go 1.21+ (goroutine-based concurrency, excellent networking)
- **Networking**: Custom TCP binary protocol built on Go's net package
- **Storage**: In-memory with event sourcing (BadgerDB planned)
- **Security**: AES-256-GCM encryption with TLS 1.3 transport
- **Serialization**: Custom binary format optimized for SwarmMQ messages

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Make (optional, for build automation)

### Build

```bash
# Clone and build
git clone <repository-url>
cd swarmMQ
go mod tidy
make build

# Or build manually
go build ./cmd/node
go build ./cmd/producer-client
go build ./cmd/consumer-client
```

### Run a Single Node

```bash
# Start a SwarmMQ node
./node --config configs/mvp-cluster.yaml --node-id node-1 --port 9001

# Check health
curl http://localhost:8080/health
```

### Send Messages

```bash
# Run example producer
./producer-client --node localhost:9001 --destination "test-queue" --message "Hello SwarmMQ"
```

### Consume Messages

```bash
# Run example consumer
./consumer-client --node localhost:9001 --destination "test-queue"
```

## Development

### Project Structure

```
swarmMQ/
├── cmd/                    # Executable applications
│   ├── node/              # Main SwarmMQ node
│   ├── producer-client/   # Example producer
│   └── consumer-client/   # Example consumer
├── internal/              # Private application code
│   ├── node/              # Core node implementation
│   ├── message/           # Message handling and encryption
│   ├── cluster/           # Cluster management
│   ├── client/            # Client libraries
│   └── config/            # Configuration management
├── pkg/                   # Public packages
├── test/                  # Test suites
├── configs/               # Configuration files
└── docs/                  # Documentation
```

### Available Make Targets

```bash
make build      # Build all binaries
make test       # Run all tests with race detection
make clean      # Clean build artifacts
make fmt        # Format all Go code
make vet        # Run static analysis
make all        # fmt + vet + build + test
```

### Running Tests

```bash
# All tests
go test -v -race ./...

# Integration tests
go test -v ./test/integration/

# Benchmarks
go test -bench=. ./...
```

## Configuration

Example cluster configuration (`configs/mvp-cluster.yaml`):

```yaml
cluster:
  name: "mvp-cluster"
  nodes:
    - id: "node-1"
      address: "localhost:9001"
      roles: ["producer", "consumer", "router"]

performance:
  max_message_size: "1MB"
  batch_size: 100
  ack_timeout: "5s"

security:
  tls_cert: "/path/to/cert.pem"
  tls_key: "/path/to/key.pem"
```

## Performance Targets

### MVP Goals
- **Throughput**: 1,000 messages/second per cluster
- **Latency**: <50ms average end-to-end delivery
- **Memory**: <512MB per node under normal load
- **Availability**: 99.9% uptime with single node failure tolerance

### Monitoring

Basic metrics exposed at `http://localhost:8080/metrics`:
- Message throughput (messages/second)
- End-to-end latency distribution
- Node resource utilization
- Error rates and queue depths

## Development Roadmap

### ✅ Phase 1: Foundation (Current MVP)
- Single cluster with 5-7 nodes
- Basic producer-to-consumer delivery
- Static hash-based routing
- Message encryption and validation

### 🚧 Phase 2: Smart Routing (Next)
- Performance-based routing decisions
- Gossip protocol for distributed coordination
- Dynamic failure detection and recovery
- Advanced load balancing algorithms

### 📋 Phase 3: Multi-Cluster
- Hierarchical cluster coordination
- Cross-cluster message routing
- Dynamic scaling and cluster management

### 📋 Phase 4: Production Features
- Advanced monitoring and observability
- Comprehensive operational tooling
- Client SDKs for multiple languages

## Contributing

### Development Workflow

1. Pick a task from the current development phase
2. Create feature branch from main
3. Implement with comprehensive tests
4. Run full test suite: `make all`
5. Submit pull request with clear description

### Code Standards

- Follow Go idioms and conventions
- Maintain test coverage >80%
- Include integration tests for new features
- Document public APIs thoroughly
- Use structured logging for observability

## Learning Resources

SwarmMQ is designed as a learning project for distributed systems concepts:

- **Go Concurrency**: Goroutines, channels, and concurrent patterns
- **Network Programming**: Custom protocols, TCP/UDP handling
- **Distributed Systems**: Consensus, failure detection, routing algorithms
- **Performance Optimization**: Profiling, benchmarking, memory management

## License

[Planned]

## Support

- Documentation: `docs/` directory
- Issues: [Planned]
- Discussions: [Planned]

---

**Note**: This is the MVP implementation focusing on single-cluster deployment. Multi-cluster features and advanced distributed systems capabilities are planned for future phases.