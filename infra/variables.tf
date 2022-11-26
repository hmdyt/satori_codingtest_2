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
