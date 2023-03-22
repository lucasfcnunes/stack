<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


/**
 * Response - Success
 * 
 * @package formance\stack\Models\Shared
 * @access public
 */
class Response
{
	#[\JMS\Serializer\Annotation\SerializedName('cursor')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\ResponseCursor')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?ResponseCursor $cursor = null;
    
    /**
     * The payload
     * 
     * @var ?array<string, mixed> $data
     */
	#[\JMS\Serializer\Annotation\SerializedName('data')]
    #[\JMS\Serializer\Annotation\Type('array<string, mixed>')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?array $data = null;
    
	public function __construct()
	{
		$this->cursor = null;
		$this->data = null;
	}
}
