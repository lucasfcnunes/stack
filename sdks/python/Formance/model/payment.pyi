# coding: utf-8

"""
    Formance Stack API

    Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />   # noqa: E501

    The version of the OpenAPI document: develop
    Contact: support@formance.com
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from Formance import schemas  # noqa: F401


class Payment(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "metadata",
            "adjustments",
            "scheme",
            "raw",
            "type",
            "reference",
            "accountID",
            "createdAt",
            "provider",
            "initialAmount",
            "id",
            "asset",
            "status",
        }
        
        class properties:
            id = schemas.StrSchema
            reference = schemas.StrSchema
            accountID = schemas.StrSchema
            
            
            class type(
                schemas.EnumBase,
                schemas.StrSchema
            ):
                
                @schemas.classproperty
                def PAYIN(cls):
                    return cls("PAY-IN")
                
                @schemas.classproperty
                def PAYOUT(cls):
                    return cls("PAYOUT")
                
                @schemas.classproperty
                def TRANSFER(cls):
                    return cls("TRANSFER")
                
                @schemas.classproperty
                def OTHER(cls):
                    return cls("OTHER")
        
            @staticmethod
            def provider() -> typing.Type['Connector']:
                return Connector
        
            @staticmethod
            def status() -> typing.Type['PaymentStatus']:
                return PaymentStatus
            
            
            class initialAmount(
                schemas.Int64Schema
            ):
                pass
            
            
            class scheme(
                schemas.EnumBase,
                schemas.StrSchema
            ):
                
                @schemas.classproperty
                def VISA(cls):
                    return cls("visa")
                
                @schemas.classproperty
                def MASTERCARD(cls):
                    return cls("mastercard")
                
                @schemas.classproperty
                def AMEX(cls):
                    return cls("amex")
                
                @schemas.classproperty
                def DINERS(cls):
                    return cls("diners")
                
                @schemas.classproperty
                def DISCOVER(cls):
                    return cls("discover")
                
                @schemas.classproperty
                def JCB(cls):
                    return cls("jcb")
                
                @schemas.classproperty
                def UNIONPAY(cls):
                    return cls("unionpay")
                
                @schemas.classproperty
                def SEPA_DEBIT(cls):
                    return cls("sepa debit")
                
                @schemas.classproperty
                def SEPA_CREDIT(cls):
                    return cls("sepa credit")
                
                @schemas.classproperty
                def SEPA(cls):
                    return cls("sepa")
                
                @schemas.classproperty
                def APPLE_PAY(cls):
                    return cls("apple pay")
                
                @schemas.classproperty
                def GOOGLE_PAY(cls):
                    return cls("google pay")
                
                @schemas.classproperty
                def A2A(cls):
                    return cls("a2a")
                
                @schemas.classproperty
                def ACH_DEBIT(cls):
                    return cls("ach debit")
                
                @schemas.classproperty
                def ACH(cls):
                    return cls("ach")
                
                @schemas.classproperty
                def RTP(cls):
                    return cls("rtp")
                
                @schemas.classproperty
                def UNKNOWN(cls):
                    return cls("unknown")
                
                @schemas.classproperty
                def OTHER(cls):
                    return cls("other")
            asset = schemas.StrSchema
            createdAt = schemas.DateTimeSchema
            
            
            class raw(
                schemas.DictBase,
                schemas.NoneBase,
                schemas.Schema,
                schemas.NoneFrozenDictMixin
            ):
            
            
                def __new__(
                    cls,
                    *_args: typing.Union[dict, frozendict.frozendict, None, ],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                    **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
                ) -> 'raw':
                    return super().__new__(
                        cls,
                        *_args,
                        _configuration=_configuration,
                        **kwargs,
                    )
            
            
            class adjustments(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['PaymentAdjustment']:
                        return PaymentAdjustment
            
                def __new__(
                    cls,
                    _arg: typing.Union[typing.Tuple['PaymentAdjustment'], typing.List['PaymentAdjustment']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'adjustments':
                    return super().__new__(
                        cls,
                        _arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'PaymentAdjustment':
                    return super().__getitem__(i)
        
            @staticmethod
            def metadata() -> typing.Type['PaymentMetadata']:
                return PaymentMetadata
            __annotations__ = {
                "id": id,
                "reference": reference,
                "accountID": accountID,
                "type": type,
                "provider": provider,
                "status": status,
                "initialAmount": initialAmount,
                "scheme": scheme,
                "asset": asset,
                "createdAt": createdAt,
                "raw": raw,
                "adjustments": adjustments,
                "metadata": metadata,
            }
    
    metadata: 'PaymentMetadata'
    adjustments: MetaOapg.properties.adjustments
    scheme: MetaOapg.properties.scheme
    raw: MetaOapg.properties.raw
    type: MetaOapg.properties.type
    reference: MetaOapg.properties.reference
    accountID: MetaOapg.properties.accountID
    createdAt: MetaOapg.properties.createdAt
    provider: 'Connector'
    initialAmount: MetaOapg.properties.initialAmount
    id: MetaOapg.properties.id
    asset: MetaOapg.properties.asset
    status: 'PaymentStatus'
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["reference"]) -> MetaOapg.properties.reference: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["accountID"]) -> MetaOapg.properties.accountID: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["type"]) -> MetaOapg.properties.type: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["provider"]) -> 'Connector': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["status"]) -> 'PaymentStatus': ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["initialAmount"]) -> MetaOapg.properties.initialAmount: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["scheme"]) -> MetaOapg.properties.scheme: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["asset"]) -> MetaOapg.properties.asset: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["createdAt"]) -> MetaOapg.properties.createdAt: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["raw"]) -> MetaOapg.properties.raw: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["adjustments"]) -> MetaOapg.properties.adjustments: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["metadata"]) -> 'PaymentMetadata': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["id", "reference", "accountID", "type", "provider", "status", "initialAmount", "scheme", "asset", "createdAt", "raw", "adjustments", "metadata", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["id"]) -> MetaOapg.properties.id: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["reference"]) -> MetaOapg.properties.reference: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["accountID"]) -> MetaOapg.properties.accountID: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["type"]) -> MetaOapg.properties.type: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["provider"]) -> 'Connector': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["status"]) -> 'PaymentStatus': ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["initialAmount"]) -> MetaOapg.properties.initialAmount: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["scheme"]) -> MetaOapg.properties.scheme: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["asset"]) -> MetaOapg.properties.asset: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["createdAt"]) -> MetaOapg.properties.createdAt: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["raw"]) -> MetaOapg.properties.raw: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["adjustments"]) -> MetaOapg.properties.adjustments: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["metadata"]) -> 'PaymentMetadata': ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["id", "reference", "accountID", "type", "provider", "status", "initialAmount", "scheme", "asset", "createdAt", "raw", "adjustments", "metadata", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        metadata: 'PaymentMetadata',
        adjustments: typing.Union[MetaOapg.properties.adjustments, list, tuple, ],
        scheme: typing.Union[MetaOapg.properties.scheme, str, ],
        raw: typing.Union[MetaOapg.properties.raw, dict, frozendict.frozendict, None, ],
        type: typing.Union[MetaOapg.properties.type, str, ],
        reference: typing.Union[MetaOapg.properties.reference, str, ],
        accountID: typing.Union[MetaOapg.properties.accountID, str, ],
        createdAt: typing.Union[MetaOapg.properties.createdAt, str, datetime, ],
        provider: 'Connector',
        initialAmount: typing.Union[MetaOapg.properties.initialAmount, decimal.Decimal, int, ],
        id: typing.Union[MetaOapg.properties.id, str, ],
        asset: typing.Union[MetaOapg.properties.asset, str, ],
        status: 'PaymentStatus',
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'Payment':
        return super().__new__(
            cls,
            *_args,
            metadata=metadata,
            adjustments=adjustments,
            scheme=scheme,
            raw=raw,
            type=type,
            reference=reference,
            accountID=accountID,
            createdAt=createdAt,
            provider=provider,
            initialAmount=initialAmount,
            id=id,
            asset=asset,
            status=status,
            _configuration=_configuration,
            **kwargs,
        )

from Formance.model.connector import Connector
from Formance.model.payment_adjustment import PaymentAdjustment
from Formance.model.payment_metadata import PaymentMetadata
from Formance.model.payment_status import PaymentStatus
