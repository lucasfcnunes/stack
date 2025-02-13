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


class Script(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "plain",
        }
        
        class properties:
            plain = schemas.StrSchema
            vars = schemas.DictSchema
            reference = schemas.StrSchema
        
            @staticmethod
            def metadata() -> typing.Type['LedgerMetadata']:
                return LedgerMetadata
            __annotations__ = {
                "plain": plain,
                "vars": vars,
                "reference": reference,
                "metadata": metadata,
            }
    
    plain: MetaOapg.properties.plain
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["plain"]) -> MetaOapg.properties.plain: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["vars"]) -> MetaOapg.properties.vars: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["reference"]) -> MetaOapg.properties.reference: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["metadata"]) -> 'LedgerMetadata': ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["plain", "vars", "reference", "metadata", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["plain"]) -> MetaOapg.properties.plain: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["vars"]) -> typing.Union[MetaOapg.properties.vars, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["reference"]) -> typing.Union[MetaOapg.properties.reference, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["metadata"]) -> typing.Union['LedgerMetadata', schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["plain", "vars", "reference", "metadata", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *_args: typing.Union[dict, frozendict.frozendict, ],
        plain: typing.Union[MetaOapg.properties.plain, str, ],
        vars: typing.Union[MetaOapg.properties.vars, dict, frozendict.frozendict, schemas.Unset] = schemas.unset,
        reference: typing.Union[MetaOapg.properties.reference, str, schemas.Unset] = schemas.unset,
        metadata: typing.Union['LedgerMetadata', schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'Script':
        return super().__new__(
            cls,
            *_args,
            plain=plain,
            vars=vars,
            reference=reference,
            metadata=metadata,
            _configuration=_configuration,
            **kwargs,
        )

from Formance.model.ledger_metadata import LedgerMetadata
