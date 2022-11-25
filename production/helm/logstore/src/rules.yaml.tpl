---
groups:
  - name: "logstore_rules"
    rules:
    - expr: histogram_quantile(0.99, sum(rate(logstore_request_duration_seconds_bucket[1m]))
        by (le, job))
      record: job:logstore_request_duration_seconds:99quantile
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: histogram_quantile(0.50, sum(rate(logstore_request_duration_seconds_bucket[1m]))
        by (le, job))
      record: job:logstore_request_duration_seconds:50quantile
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_sum[1m])) by (job) / sum(rate(logstore_request_duration_seconds_count[1m]))
        by (job)
      record: job:logstore_request_duration_seconds:avg
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_bucket[1m])) by (le, job)
      record: job:logstore_request_duration_seconds_bucket:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_sum[1m])) by (job)
      record: job:logstore_request_duration_seconds_sum:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_count[1m])) by (job)
      record: job:logstore_request_duration_seconds_count:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: histogram_quantile(0.99, sum(rate(logstore_request_duration_seconds_bucket[1m]))
        by (le, job, route))
      record: job_route:logstore_request_duration_seconds:99quantile
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: histogram_quantile(0.50, sum(rate(logstore_request_duration_seconds_bucket[1m]))
        by (le, job, route))
      record: job_route:logstore_request_duration_seconds:50quantile
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_sum[1m])) by (job, route) / sum(rate(logstore_request_duration_seconds_count[1m]))
        by (job, route)
      record: job_route:logstore_request_duration_seconds:avg
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_bucket[1m])) by (le, job, route)
      record: job_route:logstore_request_duration_seconds_bucket:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_sum[1m])) by (job, route)
      record: job_route:logstore_request_duration_seconds_sum:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_count[1m])) by (job, route)
      record: job_route:logstore_request_duration_seconds_count:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: histogram_quantile(0.99, sum(rate(logstore_request_duration_seconds_bucket[1m]))
        by (le, namespace, job, route))
      record: namespace_job_route:logstore_request_duration_seconds:99quantile
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: histogram_quantile(0.50, sum(rate(logstore_request_duration_seconds_bucket[1m]))
        by (le, namespace, job, route))
      record: namespace_job_route:logstore_request_duration_seconds:50quantile
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_sum[1m])) by (namespace, job, route)
        / sum(rate(logstore_request_duration_seconds_count[1m])) by (namespace, job, route)
      record: namespace_job_route:logstore_request_duration_seconds:avg
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_bucket[1m])) by (le, namespace, job,
        route)
      record: namespace_job_route:logstore_request_duration_seconds_bucket:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_sum[1m])) by (namespace, job, route)
      record: namespace_job_route:logstore_request_duration_seconds_sum:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
    - expr: sum(rate(logstore_request_duration_seconds_count[1m])) by (namespace, job, route)
      record: namespace_job_route:logstore_request_duration_seconds_count:sum_rate
      labels:
        cluster: "{{ include "logstore.fullname" $ }}"
