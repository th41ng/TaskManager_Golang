job "taskmanager" {
  datacenters = ["dc1"]
  type        = "service"

  #Internal Group (MySQL + gRPC services)
  group "internal-services" {
    count = 1

    network {
      mode = "host"
      port "mysql"   { static = 3306 }
      port "user"    { static = 50051 }
      port "project" { static = 50052 }
      port "task"    { static = 50053 }
    }

    #MySQL Database
    task "mysql" {
      driver = "docker"

      volume_mount {
        volume      = "mysql"
        destination = "/var/lib/mysql"
        read_only   = false
      }

      config {
        image = "mysql:8.0"
        ports = ["mysql"]
      }

      env {
        MYSQL_ROOT_PASSWORD = "Thang@2004"
        MYSQL_DATABASE      = "taskmanager"
        MYSQL_ROOT_HOST     = "%"
      }

      resources {
        cpu    = 500
        memory = 512
      }

      service {
        name = "mysql"
        port = "mysql"

        check {
          name     = "mysql-tcp"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    #User Service
    task "user-service" {
      driver = "docker"

      config {
        image = "thangmicro/user-service:latest"
        ports = ["user"]
      }

      env {
        MYSQL_URL = "root:Thang@2004@tcp(172.21.223.107:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "user-service"
        port = "user"
        tags = ["internal"]

        check {
          name     = "user-grpc"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # 🟡 Project Service
    task "project-service" {
      driver = "docker"

      config {
        image = "thangmicro/project-service:latest"
        ports = ["project"]
      }

      env {
        MYSQL_URL = "root:Thang@2004@tcp(172.21.223.107:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "project-service"
        port = "project"
        tags = ["internal"]

        check {
          name     = "project-grpc"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # 🔵 Task Service
    task "task-service" {
      driver = "docker"

      config {
        image = "thangmicro/task-service:latest"
        ports = ["task"]
      }

      env {
        MYSQL_URL = "root:Thang@2004@tcp(172.21.223.107:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "task-service"
        port = "task"
        tags = ["internal"]

        check {
          name     = "task-grpc"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    volume "mysql" {
      type      = "host"
      read_only = false
      source    = "mysql"
    }
  }

  # 🟢 Gateway (public)
  group "gateway" {
    count = 1

    network {
      mode = "host"
      port "http" { static = 8080 }
    }

    task "gateway" {
      driver = "docker"

      config {
        image = "thangmicro/gateway:latest"
        ports = ["http"]
      }

      env {
        USER_SERVICE_URL    = "172.21.223.107:50051"
        PROJECT_SERVICE_URL = "172.21.223.107:50052"
        TASK_SERVICE_URL    = "172.21.223.107:50053"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "gateway"
        port = "http"
        tags = ["public"]

        check {
          name     = "gateway-http"
          type     = "http"
          path     = "/healthz"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }
}



         