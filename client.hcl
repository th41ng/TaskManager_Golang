client {
  enabled = true

  options {
    "docker.enable"           = "true"
    "docker.cleanup.image"    = "true"
    "docker.volumes.enabled"  = "true"
  }
  cni_path = "/opt/cni/bin"

    host_volume "mysql" {
    path      = "/opt/mysql/data"
    read_only = false
  }

  # Host volumes for per-service MySQL instances
  host_volume "mysql_user" {
    path      = "/opt/mysql/user"
    read_only = false
  }

  host_volume "mysql_project" {
    path      = "/opt/mysql/project"
    read_only = false
  }

  host_volume "mysql_task" {
    path      = "/opt/mysql/task"
    read_only = false
  }

}