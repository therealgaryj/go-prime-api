aws kms decrypt --ciphertext-blob fileb://<(echo -n "$1" | base64 -d) --output text --query Plaintext | base64 -d
echo ""
