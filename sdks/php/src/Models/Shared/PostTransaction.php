<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


/**
 * PostTransaction - The request body must contain at least one of the following objects:
 * 
 *   - `postings`: suitable for simple transactions
 *   - `script`: enabling more complex transactions with Numscript
 * 
 * 
 * @package formance\stack\Models\Shared
 * @access public
 */
class PostTransaction
{
    /**
     * $metadata
     * 
     * @var ?array<string, mixed> $metadata
     */
	#[\JMS\Serializer\Annotation\SerializedName('metadata')]
    #[\JMS\Serializer\Annotation\Type('array<string, mixed>')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?array $metadata = null;
    
    /**
     * $postings
     * 
     * @var ?array<\formance\stack\Models\Shared\Posting> $postings
     */
	#[\JMS\Serializer\Annotation\SerializedName('postings')]
    #[\JMS\Serializer\Annotation\Type('array<formance\stack\Models\Shared\Posting>')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?array $postings = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('reference')]
    #[\JMS\Serializer\Annotation\Type('string')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?string $reference = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('script')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\PostTransactionScript')]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?PostTransactionScript $script = null;
    
	#[\JMS\Serializer\Annotation\SerializedName('timestamp')]
    #[\JMS\Serializer\Annotation\Type("DateTime<'Y-m-d\TH:i:s.up'>")]
    #[\JMS\Serializer\Annotation\SkipWhenEmpty]
    public ?\DateTime $timestamp = null;
    
	public function __construct()
	{
		$this->metadata = null;
		$this->postings = null;
		$this->reference = null;
		$this->script = null;
		$this->timestamp = null;
	}
}
