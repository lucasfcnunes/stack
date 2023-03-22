<?php

/**
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

declare(strict_types=1);

namespace formance\stack\Models\Shared;


/**
 * AccountsCursor - OK
 * 
 * @package formance\stack\Models\Shared
 * @access public
 */
class AccountsCursor
{
	#[\JMS\Serializer\Annotation\SerializedName('cursor')]
    #[\JMS\Serializer\Annotation\Type('formance\stack\Models\Shared\AccountsCursorCursor')]
    public AccountsCursorCursor $cursor;
    
	public function __construct()
	{
		$this->cursor = new \formance\stack\Models\Shared\AccountsCursorCursor();
	}
}
