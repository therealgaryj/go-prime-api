aws kms encrypt --key-id "arn:aws:kms:eu-west-2:693619507695:key/eb6d3ec2-bbce-4d71-8586-2172b49a1bbf" --plaintext fileb://<(echo -n "$1") --output text --query CiphertextBlob
