variable "project_id" {
  type    = string
  default = "satori-codingtest-2"
}

variable "region" {
  type    = string
  default = "asia-northeast1"
}

variable "zone" {
  type    = string
  default = "asia-northeast1-a"
}

variable "api_dir" {
  type    = string
  default = "../api"
}

variable "api_output_path" {
  type    = string
  default = "/tmp/function.zip"
}

variable "web_container_uri" {
  type    = string
  default = "asia-northeast1-docker.pkg.dev/satori-codingtest-2/satori-codingtest-2-web/web:latest"
}

variable "functions_go_runtime" {
  type    = string
  default = "go119"
}
