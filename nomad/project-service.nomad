job "project-service" {
  datacenters = ["dc1"]
  type = "service"

  group "project" {
    count = 1

    # No network stanza: run container without Nomad-managed port labels to avoid
    # port-label conflicts on local dev agents. The container will still run, but
    # Nomad won't reserve or publish ports for it.

    task "project" {
      driver = "docker"

      config {
        image = "taskmanager/project-service:latest"
        # Not declaring ports here to avoid Nomad port label validation failures
        # when the agent holds stale reservations. For local testing you can
        # still exec into the container or use Nomad alloc logs to verify startup.
      }

      env {
        MYSQL_URL = "mysql://root:123123@127.0.0.1:3306/project-db"
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