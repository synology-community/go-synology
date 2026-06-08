# dsm-api

## Running Integration Tests

Integration tests require a running DSM instance. You can use the included Docker Compose setup to spin up a virtual DSM, or run against a real NAS.

### Using Virtual DSM (Docker)

```sh
# Start the virtual DSM
docker compose up -d

# Wait for DSM to become healthy, then complete the setup wizard at:
# http://localhost:5001
# Create an admin user (e.g. realperson / Password123!)

# Set credentials and run tests
export SYNOLOGY_HOST=http://localhost:5001
export SYNOLOGY_USER=realperson
export SYNOLOGY_PASSWORD=Password123!
export SYNOLOGY_VIRTUAL_DSM=true  # skips tests known to be unsupported by virtual DSM

go test ./pkg/api/core/...
```

> **Note:** Testing against the virtual DSM image revealed that SYNO.Core.User returns error 105 and SYNO.Core.EventScheduler returns error 117 for all accounts. Setting `SYNOLOGY_VIRTUAL_DSM=true` skips these tests with a descriptive reason instead of failing.

### Using a Real NAS

```sh
export SYNOLOGY_HOST=https://<nas-ip>:5001
export SYNOLOGY_USER=<admin-user>
export SYNOLOGY_PASSWORD=<password>

go test ./pkg/api/core/...
```
