variable "version" {
  type        = string
  description = "Logstore version"
  default     = "2.7.5"
}

job "logstore" {
  datacenters = ["dc1"]

  group "compactor" {
    count = 1

    ephemeral_disk {
      size    = 1000
      sticky  = true
    }

    network {
      port "http" {}
      port "grpc" {}
    }

    service {
      name = "logstore-compactor"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "compactor"
      }

      tags = [
        "traefik.enable=true",

        "traefik.http.routers.logstore-compactor-ring.entrypoints=https",
        "traefik.http.routers.logstore-compactor-ring.rule=Host(`logstore-compactor.service.consul`) && Path(`/compactor/ring`)",
      ]

      check {
        name     = "Logstore compactor"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "compactor" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=compactor",
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

      resources {
        cpu        = 3000
        memory     = 256
        memory_max = 1024
      }
    }
  }

  group "ruler" {
    count = 1

    ephemeral_disk {
      size    = 1000
      sticky  = true
    }

    network {
      port "http" {}
      port "grpc" {}
    }

    service {
      name = "logstore-ruler"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "ruler"
      }

      tags = [
        "traefik.enable=true",

        "traefik.http.routers.logstore-ruler.entrypoints=https",
        "traefik.http.routers.logstore-ruler.rule=Host(`logstore-query-frontend.service.consul`) && (PathPrefix(`/logstore/api/v1/rules`) || PathPrefix(`/api/prom/rules`) || PathPrefix (`/prometheus/api/v1`))",

        "traefik.http.routers.logstore-ruler-ring.entrypoints=https",
        "traefik.http.routers.logstore-ruler-ring.rule=Host(`logstore-ruler.service.consul`) && Path(`/ruler/ring`)",
      ]

      check {
        name     = "Logstore ruler"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "ruler" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=ruler",
          "-config.file=/local/config.yml",
          "-config.expand-env=true",
        ]
      }

      template {
        data        = file("config.yml")
        destination = "local/config.yml"
      }

      dynamic "template" {
        for_each = fileset(".", "rules/**")

        content {
          data            = file(template.value)
          destination     = "local/${template.value}"
          left_delimiter  = "[["
          right_delimiter = "]]"
        }
      }

      template {
        data = <<-EOH
        S3_ACCESS_KEY_ID=<access_key>
        S3_SECRET_ACCESS_KEY=<secret_access_key>
        EOH

        destination = "secrets/s3.env"
        env         = true
      }

      resources {
        cpu        = 1000
        memory     = 256
        memory_max = 512
      }
    }
  }

  group "distributor" {
    count = 2

    network {
      port "http" {}
      port "grpc" {}
    }

    service {
      name = "logstore-distributor"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "distributor"
      }

      tags = [
        "traefik.enable=true",

        "traefik.http.routers.logstore-distributor.entrypoints=https",
        "traefik.http.routers.logstore-distributor.rule=Host(`logstore-distributor.service.consul`)",

        "traefik.http.routers.logstore-distributor-ring.entrypoints=https",
        "traefik.http.routers.logstore-distributor-ring.rule=Host(`logstore-distributor.cinarra.com`) && Path(`/distributor/ring`)",
      ]

      check {
        name     = "Logstore distributor"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "distributor" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=distributor",
          "-config.file=/local/config.yml",
          "-config.expand-env=true",
        ]
      }

      template {
        data        = file("config.yml")
        destination = "local/config.yml"
      }

      resources {
        cpu        = 200
        memory     = 128
        memory_max = 1024
      }
    }
  }

  group "ingester" {
    count = 2

    constraint {
      # choose your failure domain
      # must be the same as `instance_availability_zone` in config file
      distinct_property = node.unique.name
      # distinct_property = node.datacenter
      # distinct_property = attr.platform.aws.placement.availability-zone
    }

    ephemeral_disk {
      size    = 4000
      sticky  = true
    }

    network {
      port "http" {}
      port "grpc" {}
    }

    service {
      name = "logstore-ingester"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "ingester"
      }

      tags = [
        "traefik.enable=true",

        "traefik.http.routers.logstore-ingester-ring.entrypoints=https",
        "traefik.http.routers.logstore-ingester-ring.rule=Host(`logstore-ingester.service.consul`) && Path(`/ring`)",
      ]

      check {
        name     = "Logstore ingester"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "ingester" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=ingester",
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

      resources {
        cpu        = 300
        memory     = 128
        memory_max = 2048
      }
    }
  }

  group "querier" {
    count = 2

    network {
      port "http" {}
      port "grpc" {}
    }

    service {
      name = "logstore-querier"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "querier"
      }

      check {
        name     = "Logstore querier"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "50s"
        timeout  = "1s"
      }
    }

    task "querier" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=querier",
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

      resources {
        cpu        = 200
        memory     = 128
        memory_max = 2048
      }
    }
  }

  group "query-scheduler" {
    count = 2

    network {
      port "http" {}
      port "grpc" {
        to     = 9096
        static = 9096
      }
    }

    service {
      name = "logstore-query-scheduler"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "query-scheduler"
      }

      check {
        name     = "Logstore query-scheduler"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "query-scheduler" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=query-scheduler",
          "-config.file=/local/config.yml",
          "-config.expand-env=true",
        ]
      }

      template {
        data        = file("config.yml")
        destination = "local/config.yml"
      }

      resources {
        cpu        = 100
        memory     = 64
        memory_max = 128
      }
    }
  }

  group "query-frontend" {
    count = 2

    network {
      port "http" {}
      port "grpc" {}
    }

    service {
      name = "logstore-query-frontend"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "query-frontend"
      }

      tags = [
        "traefik.enable=true",

        "traefik.http.routers.logstore-query-frontend.entrypoints=https",
        "traefik.http.routers.logstore-query-frontend.rule=Host(`logstore-query-frontend.service.consul`)",
      ]

      check {
        name     = "Logstore query-frontend"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "query-frontend" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=query-frontend",
          "-config.file=/local/config.yml",
          "-config.expand-env=true",
        ]
      }

      template {
        data        = file("config.yml")
        destination = "local/config.yml"
      }

      resources {
        cpu        = 100
        memory     = 64
        memory_max = 128
      }
    }
  }

  group "index-gateway" {
    count = 1

    ephemeral_disk {
      size    = 1000
      sticky  = true
    }

    network {
      port "http" {}
      port "grpc" {
        to     = 9097
        static = 9097
      }
    }

    service {
      name = "logstore-index-gateway"
      port = "http"

      meta {
        alloc_id  = NOMAD_ALLOC_ID
        component = "index-gateway"
      }

      check {
        name     = "Logstore index-gateway"
        port     = "http"
        type     = "http"
        path     = "/ready"
        interval = "20s"
        timeout  = "1s"
      }
    }

    task "index-gateway" {
      driver       = "docker"
      user         = "nobody"
      kill_timeout = "90s"

      config {
        image = "acme/logstore:${var.version}"
        ports = [
          "http",
          "grpc",
        ]

        args = [
          "-target=index-gateway",
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

      resources {
        cpu        = 200
        memory     = 128
        memory_max = 1024
      }
    }
  }
}
