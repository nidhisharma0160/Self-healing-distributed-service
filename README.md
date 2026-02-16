## Overview
A self-healing distributed system composed of a Python Flask microservice and a Go-based watchdog. This project demonstrates automated fault detection and remediation—key components of Site Reliability Engineering (SRE).

## The SRE Implementation
Observability: The Go watchdog monitors Resident Set Size (RSS) memory of the target process.

Automation: Automatically triggers a SIGKILL and restart sequence when memory breaches the 50MB SLO (Service Level Objective).

Resilience: Designed to recover service availability without human intervention.

## Real-World Learnings (Port Conflict)
During testing on macOS, the system encountered a Port 5000: Address already in use error.

Root Cause: macOS AirPlay Receiver occupies port 5000, and the OS TIME_WAIT state prevents immediate socket reuse after a process kill.

Solution: Implemented a 2-second 'settle' delay in the Go watchdog to allow the OS to release the port before restarting the service.