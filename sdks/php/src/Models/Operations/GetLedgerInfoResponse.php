<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Operations;


class GetLedgerInfoResponse
{
	
    public string $contentType;
    
    /**
     * Error
     * 
     * @var ?\formance\stack\Models\Shared\ErrorResponse $errorResponse
     */
	
    public ?\formance\stack\Models\Shared\ErrorResponse $errorResponse = null;
    
    /**
     * OK
     * 
     * @var mixed $ledgerInfoResponse
     */
	
    public mixed $ledgerInfoResponse = null;
    
	
    public int $statusCode;
    
	
    public ?\Psr\Http\Message\ResponseInterface $rawResponse = null;
    
	public function __construct()
	{
		$this->contentType = "";
		$this->errorResponse = null;
		$this->ledgerInfoResponse = null;
		$this->statusCode = 0;
		$this->rawResponse = null;
	}
}
