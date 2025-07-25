# SwarmMQ Comprehensive Architecture Document

## Table of Contents
1. [System Overview](#system-overview)
2. [MVP Architecture](#mvp-architecture)
3. [Post-MVP Evolution](#post-mvp-evolution)
4. [Core Components](#core-components)
5. [Project Structure](#project-structure)
6. [Implementation Phases](#implementation-phases)
7. [Technical Specifications](#technical-specifications)
8. [Performance Targets](#performance-targets)

## System Overview

SwarmMQ is an intelligent, hierarchical mesh streaming platform designed for high-throughput, low-latency message delivery with adaptive routing and self-optimization capabilities. The system uses a distributed network of specialized nodes that learn and adapt routing decisions based on performance metrics.

**Core Philosophy**: Start simple with a single cluster, prove the concept, then scale to multi-cluster hierarchical mesh architecture.

## MVP Architecture

### MVP Scope
- **Single cluster deployment** with 5-7 routing nodes
- **Target throughput**: 1,000 messages/second
- **Basic producer-to-consumer delivery** with guaranteed delivery
- **Simple static routing** with round-robin load balancing
- **Message encryption** (mandatory)
- **Pull-based consumption** with acknowledgments

### MVP Components

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  Producer       │    │   SwarmMQ        │    │   Consumer      │
│  Client         │───▶│   Cluster        │───▶│   Client        │
│                 │    │  (5-7 nodes)     │    │                 │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

### MVP Node Architecture

**Single Node Type - Hybrid Nodes**
For MVP simplicity, each node handles multiple roles:
- Accept messages from producers
- Route messages to appropriate destinations
- Store messages for consumers
- Handle consumer pull requests

```
┌─────────────────────────────────────────────┐
│                SwarmMQ Node                 │
├─────────────────────────────────────────────┤
│  Producer Interface  │  Consumer Interface  │
│  - Accept messages   │  - Store messages    │
│  - Validate payload  │  - Handle pulls      │
│  - Encrypt data      │  - Send ACKs         │
├─────────────────────────────────────────────┤
│           Message Routing Engine            │
│  - Route selection   │  - Load balancing    │
│  - Health checking   │  - Circuit breaker   │
├─────────────────────────────────────────────┤
│              Local Storage                  │
│  - Message queue     │  - Routing table     │
│  - Consumer offsets  │  - Node metadata     │
└─────────────────────────────────────────────┘
```

### MVP Message Flow

1. **Producer sends message** → Random node in cluster
2. **Node encrypts and validates** → Assigns message ID
3. **Routing decision** → Hash-based destination selection
4. **Message delivery** → Store in destination node's queue
5. **Consumer pulls** → Retrieve from assigned node
6. **Acknowledgment** → Confirm delivery, cleanup message

### MVP Technology Stack

**Core Runtime**
- **Language**: Go (excellent concurrency, performance, deployment)
- **Message Transport**: gRPC (type-safe, efficient, built-in load balancing)
- **Storage**: BadgerDB (embedded, fast key-value store)
- **Encryption**: AES-256-GCM with TLS 1.3 transport

**Supporting Tools**
- **Configuration**: YAML + Viper
- **Logging**: Structured logging with zerolog
- **Metrics**: Prometheus with basic dashboards
- **Testing**: Go testing framework + Testify

## Post-MVP Evolution

### Phase 2: Smart Routing (Multi-Node Types)

**Specialized Node Types**
```
Producer Nodes ──┐
                 ├─▶ Routing Nodes ──▶ Consumer Nodes
Load Balancer ───┘
```

- **Producer Nodes**: Accept and pre-process messages
- **Routing Nodes**: Intelligent path selection and optimization
- **Consumer Nodes**: Message storage and delivery to consumers

**Smart Routing Features**
- Performance metric collection (latency, throughput, error rates)
- Dynamic route optimization based on historical data
- Adaptive load balancing with weighted round-robin
- Circuit breaker patterns for failure handling

### Phase 3: Multi-Cluster Hierarchy

**Hierarchical Architecture**
```
        Cluster Coordinator Layer
              /     |     \
    Cluster A   Cluster B   Cluster C
   (5-7 nodes) (5-7 nodes) (5-7 nodes)
```

**New Components**
- **Cluster Coordinators**: Inter-cluster routing and coordination
- **Gossip Protocol**: Distributed routing table updates
- **Consistent Hashing**: Automatic cluster assignment
- **Cross-cluster Load Balancing**: Geographic and logical distribution

### Phase 4: Production Features

- **Event Sourcing System**: Complete state recovery capabilities
- **Dynamic Scaling**: Automatic cluster creation/destruction
- **Advanced Monitoring**: Distributed tracing and anomaly detection
- **Client SDKs**: Multiple language support with connection pooling

## Core Components

### Message Structure

**MVP Message Format**
```go
type Message struct {
    ID          string    `json:"id"`           // UUID for deduplication
    Payload     []byte    `json:"payload"`      // Encrypted data
    Destination string    `json:"destination"`  // Topic/queue name
    Priority    int       `json:"priority"`     // 1-10 (10 highest)
    Timestamp   time.Time `json:"timestamp"`    // Creation time
    ProducerID  string    `json:"producer_id"`  // Source identification
}
```

**Post-MVP Enhancements**
```go
type SmartMessage struct {
    Message                    // Embed basic message
    RouteHistory []string     `json:"route_history"` // Path taken
    TTL         time.Duration `json:"ttl"`           // Time to live
    Retries     int          `json:"retries"`       // Attempt count
    Checksum    string       `json:"checksum"`      // Integrity verification
}
```

### Routing Algorithm

**MVP: Static Hash-Based Routing**
```go
func (c *Cluster) selectDestinationNode(message *Message) *Node {
    hash := sha256.Sum256([]byte(message.Destination))
    index := binary.BigEndian.Uint64(hash[:8]) % uint64(len(c.nodes))
    return c.nodes[index]
}
```

**Post-MVP: Performance-Based Routing**
```go
func (r *SmartRouter) selectOptimalPath(message *Message) []*Node {
    candidates := r.getHealthyNodes(message.Destination)
    return r.optimizeRoute(candidates, message.Priority)
}
```

### Storage Layer

**MVP: Simple Key-Value Storage**
- **Message Storage**: `message_id` → `encrypted_message`
- **Consumer Offsets**: `consumer_id` → `last_processed_message_id`
- **Routing Table**: `destination` → `responsible_node_id`

**Post-MVP: Event Sourcing**
- **Event Log**: Immutable sequence of all state changes
- **Snapshots**: Periodic state checkpoints for faster recovery
- **Vector Clocks**: Distributed consistency and ordering

## Project Structure

```
swarmMQ/
├── cmd/
│   ├── node/                 # Node executable
│   ├── producer-client/      # Example producer
│   └── consumer-client/      # Example consumer
├── internal/
│   ├── node/                 # Core node implementation
│   │   ├── server.go
│   │   ├── router.go
│   │   ├── storage.go
│   │   └── metrics.go
│   ├── message/              # Message handling
│   │   ├── message.go
│   │   ├── encryption.go
│   │   └── validation.go
│   ├── cluster/              # Cluster management
│   │   ├── cluster.go
│   │   ├── discovery.go
│   │   └── health.go
│   ├── client/               # Client libraries
│   │   ├── producer.go
│   │   └── consumer.go
│   └── config/               # Configuration
│       ├── config.go
│       └── defaults.go
├── api/
│   ├── proto/                # gRPC definitions
│   │   ├── swarmMQ.proto
│   │   └── generated/
│   └── rest/                 # HTTP API (future)
├── pkg/                      # Public packages
│   ├── client/               # Client SDK
│   └── types/                # Shared types
├── test/
│   ├── integration/          # End-to-end tests
│   ├── performance/          # Load testing
│   └── unit/                 # Unit tests
├── configs/                  # Configuration files
│   ├── mvp-cluster.yaml
│   └── production.yaml
├── scripts/                  # Deployment scripts
│   ├── start-cluster.sh
│   └── benchmark.sh
├── docs/                     # Documentation
├── Makefile                  # Build automation
└── go.mod                    # Go modules
```

## Implementation Phases

### Phase 1: MVP Foundation (4-6 weeks)

**Week 1-2: Core Infrastructure**
- [ ] Project setup and basic gRPC services
- [ ] Message encryption and validation
- [ ] BadgerDB integration for local storage
- [ ] Basic node health checking

**Week 3-4: Routing and Delivery**
- [ ] Static hash-based routing implementation
- [ ] Producer client with connection management
- [ ] Consumer client with pull-based consumption
- [ ] Message acknowledgment system

**Week 5-6: MVP Completion**
- [ ] Cluster formation and discovery
- [ ] End-to-end integration testing
- [ ] Basic monitoring and metrics
- [ ] Performance benchmarking setup

**MVP Success Criteria**
- 5-node cluster handling 1,000 messages/second
- Sub-10ms average message delivery latency
- Zero message loss under normal operations
- Graceful handling of single node failures

### Phase 2: Smart Routing (6-8 weeks)

**Weeks 7-10: Node Specialization**
- [ ] Separate producer/routing/consumer node types
- [ ] Performance metric collection system
- [ ] Dynamic routing based on node health
- [ ] Circuit breaker implementation

**Weeks 11-14: Optimization**
- [ ] Weighted load balancing algorithms
- [ ] Route optimization based on historical data
- [ ] Adaptive timeout handling
- [ ] Advanced testing and chaos engineering

### Phase 3: Multi-Cluster (8-10 weeks)

**Weeks 15-20: Hierarchy Implementation**
- [ ] Cluster coordinator design and implementation
- [ ] Gossip protocol for routing table distribution
- [ ] Inter-cluster message routing
- [ ] Consistent hashing for cluster assignment

**Weeks 21-24: Scaling Features**
- [ ] Dynamic cluster creation/destruction
- [ ] Geographic routing preferences
- [ ] Cross-cluster load balancing
- [ ] Advanced monitoring dashboards

## Technical Specifications

### Network Protocol

**gRPC Service Definition**
```protobuf
service SwarmMQ {
    rpc SendMessage(SendMessageRequest) returns (SendMessageResponse);
    rpc PullMessages(PullMessagesRequest) returns (stream PullMessagesResponse);
    rpc AckMessage(AckMessageRequest) returns (AckMessageResponse);
    rpc GetNodeHealth(HealthRequest) returns (HealthResponse);
}
```

**Connection Management**
- **Client Connections**: Persistent gRPC connections with automatic reconnection
- **Inter-Node Communication**: Connection pooling with circuit breakers
- **TLS Configuration**: Mutual TLS for all inter-node communication

### Security Model

**Message Encryption**
- **Algorithm**: AES-256-GCM with randomly generated keys per message
- **Key Management**: Keys derived from cluster-wide master key using HKDF
- **Transport Security**: TLS 1.3 for all network communication

**Authentication & Authorization**
- **Node Authentication**: Certificate-based mutual TLS
- **Client Authentication**: API keys with configurable permissions
- **Message Integrity**: HMAC-SHA256 signatures on all messages

### Configuration System

**Cluster Configuration (YAML)**
```yaml
cluster:
  name: "mvp-cluster"
  nodes:
    - id: "node-1"
      address: "localhost:9001"
      roles: ["producer", "consumer", "router"]
    - id: "node-2"
      address: "localhost:9002"
      roles: ["producer", "consumer", "router"]
  
performance:
  max_message_size: "1MB"
  batch_size: 100
  ack_timeout: "5s"
  
storage:
  data_dir: "/var/lib/swarmMQ"
  max_disk_usage: "10GB"
  
security:
  tls_cert: "/path/to/cert.pem"
  tls_key: "/path/to/key.pem"
  master_key: "base64-encoded-key"
```

## Performance Targets

### MVP Targets
- **Throughput**: 1,000 messages/second per cluster
- **Latency**: <10ms average end-to-end delivery
- **Availability**: 99.9% uptime with single node failure tolerance
- **Memory Usage**: <512MB per node under normal load
- **Disk Usage**: <100MB per million messages stored

### Post-MVP Targets
- **Throughput**: 100,000+ messages/second per cluster
- **Latency**: <5ms average with 99th percentile <50ms
- **Availability**: 99.99% uptime with multi-node failure tolerance
- **Scalability**: Support for 100+ clusters in hierarchy
- **Recovery Time**: <30 seconds for complete cluster recovery

### Monitoring Metrics

**Key Performance Indicators**
- Message throughput (messages/second)
- End-to-end latency distribution
- Node resource utilization (CPU, memory, disk)
- Error rates and failure modes
- Routing efficiency metrics

**Alerting Thresholds**
- Latency >50ms for 95th percentile
- Error rate >0.1%
- Node resource usage >80%
- Message queue depth >10,000

---

**Last Updated**: July 25, 2025  
**Document Version**: 1.0  
**Next Review**: After MVP completion