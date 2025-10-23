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

}