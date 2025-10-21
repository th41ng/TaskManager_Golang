Nomad + Consul quick start for TaskManager repo

Overview
--------
This folder contains minimal Nomad job files for local dev: `gateway.nomad`, `user-service.nomad`, `task-service.nomad`, `project-service.nomad` and a helper `run_nomad_dev.ps1`.

What these do
--------------
- Start Consul and Nomad in dev mode (via `run_nomad_dev.ps1`)
- Build local Docker images for each service
- Run jobs which register services in Consul
- Gateway job uses Consul DNS names like `user-service.service.consul:50051` to reach internal services

Notes
-----
- The job examples use static ports (8080, 50051, 50052, 50053) to make local testing straightforward. On a real multi-node cluster prefer dynamic ports and an external ingress/load-balancer (Traefik, HAProxy) or Consul Connect + Envoy.
- For multi-node, push images to a registry and update the `image` fields accordingly.
- Use Vault for secrets in production; here sample MYSQL_URLs are present for dev only.

Run locally (Windows PowerShell)
-------------------------------
1. Install Consul and Nomad (choco or download binaries) and Docker.
2. From repo root in PowerShell run:

```powershell
Set-ExecutionPolicy -Scope Process -ExecutionPolicy Bypass
.
\nomad\run_nomad_dev.ps1
```

3. Open Consul UI: http://127.0.0.1:8500 and Nomad UI: http://127.0.0.1:4646
4. Test gateway in browser/Postman: http://localhost:8080

Production notes
----------------
- For multi-node clusters use dynamic ports and a load-balancer. Traefik has a Consul provider and fits well. Alternatively use Consul Connect + Envoy for mTLS.
- Register health checks in each Nomad `service` stanza so Consul and load balancers can avoid unhealthy instances.
