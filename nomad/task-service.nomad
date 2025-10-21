job "task-service" {
  datacenters = ["dc1"]
  type = "service"

  group "task" {
    count = 1

    # No network stanza: run task without Nomad-managed ports to avoid local dev
    # port-label validation errors.

    task "task" {
      driver = "docker"

      config {
        image = "taskmanager/task-service:latest"
        # Not declaring ports to avoid port-label validation failures
      }

      env {
        MYSQL_URL = "mysql://root:123123@127.0.0.1:3306/task-db"
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