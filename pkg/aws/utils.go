package aws


/*
// EstablishAwsSession configures the base parameters for interacting with AWS APIs
func EstablishAwsSession() *session.Session {
	session, sessionError := session.NewSession(&aws2.Config{
		Region: aws2.String("eu-west-2"),
	})

	if sessionError != nil {
		panic(fmt.Sprintf("Failed to instantiate Lambda: %v", sessionError.Error()))
	}

	return session
}

// DecryptValue hands of decryption of the cipher text to KMS
func DecryptValue(kmsClient *kms.KMS, cipherText string) string {
	decodedByteArray, decodingError := base64.StdEncoding.DecodeString(cipherText)

	if decodingError != nil {
		panic(fmt.Sprintf("Failed to instantiate Lambda: %v", decodingError.Error()))
	}

	result, decryptionError := kmsClient.Decrypt(&kms.DecryptInput{
		CiphertextBlob: decodedByteArray,
	})

	if decryptionError != nil {
		panic(fmt.Sprintf("Failed to instantiate Lambda: %v", decryptionError.Error()))
	}

	log.Debug("Decrypted value with KMS")

	return string(result.Plaintext)
}

// ConvertRequest takes an AWS Lambda request and converts to a FusionEnergy request
func ConvertRequest(request events.APIGatewayProxyRequest) types.FeApiRequest {
	return types.FeApiRequest{Method: request.HTTPMethod, Path: request.Path, Body: request.Body, Headers: request.Headers}
}

// ConvertResponse takes a FusionEnergy Response and converts to an AWS Lambda response
func ConvertResponse(feApiResponse types.FeApiResponse) (events.APIGatewayProxyResponse, error) {

	serializedStringBody := ""

	if feApiResponse.IsJsonEncoded {
		serializedStringBody = fmt.Sprintf("%v", feApiResponse.Body)
	} else if feApiResponse.Body != nil && feApiResponse.Body != "" {
		serializedBody, err := json.Marshal(feApiResponse.Body)

		if err != nil {
			return events.APIGatewayProxyResponse{StatusCode: 500, Body: "Failed to generate response"}, err
		}

		serializedStringBody = string(serializedBody)
	}

	responseHeaders := feApiResponse.Headers
	if responseHeaders == nil {
		responseHeaders = make(map[string]string)
	}

	responseHeaders["Content-Type"] = "application/json; charset=utf-8"

	response := events.APIGatewayProxyResponse{
		StatusCode:      feApiResponse.StatusCode,
		IsBase64Encoded: false,
		Body:            serializedStringBody,
		Headers: responseHeaders,
	}

	log.Info(fmt.Sprintf("Response to Lambda: %v", response))

	return response, nil
}

// ConvertMessage takes an AWS Lambda request from SQS and converts to a FusionEnergy request
func ConvertEvent(request events.SQSEvent) types.FeApiRequest {
	return types.FeApiRequest{}
}

// IsCorsPreFlightRequest takes an API Gateway request and checks to see if it matches the CORS pre-flight rules
func IsCorsPreFlightRequest(request events.APIGatewayProxyRequest) bool {
	return strings.ToLower(request.HTTPMethod) == "options"
}

// DebugLogSQSEvent If debug is enabled in the env properties, log the full inbound request
func DebugLogSQSEvent(request events.SQSEvent) {
	log.Debug(fmt.Sprintf("Number of Events in Request: %d", len(request.Records)))

	for _, event := range request.Records {
		log.Debug(event.Body)
	}
}

// DebugLogGatewayRequest If debug is enabled in the env properties, log the full inbound request
func DebugLogGatewayRequest(request events.APIGatewayProxyRequest) {
	log.Debug("Printing APIGateway Request")

	queryString := ""
	for queryName, queryValue := range request.QueryStringParameters {
		queryString += queryName + "=" + queryValue
	}

	if queryString != "" {
		queryString = "?" + queryString
	}

	log.Debug(fmt.Sprintf("%s %s%s", request.HTTPMethod, request.Path, queryString))

	for headerName, headerValue := range request.Headers {
		log.Debug(fmt.Sprintf("%s: %s", headerName, headerValue))
	}

	log.Debug(fmt.Sprintf("%s", request.Body))
}*/