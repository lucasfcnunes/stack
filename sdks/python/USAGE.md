<!-- Start SDK Example Usage -->
```python
import sdk
from sdk.models import operations, shared

s = sdk.SDK(
    security=shared.Security(
        authorization="Bearer YOUR_ACCESS_TOKEN_HERE",
    ),
)


req = operations.AddScopeToClientRequest(
    client_id="corrupti",
    scope_id="provident",
)
    
res = s.auth.add_scope_to_client(req)

if res.status_code == 200:
    # handle response
```
<!-- End SDK Example Usage -->