#!/usr/bin/env bats

@test "accept when object is not a Pod" {
  run kwctl run -r test_data/ingress.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is accepted
  [ $(expr "$output" : '.*"allowed":true.*') -ne 0 ]
}
@test "accept when object is not a Pod and has palindrome labels" {
  run kwctl run -r test_data/ingressPalindromeLabel.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is accepted
  [ $(expr "$output" : '.*"allowed":true.*') -ne 0 ]
}
@test "accept Pod with no palindrome labels" {
  run kwctl run -r test_data/pod.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is accepted
  [ $(expr "$output" : '.*"allowed":true.*') -ne 0 ]
}
@test "accept Pod with no labels" {
  run kwctl run -r test_data/podNoLabels.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is accepted
  [ $(expr "$output" : '.*"allowed":true.*') -ne 0 ]
}
@test "accept Pod with label with a prefix" {
  run kwctl run -r test_data/podNoPalindromeLabelWithPrefix.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is accepted
  [ $(expr "$output" : '.*"allowed":true.*') -ne 0 ]
}
@test "accept Pod with one char label" {
  run kwctl run -r test_data/podOneCharLabel.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is accepted
  [ $(expr "$output" : '.*"allowed":true.*') -ne 0 ]
}
@test "reject pod with palindrome label" {
  run kwctl run -r test_data/podPalindromeLabel.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is rejected
  [ $(expr "$output" : '.*"allowed":false.*') -ne 0 ]
  [[ "$output" == *"Label level is palindrome"* ]]
}
@test "reject pod with palindrome label which has upper and lower case letters" {
  run kwctl run -r test_data/podPalindromeLabelCapitalLetters.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is rejected
  [ $(expr "$output" : '.*"allowed":false.*') -ne 0 ]
  [[ "$output" == *"Label levEL is palindrome"* ]]
}
@test "reject pod with palindrome label (even length)" {
  run kwctl run -r test_data/podPalindromeLabelEven.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is rejected
  [ $(expr "$output" : '.*"allowed":false.*') -ne 0 ]
  [[ "$output" == *"Label abccba is palindrome"* ]]
}

@test "reject pod with palindrome label with prefix" {
  run kwctl run -r test_data/podPalindromeLabelWithPrefix.json policy.wasm

  # this prints the output when one the checks below fails
  echo "output = ${output}"

  # request is rejected
  [ $(expr "$output" : '.*"allowed":false.*') -ne 0 ]
  [[ "$output" == *"Label abc/cba is palindrome"* ]]
}

