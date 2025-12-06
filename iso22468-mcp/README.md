# ISO 22468:2020 VSM MCP Server

## Document Information

**Document Title**: ISO 22468:2020 VSM MCP Server User Guide  
**Document Version**: 1.0  
**Document Date**: December 2025  
**Software Version**: v1.0.0  
**Compliance Standard**: ISO/IEC/IEEE 26514:2017  
**MCP Protocol Version**: 2025-06-18  

## Table of Contents

| Section | Topic | Description |
|---------|-------|-------------|
| **1** | **[Introduction](#1-introduction)** | Document purpose, scope, audience, and requirements |
| **1.1** | **[Purpose](#11-purpose)** | Goals and objectives of this MCP server |
| **1.2** | **[Scope](#12-scope)** | What is covered and what is not |
| **1.3** | **[Target Audience](#13-target-audience)** | Who this documentation is for |
| **1.4** | **[Prerequisites](#14-prerequisites)** | Required knowledge and software |
| **2** | **[Concept of Operations](#2-concept-of-operations)** | Understanding MCP and VSM integration |
| **2.1** | **[MCP Protocol Overview](#21-mcp-protocol-overview)** | How MCP enables AI-tool integration |
| **2.2** | **[VSM Server Architecture](#22-vsm-server-architecture)** | How the server components work together |
| **2.3** | **[Core Capabilities](#23-core-capabilities)** | Tools, resources, and prompts provided |
| **3** | **[Procedures](#3-procedures)** | Step-by-step usage instructions |
| **3.1** | **[Installation](#31-installation)** | How to install and set up the server |
| **3.2** | **[Getting Started](#32-getting-started)** | Basic setup and first VSM analysis |
| **3.3** | **[Basic Usage](#33-basic-usage)** | Essential operations and workflows |
| **3.4** | **[Advanced Usage](#34-advanced-usage)** | Complex scenarios and optimization |
| **4** | **[Reference Information](#4-reference-information)** | Detailed technical specifications |
| **4.1** | **[Tool Reference](#41-tool-reference)** | Complete list of available tools |
| **4.2** | **[Resource Reference](#42-resource-reference)** | Available resources and their content |
| **4.3** | **[Prompt Reference](#43-prompt-reference)** | Interactive prompts for guidance |
| **4.4** | **[Configuration](#44-configuration)** | Server configuration options |
| **5** | **[Troubleshooting](#5-troubleshooting)** | Common issues and solutions |
| **6** | **[Glossary](#6-glossary)** | Key terms and definitions |
| **7** | **[Standards Compliance](#7-standards-compliance)** | ISO compliance and validation |
| **8** | **[Contributing](#8-contributing)** | How to contribute to the project |
| **9** | **[License](#9-license)** | Legal information and permissions |

### Quick Navigation Guide

| **Getting Started** | **Tool Reference** | **Troubleshooting** |
|-------------------|------------------|-------------------|
| [Installation](#31-installation) | [Tool Reference](#41-tool-reference) | [Troubleshooting](#5-troubleshooting) |
| [Getting Started](#32-getting-started) | [Resources](#42-resource-reference) | [Common Issues](#5-troubleshooting) |
| [Basic Usage](#33-basic-usage) | [Prompts](#43-prompt-reference) | [Error Solutions](#5-troubleshooting) |
| [Advanced Usage](#34-advanced-usage) | [Configuration](#44-configuration) | [Help & Support](#5-troubleshooting) |

### Document Statistics
- **Total Sections**: 9 main sections
- **Total Subsections**: 14 subsections
- **Code Examples**: 15+ practical examples
- **Compliance**: ISO/IEC/IEEE 26514:2017 and ISO 22468:2020
- **MCP Compliance**: Model Context Protocol specification

---

<a name="1-introduction"></a>
## 1. Introduction

### 1.1 Purpose

This document provides comprehensive guidance for users implementing Value Stream Management (VSM) solutions using the ISO 22468:2020 VSM MCP Server. The server offers a complete MCP-compatible implementation that enables AI assistants and applications to perform systematic VSM analysis, design, planning, and assessment operations.

The primary purposes of this MCP server are to:
- Provide standardized VSM operations compliant with ISO 22468:2020
- Enable AI-assisted value stream analysis and improvement
- Support complete VSM workflow automation
- Facilitate integration with MCP-compatible AI applications

### 1.2 Scope

This documentation covers:
- Installation and configuration of the VSM MCP Server
- Core concepts of MCP protocol integration with VSM
- Step-by-step procedures for VSM operations via MCP
- Tool, resource, and prompt references
- Troubleshooting common integration issues

This documentation does **not** cover:
- General MCP protocol implementation details (refer to MCP specification)
- ISO 22468:2020 standard interpretation (refer to the standard itself)
- AI application development beyond MCP integration
- Third-party MCP client implementations

### 1.3 Target Audience

This guide is intended for:
- **AI Application Developers**: Developers integrating MCP clients with VSM capabilities
- **Process Engineers**: Engineers requiring AI-assisted VSM analysis
- **System Architects**: Architects designing MCP-based VSM solutions
- **DevOps Engineers**: Engineers automating VSM workflows
- **Quality Assurance Specialists**: QA professionals validating VSM MCP integrations

**Assumed Knowledge**:
- Basic understanding of Model Context Protocol (MCP)
- Familiarity with JSON data structures and REST APIs
- Understanding of manufacturing processes and value streams
- Knowledge of command-line interfaces and build tools

### 1.4 Prerequisites

Before using this MCP server, ensure you have:

**Software Requirements**:
- Go 1.19 or later
- Operating system: Windows, macOS, or Linux
- MCP-compatible client application (e.g., Claude Desktop, custom MCP client)

**Knowledge Prerequisites**:
- Go programming fundamentals (for custom integrations)
- Basic understanding of JSON serialization
- Familiarity with manufacturing process terminology
- Understanding of value stream mapping concepts

---

<a name="2-concept-of-operations"></a>
## 2. Concept of Operations

### 2.1 MCP Protocol Overview

The Model Context Protocol (MCP) is an open protocol that enables seamless communication between AI applications and external tools and data sources. This VSM MCP Server implements the MCP specification to provide standardized access to ISO 22468:2020 compliant VSM operations.

The server operates in stdio mode, communicating with MCP clients through standard input/output streams using JSON-RPC 2.0 messages. This design enables:
- Secure, sandboxed execution of VSM operations
- Standardized tool and resource discovery
- Consistent error handling and progress reporting
- Integration with any MCP-compatible AI application

### 2.2 VSM Server Architecture

The VSM MCP Server implements a modular architecture that mirrors the ISO 22468:2020 VSM methodology:

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   MCP Client    │◄──►│  VSM MCP Server │◄──►│  VSM SDK Core   │
│ (AI Assistant)  │    │                 │    │                 │
└─────────────────┘    │ ┌─────────────┐ │    │ ┌─────────────┐ │
                       │ │ MCP Tools   │ │    │ │ Analyzer    │ │
                       │ │ MCP Resources│ │    │ │ Designer    │ │
                       │ │ MCP Prompts │ │    │ │ Planner     │ │
                       │ └─────────────┘ │    │ │ Assessor    │ │
                       └─────────────────┘    │ │ Serializer  │ │
                                              │ └─────────────┘ │
                                              └─────────────────┘
```

**Key Components**:
- **MCP Tools**: Executable VSM operations (analysis, design, planning, assessment)
- **MCP Resources**: Static data access (symbols, standards, templates)
- **MCP Prompts**: Interactive guidance for complex workflows
- **VSM SDK Core**: ISO 22468:2020 compliant VSM implementation

### 2.3 Core Capabilities

#### Tools
The server provides 13 specialized VSM tools organized by methodology phase:

**Analysis Tools**:
- `vsm_analyze_current_state`: Comprehensive current state analysis
- `vsm_detect_bottlenecks`: Bottleneck identification and analysis
- `vsm_analyze_waste`: 8 types of waste analysis (Muda)

**Design Tools**:
- `vsm_design_future_state`: Future state design with improvement potentials
- `vsm_generate_ideal_state`: Generate ideal state value streams

**Planning Tools**:
- `vsm_plan_implementation`: Implementation planning with timelines
- `vsm_generate_smart_measures`: SMART measure generation (removed in current version)

**Assessment Tools**:
- `vsm_assess_performance`: Performance assessment against targets

**Serialization Tools**:
- `vsm_serialize_value_stream`: JSON serialization
- `vsm_deserialize_value_stream`: JSON deserialization
- `vsm_save_to_file`: Save to file system
- `vsm_load_from_file`: Load from file system

**Utility Tools**:
- `vsm_calculate_kpis`: KPI calculation
- `vsm_create_value_stream`: Create new value streams

#### Resources
Static resources providing reference information:
- `vsm://symbols`: VSM symbol reference library
- `vsm://standards`: ISO compliance information
- `vsm://templates`: Pre-defined value stream templates

#### Prompts
Interactive guidance prompts:
- `vsm_analysis_workflow`: Current state analysis guidance
- `vsm_implementation_guide`: Implementation planning guidance

---

<a name="3-procedures"></a>
## 3. Procedures

<a name="31-installation"></a>
### 3.1 Installation

To install the VSM MCP Server in your environment:

#### Step 1: Prerequisites Check
```bash
# Verify Go installation
go version

# Expected output: go version go1.19 or later
```

#### Step 2: Clone or Download
```bash
# Navigate to your MCP servers directory
cd /path/to/mcp/servers

# Clone the repository (if available) or ensure the binary is available
# The server binary should be named 'iso22468-mcp-server'
```

#### Step 3: Verify Installation
```bash
# Test server execution
./iso22468-mcp-server --help

# Expected: Server starts and shows MCP protocol messages
# Use Ctrl+C to stop
```

**Expected Result**: The VSM MCP Server is installed and ready for MCP client integration.

<a name="32-getting-started"></a>
### 3.2 Getting Started

Follow these steps to integrate the VSM MCP Server with your MCP client:

#### Step 1: Configure MCP Client
Add the server to your MCP client configuration (e.g., Claude Desktop):

```json
{
  "mcpServers": {
    "iso22468-vsm": {
      "command": "path/to/iso22468-mcp-server",
      "args": [],
      "env": {}
    }
  }
}
```

#### Step 2: Start MCP Client
Launch your MCP-compatible application with the server configured.

#### Step 3: Verify Connection
The server will automatically initialize and register its capabilities with the MCP client.

#### Step 4: Test Basic Functionality
Use the MCP client to request available tools:

```
List available VSM tools and verify the server is responding.
```

**Expected Result**: MCP client shows 13 VSM tools, 3 resources, and 2 prompts available.

<a name="33-basic-usage"></a>
### 3.3 Basic Usage

#### Creating a Value Stream
```json
{
  "tool": "vsm_create_value_stream",
  "parameters": {
    "id": "production_line_1",
    "name": "Main Production Line",
    "product_family": "Widget Family A",
    "customer_json": "{\"id\":\"customer_1\",\"name\":\"Manufacturing Corp\",\"demand\":1000,\"takt_time\":5.0}"
  }
}
```

#### Analyzing Current State
```json
{
  "tool": "vsm_analyze_current_state",
  "parameters": {
    "value_stream_json": "{...value stream data...}"
  }
}
```

#### Calculating KPIs
```json
{
  "tool": "vsm_calculate_kpis",
  "parameters": {
    "value_stream_json": "{...value stream data...}"
  }
}
```

#### Saving Results
```json
{
  "tool": "vsm_save_to_file",
  "parameters": {
    "value_stream_json": "{...value stream data...}",
    "filename": "production_line_vsm.json"
  }
}
```

<a name="34-advanced-usage"></a>
### 3.4 Advanced Usage

#### Complete VSM Workflow
```json
// 1. Create value stream
{
  "tool": "vsm_create_value_stream",
  "parameters": {...}
}

// 2. Analyze current state
{
  "tool": "vsm_analyze_current_state",
  "parameters": {...}
}

// 3. Design future state
{
  "tool": "vsm_design_future_state",
  "parameters": {...}
}

// 4. Plan implementation
{
  "tool": "vsm_plan_implementation",
  "parameters": {...}
}

// 5. Assess performance
{
  "tool": "vsm_assess_performance",
  "parameters": {...}
}
```

#### Using Resources
```json
// Access VSM symbols
{
  "resource": "vsm://symbols"
}

// Get standards information
{
  "resource": "vsm://standards"
}
```

#### Interactive Guidance
```json
// Get analysis workflow guidance
{
  "prompt": "vsm_analysis_workflow",
  "parameters": {
    "industry": "automotive",
    "complexity": "medium"
  }
}
```

---

<a name="4-reference-information"></a>
## 4. Reference Information

<a name="41-tool-reference"></a>
### 4.1 Tool Reference

#### vsm_create_value_stream
Creates a new value stream with basic structure.

**Parameters**:
- `id` (string, required): Unique identifier
- `name` (string, required): Display name
- `product_family` (string, required): Product family name
- `customer_json` (string, optional): Customer information as JSON

#### vsm_analyze_current_state
Performs comprehensive current state analysis.

**Parameters**:
- `value_stream_json` (string, required): Value stream data as JSON

**Returns**: Analysis results including bottlenecks, waste analysis, and flow metrics.

#### vsm_detect_bottlenecks
Identifies bottlenecks in the value stream.

**Parameters**:
- `value_stream_json` (string, required): Value stream data as JSON

**Returns**: List of identified bottlenecks with severity and location.

#### vsm_analyze_waste
Analyzes waste according to the 8 types of muda.

**Parameters**:
- `value_stream_json` (string, required): Value stream data as JSON

**Returns**: Waste analysis by category with reduction potentials.

#### vsm_design_future_state
Creates future state design with improvements.

**Parameters**:
- `current_state_json` (string, required): Current value stream
- `analysis_result_json` (string, required): Analysis results

**Returns**: Future state design with improvement potentials.

#### vsm_generate_ideal_state
Generates ideal state value stream design.

**Parameters**:
- `current_state_json` (string, required): Current value stream

**Returns**: Ideal state value stream with minimal waste.

#### vsm_plan_implementation
Creates implementation plan with measures.

**Parameters**:
- `design_result_json` (string, required): Design results
- `project_json` (string, required): Project information

**Returns**: Implementation plan with timelines and measures.

#### vsm_assess_performance
Assesses performance against targets.

**Parameters**:
- `value_stream_json` (string, required): Value stream to assess
- `targets_json` (string, required): KPI targets as JSON
- `history_json` (string, optional): Historical data

**Returns**: Performance assessment with scores and recommendations.

#### vsm_calculate_kpis
Calculates key performance indicators.

**Parameters**:
- `value_stream_json` (string, required): Value stream data

**Returns**: Map of calculated KPIs.

#### Serialization Tools
- `vsm_serialize_value_stream`: Convert to JSON
- `vsm_deserialize_value_stream`: Parse from JSON
- `vsm_save_to_file`: Save to file system
- `vsm_load_from_file`: Load from file system

<a name="42-resource-reference"></a>
### 4.2 Resource Reference

#### vsm://symbols
Complete VSM symbol reference library.

**Content**: All ISO 22468:2020 VSM symbols with descriptions, categories, and usage information.

#### vsm://standards
ISO compliance and standards information.

**Content**: Compliance matrix, Rother & Shook principles, and validation information.

#### vsm://templates
Pre-defined value stream templates.

**Content**: Ready-to-use templates for common manufacturing scenarios (production lines, service processes, order fulfillment).

<a name="43-prompt-reference"></a>
### 4.3 Prompt Reference

#### vsm_analysis_workflow
Guides through current state analysis process.

**Parameters**:
- `industry` (string): Industry sector for context
- `complexity` (string): Value stream complexity level

**Returns**: Step-by-step analysis guidance tailored to context.

#### vsm_implementation_guide
Guides through implementation planning.

**Parameters**:
- `timeline` (string): Available implementation timeline
- `resources` (string): Available resources and budget

**Returns**: Implementation planning guidance with risk assessment.

<a name="44-configuration"></a>
### 4.4 Configuration

The server is configured through the MCP client configuration. No additional configuration files are required.

**Supported Transports**:
- stdio (standard input/output) - primary transport
- Future: HTTP/SSE transport support

**Logging**: Server logs are sent through MCP protocol logging messages.

**Error Handling**: All errors are returned as structured MCP error responses.

---

<a name="5-troubleshooting"></a>
## 5. Troubleshooting

### Common Issues and Solutions

#### Issue: Server fails to start
**Symptoms**: MCP client reports connection failure.

**Solutions**:
1. Verify server binary exists and is executable
2. Check file permissions
3. Ensure Go runtime is available
4. Review MCP client configuration

#### Issue: Tools not appearing in MCP client
**Symptoms**: Server starts but tools are not available.

**Solutions**:
1. Verify MCP protocol version compatibility
2. Check server initialization logs
3. Restart MCP client after server configuration
4. Ensure stdio transport is properly configured

#### Issue: Invalid value stream JSON
**Symptoms**: Tools return JSON parsing errors.

**Solutions**:
1. Validate JSON structure against ISO 22468:2020 schema
2. Ensure all required fields are present
3. Check data types and ranges
4. Use `vsm_create_value_stream` for basic structure

#### Issue: Analysis returns empty results
**Symptoms**: Analysis tools return no meaningful data.

**Solutions**:
1. Verify value stream has processes defined
2. Check process parameters are valid
3. Ensure customer demand is specified
4. Validate process connections and flows

#### Issue: Memory usage during large analyses
**Symptoms**: Server consumes excessive memory.

**Solutions**:
1. Process large value streams in batches
2. Use streaming serialization for big datasets
3. Monitor MCP client resource limits
4. Consider value stream size optimization

### Error Codes and Meanings

| Error Code | Description | Solution |
|------------|-------------|----------|
| `INVALID_JSON` | Malformed JSON input | Validate JSON structure and syntax |
| `MISSING_FIELD` | Required field not provided | Check tool parameter requirements |
| `CALCULATION_ERROR` | Mathematical calculation failed | Verify input data ranges and types |
| `SERIALIZATION_ERROR` | Data persistence failed | Check file system permissions |
| `VALIDATION_ERROR` | Data validation failed | Review ISO 22468:2020 compliance |

### Performance Tuning

#### Memory Optimization
- Use streaming operations for large datasets
- Process value streams in logical chunks
- Monitor MCP client memory limits

#### Response Time Optimization
- Cache frequently accessed resources
- Use incremental analysis for large streams
- Optimize JSON serialization settings

### Getting Help

If you encounter issues not covered here:

1. Check the server logs through MCP protocol logging
2. Review MCP specification for protocol-level issues
3. Validate against ISO 22468:2020 standard for VSM methodology
4. Test with minimal value stream examples
5. Report issues with complete error messages and configuration

---

<a name="6-glossary"></a>
## 6. Glossary

| Term | Definition |
|------|------------|
| **MCP (Model Context Protocol)** | Open protocol for AI-tool integration enabling secure communication between applications |
| **Value Stream** | Sequence of activities required to design, produce, and deliver a product to a customer |
| **VSM (Value Stream Management)** | Methodology for analyzing and improving value streams by identifying waste and bottlenecks |
| **Tool** | Executable function exposed through MCP for performing specific operations |
| **Resource** | Static data accessible through MCP for reference and templates |
| **Prompt** | Interactive guidance template for complex workflows |
| **Bottleneck** | Process that limits the overall capacity of the value stream |
| **Waste (Muda)** | Any activity that consumes resources but creates no value |
| **Takt Time** | Rate at which customers require product delivery |
| **KPI (Key Performance Indicator)** | Quantifiable measure of performance against goals |
| **PDCA Cycle** | Plan-Do-Check-Act continuous improvement methodology |
| **SMART Goals** | Specific, Measurable, Achievable, Relevant, Time-bound objectives |

---

<a name="7-standards-compliance"></a>
## 7. Standards Compliance

This MCP server is fully compliant with:

### MCP Protocol Compliance
- **Protocol Version**: 2025-06-18
- **Transport Support**: stdio (primary), HTTP/SSE (future)
- **Message Format**: JSON-RPC 2.0
- **Security**: Sandboxed execution through stdio transport

### ISO 22468:2020 Compliance
- **Core Methodology**: Complete VSM 4-phase implementation
- **Calculation Procedures**: All Annex A formulas implemented
- **Data Structures**: ISO-compliant value stream models
- **Symbol Library**: Complete VSM symbol set
- **Waste Analysis**: 8 types of muda identification

### ISO/IEC/IEEE 26514:2017 Compliance
- **Document Structure**: Front matter, introduction, procedures, reference
- **Content Quality**: Clear, complete, accurate information
- **Accessibility**: Navigable structure with TOC and cross-references
- **Consistency**: Standardized terminology and formatting
- **Maintainability**: Version-controlled documentation

### Validation and Testing
- **Unit Tests**: All VSM calculations validated
- **Integration Tests**: MCP protocol compliance verified
- **Performance Tests**: Memory and response time validated
- **Compatibility Tests**: Multiple MCP clients tested

---

<a name="8-contributing"></a>
## 8. Contributing

We welcome contributions to the ISO 22468:2020 VSM MCP Server project.

### Ways to Contribute
- **Bug Reports**: Report MCP integration issues
- **Feature Requests**: Suggest new VSM tools or capabilities
- **Code Contributions**: Submit improvements to VSM algorithms
- **Documentation**: Improve this user guide
- **Testing**: Add test cases for MCP integration

### Development Setup
1. Install Go 1.19 or later
2. Clone the repository
3. Run `go mod tidy` to install dependencies
4. Build with `go build -o iso22468-mcp-server .`
5. Test with MCP client integration

### Contribution Guidelines
- Follow Go coding standards
- Add tests for new functionality
- Update documentation for API changes
- Maintain ISO 22468:2020 compliance
- Test MCP protocol compatibility

---

<a name="9-license"></a>
## 9. License

This project implements the ISO 22468:2020 Value Stream Management standard and is licensed under the MIT License.

**MIT License**

Copyright (c) 2025 ISO 22468:2020 VSM MCP Server Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

**Document History**
- v1.0 (December 2025) - Initial ISO/IEC/IEEE 26514 compliant documentation for VSM MCP Server

**Contact Information**
For questions or support, please refer to the MCP protocol documentation or ISO 22468:2020 standard.
