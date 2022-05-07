#!/bin/sh

function dry() {
  ./sqldef -h localhost -u root goca -p root --dry-run < schema.sql
}

function apply() {
  ./sqldef -h localhost -u root goca -p root < schema.sql
}

$1
