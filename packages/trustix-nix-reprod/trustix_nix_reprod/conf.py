from pydantic import BaseModel
import typing
import os.path
import os


_default_attrs: typing.List[str] = (
    os.environ["DEFAULT_ATTRS"].split(":")
    if "DEFAULT_ATTRS" in os.environ
    else [
        "hello.x86_64-linux",
    ]
)


class CacheTTLSettingsModel(BaseModel):
    # Three days default
    # These documents are relatively large but should be few,
    # they are however very expensive to construct
    diff: int = 3 * 24 * 60 * 60

    # 30 minutes default
    drv_reprod: int = 30 * 60
    suggest: int = 30 * 60
    search: int = 30 * 60


class SettingsModel(BaseModel):
    trustix_rpc: str = os.environ["TRUSTIX_RPC"]
    binary_cache_proxy: str = os.environ["TRUSTIX_BINARY_CACHE_PROXY"]
    db_uri: str = os.environ["DB_URI"]

    redis_uri: str = os.environ.get("REDIS_URI", "redis://localhost")

    default_attrs: typing.List[str] = _default_attrs

    placeholder_attr: str = _default_attrs[0] if _default_attrs else "hello.x86_64-linux"

    # Npm managed
    js_store: str = os.environ.get(
        "EXTERNAL_STORE",
        os.path.join(
            os.path.dirname(os.path.abspath(__file__)),
            "js",
            "dist",
        ),
    )

    supported_systems: typing.List[str] = (
        os.environ["SUPPORTED_SYSTEMS"].split(":")
        if "SUPPORTED_SYSTEMS" in os.environ
        else [
            "aarch64-linux",
            "x86_64-linux",
            "x86_64-darwin",
        ]
    )

    cache_ttl: CacheTTLSettingsModel = CacheTTLSettingsModel()

    @property
    def tortoise_config(self) -> typing.Dict:
        return {
            "connections": {
                "default": self.db_uri,
            },
            "apps": {
                "trustix_nix_reprod": {
                    "models": ["trustix_nix_reprod.models", "aerich.models"],
                }
            },
            "use_tz": False,
            "timezone": "UTC",
        }


settings = SettingsModel()


TORTOISE_ORM = settings.tortoise_config