<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack;

class Webhooks 
{

	// SDK private variables namespaced with _ to avoid conflicts with API models
	private \GuzzleHttp\ClientInterface $_defaultClient;
	private \GuzzleHttp\ClientInterface $_securityClient;
	private string $_serverUrl;
	private string $_language;
	private string $_sdkVersion;
	private string $_genVersion;	

	/**
	 * @param \GuzzleHttp\ClientInterface $defaultClient
	 * @param \GuzzleHttp\ClientInterface $securityClient
	 * @param string $serverUrl
	 * @param string $language
	 * @param string $sdkVersion
	 * @param string $genVersion
	 */
	public function __construct(\GuzzleHttp\ClientInterface $defaultClient, \GuzzleHttp\ClientInterface $securityClient, string $serverUrl, string $language, string $sdkVersion, string $genVersion)
	{
		$this->_defaultClient = $defaultClient;
		$this->_securityClient = $securityClient;
		$this->_serverUrl = $serverUrl;
		$this->_language = $language;
		$this->_sdkVersion = $sdkVersion;
		$this->_genVersion = $genVersion;
	}
	
    /**
     * Activate one config
     * 
     * Activate a webhooks config by ID, to start receiving webhooks to its endpoint.
     * 
     * @param \formance\stack\Models\Operations\ActivateConfigRequest $request
     * @return \formance\stack\Models\Operations\ActivateConfigResponse
     */
	public function activateConfig(
        \formance\stack\Models\Operations\ActivateConfigRequest $request,
    ): \formance\stack\Models\Operations\ActivateConfigResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs/{id}/activate', \formance\stack\Models\Operations\ActivateConfigRequest::class, $request);
        
        $options = ['http_errors' => false];
        
        $httpResponse = $this->_securityClient->request('PUT', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\ActivateConfigResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
            if (Utils\Utils::matchContentType($contentType, 'application/json')) {
                $serializer = Utils\JSON::createSerializer();
                $response->configResponse = $serializer->deserialize((string)$httpResponse->getBody(), 'formance\stack\Models\Shared\ConfigResponse', 'json');
            }
        }
        else if ($httpResponse->getStatusCode() === 304) {
        }

        return $response;
    }
	
    /**
     * Change the signing secret of a config
     * 
     * Change the signing secret of the endpoint of a webhooks config.
     * 
     * If not passed or empty, a secret is automatically generated.
     * The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)
     * 
     * 
     * @param \formance\stack\Models\Operations\ChangeConfigSecretRequest $request
     * @return \formance\stack\Models\Operations\ChangeConfigSecretResponse
     */
	public function changeConfigSecret(
        \formance\stack\Models\Operations\ChangeConfigSecretRequest $request,
    ): \formance\stack\Models\Operations\ChangeConfigSecretResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs/{id}/secret/change', \formance\stack\Models\Operations\ChangeConfigSecretRequest::class, $request);
        
        $options = ['http_errors' => false];
        $body = Utils\Utils::serializeRequestBody($request, "configChangeSecret", "json");
        $options = array_merge_recursive($options, $body);
        
        $httpResponse = $this->_securityClient->request('PUT', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\ChangeConfigSecretResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
            if (Utils\Utils::matchContentType($contentType, 'application/json')) {
                $serializer = Utils\JSON::createSerializer();
                $response->configResponse = $serializer->deserialize((string)$httpResponse->getBody(), 'formance\stack\Models\Shared\ConfigResponse', 'json');
            }
        }

        return $response;
    }
	
    /**
     * Deactivate one config
     * 
     * Deactivate a webhooks config by ID, to stop receiving webhooks to its endpoint.
     * 
     * @param \formance\stack\Models\Operations\DeactivateConfigRequest $request
     * @return \formance\stack\Models\Operations\DeactivateConfigResponse
     */
	public function deactivateConfig(
        \formance\stack\Models\Operations\DeactivateConfigRequest $request,
    ): \formance\stack\Models\Operations\DeactivateConfigResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs/{id}/deactivate', \formance\stack\Models\Operations\DeactivateConfigRequest::class, $request);
        
        $options = ['http_errors' => false];
        
        $httpResponse = $this->_securityClient->request('PUT', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\DeactivateConfigResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
            if (Utils\Utils::matchContentType($contentType, 'application/json')) {
                $serializer = Utils\JSON::createSerializer();
                $response->configResponse = $serializer->deserialize((string)$httpResponse->getBody(), 'formance\stack\Models\Shared\ConfigResponse', 'json');
            }
        }
        else if ($httpResponse->getStatusCode() === 304) {
        }

        return $response;
    }
	
    /**
     * Delete one config
     * 
     * Delete a webhooks config by ID.
     * 
     * @param \formance\stack\Models\Operations\DeleteConfigRequest $request
     * @return \formance\stack\Models\Operations\DeleteConfigResponse
     */
	public function deleteConfig(
        \formance\stack\Models\Operations\DeleteConfigRequest $request,
    ): \formance\stack\Models\Operations\DeleteConfigResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs/{id}', \formance\stack\Models\Operations\DeleteConfigRequest::class, $request);
        
        $options = ['http_errors' => false];
        
        $httpResponse = $this->_securityClient->request('DELETE', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\DeleteConfigResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
        }

        return $response;
    }
	
    /**
     * Get many configs
     * 
     * Sorted by updated date descending
     * 
     * @param \formance\stack\Models\Operations\GetManyConfigsRequest $request
     * @return \formance\stack\Models\Operations\GetManyConfigsResponse
     */
	public function getManyConfigs(
        \formance\stack\Models\Operations\GetManyConfigsRequest $request,
    ): \formance\stack\Models\Operations\GetManyConfigsResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs');
        
        $options = ['http_errors' => false];
        $options = array_merge_recursive($options, Utils\Utils::getQueryParams(\formance\stack\Models\Operations\GetManyConfigsRequest::class, $request, null));
        
        $httpResponse = $this->_securityClient->request('GET', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\GetManyConfigsResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
            if (Utils\Utils::matchContentType($contentType, 'application/json')) {
                $serializer = Utils\JSON::createSerializer();
                $response->configsResponse = $serializer->deserialize((string)$httpResponse->getBody(), 'formance\stack\Models\Shared\ConfigsResponse', 'json');
            }
        }

        return $response;
    }
	
    /**
     * Insert a new config
     * 
     * Insert a new webhooks config.
     * 
     * The endpoint should be a valid https URL and be unique.
     * 
     * The secret is the endpoint's verification secret.
     * If not passed or empty, a secret is automatically generated.
     * The format is a random string of bytes of size 24, base64 encoded. (larger size after encoding)
     * 
     * All eventTypes are converted to lower-case when inserted.
     * 
     * 
     * @param \formance\stack\Models\Shared\ConfigUser $request
     * @return \formance\stack\Models\Operations\InsertConfigResponse
     */
	public function insertConfig(
        \formance\stack\Models\Shared\ConfigUser $request,
    ): \formance\stack\Models\Operations\InsertConfigResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs');
        
        $options = ['http_errors' => false];
        $body = Utils\Utils::serializeRequestBody($request, "request", "json");
        if ($body === null) {
            throw new \Exception('Request body is required');
        }
        $options = array_merge_recursive($options, $body);
        
        $httpResponse = $this->_securityClient->request('POST', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\InsertConfigResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
            if (Utils\Utils::matchContentType($contentType, 'application/json')) {
                $serializer = Utils\JSON::createSerializer();
                $response->configResponse = $serializer->deserialize((string)$httpResponse->getBody(), 'formance\stack\Models\Shared\ConfigResponse', 'json');
            }
        }
        else if ($httpResponse->getStatusCode() === 400) {
            if (Utils\Utils::matchContentType($contentType, 'text/plain')) {
                $response->insertConfig400TextPlainString = $httpResponse->getBody()->getContents();
            }
        }

        return $response;
    }
	
    /**
     * Test one config
     * 
     * Test a config by sending a webhook to its endpoint.
     * 
     * @param \formance\stack\Models\Operations\TestConfigRequest $request
     * @return \formance\stack\Models\Operations\TestConfigResponse
     */
	public function testConfig(
        \formance\stack\Models\Operations\TestConfigRequest $request,
    ): \formance\stack\Models\Operations\TestConfigResponse
    {
        $baseUrl = $this->_serverUrl;
        $url = Utils\Utils::generateUrl($baseUrl, '/api/webhooks/configs/{id}/test', \formance\stack\Models\Operations\TestConfigRequest::class, $request);
        
        $options = ['http_errors' => false];
        
        $httpResponse = $this->_securityClient->request('GET', $url, $options);
        
        $contentType = $httpResponse->getHeader('Content-Type')[0] ?? '';

        $response = new \formance\stack\Models\Operations\TestConfigResponse();
        $response->statusCode = $httpResponse->getStatusCode();
        $response->contentType = $contentType;
        $response->rawResponse = $httpResponse;
        
        if ($httpResponse->getStatusCode() === 200) {
            if (Utils\Utils::matchContentType($contentType, 'application/json')) {
                $serializer = Utils\JSON::createSerializer();
                $response->attemptResponse = $serializer->deserialize((string)$httpResponse->getBody(), 'formance\stack\Models\Shared\AttemptResponse', 'json');
            }
        }

        return $response;
    }
}