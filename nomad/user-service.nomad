job "user-service" {
  datacenters = ["dc1"]
  type = "service"

  group "user" {
    count = 1

    # No network stanza: run task without Nomad-managed ports to avoid local dev
    # port-label validation errors.

    task "user" {
      driver = "docker"

      config {
        image = "taskmanager/user-service:latest"
        # Not declaring ports to avoid port-label validation failures
      }

      env {
        MYSQL_URL = "mysql://root:123123@127.0.0.1:3306/user-db"
      }

      resources {
        cpu    = 500
        memory = 256
        # No network stanza here
      }

      # No Consul service registration for this task to avoid port-label conflicts.
    }
  }
}