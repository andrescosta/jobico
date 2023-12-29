# Jobico

## Introduction

Welcome to Jobico – An Experimental, Multi-Tenancy Job Execution Platform!

Jobico is a project in the realm of experimental development with no commercial objectives. Specifically crafted for exploring job execution using WebAssembly (WASM) in a multi-tenancy environment, Jobico emphasizes flexibility and customization. It empowers tenants to define events, validate incoming data using JSON Schema, and execute custom programs written in any WASM-compatible language.

## Key Characteristics

- **Exploratory Nature**: Jobico is an exploratory project, offering a platform to delve into innovative approaches to job execution technologies.
- **Non-Commercial Objective**: Our primary goal is not commercial gain but rather fostering exploration and creativity in the realm of job execution and event processing.
- **Multi-Tenancy Focus**: Jobico is architected to support multi-tenancy, allowing for the simultaneous operation of multiple isolated tenants on the platform.
- **Event Definition with JSON Schema**: Tenants can define events by providing JSON Schema, enabling structured and dynamic event handling. Incoming requests are validated against the defined schema.
- **WASM-Compatible Language Support**: Tenants have the freedom to create custom programs in any WASM-compatible language, promoting flexibility and diversity in job execution.


# Architecture
![alt](docs/img/Jobico.svg?)

## Components

### Listener
The **Listener** component serves as the entry point for external events, providing a REST API that functions as a webhook. Its primary responsibilities include receiving events, validating them against pre-defined JSON schemas, and subsequently enqueueing them for further processing. This component acts as the bridge between external sources triggering events and the internal processing pipeline.

### Queue
The **Queue** component acts as a chronological buffer for events, temporarily storing them until they can be processed by the Job Executors. Events are maintained in the queue in the order they are received, ensuring a sequential flow of processing. This component plays a crucial role in decoupling the event reception from the actual event processing, allowing for scalability and efficient handling of bursts of incoming events.

### Job Executors
**Job Executors** are responsible for consuming events from the Queue and providing a controlled environment for the execution of WebAssembly (WASM) programs that process these events. This component manages the execution context, ensuring isolation and security for running custom WASM programs written by tenants. It plays a key role in the dynamic and scalable execution of programmed jobs in response to events.

### Control Service
The **Control Service** is a centralized hub where Job definitions are stored and can be queried by other components in the system. It acts as the authoritative source for managing job configurations, allowing dynamic adjustments to the processing logic without interrupting the overall system operation. This service facilitates coordination and control over the execution of jobs across the entire platform.

### Job Repository
The **Job Repository** serves as a storage facility for WebAssembly (WASM) programs and JSON schema files. It provides a dedicated API for storing and retrieving these essential files, ensuring accessibility for the Job Executors and enabling tenants to manage their custom program logic efficiently. This component acts as a repository for the building blocks required for event processing.

### Executions Recorder
The **Executions Recorder** is a service designed to capture and store log information generated by the Job Executors during the execution of jobs. This component acts as a centralized logging system, allowing for post-execution analysis, troubleshooting, and performance monitoring. The recorded information can be queried using both the Command Line Interface (CLI) and Dashboard tools, providing visibility into the historical execution details of jobs.

### Command Line Tool
The **Command Line Tool** serves as the primary management interface, providing a comprehensive set of commands for tenants to interact with the Jobico platform. This tool empowers tenants to deploy job definitions, upload associated WASM and JSON schema files, and query the Executions Recorder for log information. Its command-line interface offers flexibility and efficiency in managing and overseeing job-related activities within the system.

### Dashboard
The **Dashboard** is a terminal-based application designed for visualizing Job definitions and execution information. Sporting a colorful and intuitive interface, the Dashboard provides a user-friendly experience for tenants to monitor and analyze the status of their jobs. It serves as a graphical representation of the Jobico platform, offering insights into the dynamic execution of jobs, the current state of the system, and facilitating quick decision-making through its visually appealing display. The Dashboard enhances the overall user experience by providing a vivid and informative overview of the platform's activity.

# Job Definition

#### `id`

- **Description:** The "id" attribute represents the unique identifier for the job. It serves as a distinct reference to identify and manage the job within the Jobico platform.

- **Example:**

  ```yaml
  id: 12345
  ```

  In this example, "12345" is the unique identifier assigned to the job.

#### `tenant`

- **Description:** The "tenant" attribute represents the ID of the tenant associated with the job. It ensures that the job is attributed to a specific tenant within the multi-tenancy architecture of Jobico.

- **Example:**

  ```yaml
  tenant: 56789
  ```

  In this example, "56789" is the ID of the tenant associated with the job.

#### `queues`

- **Description:** The "queues" section describes the queues associated with the job. This section allows for future expansion where queue environments and capabilities can be defined.

  - `queues.id`: ID of the queue.
  - `queues.name`: Friendly name of the queue.

- **Example:**

  ```yaml
  queues:
    - id: 1
      name: high-priority-queue
    - id: 2
      name: default-queue
  ```

  In this example, two queues, "high-priority-queue" and "default-queue," are defined with their respective IDs.

#### `jobs`

- **Description:** The "jobs" section is where the events are defined and how they will be validated and processed.

  - `jobs.event`: An event definition.
  - `jobs.event.name`: Friendly name for the event.
  - `jobs.event.id`: ID for the event, used by the REST API and executors to determine the schema and WASM file.
  - `jobs.event.datatype`: Specifies the data type of the event. "0" represents JSON.

  - **`jobs.event.schema`: Schema file definition:**

    - `jobs.event.schema.id`: ID of the schema file.
    - `jobs.event.schema.name`: Name of the schema file.
    - `jobs.event.schema.schemaref`: Reference used to retrieve the file from the repository.

  - `jobs.event.supplierqueue`: Specifies the ID of the queue where this event will be published.
  - `jobs.event.runtime`: ID of the runtime that will process this event.
  - `jobs.event.result`: Specifies how the result of the execution will be treated (Under Construction).

- **Example:**

  ```yaml
  jobs:
    - event:
        name: user-registration
        id: 6789
        datatype: 0
        schema:
          id: 9876
          name: user-registration-schema
          schemaref: /schemas/user-registration-schema.json
        supplierqueue: 1
        runtime: 1
        result: under-construction
  ```

  In this example, a job is defined for processing "user-registration" events with the associated schema and runtime details.

#### `runtimes`

- **Description:** The "runtimes" section specifies the runtimes available to process the events.

  - `runtimes.id`: ID of the runtime, used to reference a specific runtime.
  - `runtimes.name`: Friendly name.
  - `runtimes.moduleref`: Reference used to retrieve the file from the repository.
  - `runtimes.mainfuncname`: Not used at the moment.
  - `runtimes.type`: "0" represents WASM as the runtime type.

- **Example:**

  ```yaml
  runtimes:
    - id: 1
      name: wasm-runtime
      moduleref: /runtimes/wasm-runtime.wasm
      mainfuncname: main
      type: 0
  ```

  In this example, a runtime named "wasm-runtime" is defined with the associated WASM file and runtime type.

These attributes collectively form a comprehensive YAML file, capturing the essential details for defining and deploying jobs within the Jobico platform. If you have additional attributes or specific details you'd like to cover, feel free to provide them, and I'll be happy to assist further.

### Example
  ```yaml
name: test
id: test
tenant: m1 
queues:
  - id: queue
    name: queue
jobs:
  - event:
      name: an event
      id: ev1
      datatype: 0
      schema:
        id: schema
        name: the schema
        schemaref: schema.json
      supplierqueue: queue
      runtime: runtime1
    result:
      ok:
        name: ev1_ok
        id: ev1_ok
        datatype: 0
        schema:
          id: sche1
          name: sche1
          schemaref: schema.json
        supplierqueue: queue1
        runtime: runtime1
      error:
        name: ev1_err
        id: ev1_err
        datatype: 0
        schema:
          id: sche1
          name: sche1
          schemaref: schema.json
        supplierqueue: queue1
        runtime: runtime1

runtimes:
  - id: runtime1
    name: greet.wasm
    moduleref: greet.wasm
    mainfuncname: event
    type: 0

  ```


# Jobicolet


# Tools

## Command Line
## Terminal Dashboard

## Getting Started

To embark on your journey with Jobico, follow our [getting started guide](link/to/getting/started) to set up your multi-tenancy environment and define your first event.

## Installation
### Docker

## Goico
### Packages


# Usage
## Roadmap

### Acknowledgements

## Support and Contact

For questions, feedback, or assistance, reach out to us at [jobicowasm@gmail.com].
