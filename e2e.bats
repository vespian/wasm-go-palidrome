#!/usr/bin/env bats

@test "reject because label key is palidrome" {
  run kwctl run policy.wasm -r test_data/pod-with-labels-palidrome.json

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request rejected
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*false') -ne 0 ]
}

@test "accept because there are no palidrome labels" {
  run kwctl run policy.wasm -r test_data/pod-with-labels-no-palidrome.json
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}

@test "accept because there are no labels" {
  run kwctl run policy.wasm -r test_data/pod-no-labels.json
  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request accepted
  [ "$status" -eq 0 ]
  [ $(expr "$output" : '.*allowed.*true') -ne 0 ]
}
