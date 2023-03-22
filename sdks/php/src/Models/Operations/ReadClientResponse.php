<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;


class ReadClientResponse
{
	
    public string $contentType;
    
    /**
     * Retrieved client
     * 
     * @var ?\formance\stack\Models\Shared\ReadClientResponse $readClientResponse
     */
	
    public ?\formance\stack\Models\Shared\ReadClientResponse $readClientResponse = null;
    
	
    public int $statusCode;
    
	
    public ?\Psr\Http\Message\ResponseInterface $rawResponse = null;
    
	public function __construct()
	{
		$this->contentType = "";
		$this->readClientResponse = null;
		$this->statusCode = 0;
		$this->rawResponse = null;
	}
}
