job "taskmanager" {
  datacenters = ["dc1"]
  type        = "service"

  #Internal Group (MySQL + gRPC services)
  group "internal-services" {
    count = 1

    network {
      mode = "host"
      port "mysql_user"    { static = 3306 }
      port "mysql_project" { static = 3307 }
      port "mysql_task"    { static = 3308 }
      port "user"    { static = 50051 }
      port "project" { static = 50053 }
      port "task"    { static = 50052 }
    }

    # MySQL for User Service
    task "mysql-user" {
      driver = "docker"

      config {
        image = "mysql:8.0"
        ports = ["mysql_user"]
        volumes = [
          "/opt/mysql/user:/var/lib/mysql",
        ]
        args = ["--port=3306", "--mysqlx=0"]
      }

      env {
        MYSQL_ROOT_PASSWORD = "Thang@2004"
        MYSQL_ROOT_HOST     = "%"
        MYSQL_DATABASE      = "user_db"
        MYSQL_USER          = "user_svc"
        MYSQL_PASSWORD      = "Thang412004"
      }

      resources {
        cpu    = 300
        memory = 384
      }

      service {
        name = "mysql-user"
        port = "mysql_user"

        check {
          name     = "mysql-user-tcp"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # MySQL for Project Service
    task "mysql-project" {
      driver = "docker"

      config {
        image = "mysql:8.0"
        ports = ["mysql_project"]
        volumes = [
          "/opt/mysql/project:/var/lib/mysql",
        ]
        args = ["--port=3307", "--mysqlx=0"]
      }

      env {
        MYSQL_ROOT_PASSWORD = "Thang@2004"
        MYSQL_ROOT_HOST     = "%"
        MYSQL_DATABASE      = "project_db"
        MYSQL_USER          = "project_svc"
        MYSQL_PASSWORD      = "Thang412004"
      }

      resources {
        cpu    = 300
        memory = 384
      }

      service {
        name = "mysql-project"
        port = "mysql_project"

        check {
          name     = "mysql-project-tcp"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # MySQL for Task Service
    task "mysql-task" {
      driver = "docker"

      config {
        image = "mysql:8.0"
        ports = ["mysql_task"]
        volumes = [
          "/opt/mysql/task:/var/lib/mysql",
        ]
        args = ["--port=3308", "--mysqlx=0"]
      }

      env {
        MYSQL_ROOT_PASSWORD = "Thang@2004"
        MYSQL_ROOT_HOST     = "%"
        MYSQL_DATABASE      = "task_db"
        MYSQL_USER          = "task_svc"
        MYSQL_PASSWORD      = "Thang412004"
      }

      resources {
        cpu    = 300
        memory = 384
      }

      service {
        name = "mysql-task"
        port = "mysql_task"

        check {
          name     = "mysql-task-tcp"
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
        MYSQL_URL = "user_svc:Thang412004@tcp(172.21.223.107:3306)/user_db?charset=utf8mb4&parseTime=True&loc=Local"
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

    #Project Service
    task "project-service" {
      driver = "docker"

      config {
        image = "thangmicro/project-service:latest"
        ports = ["project"]
      }

      env {
        MYSQL_URL = "project_svc:Thang412004@tcp(172.21.223.107:3307)/project_db?charset=utf8mb4&parseTime=True&loc=Local"
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
        MYSQL_URL = "task_svc:Thang412004@tcp(172.21.223.107:3308)/task_db?charset=utf8mb4&parseTime=True&loc=Local"
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

    # Using Docker bind mounts above; no host volumes required
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
        PROJECT_SERVICE_URL = "172.21.223.107:50053"
        TASK_SERVICE_URL    = "172.21.223.107:50052"
      }

      resources {
        cpu    = 300
        memory = 256
      }

      service {
        name = "gateway"
        port = "http"
        tags = ["public", "api"]

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

  # 🎨 Frontend Services
  group "frontend" {
    count = 1

    network {
      mode = "host"
      port "shell"   { static = 5170 }
      port "user"    { static = 5001 }
      port "project" { static = 5002 }
      port "task"    { static = 5003 }
    }

    # Shell App (Main Frontend)
    task "shell-app" {
      driver = "docker"

      config {
        image = "thangmicro/shell-frontend:latest"
        ports = ["shell"]
      }

      env {
        VITE_API_BASE = "http://172.21.223.107:8080"
      }

      resources {
        cpu    = 200
        memory = 128
      }

      service {
        name = "shell-app"
        port = "shell"
        tags = ["frontend", "public"]

        check {
          name     = "shell-app-http"
          type     = "http"
          path     = "/"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # User App
    task "user-app" {
      driver = "docker"

      config {
        image = "thangmicro/user-frontend:latest"
        ports = ["user"]
      }

      env {
        VITE_API_BASE = "http://172.21.223.107:8080"
      }

      resources {
        cpu    = 200
        memory = 128
      }

      service {
        name = "user-app"
        port = "user"
        tags = ["frontend"]

        check {
          name     = "user-app-http"
          type     = "http"
          path     = "/"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # Project App
    task "project-app" {
      driver = "docker"

      config {
        image = "thangmicro/project-frontend:latest"
        ports = ["project"]
      }

      env {
        VITE_API_BASE = "http://172.21.223.107:8080"
      }

      resources {
        cpu    = 200
        memory = 128
      }

      service {
        name = "project-app"
        port = "project"
        tags = ["frontend"]

        check {
          name     = "project-app-http"
          type     = "http"
          path     = "/"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }

    # Task App
    task "task-app" {
      driver = "docker"

      config {
        image = "thangmicro/task-frontend:latest"
        ports = ["task"]
      }

      env {
        VITE_API_BASE = "http://172.21.223.107:8080"
      }

      resources {
        cpu    = 200
        memory = 128
      }

      service {
        name = "task-app"
        port = "task"
        tags = ["frontend"]

        check {
          name     = "task-app-http"
          type     = "http"
          path     = "/"
          interval = "10s"
          timeout  = "2s"
        }
      }
    }
  }
}



         