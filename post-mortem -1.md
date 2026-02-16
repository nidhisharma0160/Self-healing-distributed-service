Incident Post-Mortem: Service Recovery Failure (Port Conflict)
Date: 2026-02-15

Author: Nidhi Sharma

Status: Completed

Executive Summary
During the testing of the Automated Remediation Watchdog, the system successfully identified a memory leak but failed to restart the service immediately. The service remained down for several seconds due to an Address already in use error.

The Incident
Detection: The Go watchdog correctly identified that app.py exceeded the 50MB RSS memory threshold.

Action: The watchdog issued a SIGKILL to the Python process.

Failure: The subsequent restart command failed because Port 5000 was still occupied.

Root Cause Analysis:
TCP Socket Lifecycle: When a process is killed, the operating system keeps the socket in a TIME_WAIT state for a brief period to ensure all packets are received.

macOS Specific Conflict: On macOS Monterey and later, the AirPlay Receiver service defaults to port 5000, creating a systemic conflict with the Flask default port.

Resolution and Prevention
Immediate Fix: Manually disabled "AirPlay Receiver" in macOS System Settings.

Systemic Fix: Refactored the Go watchdog to include a 2-second "Cool-down" delay between the Kill and Restart actions to allow the OS to release the socket.

Future Action: Transition the microservice to use dynamic port allocation or an environment-variable-defined port to avoid hardcoded conflicts.
