variable "one_flag" {
  description = "One random flag"
  type        = string
  default     = "this_default_value"
}


variable "execution_timestamp" {
  description = "Timestamp when the code was executed"
  type        = string
  default     = ""
}

locals {
  time_stamp = var.execution_timestamp != "" ? var.execution_timestamp : timestamp()
}


variable "tag_bucket_name" {
  description = "S3 bucket unique name"
  type        = string
  default     = "flugel-bucket-3ll34ndr0"
}

variable "tag_bucket_environment" {
  description = "The Environment tag to set for the S3 Bucket."
  type        = string
  default     = "Test"
}

variable "region" {
  description = "AWS Region"
  type        = string
  default     = "us-east-1"
}

variable "s3_bucket" {
 description = "AWS Bucket"
  type        = string
  default     = "random-tests-3ll34ndr0"
}
