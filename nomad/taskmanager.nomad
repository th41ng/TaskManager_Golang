job "taskmanager" {
  datacenters = ["dc1"]
  type        = "service"

  ##########################
  # 🧱 MySQL Database (internal)
  ##########################
  group "mysql" {
    count = 1

    network {
      port "db" {
        static = 3306
      }
    }

    volume "mysql" {
      type      = "host"
      read_only = false
      source    = "mysql"
    }

    task "mysql" {
      driver = "docker"

      volume_mount {
        volume      = "mysql"
        destination = "/var/lib/mysql"
        read_only   = false
      }

      env {
        MYSQL_ROOT_PASSWORD = "Thang@2004"
        MYSQL_DATABASE      = "taskmanager"
        MYSQL_ROOT_HOST     = "%"
      }

      config {
        image = "mysql:8.0"
        ports = ["db"]
      }

      resources {
        cpu    = 500
        memory = 512
      }

      service {
        name = "mysql"
        port = "db"
        tags = ["internal"]
        check {
          name     = "mysql-tcp"
          type     = "tcp"
          port     = "db"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }

  ##########################
  # 🟣 User Service (internal)
  ##########################
  group "user-service" {
    count = 1

    network {
      port "grpc" {
        to = 50051
      }
    }

    task "user" {
      driver = "docker"

      config {
        image = "thangmicro/user-service:latest"
        ports = ["grpc"]
      }

      env {
        DB_HOST   = "mysql.service.consul"
        DB_PORT   = "3306"
        DB_USER   = "root"
        DB_PASS   = "Thang@2004"
        DB_NAME   = "taskmanager"
        MYSQL_URL = "root:Thang@2004@tcp(172.21.223.107:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "user-service"
        port = "grpc"
        tags = ["internal"]
        check {
          name     = "user-grpc"
          type     = "tcp"
          port     = "grpc"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }

  ##########################
  # 🟡 Project Service (internal)
  ##########################
  group "project-service" {
    count = 1

    network {
      port "grpc" {
        to = 50052
      }
    }

    task "project" {
      driver = "docker"

      config {
        image = "thangmicro/project-service:latest"
        ports = ["grpc"]
      }

      env {
        DB_HOST   = "mysql.service.consul"
        DB_PORT   = "3306"
        DB_USER   = "root"
        DB_PASS   = "Thang@2004"
        DB_NAME   = "taskmanager"
        MYSQL_URL = "root:Thang@2004@tcp(172.21.223.107:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "project-service"
        port = "grpc"
        tags = ["internal"]
        check {
          name     = "project-grpc"
          type     = "tcp"
          port     = "grpc"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }

  ##########################
  # 🔵 Task Service (internal)
  ##########################
  group "task-service" {
    count = 1

    network {
      port "grpc" {
        to = 50053
      }
    }

    task "task" {
      driver = "docker"

      config {
        image = "thangmicro/task-service:latest"
        ports = ["grpc"]
      }

      env {
        DB_HOST   = "mysql.service.consul"
        DB_PORT   = "3306"
        DB_USER   = "root"
        DB_PASS   = "Thang@2004"
        DB_NAME   = "taskmanager"
        MYSQL_URL = "root:Thang@2004@tcp(172.21.223.107:3306)/taskmanager?charset=utf8mb4&parseTime=True&loc=Local"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "task-service"
        port = "grpc"
        tags = ["internal"]
        check {
          name     = "task-grpc"
          type     = "tcp"
          port     = "grpc"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }

  ##########################
  # 🟢 Gateway (Public)
  ##########################
  group "gateway" {
    count = 1

    network {
      port "http" {
        static = 8080
        to     = 8080
      }
    }

    task "gateway" {
      driver = "docker"

      config {
        image = "thangmicro/gateway:latest"
        ports = ["http"]
      }

      env {
        USER_SERVICE_URL    = "172.21.223.107:27898"
        PROJECT_SERVICE_URL = "172.21.223.107:27196"
        TASK_SERVICE_URL    = "172.21.223.107:26735"
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
          name     = "gateway-tcp"
          type     = "tcp"
          port     = "http"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }
}
