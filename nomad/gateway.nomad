job "gateway" {
  datacenters = ["dc1"]
  type = "service"

  group "gateway" {
    count = 1

    network {
        # No network stanza here to avoid port-label reservation issues on local dev.
    }

    task "gateway" {
      driver = "docker"

      config {
        image = "taskmanager/gateway:latest"
          # Not declaring ports to avoid port-label validation
      }

      env {
        # Use Consul DNS names to reach other services inside the cluster
          USER_SERVICE_URL    = "user-service.service.consul"
          TASK_SERVICE_URL    = "task-service.service.consul"
          PROJECT_SERVICE_URL = "project-service.service.consul"
      }

      resources {
        cpu    = 500
        memory = 256
      }

        # No Consul service registration for the gateway in this local-dev workaround.
    }
  }
}