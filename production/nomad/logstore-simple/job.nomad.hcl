variable "version" {
  type        = string
  description = "Logstore version"
  default     = "2.7.5"
}

job "logstore" {
  datacenters = ["dc1"]

  group "read" {
    count = 1

    ephemeral_disk {
      size   = 1000
      sticky = true
    }

    network {
      port "http" {}
      port "grpc" {}
    }

    task "read" {
      driver = "docker"
      user   = "nobody"

      config {
        image = "acme/logstore:${var.version}"

        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=read",
          "-config.file=/local/config.yml",
          "-config.expand-env=true",
        ]
      }

      template {
        data        = file("config.yml")
        destination = "local/config.yml"
      }

      template {
        data = <<-EOH
        S3_ACCESS_KEY_ID=<access_key>
        S3_SECRET_ACCESS_KEY=<secret_access_key>
        EOH

        destination = "secrets/s3.env"
        env         = true
      }

      service {
        name = "logstore-read"
        port = "http"

        tags = [
          "traefik.enable=true",
          "traefik.http.routers.logstore-read.entrypoints=https",
          "traefik.http.routers.logstore-read.rule=Host(`logstore-read.service.consul`)",
        ]

        check {
          name     = "Logstore read"
          port     = "http"
          type     = "http"
          path     = "/ready"
          interval = "20s"
          timeout  = "1s"

          initial_status = "passing"
        }
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }

  group "write" {
    count = 2

    ephemeral_disk {
      size   = 1000
      sticky = true
    }

    network {
      port "http" {}
      port "grpc" {}
    }

    task "write" {
      driver = "docker"
      user   = "nobody"

      config {
        image = "acme/logstore:${var.version}"

        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=write",
          "-config.file=/local/config.yml",
          "-config.expand-env=true",
        ]
      }

      template {
        data        = file("config.yml")
        destination = "local/config.yml"
      }

      template {
        data = <<-EOH
        S3_ACCESS_KEY_ID=<access_key>
        S3_SECRET_ACCESS_KEY=<secret_access_key>
        EOH

        destination = "secrets/s3.env"
        env         = true
      }

      service {
        name = "logstore-write"
        port = "http"

        tags = [
          "traefik.enable=true",
          "traefik.http.routers.logstore-write.entrypoints=https",
          "traefik.http.routers.logstore-write.rule=Host(`logstore-write.service.consul`)",
        ]

        check {
          name     = "Logstore write"
          port     = "http"
          type     = "http"
          path     = "/ready"
          interval = "20s"
          timeout  = "1s"

          initial_status = "passing"
        }
      }

      resources {
        cpu    = 500
        memory = 256
      }
    }
  }
}
