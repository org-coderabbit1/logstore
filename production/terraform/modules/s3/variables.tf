variable "region" {
  description = "AWS region"
  type        = string
  default     = "us-east-2"
}

variable "bucket_name" {
  description = "Bucket name for Logstore storage"
  type    = string
}

variable "cluster_name" {
  description = "Name of EKS cluster"
  type = string
}

variable "namespace" {
  description = "Namespace of Logstore installation"
  type        = string
}

variable "serviceaccount" {
  description = "Service account of Logstore installation"
  type        = string
  default     = "logstore"
}
