<!-- Start SDK Example Usage -->
```php
<?php

declare(strict_types=1);

use formance\stack\SDK;
use \formance\stack\Models\Shared\Security;
use \formance\stack\Models\Operations\AddScopeToClientRequest;

$security = new Security();
$security->authorization = 'Bearer YOUR_ACCESS_TOKEN_HERE';

$sdk = SDK::builder()
    ->setSecurity($security);
    ->build();

try {
    $request = new AddScopeToClientRequest();
    $request->clientId = 'corrupti';
    $request->scopeId = 'provident';

    $response = $sdk->auth->addScopeToClient($request);

    if ($response->statusCode === 200) {
        // handle response
    }
} catch (Exception $e) {
    // handle exception
}
```
<!-- End SDK Example Usage -->