#!/usr/bin/env bash
read -p 'Question ID: ' id
read -p 'Slug: ' slug
go run github.com/colin-valentini/go/cmd/leetcode init $id "$slug" $1
