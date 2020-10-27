terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }

}


provider "aws" {
  region = var.region
}

resource "aws_s3_bucket" "flugel-it-bucket-3ll34ndr0" {
  bucket = var.s3_bucket
  acl    = "private"
  versioning {
    enabled = true
  }
  tags = {
    Name        = var.tag_bucket_name
    Environment = var.tag_bucket_environment
  }
}

resource "local_file" "test1" {
  content  = local.time_stamp
  filename = "test1.txt"
}

resource "local_file" "test2" {
  content  = local.time_stamp
  filename = "test2.txt"
}

resource "aws_s3_bucket_object" "test1" {

  bucket       = aws_s3_bucket.flugel-it-bucket-3ll34ndr0.id
  key          = "test1.txt"
  content_type = "text/plain"
  acl          = "private" # or can be "public-read"
  source       = local_file.test1.filename
}


resource "aws_s3_bucket_object" "test2" {

  bucket       = aws_s3_bucket.flugel-it-bucket-3ll34ndr0.id
  key          = "test2.txt"
  content_type = "text/plain"
  acl          = "private"
  source       = local_file.test2.filename
}



